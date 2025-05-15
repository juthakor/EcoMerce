exports.getAllProducts = (req, res) => {
  res.json([{ name: "Sample Product" }]); // stub
};

exports.getProductById = (req, res) => {
  res.json({ id: req.params.id, name: "Product Detail" }); // stub
};
