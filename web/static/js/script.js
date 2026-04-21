/* ── Data ─────────────────────────────────────────── */
let state = {
  drinks: [], // { id, name, type, measure, qty, time, icon }
  foods: [], // { name, icon, time }
  sessionStart: null,
  selectedDrink: null,
  selectedFoods: new Set(),
  drinkType: "alcoholic",
};

const ALCO_DRINKS = [
  { label: "Beer", icon: "🍺", points: 10 },
  { label: "Whiskey", icon: "🥃", points: 15 },
  { label: "Rum", icon: "🍹", points: 14 },
  { label: "Shots", icon: "🥂", points: 20 },
  { label: "Wine", icon: "🍷", points: 12 },
  { label: "Vodka", icon: "🫗", points: 16 },
];
const NON_DRINKS = [
  { label: "Water", icon: "💧", points: 0 },
  { label: "Juice", icon: "🧃", points: 2 },
  { label: "Soda", icon: "🥤", points: 1 },
  { label: "Tea", icon: "🍵", points: 1 },
  { label: "Coffee", icon: "☕", points: 1 },
  { label: "Mocktail", icon: "🍹", points: 3 },
];
const FOODS = [
  { label: "Peanuts", icon: "🥜" },
  { label: "Tandoori Chicken", icon: "🍗" },
  { label: "Chips", icon: "🍟" },
  { label: "Nachos", icon: "🌮" },
  { label: "Paneer Tikka", icon: "🧆" },
  { label: "Momos", icon: "🥟" },
  { label: "Peri Peri Fries", icon: "🍟" },
  { label: "Biryani", icon: "🍛" },
  { label: "Pizza", icon: "🍕" },
  { label: "Kebab", icon: "🍢" },
];

/* ── Tab switching ───────────────────────────────── */
function switchTab(name) {
  document
    .querySelectorAll(".page")
    .forEach((p) => p.classList.remove("active"));
  document
    .querySelectorAll(".tab")
    .forEach((t) => t.classList.remove("active"));
  document
    .querySelectorAll(".nav-item")
    .forEach((n) => n.classList.remove("active"));

  document.getElementById("page-" + name).classList.add("active");
  document
    .querySelector(`.tab[onclick="switchTab('${name}')"]`)
    .classList.add("active");
  document.getElementById("nav-" + name).classList.add("active");
}

/* ── Drink type toggle ───────────────────────────── */
function setType(type) {
  state.drinkType = type;
  state.selectedDrink = null;
  document
    .getElementById("typeAlc")
    .classList.toggle("active", type === "alcoholic");
  document
    .getElementById("typeNon")
    .classList.toggle("active", type === "non-alcoholic");
  renderDrinkChips();
}

function renderDrinkChips() {
  const arr = state.drinkType === "alcoholic" ? ALCO_DRINKS : NON_DRINKS;
  const el = document.getElementById("drinkChips");
  el.innerHTML = arr
    .map(
      (d, i) => `
    <div class="drink-chip ${state.selectedDrink === i ? "selected" : ""}" onclick="selectDrink(${i})">
      <div class="drink-chip-icon">${d.icon}</div>
      <div class="drink-chip-label">${d.label}</div>
    </div>
  `,
    )
    .join("");
}

function selectDrink(i) {
  state.selectedDrink = state.selectedDrink === i ? null : i;
  const arr = state.drinkType === "alcoholic" ? ALCO_DRINKS : NON_DRINKS;
  if (state.selectedDrink !== null) {
    document.getElementById("drinkName").value = arr[i].label;
  }
  renderDrinkChips();
}

