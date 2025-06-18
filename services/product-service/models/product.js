const mongoose = require("mongoose");

const ProductSchema = new mongoose.Schema({
  name: String,
  description: String,
  price: Number,
  stock: Number,
  eco_tags: [String],
  supplier: {
    name: String,
    country: String,
    certifications: [String],
  },
  created_at: {
    type: Date,
    default: Date.now,
  },
  images: [String],
  rating: Number,
});

module.exports = mongoose.model("Product", ProductSchema);
