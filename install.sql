/*
SQLyog Enterprise - MySQL GUI v8.05 
MySQL - 8.0.31 : Database - go_microservice
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- CREATE DATABASE /*!32312 IF NOT EXISTS*/`go_microservice` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `go_microservice`;

/*Table structure for table `logs` */

DROP TABLE IF EXISTS `logs`;

CREATE TABLE `logs` (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `created_datetime` datetime NOT NULL,
  `url` varchar(255) DEFAULT NULL,
  `remote_address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB AUTO_INCREMENT=218284 DEFAULT CHARSET=latin1;

/*Data for the table `logs` */

insert  into `logs`(`log_id`,`created_datetime`,`url`,`remote_address`) values (218203,'2022-12-25 17:45:31','/product?id=7','127.0.0.1'),(218204,'2022-12-25 17:47:29','/users','127.0.0.1'),(218205,'2022-12-25 19:08:29','/login/','127.0.0.1'),(218206,'2022-12-25 19:09:59','/login/','127.0.0.1'),(218207,'2022-12-25 19:10:07','/login/','127.0.0.1'),(218208,'2022-12-25 19:10:37','/login/','127.0.0.1'),(218209,'2022-12-25 19:10:46','/login/','127.0.0.1'),(218210,'2022-12-25 19:12:28','/login/','127.0.0.1'),(218211,'2022-12-25 19:13:50','/login/','127.0.0.1'),(218212,'2022-12-25 19:14:37','/login/','127.0.0.1'),(218213,'2022-12-25 19:17:43','/user/','127.0.0.1'),(218214,'2022-12-25 19:18:35','/login/','127.0.0.1'),(218215,'2022-12-25 19:21:24','/login/','127.0.0.1'),(218216,'2022-12-25 19:26:05','/login/','127.0.0.1'),(218217,'2022-12-25 19:28:00','/login/','127.0.0.1'),(218218,'2022-12-25 19:57:52','/login/','127.0.0.1'),(218219,'2022-12-25 20:00:41','/login/','127.0.0.1'),(218220,'2022-12-25 20:07:53','/login/','127.0.0.1'),(218221,'2022-12-25 20:08:20','/login/','127.0.0.1'),(218222,'2022-12-25 20:08:30','/login/','127.0.0.1'),(218223,'2022-12-25 20:26:48','/logout/','127.0.0.1'),(218224,'2022-12-25 20:27:04','/logout/','127.0.0.1'),(218225,'2022-12-25 20:27:33','/logout/','127.0.0.1'),(218226,'2022-12-25 20:57:12','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218227,'2022-12-25 20:57:46','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218228,'2022-12-25 20:58:34','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218229,'2022-12-25 21:05:20','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218230,'2022-12-25 21:05:36','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218231,'2022-12-25 21:05:37','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218232,'2022-12-25 21:06:00','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218233,'2022-12-25 21:10:04','/order_detail?id=1&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218234,'2022-12-25 21:10:26','/order?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218235,'2022-12-25 21:10:49','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218236,'2022-12-25 21:10:54','/product/?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218237,'2022-12-25 21:25:45','/user?id=9&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218238,'2022-12-25 21:25:52','/user?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218239,'2022-12-25 21:26:57','/user?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218240,'2022-12-25 21:27:24','/user?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218241,'2022-12-25 21:27:29','/order?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218242,'2022-12-25 21:27:39','/order?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218243,'2022-12-25 21:30:05','/order?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218244,'2022-12-25 21:30:09','/order_detail?id=1&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218245,'2022-12-25 21:30:22','/get_cart_items?user_id=10&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218246,'2022-12-26 08:06:04','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218247,'2022-12-26 08:06:12','/user?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218248,'2022-12-26 11:16:39','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218249,'2022-12-26 11:21:37','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218250,'2022-12-26 11:23:06','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218251,'2022-12-26 11:23:20','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218252,'2022-12-26 11:25:40','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218253,'2022-12-26 11:26:24','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218254,'2022-12-26 11:28:46','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218255,'2022-12-26 11:29:24','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218256,'2022-12-26 11:32:15','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218257,'2022-12-26 11:32:18','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218258,'2022-12-26 11:40:33','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218259,'2022-12-26 11:42:59','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218260,'2022-12-26 11:43:48','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218261,'2022-12-26 11:43:53','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218262,'2022-12-26 11:43:56','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218263,'2022-12-26 11:44:01','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218264,'2022-12-26 11:44:06','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218265,'2022-12-26 11:45:45','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218266,'2022-12-26 11:45:48','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218267,'2022-12-26 11:45:51','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218268,'2022-12-26 11:46:02','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218269,'2022-12-26 11:46:07','/product?id=7&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218270,'2022-12-26 11:46:21','/user?id=12&session=533be427-9800-48be-8c72-d9efaffd6d2f','127.0.0.1'),(218271,'2022-12-26 11:46:47','/product?id=7','127.0.0.1'),(218272,'2022-12-26 11:46:50','/product?id=7','127.0.0.1'),(218273,'2022-12-26 11:46:53','/product?id=7','127.0.0.1'),(218274,'2022-12-26 11:47:04','/product?id=11','127.0.0.1'),(218275,'2022-12-26 11:47:18','/products','127.0.0.1'),(218276,'2022-12-26 11:48:18','/order_details','127.0.0.1'),(218277,'2022-12-26 11:48:26','/order_template','127.0.0.1'),(218278,'2022-12-26 11:50:39','/login','127.0.0.1'),(218279,'2022-12-26 11:51:21','/logout','127.0.0.1'),(218280,'2022-12-26 11:52:30','/product?id=7','127.0.0.1'),(218281,'2022-12-26 11:54:37','/product?id=7','127.0.0.1'),(218282,'2022-12-26 11:54:39','/product?id=7','127.0.0.1'),(218283,'2022-12-26 11:54:42','/product?id=7','127.0.0.1');

/*Table structure for table `order_details` */

DROP TABLE IF EXISTS `order_details`;

CREATE TABLE `order_details` (
  `order_detail_id` int NOT NULL AUTO_INCREMENT,
  `order_id` int DEFAULT '0',
  `product_id` int DEFAULT NULL,
  `quantity` int DEFAULT '0',
  `price` double DEFAULT '0',
  `created_datetime` datetime DEFAULT NULL,
  `updated_datetime` datetime DEFAULT NULL,
  PRIMARY KEY (`order_detail_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `order_details` */

insert  into `order_details`(`order_detail_id`,`order_id`,`product_id`,`quantity`,`price`,`created_datetime`,`updated_datetime`) values (1,4,7,2,599.98,'2022-12-24 13:36:22','2022-12-25 10:43:44'),(4,11,15,6,600,'2022-12-25 14:35:52','2022-12-25 14:35:52'),(5,11,16,5,500,'2022-12-25 14:35:52','2022-12-25 14:35:52'),(6,11,17,3,3000,'2022-12-25 14:35:52','2022-12-25 14:35:52'),(7,11,19,3,600,'2022-12-25 14:35:52','2022-12-25 14:35:52'),(8,12,15,6,600,'2022-12-25 14:36:21','2022-12-25 14:36:21'),(9,12,16,5,500,'2022-12-25 14:36:21','2022-12-25 14:36:21'),(10,12,17,3,3000,'2022-12-25 14:36:21','2022-12-25 14:36:21'),(11,12,19,3,600,'2022-12-25 14:36:21','2022-12-25 14:36:21');

/*Table structure for table `order_templates` */

DROP TABLE IF EXISTS `order_templates`;

CREATE TABLE `order_templates` (
  `order_template_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'will be order_id',
  `product_id` int DEFAULT NULL,
  `order_id` int DEFAULT '0',
  `quantity` int DEFAULT '1' COMMENT 'will take from stock if checkout',
  `price` double DEFAULT '0',
  `user_id` int DEFAULT NULL,
  `cart_phase_type` tinyint(1) DEFAULT '0' COMMENT '0 = cart, 1 = checkout, 2 = canceled',
  `created_datetime` datetime DEFAULT NULL,
  `updated_datetime` datetime DEFAULT NULL,
  PRIMARY KEY (`order_template_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `order_templates` */

insert  into `order_templates`(`order_template_id`,`product_id`,`order_id`,`quantity`,`price`,`user_id`,`cart_phase_type`,`created_datetime`,`updated_datetime`) values (1,7,0,1,299.99,5,0,'2022-12-24 19:33:48','2022-12-25 11:04:30'),(3,9,0,5,2500,5,0,'0000-00-00 00:00:00','2022-12-25 11:05:55'),(4,10,0,1,100,5,0,'2022-12-24 19:33:48','2022-12-25 11:06:15'),(5,15,12,6,600,10,1,'2022-12-24 21:53:46','2022-12-25 12:33:25'),(7,16,12,5,500,10,1,'2022-12-25 11:03:34','2022-12-25 12:33:25'),(9,18,0,1,279.99,10,2,'2022-12-25 11:22:02','2022-12-25 12:33:25'),(10,17,12,3,3000,10,1,'2022-12-25 11:30:45','2022-12-25 11:44:44'),(11,18,0,3,839.97,10,1,'2022-12-25 11:31:09','2022-12-25 11:31:09'),(12,19,12,3,600,10,1,'2022-12-25 11:42:45','2022-12-25 12:33:25'),(15,12,0,3,3000,10,0,'2022-12-26 11:48:26','2022-12-26 11:48:26');

/*Table structure for table `orders` */

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `order_datetime` datetime DEFAULT NULL,
  `checkout_datetime` datetime DEFAULT NULL COMMENT 'if blank not affect stocks',
  `user_id` int DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `orders` */

insert  into `orders`(`order_id`,`order_datetime`,`checkout_datetime`,`user_id`) values (4,'2022-12-24 10:50:07','2022-12-24 10:50:07',5),(5,'2022-12-24 10:51:24','2022-12-24 10:51:24',5),(6,'2022-12-24 10:52:19','2022-12-24 10:52:19',5),(12,'2022-12-25 14:36:21','2022-12-25 14:36:21',10);

/*Table structure for table `products` */

DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `product_id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '',
  `price` double DEFAULT '0',
  `created_datetime` datetime DEFAULT NULL,
  `updated_datetime` datetime DEFAULT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `products` */

insert  into `products`(`product_id`,`product_name`,`price`,`created_datetime`,`updated_datetime`) values (7,'Samsung Galaxy Tab 10.2',299.99,'2022-12-23 21:47:14','2022-12-23 22:01:02'),(9,'MacBook',500,'2022-12-23 21:55:17','2022-12-23 21:55:17'),(10,'iPhone Cinema 30',100,'2022-12-24 16:59:41','2022-12-24 16:59:41'),(12,'Sony VAIO',1000,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(13,'HP LP3065',100,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(14,'Canon EOS 5D',100,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(15,'HTC Touch HD',100,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(16,'iPod Classic',100,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(17,'MacBook Air',1000,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(18,'Palm Treo Pro',279.99,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(19,'Samsung SyncMaster 941BW',200,'2022-12-24 20:05:07','2022-12-24 20:05:07'),(20,'iMac',100,'2022-12-24 20:05:07','2022-12-24 20:05:07');

/*Table structure for table `sessions` */

DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions` (
  `session_id` int NOT NULL AUTO_INCREMENT,
  `session_code` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_id` int DEFAULT NULL,
  `created_datetime` datetime DEFAULT NULL,
  `is_expired` tinyint DEFAULT '0',
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `sessions` */

insert  into `sessions`(`session_id`,`session_code`,`user_id`,`created_datetime`,`is_expired`) values (1,'c3f396e9-516c-4ab4-93db-26cfa03d01d6',12,'2022-12-25 19:00:00',1),(2,'f2e49dfa-da7d-4a2d-8159-56c0c9dc9d73',12,'2022-12-25 20:00:41',1),(3,'533be427-9800-48be-8c72-d9efaffd6d2f',12,'2022-12-26 11:43:24',0);

/*Table structure for table `stocks` */

DROP TABLE IF EXISTS `stocks`;

CREATE TABLE `stocks` (
  `stock_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  `created_datetime` datetime DEFAULT NULL,
  `updated_datetime` datetime DEFAULT NULL,
  PRIMARY KEY (`stock_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `stocks` */

insert  into `stocks`(`stock_id`,`product_id`,`quantity`,`created_datetime`,`updated_datetime`) values (1,7,150,'2022-12-24 13:56:35','2022-12-24 14:03:35'),(4,9,50,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(5,10,60,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(6,12,70,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(7,13,80,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(8,14,90,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(9,15,94,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(10,16,15,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(11,17,12,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(12,18,10,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(13,19,37,'2022-12-24 13:56:35','2022-12-24 13:56:35'),(14,20,60,'2022-12-24 13:56:35','2022-12-24 13:56:35');

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_datetime` datetime DEFAULT NULL,
  `updated_datetime` datetime DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `users` */

insert  into `users`(`user_id`,`username`,`email`,`created_datetime`,`updated_datetime`,`password`) values (5,'abc','abc@def.com','2022-12-24 12:15:05','2022-12-24 12:15:05','$2a$04$kFL0xWORi.CJT.tzr8bhrOXmiNf.pfGSCFqmjg9rk2mOrYiljCoIW'),(10,'def','def@def.com','2022-12-24 13:56:35','2022-12-24 13:56:35','$2a$04$kFL0xWORi.CJT.tzr8bhrOXmiNf.pfGSCFqmjg9rk2mOrYiljCoIW'),(11,'absolute','absolute@abc.com','2022-12-25 16:16:00','2022-12-25 16:16:00','$2a$04$QauOBNHTqeFT7uFSWJ5dPOQ28MLAgAnAGC7oINd/pWO3VjhEJ6TXW'),(12,'simple','simple@abc.com','2022-12-25 19:17:43','2022-12-25 19:17:43','$2a$04$2ban6PdNxuahmKAio7NdD.sVwwo8VNA.WrruGmO/jE0pVjwlLF0Da');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
