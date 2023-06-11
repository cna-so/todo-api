CREATE TABLE users (
  user_uid UUID PRIMARY KEY ,
  first_name VARCHAR(24),
  last_name VARCHAR(64),
  email VARCHAR(150) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  role VARCHAR(1) DEFAULT 'U',
  created_at TIMESTAMP DEFAULT Now(),
  updated_at TIMESTAMP DEFAULT Now()
);
