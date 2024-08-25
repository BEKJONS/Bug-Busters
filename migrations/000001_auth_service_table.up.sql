-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the driver_licenses table
CREATE TABLE driver_licenses
(
    id              UUID default gen_random_uuid(),
    first_name      VARCHAR(100),
    last_name       VARCHAR(100),
    father_name     VARCHAR(100),
    birth_date      DATE,
    address         TEXT,
    issue_date      DATE,
    expiration_date DATE,
    category        VARCHAR(50),
    issued_by       VARCHAR(100),
    license_number  VARCHAR PRIMARY KEY
);

-- Insert mock data into driver_licenses
INSERT INTO driver_licenses (id, first_name, last_name, father_name, birth_date, address, issue_date, expiration_date, category, issued_by, license_number)
VALUES
    (gen_random_uuid(), 'John', 'Doe', 'Edward', '1985-06-15', '123 Elm St, Springfield', '2020-06-15', '2030-06-15', 'B', 'Department of Motor Vehicles', 'DL123456789'),
    (gen_random_uuid(), 'Jane', 'Smith', 'Anne', '1990-04-25', '456 Oak St, Springfield', '2021-04-25', '2031-04-25', 'A', 'Department of Motor Vehicles', 'DL987654321');

-- Create the users table
CREATE TABLE users
(
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    driver_license VARCHAR REFERENCES driver_licenses (license_number),
    email          VARCHAR UNIQUE,
    password       VARCHAR,
    role           VARCHAR,
    created_at     TIMESTAMP        DEFAULT now(),
    updated_at     TIMESTAMP        DEFAULT now(),
    deleted_at     BIGINT default 0
);

-- Insert mock data into users
INSERT INTO users (id, driver_license, email, password, role, created_at, updated_at, deleted_at)
VALUES
    (gen_random_uuid(), 'DL123456789', 'john.doe@example.com', 'hashedpassword123', 'user', now(), now(), NULL),
    (gen_random_uuid(), 'DL987654321', 'jane.smith@example.com', 'hashedpassword456', 'ichki_ishlar', now(), now(), NULL),
    (gen_random_uuid(), NULL, 'admin@example.com', 'hashedpassword789', 'yagona_darcha', now(), now(), NULL),
    (gen_random_uuid(), NULL, 'services@example.com', 'hashedpassword012', 'services', now(), now(), NULL);

-- Create the cars table
CREATE TABLE cars
(
    id                   UUID UNIQUE DEFAULT gen_random_uuid(),
    user_id uuid references users(id),
    type                 VARCHAR(100),
    model                VARCHAR(100),
    color                VARCHAR(50),
    year                 INT,
    body_number          VARCHAR(100) UNIQUE,
    engine_number        VARCHAR(100) UNIQUE,
    horsepower           INT,
    image_url            VARCHAR,
    license_plate        VARCHAR(100) UNIQUE,
    tech_passport_number UUID primary key
);

-- Insert mock data into cars
INSERT INTO cars (id, user_id, type, model, color, year, body_number, engine_number, horsepower, image_url, license_plate, tech_passport_number)
VALUES
        (gen_random_uuid(), (SELECT id FROM users WHERE email = 'john.doe@example.com'), 'Sedan', 'Toyota Camry', 'Blue', 2021, 'XYZ123456789', 'ENG123456789', 200, 'http://example.com/image1.jpg', 'ABC123', gen_random_uuid()),
    (gen_random_uuid(), (SELECT id FROM users WHERE email = 'jane.smith@example.com'), 'SUV', 'Honda CR-V', 'Red', 2022, 'ABC987654321', 'ENG987654321', 250, 'http://example.com/image2.jpg', 'XYZ987', gen_random_uuid());

-- Create the services table
CREATE TABLE Services
(
    id                 UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type               VARCHAR(100),
    name               VARCHAR(100),
    certificate_number VARCHAR(100) UNIQUE,
    manager_name       VARCHAR(150),
    address            TEXT,
    phone_number       VARCHAR(13)
);

-- Insert mock data into Services
INSERT INTO Services (id, type, name, certificate_number, manager_name, address, phone_number)
VALUES
    (gen_random_uuid(), 'Repair', 'Auto Repair Shop', 'CERT123456', 'Alice Johnson', '789 Maple St, Springfield', '+1234567890'),
    (gen_random_uuid(), 'Maintenance', 'Quick Oil Change', 'CERT654321', 'Bob Williams', '321 Birch St, Springfield', '+0987654321');

-- Create the fines table
CREATE TABLE Fines
(
    id                   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tech_passport_number UUID REFERENCES cars (tech_passport_number),
    license_plate        VARCHAR(100) REFERENCES cars (license_plate),
    officer_id           UUID REFERENCES users (id),
    fine_owner           UUID REFERENCES users (id),
    fine_date            TIMESTAMP,
    price                INT,
    payment_date         TIMESTAMP
);

-- Insert mock data into Fines
INSERT INTO Fines (id, tech_passport_number, license_plate, officer_id, fine_owner, fine_date, price, payment_date)
VALUES
    (gen_random_uuid(), (SELECT tech_passport_number FROM cars WHERE license_plate = 'ABC123'), 'ABC123', (SELECT id FROM users WHERE email = 'jane.smith@example.com'), (SELECT id FROM users WHERE email = 'john.doe@example.com'), '2024-08-01 15:00:00', 150, '2024-08-10 10:00:00'),
    (gen_random_uuid(), (SELECT tech_passport_number FROM cars WHERE license_plate = 'XYZ987'), 'XYZ987', (SELECT id FROM users WHERE email = 'john.doe@example.com'), (SELECT id FROM users WHERE email = 'jane.smith@example.com'), '2024-08-15 09:00:00', 200, NULL);

-- Create the services_provided table
CREATE TABLE Services_Provided
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    license_plate   VARCHAR(100) REFERENCES cars (license_plate),
    service_type    VARCHAR(100),
    service_date    DATE,
    expiration_date DATE,
    user_id         UUID REFERENCES users (id)
);

-- Insert mock data into Services_Provided
INSERT INTO Services_Provided (id, license_plate, service_type, service_date, expiration_date, user_id)
VALUES
    (gen_random_uuid(), 'ABC123', 'Oil Change', '2024-07-01', '2025-07-01', (SELECT id FROM users WHERE email = 'john.doe@example.com')),
    (gen_random_uuid(), 'XYZ987', 'Tire Rotation', '2024-08-01', '2025-08-01', (SELECT id FROM users WHERE email = 'jane.smith@example.com'));
