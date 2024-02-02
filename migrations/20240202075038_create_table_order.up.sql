CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    customer_id INT,
    product_id INT,
    quantity INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES Customer(id),
    FOREIGN KEY (product_id) REFERENCES Product(id)
);

INSERT INTO Orders (customer_id, product_id, quantity, total, created_at, updated_at)
VALUES
  (101, 1, 2, 1600, '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (102, 2, 1, 500, '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (103, 3, 3, 600, '2024-01-22 12:00:00', '2024-01-22 12:00:00');
