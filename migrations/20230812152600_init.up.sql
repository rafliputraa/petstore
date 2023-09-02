CREATE TABLE IF NOT EXISTS customers (
  customer_id serial PRIMARY KEY,
  first_name varchar(50),
  last_name varchar(50),
  email varchar(100),
  phone varchar(15)
);

CREATE TABLE IF NOT EXISTS pets (
  pet_id serial PRIMARY KEY,
  name varchar(100),
  species varchar(50),
  age integer,
  available boolean,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE IF NOT EXISTS orders (
  order_id serial PRIMARY KEY,
  customer_id integer,
  order_date timestamp,
  total_amount integer,
  FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

CREATE TABLE IF NOT EXISTS order_items (
  order_item_id serial PRIMARY KEY,
  order_id integer,
  pet_id integer,
  quantity integer,
  FOREIGN KEY (order_id) REFERENCES orders(order_id),
  FOREIGN KEY (pet_id) REFERENCES pets(pet_id)
);

CREATE INDEX idx_order_item_order ON order_items(order_id);
CREATE INDEX idx_order_item_pet ON order_items(pet_id);

-- Insert 50 rows into the customers table
INSERT INTO customers (first_name, last_name, email, phone)
SELECT
  'First_Name_' || generate_series(1, 50),
  'Last_Name_' || generate_series(1, 50),
  'email_' || generate_series(1, 50) || '@example.com',
  '123-456-7890'
FROM generate_series(1, 50);

-- Insert 50 rows into the pets table
INSERT INTO pets (name, species, age, available, created_at, updated_at)
SELECT
  'Pet_Name_' || generate_series(1, 50),
  'Species_' || generate_series(1, 50),
  (random() * 10 + 1)::integer, -- Random age between 1 and 10
  random() < 0.5, -- Randomly set available to true or false
  NOW() - interval '1 day' * (random() * 365), -- Random created_at timestamp within the last year
  NOW() - interval '1 day' * (random() * 365) -- Random updated_at timestamp within the last year
FROM generate_series(1, 50);