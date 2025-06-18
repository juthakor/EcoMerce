CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  full_name TEXT,
  role TEXT DEFAULT 'customer',
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS eco_badges (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id),
  badge_type TEXT,
  awarded_at TIMESTAMP DEFAULT now()
);
