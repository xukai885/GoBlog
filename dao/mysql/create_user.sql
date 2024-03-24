SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Appuser
-- ----------------------------
DROP TABLE IF EXISTS `Appuser`;
CREATE TABLE `Appuser` (
                           `id` int(11) NOT NULL AUTO_INCREMENT,
                           `user_id` bigint(20) DEFAULT NULL,
                           `username` varchar(255) NOT NULL,
                           `password` varchar(255) NOT NULL,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for boke
-- ----------------------------
DROP TABLE IF EXISTS `boke`;
CREATE TABLE `boke` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `titlename` varchar(100) DEFAULT NULL,
                        `filename` varchar(100) DEFAULT NULL,
                        `filepath` varchar(255) DEFAULT NULL,
                        `introduction` varchar(255) DEFAULT NULL,
                        `type` int(11) DEFAULT NULL,
                        `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `md5filename` varchar(255) DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `titlename` (`titlename`),
                        KEY `fk_boke_type` (`type`),
                        CONSTRAINT `fk_boke_type` FOREIGN KEY (`type`) REFERENCES `boketype` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for boketype
-- ----------------------------
DROP TABLE IF EXISTS `boketype`;
CREATE TABLE `boketype` (
                            `id` int(11) NOT NULL AUTO_INCREMENT,
                            `typename` varchar(50) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `typename` (`typename`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for images
-- ----------------------------
DROP TABLE IF EXISTS `images`;
CREATE TABLE `images` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) DEFAULT NULL,
                          `path` varchar(255) DEFAULT NULL,
                          `createtime` datetime DEFAULT NULL,
                          `size` double DEFAULT NULL,
                          `imagesUrl` varchar(255) DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for share
-- ----------------------------
DROP TABLE IF EXISTS `share`;
CREATE TABLE `share` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `title` varchar(255) NOT NULL,
                         `url` varchar(255) NOT NULL,
                         `type` varchar(255) DEFAULT NULL,
                         `date` datetime NOT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for yiyan
-- ----------------------------
DROP TABLE IF EXISTS `yiyan`;
CREATE TABLE `yiyan` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `hitokoto` varchar(255) DEFAULT NULL,
                         `from_source` varchar(255) DEFAULT NULL,
                         `from_who` varchar(255) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `hitokoto` (`hitokoto`)
) ENGINE=InnoDB AUTO_INCREMENT=591 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
