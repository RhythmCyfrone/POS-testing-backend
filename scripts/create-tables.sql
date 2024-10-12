DROP TABLE IF EXISTS tables;
CREATE TABLE tables (
  TableId    VARCHAR(128) NOT NULL,
  Floor      VARCHAR(128) NOT NULL,
  Siting     INT NOT NULL,
  Curr_status  VARCHAR(128) NOT NULL,
  BranchId  VARCHAR(128) NOT NULL,
  PRIMARY KEY (TableID)
);

INSERT INTO tables (TableId, Floor, Siting, Curr_status, BranchId) VALUES
('A1', '1', 2, 'Free', '1'),
('A2', '1', 4, 'Reserved', '1'),
('A3', '1', 2, 'Assigned', '1'),
('A4', '1', 4, 'Ordered', '1'),
('A5', '1', 2, 'Billed', '1'),
('A6', '1', 4, 'Paid', '1'),
('A7', '1', 2, 'Free', '1'),
('A8', '1', 4, 'Reserved', '1'),
('A9', '1', 2, 'Assigned', '1'),
('A10', '1', 4, 'Ordered', '1'),
('A11', '1', 2, 'Billed', '1'),
('A12', '1', 4, 'Paid', '1'),
('A13', '1', 2, 'Free', '1'),
('A14', '1', 4, 'Reserved', '1'),
('A15', '1', 2, 'Assigned', '1'),
('A16', '1', 4, 'Ordered', '1'),
('A17', '1', 2, 'Billed', '1'),
('A18', '1', 4, 'Paid', '1'),
('A19', '1', 2, 'Free', '1'),
('A20', '1', 4, 'Reserved', '1');

DROP TABLE IF EXISTS orders;
-- Create table for 'Order'
CREATE TABLE orders (
    OrderId VARCHAR(255) PRIMARY KEY,
    TableId VARCHAR(255) NOT NULL,
    CustomerId VARCHAR(255),
    CustomerName VARCHAR(255),
    OrderTakerId VARCHAR(255),
    OrderTakerName VARCHAR(255),
    OrderStatusId int,
    BillId INT,
    FOREIGN KEY (TableId) REFERENCES tables(TableId)
);

-- Insert random data into Orders
INSERT INTO orders (OrderId, TableId, CustomerId, CustomerName, OrderTakerId, OrderTakerName, OrderStatusId, BillId)
VALUES 
('ORD001', 'A1', 'CUST001', 'John Doe', 'TAKER001', 'Alice', 1, 101),
('ORD002', 'A2', 'CUST002', 'Jane Smith', 'TAKER002', 'Bob', 2, 102),
('ORD003', 'A3', 'CUST003', 'Michael Johnson', 'TAKER003', 'Charlie', 1, 103),
('ORD004', 'A4', 'CUST004', 'Emily Davis', 'TAKER004', 'David', 3, 104),
('ORD005', 'A5', 'CUST005', 'Chris Martin', 'TAKER001', 'Alice', 2, 105),
('ORD006', 'A6', 'CUST006', 'Laura Brown', 'TAKER002', 'Bob', 1, 106),
('ORD007', 'A7', 'CUST007', 'Ethan Wilson', 'TAKER003', 'Charlie', 3, 107),
('ORD008', 'A8', 'CUST008', 'Olivia Taylor', 'TAKER004', 'David', 2, 108),
('ORD009', 'A9', 'CUST009', 'Sophia Thomas', 'TAKER001', 'Alice', 1, 109),
('ORD010', 'A10', 'CUST010', 'Liam Clark', 'TAKER002', 'Bob', 3, 110);

DROP TABLE IF EXISTS tableStatus;

CREATE TABLE tableStatus (
    StatusId INT PRIMARY KEY,
    StatusName VARCHAR(255),
    StatusColor VARCHAR(255)
);

INSERT INTO tableStatus (StatusId, StatusName, StatusColor) VALUES
(1, 'Free', '#41C5FF'),
(2, 'Reserved', '#FFB800'),
(3, 'Assigned', '#3B82F6'),
(4, 'Ordered', '#E64980'),
(5, 'Served', '#00D930'),
(6, 'Billed', '#A1007E'),
(7, 'Paid', '#40E0D0');

DROP TABLE IF EXISTS orderStatus;

CREATE TABLE orderStatus (
    StatusId INT PRIMARY KEY,
    StatusName VARCHAR(255),
    StatusColor VARCHAR(255)
);

INSERT INTO orderStatus (StatusId, StatusName, StatusColor) VALUES
(1, 'Ordered', '#E64980'),
(2, 'Served', '#00D930'),
(3, 'Billed', '#A1007E'),
(4, 'Paid', '#40E0D0');

DROP TABLE IF EXISTS takeaways;

CREATE TABLE takeaways (
    OrderId INT PRIMARY KEY AUTO_INCREMENT,
    CustomerId VARCHAR(255),
    CustomerName VARCHAR(255),
    CustomerPhone VARCHAR(255),
    OrderStatusId INT,
    CurrentStatus VARCHAR(255),
    BillId INT,
    BranchId VARCHAR(255)
);