/* ── Log a drink ─────────────────────────────────── */
function logDrink() {
  const name = document.getElementById("drinkName").value.trim() || "Drink";
  const qty = parseInt(document.getElementById("drinkQty").value) || 1;
  const measure = document.getElementById("drinkMeasure").value;
  const type = state.drinkType;
  const arr = type === "alcoholic" ? ALCO_DRINKS : NON_DRINKS;
  const sel =
    state.selectedDrink !== null
      ? arr[state.selectedDrink]
      : { icon: type === "alcoholic" ? "🍺" : "🥤", points: 5 };

  if (!state.sessionStart) state.sessionStart = Date.now();

  const drink = {
    id: Date.now(),
    name,
    type,
    measure,
    qty,
    icon: sel.icon,
    points: sel.points * qty,
    time: new Date(),
  };
  state.drinks.push(drink);

  updateStats();
  renderHistory();
  showToast(`${sel.icon} ${name} logged!`);

  // Reset
  document.getElementById("drinkName").value = "";
  document.getElementById("drinkQty").value = "1";
  state.selectedDrink = null;
  renderDrinkChips();
}

/* ── Log food ────────────────────────────────────── */
function renderFoodChips() {
  FOODS;
  const el = document.getElementById("foodChips");
  el.innerHTML = FOODS.map(
    (f, i) => `
    <div class="food-chip ${state.selectedFoods.has(i) ? "selected" : ""}" onclick="toggleFood(${i})">
      <span>${f.icon}</span>${f.label}
    </div>
  `,
  ).join("");
}

function toggleFood(i) {
  if (state.selectedFoods.has(i)) state.selectedFoods.delete(i);
  else state.selectedFoods.add(i);
  renderFoodChips();
}

function logFood() {
  const custom = document.getElementById("foodName").value.trim();
  const toAdd = [];

  state.selectedFoods.forEach((i) =>
    toAdd.push({ name: FOODS[i].label, icon: FOODS[i].icon, time: new Date() }),
  );
  if (custom) toAdd.push({ name: custom, icon: "🍽️", time: new Date() });

  if (!toAdd.length) {
    showToast("Pick or type a food first");
    return;
  }

  state.foods.push(...toAdd);
  renderFoodLog();
  showToast("🍽️ Food logged!");
  document.getElementById("foodName").value = "";
  state.selectedFoods.clear();
  renderFoodChips();
  updateStats();
}

function renderFoodLog() {
  const el = document.getElementById("foodList");
  if (!state.foods.length) {
    el.innerHTML = `<div class="empty-state"><div class="empty-icon">🍟</div><div class="empty-text">No food logged</div><div class="empty-sub">Add your snacks and sides</div></div>`;
    return;
  }
  el.innerHTML = [...state.foods]
    .reverse()
    .map(
      (f) => `
    <div class="food-log-item">
      <div class="food-log-icon">${f.icon}</div>
      <div class="food-log-name">${f.name}</div>
      <div class="food-log-time">${formatTime(f.time)}</div>
    </div>
  `,
    )
    .join("");
}

/* ── Stats ───────────────────────────────────────── */
function updateStats() {
  const total = state.drinks.reduce((s, d) => s + d.qty, 0);
  const points = state.drinks.reduce((s, d) => s + d.points, 0);
  const now = Date.now();
  const mins = state.sessionStart
    ? Math.round((now - state.sessionStart) / 60000)
    : 0;
  const pace = total > 0 ? (mins / total).toFixed(1) : "–";
  const pr = state.drinks.length >= 2 ? computePR() : "–";

  // Hero grid
  document.getElementById("stat-total").innerHTML = total;
  document.getElementById("stat-points").innerHTML = points;
  document.getElementById("stat-time").innerHTML =
    `${formatDuration(mins)}<span class="stat-unit"></span>`;
  document.getElementById("stat-pace").innerHTML =
    `${pace}<span class="stat-unit"> min/drink</span>`;
  document.getElementById("stat-pr").innerHTML =
    `${pr}<span class="stat-unit"> min</span>`;

  // Share card
  document.getElementById("share-total").innerHTML = total;
  document.getElementById("share-time").innerHTML =
    `${formatDuration(mins)}<span class="unit"></span>`;
  document.getElementById("share-pace").innerHTML =
    `${pace}<span class="unit"> min/drink</span>`;

  renderBreakdown();
  renderTimeline();
}

function computePR() {
  let min = Infinity;
  for (let i = 1; i < state.drinks.length; i++) {
    const diff = (state.drinks[i].time - state.drinks[i - 1].time) / 60000;
    if (diff < min) min = diff;
  }
  return min === Infinity ? "–" : min.toFixed(1);
}

