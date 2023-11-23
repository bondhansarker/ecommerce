CREATE TABLE IF NOT EXISTS brands(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    status_id BOOLEAN DEFAULT true NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    parent_id INTEGER,
    sequence INTEGER,
    status_id BOOLEAN DEFAULT true NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES categories(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS suppliers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    status_id BOOLEAN DEFAULT true NOT NULL,
    is_verified_supplier BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    specifications TEXT,
    brand_id INT,
    category_id INT,
    supplier_id INT,
    unit_price DOUBLE PRECISION,
    discount_price DOUBLE PRECISION,
    tags VARCHAR(255),
    status_id BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (brand_id) REFERENCES brands(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    UNIQUE (name, supplier_id)
);

CREATE TABLE IF NOT EXISTS product_stocks (
    id SERIAL PRIMARY KEY,
    product_id INT,
    stock_quantity INT DEFAULT 0 CHECK (stock_quantity >= 0) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

-- Insert data into brands table
INSERT INTO brands (name, status_id, created_at) VALUES
                                                     ('Brand A', true, NOW()),
                                                     ('Brand B', true, NOW()),
                                                     ('Brand C', true, NOW());

-- Insert data into categories table
INSERT INTO categories (name, parent_id, sequence, status_id, created_at) VALUES
                                                                              ('Category A', NULL, 1, true, NOW()),
                                                                              ('Category B', NULL, 2, true, NOW()),
                                                                              ('Subcategory A1', 1, 1, true, NOW()),
                                                                              ('Subcategory A2', 1, 2, true, NOW());

-- Insert data into suppliers table
INSERT INTO suppliers (name, email, phone, status_id, is_verified_supplier, created_at) VALUES
                                                                                            ('Supplier 1', 'supplier1@example.com', '1234567890', true, true, NOW()),
                                                                                            ('Supplier 2', 'supplier2@example.com', '9876543210', true, true, NOW());

-- Insert data into products table
INSERT INTO products (name, description, specifications, brand_id, category_id, supplier_id, unit_price, discount_price, tags, status_id, created_at) VALUES
                                                                                                                                                          ('Product 1', 'Description 1', 'Specs 1', 1, 1, 1, 10.99, 9.99, 'Tag1, Tag2', true, NOW()),
                                                                                                                                                          ('Product 2', 'Description 2', 'Specs 2', 2, 1, 2, 19.99, 15.99, 'Tag3, Tag4', true, NOW()),
                                                                                                                                                          ('Product 3', 'Description 3', 'Specs 3', 3, 2, 1, 29.99, 24.99, 'Tag5, Tag6', true, NOW()),
                                                                                                                                                          ('Product 4', 'Description 4', 'Specs 4', 1, 1, 2, 14.99, 11.99, 'Tag7, Tag8', true, NOW()),
                                                                                                                                                          ('Product 5', 'Description 5', 'Specs 5', 2, 2, 1, 24.99, 19.99, 'Tag9, Tag10', true, NOW()),
                                                                                                                                                          ('Product 6', 'Description 6', 'Specs 6', 3, 1, 2, 34.99, 29.99, 'Tag11, Tag12', true, NOW()),
                                                                                                                                                          ('Product 7', 'Description 7', 'Specs 7', 1, 2, 1, 44.99, 39.99, 'Tag13, Tag14', true, NOW()),
                                                                                                                                                          ('Product 8', 'Description 8', 'Specs 8', 2, 1, 2, 54.99, 49.99, 'Tag15, Tag16', true, NOW()),
                                                                                                                                                          ('Product 9', 'Description 9', 'Specs 9', 3, 1, 1, 64.99, 59.99, 'Tag17, Tag18', true, NOW()),
                                                                                                                                                          ('Product 10', 'Description 10', 'Specs 10', 1, 2, 2, 74.99, 69.99, 'Tag19, Tag20', true, NOW()),
                                                                                                                                                          ('Product 11', 'Description 11', 'Specs 11', 2, 1, 1, 84.99, 79.99, 'Tag21, Tag22', true, NOW()),
                                                                                                                                                          ('Product 12', 'Description 12', 'Specs 12', 3, 2, 2, 94.99, 89.99, 'Tag23, Tag24', true, NOW()),
                                                                                                                                                          ('Product 13', 'Description 13', 'Specs 13', 1, 1, 1, 104.99, 99.99, 'Tag25, Tag26', true, NOW()),
                                                                                                                                                          ('Product 14', 'Description 14', 'Specs 14', 2, 2, 2, 114.99, 109.99, 'Tag27, Tag28', true, NOW()),
                                                                                                                                                          ('Product 15', 'Description 15', 'Specs 15', 3, 1, 1, 124.99, 119.99, 'Tag29, Tag30', true, NOW()),
                                                                                                                                                          ('Product 16', 'Description 16', 'Specs 16', 1, 2, 2, 134.99, 129.99, 'Tag31, Tag32', true, NOW()),
                                                                                                                                                          ('Product 17', 'Description 17', 'Specs 17', 2, 1, 1, 144.99, 139.99, 'Tag33, Tag34', true, NOW()),
                                                                                                                                                          ('Product 18', 'Description 18', 'Specs 18', 3, 2, 2, 154.99, 149.99, 'Tag35, Tag36', true, NOW()),
                                                                                                                                                          ('Product 19', 'Description 19', 'Specs 19', 1, 1, 1, 164.99, 159.99, 'Tag37, Tag38', true, NOW()),
                                                                                                                                                          ('Product 20', 'Description 20', 'Specs 20', 2, 2, 2, 174.99, 169.99, 'Tag39, Tag40', true, NOW());

-- Insert data into product_stocks table
INSERT INTO product_stocks (product_id, stock_quantity, updated_at) VALUES
                                                                        (1, 100, NOW()),
                                                                        (2, 50, NOW()),
                                                                        (3, 75, NOW()),
                                                                        (4, 60, NOW()),
                                                                        (5, 45, NOW()),
                                                                        (6, 70, NOW()),
                                                                        (7, 30, NOW()),
                                                                        (8, 55, NOW()),
                                                                        (9, 80, NOW()),
                                                                        (10, 25, NOW()),
                                                                        (11, 40, NOW()),
                                                                        (12, 65, NOW()),
                                                                        (13, 50, NOW()),
                                                                        (14, 75, NOW()),
                                                                        (15, 35, NOW()),
                                                                        (16, 60, NOW()),
                                                                        (17, 45, NOW()),
                                                                        (18, 70, NOW()),
                                                                        (19, 30, NOW()),
                                                                        (20, 55, NOW());
