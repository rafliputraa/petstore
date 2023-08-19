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

CREATE INDEX idx_order_customer ON orders(customer_id);
CREATE INDEX idx_order_item_order ON order_items(order_id);
CREATE INDEX idx_order_item_pet ON order_items(pet_id);