function renderBreakdown() {
  const el = document.getElementById("breakdown-bars");
  if (!state.drinks.length) {
    el.innerHTML = `<div class="empty-state" style="padding:20px 0"><div class="empty-icon" style="font-size:32px">📊</div><div class="empty-sub">Log drinks to see breakdown</div></div>`;
    return;
  }
  const counts = {};
  state.drinks.forEach((d) => {
    counts[d.name] = (counts[d.name] || 0) + d.qty;
  });
  const maxVal = Math.max(...Object.values(counts));
  el.innerHTML = Object.entries(counts)
    .sort((a, b) => b[1] - a[1])
    .map(([name, cnt]) => {
      const pct = ((cnt / maxVal) * 100).toFixed(0);
      const drink = state.drinks.find((d) => d.name === name);
      const isNon = drink && drink.type === "non-alcoholic";
      return `
      <div class="bar-row">
        <div class="bar-label">
          <span class="bar-label-name">${drink ? drink.icon : ""} ${name}</span>
          <span class="bar-label-count">${cnt} drink${cnt > 1 ? "s" : ""}</span>
        </div>
        <div class="bar-track"><div class="bar-fill ${isNon ? "green-bar" : ""}" style="width:${pct}%"></div></div>
      </div>`;
    })
    .join("");
}

function renderTimeline() {
  const el = document.getElementById("timeline");
  if (!state.drinks.length && !state.foods.length) {
    el.innerHTML = `<div class="empty-state" style="padding:20px 0"><div class="empty-icon" style="font-size:32px">🕐</div><div class="empty-sub">Your timeline will appear here</div></div>`;
    return;
  }
  const all = [
    ...state.drinks.map((d) => ({ ...d, category: "drink" })),
    ...state.foods.map((f) => ({ ...f, category: "food" })),
  ];
  all.sort((a, b) => a.time - b.time);
  el.innerHTML = all
    .map(
      (item) => `
    <div class="timeline-item">
      <div class="timeline-dot ${item.category === "food" ? "green-dot" : ""}"></div>
      <div class="timeline-time">${formatTime(item.time)}</div>
      <div class="timeline-text">${item.icon} ${item.name}</div>
      <div class="timeline-sub">${item.category === "drink" ? `${item.qty} × ${item.measure}` : "Food / Side"}</div>
    </div>
  `,
    )
    .join("");
}

function renderHistory() {
  const el = document.getElementById("historyList");
  if (!state.drinks.length) {
    el.innerHTML = `<div class="empty-state"><div class="empty-icon">🍹</div><div class="empty-text">Nothing yet</div><div class="empty-sub">Log your first drink</div></div>`;
    return;
  }
  el.innerHTML = [...state.drinks]
    .reverse()
    .map(
      (d) => `
    <div class="log-item">
      <div class="log-item-icon ${d.type === "non-alcoholic" ? "non-alc" : ""}">${d.icon}</div>
      <div class="log-item-info">
        <div class="log-item-name">${d.name}</div>
        <div class="log-item-meta">${d.qty} × ${d.measure} &nbsp;·&nbsp; <span class="log-item-alc-tag ${d.type === "alcoholic" ? "alc-tag" : "non-alc-tag"}">${d.type === "alcoholic" ? "ALC" : "NON-ALC"}</span></div>
      </div>
      <div class="log-item-time">${formatTime(d.time)}</div>
    </div>
  `,
    )
    .join("");
}

/* ── Helpers ─────────────────────────────────────── */
function formatTime(d) {
  return d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
}
function formatDuration(mins) {
  if (mins < 60) return `${mins}m`;
  return `${Math.floor(mins / 60)}h ${mins % 60}m`;
}

let toastTimer;
function showToast(msg) {
  const el = document.getElementById("toast");
  el.textContent = msg;
  el.classList.add("show");
  clearTimeout(toastTimer);
  toastTimer = setTimeout(() => el.classList.remove("show"), 2200);
}

/* Live timer */
setInterval(() => {
  if (state.sessionStart) updateStats();
}, 30000);

/* ── Init ────────────────────────────────────────── */
renderDrinkChips();
renderFoodChips();
