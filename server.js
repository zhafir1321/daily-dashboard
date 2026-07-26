// Daily Dashboard — backend server (Express + MySQL)
require("dotenv").config();
const express = require("express");
const mysql = require("mysql2/promise");
const path = require("path");

const app = express();
app.use(express.json());
app.use(express.static(__dirname)); // serves dashboard.html

const pool = mysql.createPool({
  host: process.env.DB_HOST || "localhost",
  port: process.env.DB_PORT || 3306,
  user: process.env.DB_USER || "root",
  password: process.env.DB_PASSWORD || "",
  database: process.env.DB_NAME || "daily_dashboard",
  waitForConnections: true,
  connectionLimit: 10,
});

const uid = () =>
  Date.now().toString(36) + Math.random().toString(36).slice(2, 6);

// tiny async wrapper so errors return JSON instead of crashing
const h = (fn) => (req, res) =>
  fn(req, res).catch((e) => {
    console.error(e);
    res.status(500).json({ error: e.message });
  });

/* ---------------- TODOS ---------------- */
app.get("/api/todos", h(async (req, res) => {
  const [rows] = await pool.query("SELECT * FROM todos ORDER BY created_at DESC");
  res.json(rows.map(r => ({ ...r, done: !!r.done })));
}));

app.post("/api/todos", h(async (req, res) => {
  const { text, due, priority } = req.body;
  if (!text || !text.trim()) return res.status(400).json({ error: "text required" });
  const id = uid();
  await pool.query(
    "INSERT INTO todos (id, text, due, priority, done) VALUES (?,?,?,?,0)",
    [id, text.trim(), due || null, priority || "medium"]
  );
  res.json({ id });
}));

app.patch("/api/todos/:id", h(async (req, res) => {
  const { done } = req.body;
  await pool.query("UPDATE todos SET done=? WHERE id=?", [done ? 1 : 0, req.params.id]);
  res.json({ ok: true });
}));

app.delete("/api/todos/:id", h(async (req, res) => {
  await pool.query("DELETE FROM todos WHERE id=?", [req.params.id]);
  res.json({ ok: true });
}));

/* ---------------- EVENTS ---------------- */
app.get("/api/events", h(async (req, res) => {
  const [rows] = await pool.query("SELECT * FROM events");
  res.json(rows.map(r => ({ ...r, date: iso(r.date) })));
}));

app.post("/api/events", h(async (req, res) => {
  const { date, time, text } = req.body;
  if (!date || !text || !text.trim()) return res.status(400).json({ error: "date and text required" });
  const id = uid();
  await pool.query(
    "INSERT INTO events (id, date, time, text) VALUES (?,?,?,?)",
    [id, date, time || null, text.trim()]
  );
  res.json({ id });
}));

app.delete("/api/events/:id", h(async (req, res) => {
  await pool.query("DELETE FROM events WHERE id=?", [req.params.id]);
  res.json({ ok: true });
}));

/* ---------------- TRANSACTIONS ---------------- */
app.get("/api/txns", h(async (req, res) => {
  const [rows] = await pool.query("SELECT * FROM transactions ORDER BY date DESC, created_at DESC");
  res.json(rows.map(r => ({
    id: r.id, type: r.type, desc: r.descr, category: r.category,
    amount: Number(r.amount), date: iso(r.date),
    recurring: !!r.recurring, freq: r.freq
  })));
}));

app.post("/api/txns", h(async (req, res) => {
  const { type, desc, category, amount, date, recurring, freq } = req.body;
  if (!desc || !desc.trim() || !(amount > 0)) return res.status(400).json({ error: "desc and positive amount required" });
  const id = uid();
  await pool.query(
    "INSERT INTO transactions (id, type, descr, category, amount, date, recurring, freq) VALUES (?,?,?,?,?,?,?,?)",
    [id, type === "income" ? "income" : "expense", desc.trim(),
     (category && category.trim()) || "Uncategorized", amount,
     date || new Date().toISOString().slice(0, 10),
     recurring ? 1 : 0, recurring ? (freq || "monthly") : null]
  );
  res.json({ id });
}));

app.delete("/api/txns/:id", h(async (req, res) => {
  await pool.query("DELETE FROM transactions WHERE id=?", [req.params.id]);
  res.json({ ok: true });
}));

// MySQL DATE comes back as a JS Date; normalize to YYYY-MM-DD
function iso(d) {
  if (!d) return null;
  if (typeof d === "string") return d.slice(0, 10);
  const t = new Date(d.getTime() - d.getTimezoneOffset() * 60000);
  return t.toISOString().slice(0, 10);
}

const PORT = process.env.PORT || 3000;
pool.getConnection()
  .then((c) => {
    c.release();
    app.listen(PORT, () => console.log(`\n✅ Dashboard running: http://localhost:${PORT}\n`));
  })
  .catch((e) => {
    console.error("\n❌ Could not connect to MySQL:", e.message);
    console.error("   Check your .env settings and that MySQL is running.\n");
    process.exit(1);
  });
