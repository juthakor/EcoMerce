const { MongoClient } = require("mongodb");

async function setup() {
  const client = await MongoClient.connect("mongodb://localhost:27017");
  const db = client.db("eco_db");

  const collections = await db.listCollections({ name: "products" }).toArray();
  if (collections.length === 0) {
    await db.createCollection("products", {
      validator: {
        $jsonSchema: {
          bsonType: "object",
          required: ["name", "price"],
          properties: {
            name: { bsonType: "string" },
            price: { bsonType: "double" },
            stock: { bsonType: "int" },
            eco_tags: {
              bsonType: "array",
              items: { bsonType: "string" }
            },
            supplier: {
              bsonType: "object",
              properties: {
                name: { bsonType: "string" },
                country: { bsonType: "string" }
              }
            },
            created_at: { bsonType: "date" }
          }
        }
      }
    });

    console.log("Created products collection with schema validator");
  } else {
    console.log("Products collection already exists");
  }

  await client.close();
}

setup();