DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` varchar(36) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `publication_date` datetime DEFAULT NULL,
  `library_id` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `libraries`;

CREATE TABLE `libraries` (
  `id` varchar(36) NOT NULL DEFAULT '',
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;