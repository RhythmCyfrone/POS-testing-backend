DROP TABLE IF EXISTS tables;
CREATE TABLE tables (
  TableId    INT NOT NULL,
  TableName  VARCHAR(128),
  TableMapperId INT,
  TableMaxPax INT,
  DateTimeSinceLastStatusChange VARCHAR(128),
  TableTrackingStatusId INT,
  TableTrackingStatusName VARCHAR(128),
  PRIMARY KEY (TableId)
);

INSERT INTO tables (TableId, TableName, TableMapperId, TableMaxPax, DateTimeSinceLastStatusChange, TableTrackingStatusId, TableTrackingStatusName) VALUES
(7, 'A1', 8, 2, '2024-10-12T08:10:00', 1, 'Free'),
(8, 'A2', 9, 4, '2024-10-12T08:15:00', 2, 'Reserved'),
(9, 'A3', 10, 2, '2024-10-12T08:20:00', 3, 'Assigned'),
(10, 'A4', 11, 6, '2024-10-12T08:25:00', 4, 'Ordered'),
(11, 'A5', 12, 4, '2024-10-12T08:30:00', 5, 'Served'),
(12, 'A6', 13, 2, '2024-10-12T08:35:00', 6, 'Billed'),
(13, 'A7', 14, 8, '2024-10-12T08:40:00', 7, 'Paid'),
(14, 'A8', 15, 2, '2024-10-12T08:45:00', 1, 'Free'),
(15, 'A9', 16, 4, '2024-10-12T08:50:00', 2, 'Reserved'),
(16, 'A10', 17, 6, '2024-10-12T08:55:00', 3, 'Assigned'),
(17, 'A11', 18, 4, '2024-10-12T09:00:00', 4, 'Ordered'),
(18, 'A12', 19, 2, '2024-10-12T09:05:00', 5, 'Served'),
(19, 'A13', 20, 4, '2024-10-12T09:10:00', 6, 'Billed'),
(20, 'A14', 21, 6, '2024-10-12T09:15:00', 7, 'Paid'),
(21, 'A15', 22, 8, '2024-10-12T09:20:00', 1, 'Free'),
(22, 'A16', 23, 4, '2024-10-12T09:25:00', 2, 'Reserved'),
(23, 'A17', 24, 2, '2024-10-12T09:30:00', 3, 'Assigned'),
(24, 'A18', 25, 6, '2024-10-12T09:35:00', 4, 'Ordered'),
(25, 'A19', 26, 4, '2024-10-12T09:40:00', 5, 'Served'),
(26, 'A20', 27, 2, '2024-10-12T09:45:00', 6, 'Billed');


DROP TABLE IF EXISTS orders;
-- Create table for 'Order'
CREATE TABLE orders (
    OrderId VARCHAR(255) PRIMARY KEY,
    TableId INT NOT NULL,
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
('ORD001', 9, 'CUST001', 'John Doe', 'TAKER001', 'Alice', 1, 101),
('ORD002', 10, 'CUST002', 'Jane Smith', 'TAKER002', 'Bob', 2, 102),
('ORD003', 11, 'CUST003', 'Michael Johnson', 'TAKER003', 'Charlie', 1, 103),
('ORD004', 12, 'CUST004', 'Emily Davis', 'TAKER004', 'David', 3, 104),
('ORD005', 13, 'CUST005', 'Chris Martin', 'TAKER001', 'Alice', 2, 105),
('ORD006', 15, 'CUST006', 'Laura Brown', 'TAKER002', 'Bob', 1, 106),
('ORD007', 16, 'CUST007', 'Ethan Wilson', 'TAKER003', 'Charlie', 3, 107),
('ORD008', 17, 'CUST008', 'Olivia Taylor', 'TAKER004', 'David', 2, 108),
('ORD009', 18, 'CUST009', 'Sophia Thomas', 'TAKER001', 'Alice', 1, 109),
('ORD010', 19, 'CUST010', 'Liam Clark', 'TAKER002', 'Bob', 3, 110);

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

