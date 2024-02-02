-- DDL
CREATE TABLE Product (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- DML
INSERT INTO Product (id, name, price, stock, created_at, updated_at)
VALUES
  (1, 'Laptop', 800, 10, '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (2, 'Smartphone', 500, 20, '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (3, 'Printer', 200, 5, '2024-01-22 12:00:00', '2024-01-22 12:00:00');
