## Database
1. First create database with name `belajar_golang`

```go
create database belajar_golang;
```

2. Create table `customer`
```go
CREATE TABLE `customer` (
  `id` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `balance` int DEFAULT '0',
  `rating` double DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `birth_date` date DEFAULT NULL,
  `married` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
```

3. Create table `comments`
```go
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(100) NOT NULL,
  `comment` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
```

4. Create tabel `user`
```go
CREATE TABLE `user` (
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB;
```

## Run Test
```go
go test -v ./...
```