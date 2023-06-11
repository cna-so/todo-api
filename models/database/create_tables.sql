CREATE TABLE users (
  user_uid UUID PRIMARY KEY ,
  first_name VARCHAR(24),
  last_name VARCHAR(64),
  email VARCHAR(150) NOT NULL UNIQUE,
  password VARCHAR(32) NOT NULL,
  role VARCHAR(1),
  created_at DATE,
  updated_at DATE
);