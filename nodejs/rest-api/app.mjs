import express from "express";
import cache from "./cache.mjs";

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/api/keys/:key", (req, res) => {
  const key = req.params.key;
  if (!key || typeof key !== "string") {
    return res.status(400).json({ error: "missing or invalid key" });
  }
  if (!cache.has(key)) {
    return res.status(404).json({ error: "key does not exist" });
  }
  const value = cache.get(key);
  return res.json({ key: id, value });
});

app.get("/api/keys", (_, res) => {
  return res.json(cache.keys());
});

app.post("/api/keys", (req, res) => {
  const { key, value, ttl } = req.body;
  if (!key || !value) {
    return res.status(400).json({ error: "invalid request body" });
  }
  cache.set(key, value, ttl ?? 0);
  return res.status(201).json({ key });
});

app.delete("/api/keys/:key", (req, res) => {
  const key = req.params.key;
  if (!key || typeof key !== "string") {
    return res.status(400).json({ error: "missing or invalid key" });
  }
  if (!cache.has(key)) {
    return res.status(404).json({ error: "key does not exist" });
  }
  cache.del(key);
  return res.json({ key });
});

app.get("/healthz", (_, res) => {
  return res.sendStatus(200);
});

app.use((err, _req, res, next) => {
  if (res.headersSent) {
    return next(err);
  }
  console.error(err);
  res.status(500);
  res.json({ error: err.message });
});

app.use("*", (_, res) => {
  return res
    .status(404)
    .json({ error: "the requested resource does not exist on this server" });
});

export default app;
