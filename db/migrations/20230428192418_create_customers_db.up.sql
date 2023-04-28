CREATE TABLE IF NOT EXISTS "customers" (
  id serial PRIMARY KEY,
  first_name varchar NOT NULL,
  last_name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
);