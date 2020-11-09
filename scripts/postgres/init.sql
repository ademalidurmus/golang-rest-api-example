CREATE TABLE IF NOT EXISTS posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR (255),
  content TEXT,
  status VARCHAR (10),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);