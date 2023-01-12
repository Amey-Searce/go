-- name: create-product-table
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `specs` varchar(255) NOT NULL,
  `sku` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `price` float NOT NULL,
  `productid` varchar(255) NOT NULL,
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- name: create-category-table
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `productid` varchar(255) NOT NULL,
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- name: create-inventory-table
CREATE TABLE `inventory` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product` varchar(255) NOT NULL,
  `quantity` float NOT NULL,
  `productid` varchar(255) NOT NULL,
  primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;