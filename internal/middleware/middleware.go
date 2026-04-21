package middleware

import (
	"drinklytics/internal/config"
	"drinklytics/internal/models"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"

	"github.com/flosch/pongo2/v6"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func HashPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TokenGenerator(user *models.User) (string, error) {

	claims := Claims{
		Role:   user.Role,
		UserID: int64(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "drinklytics_fin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TokenValidator(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		log.Println("Authorization Failed: ", fmt.Errorf("Missing token"))
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.BuildErrorResponse("Bad Request", nil, fmt.Errorf("Missing token")))
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Authorization Failed: ", fmt.Errorf("Invalid token"))
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.BuildErrorResponse("Bad Request", nil, fmt.Errorf("Invalid token")))
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		log.Println("Authorization Failed: ", fmt.Errorf("Invalid claims"))
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.BuildErrorResponse("Bad Request", nil, fmt.Errorf("Invalid claims")))
		return

	}

	c.Set("userId", claims.UserID)
	c.Set("role", claims.Role)

	c.Next()

}

func Authorize(allowedRoles ...config.Role) gin.HandlerFunc {
	return func(c *gin.Context) {

		role := c.GetString("role")
		for _, r := range allowedRoles {
			if string(r) == role {
				c.Next()
				return
			}
		}
		log.Println("Authorization Failed: ", fmt.Errorf("Access Denied"))
		c.AbortWithStatusJSON(http.StatusForbidden, models.BuildErrorResponse("Bad Request", nil, fmt.Errorf("Access Denied")))
	}
}

// Rate limitter
type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var clients = make(map[string]*client)
var mu sync.Mutex

// Cleanup old and inactive clients
func CleanClients() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, c := range clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		defer mu.Unlock()

		var limiter *rate.Limiter

		if c, exists := clients[ip]; exists {
			c.lastSeen = time.Now()
			limiter = c.limiter
		} else {
			// 5 requests per second, burst up to 10.
			limiter = rate.NewLimiter(10, 20)

			// 300 requests per minute no burst
			// limiter = rate.NewLimiter(rate.Every(time.Minute/300), 1)
			clients[ip] = &client{
				limiter:  limiter,
				lastSeen: time.Now(),
			}
		}

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, models.BuildErrorResponse("Bad Request", nil, fmt.Errorf("Too many requests")))
			return
		}

		c.Next()
	}
}

var tplSet = pongo2.NewSet("default", pongo2.MustNewLocalFileSystemLoader("web/templates"))

func Render(c *gin.Context, name string, ctx pongo2.Context) {
	tpl, err := tplSet.FromFile(name)
	if err != nil {
		c.String(500, "Template load error: %v", err)
		return
	}
	fmt.Println("ctx", ctx)
	out, err := tpl.Execute(ctx)
	if err != nil {
		c.String(500, "Template Render error: %v", err)
		return
	}

	c.Data(200, "text/html; charset=utf-8", []byte(out))
}
