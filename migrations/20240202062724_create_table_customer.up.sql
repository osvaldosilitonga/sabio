CREATE TABLE Customer (
    id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO Customer (id, name, email, created_at, updated_at)
VALUES
  (101, 'John Doe', 'john@example.com', '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (102, 'Jane Doe', 'jane@example.com', '2024-01-22 12:00:00', '2024-01-22 12:00:00'),
  (103, 'Bob Smith', 'bob@example.com', '2024-01-22 12:00:00', '2024-01-22 12:00:00');
