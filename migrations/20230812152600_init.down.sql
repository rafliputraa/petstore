DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS pets;

DROP INDEX IF EXISTS idx_order_customer;
DROP INDEX IF EXISTS idx_order_item_order;
DROP INDEX IF EXISTS idx_order_item_pet;