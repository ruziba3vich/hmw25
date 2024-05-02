CREATE TABLE IF NOT EXISTS Products (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    price INT NOT NULL,
    unit_id INT NOT NULL,
    Description TEXT,
    FOREIGN KEY (category_id) REFERENCES Categories(id),
    FOREIGN KEY (unit_id) REFERENCES Units(Id)
);
