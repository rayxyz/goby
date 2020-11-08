
-- create schema

CREATE SCHEMA `goby` DEFAULT CHARACTER SET utf8 ;

-- create the test table

DROP TABLE IF EXISTS `advice`;
CREATE TABLE `advice` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` varchar(500) NOT NULL,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

LOCK TABLES `advice` WRITE;
INSERT INTO `advice` VALUES ('Hello, world!','2019-12-18 23:02:08'),('hahahahahaha','2020-01-03 23:07:49'),('uuuuuuuuu','2020-01-03 23:49:27'),('This is the advice you must take!','2020-08-04 18:42:10'),('Great idea for doing that thing.','2020-08-09 11:27:05');
UNLOCK TABLES;