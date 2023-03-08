import express from "express";
import { dirname } from "path";
import { fileURLToPath } from 'url';

const app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

app.use(express.static(__dirname));

app.get("/", (req, res) => {
  return res.sendFile(__dirname + "/web.html");
});

app.use((err, _req, res, next) => {
  if (res.headersSent) {
    return next(err);
  }
  console.error(err);
  res.status(500);
  res.json({ error: err.message });
});

export default app;
