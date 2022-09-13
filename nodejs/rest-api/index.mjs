import app from "./app.mjs";

const PORT = 8080 || parseInt(process.env.PORT);

app.listen(PORT, () => {
  console.log(`listening on http://localhost:${PORT}`);
});
