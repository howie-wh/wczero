/*
 Data Transfer
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_tab
-- ----------------------------
DROP TABLE IF EXISTS `user_tab`;
CREATE TABLE `user_tab` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_name` varchar(128) NOT NULL COMMENT 'username',
    `nick_name` varchar(128) NOT NULL DEFAULT '' COMMENT 'nickname',
    `email` varchar(128) NOT NULL DEFAULT '' COMMENT 'email',
    `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT 'avatar',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uniq_username` (`user_name`) USING BTREE COMMENT 'uniq_username',
    KEY `idx_ctime` (`create_time`) USING BTREE COMMENT 'idx_ctime'
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='user table';


-- ----------------------------
-- Table structure for user_tab
-- ----------------------------
DROP TABLE IF EXISTS `user_admin_tab`;
CREATE TABLE `user_admin_tab` (
    `user_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
    `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
    `mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='user admin table';

BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;