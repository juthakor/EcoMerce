# 📊 Entity Relationship Diagram (ERD)

This document outlines the database design for the EcoMerce.


## 🏛 PostgreSQL (Relational Data)

```mermaid
erDiagram
    USERS {
        UUID id PK
        STRING email
        STRING password_hash
        STRING full_name
        STRING role
        TIMESTAMP created_at
    }
    ORDERS {
        UUID id PK
        UUID user_id FK
        STRING status
        DECIMAL total_price
        TIMESTAMP created_at
    }
    ORDER_ITEMS {
        UUID id PK
        UUID order_id FK
        STRING product_id
        INTEGER quantity
        DECIMAL unit_price
    }
    PAYMENTS {
        UUID id PK
        UUID order_id FK
        STRING status
        STRING payment_method
        STRING transaction_id
        TIMESTAMP paid_at
    }
    ECO_BADGES {
        UUID id PK
        UUID user_id FK
        STRING badge_type
        TIMESTAMP awarded_at
    }

    USERS ||--o{ ORDERS : places
    ORDERS ||--o{ ORDER_ITEMS : contains
    ORDERS ||--|{ PAYMENTS : has
    USERS ||--o{ ECO_BADGES : earns
```

### Notes

- `USERS.role` → enum: customer, admin  
- `ORDERS.status` → enum: pending, paid, shipped, canceled  
- `PAYMENTS.status` → enum: pending, success, failed  
- `ORDER_ITEMS.product_id` → refers to MongoDB product ID  
- Foreign keys:
    - `ORDERS.user_id` → USERS.id
    - `ORDER_ITEMS.order_id` → ORDERS.id
    - `PAYMENTS.order_id` → ORDERS.id
    - `ECO_BADGES.user_id` → USERS.id


## 🍃 MongoDB (Product Catalog)
Example document structure:

```json
{
  "_id": "ObjectId",
  "name": "Eco Tote Bag",
  "description": "Made from 100% recycled cotton.",
  "price": 19.99,
  "stock": 120,
  "eco_tags": ["recycled", "zero-waste", "local"],
  "supplier": {
    "name": "GreenMakers Co.",
    "country": "Thailand",
    "certifications": ["Fair Trade", "Organic Cotton"]
  },
  "created_at": "ISODate",
  "images": ["https://..."],
  "rating": 4.8
}
```