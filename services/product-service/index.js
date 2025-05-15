const express = require("express");
const app = express();
const productRoutes = require("./routes/products");

app.use(express.json());
app.get("/health", (req, res) => res.send("product-service OK"));
app.use("/products", productRoutes);

const PORT = process.env.PORT || 8082;
app.listen(PORT, () => console.log(`Product Service running on port ${PORT}`));
