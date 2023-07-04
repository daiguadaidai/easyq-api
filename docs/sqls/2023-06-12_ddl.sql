DROP TABLE IF EXISTS mysql_db_priv;
CREATE TABLE `mysql_db_priv` (
                                 `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                 `order_uuid` VARCHAR(20) NOT NULL DEFAULT '订单编号',
                                 `user_id` bigint NOT NULL COMMENT '用户id',
                                 `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                                 `name_zh` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
                                 `meta_cluster_id` bigint NOT NULL COMMENT '集群id',
                                 `cluster_name` varchar(100) NOT NULL DEFAULT '' COMMENT '集群名称',
                                 `db_name` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库名',
                                 `vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VIP 和端口',
                                 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `idx_username_cluster_db` (`username`, `meta_cluster_id`,`db_name`),
                                 KEY `idx_order_uuid` (`order_uuid`),
                                 KEY `idx_user_id` (`user_id`),
                                 KEY `idx_db_name` (`db_name`),
                                 KEY `idx_user_db_name` (`user_id`,`db_name`),
                                 KEY `idx_created_at` (`created_at`),
                                 KEY `idx_updated_at` (`updated_at`)
) COMMENT='MySQL数据库查询权限';

DROP TABLE IF EXISTS mysql_db_priv_apply_order;
CREATE TABLE `mysql_db_priv_apply_order` (
                                             `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                             `order_uuid` VARCHAR(20) NOT NULL DEFAULT '订单编号',
                                             `user_id` bigint NOT NULL COMMENT '用户id',
                                             `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                                             `name_zh` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
                                             `apply_status` TINYINT NOT NULL DEFAULT 1 COMMENT '申请状态: 0:未知, 1:申请中, 2:申请成功, 3:申请失败',
                                             `apply_reason` TEXT COMMENT '申请原因',
                                             `error_message` TEXT COMMENT '错误信息',
                                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                             PRIMARY KEY (`id`),
                                             UNIQUE KEY `Udx_order_uuid` (`order_uuid`),
                                             KEY `idx_user_id` (`user_id`),
                                             KEY `idx_username` (`username`),
                                             KEY `idx_apply_status` (`apply_status`),
                                             KEY `idx_created_at` (`created_at`),
                                             KEY `idx_updated_at` (`updated_at`)
) COMMENT='MySQL数据库查询申请单';

DROP TABLE IF EXISTS mysql_db_priv_apply;
CREATE TABLE `mysql_db_priv_apply` (
                                       `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                       `order_uuid` VARCHAR(20) NOT NULL DEFAULT '订单编号',
                                       `user_id` bigint NOT NULL COMMENT '用户id',
                                       `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                                       `name_zh` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
                                       `meta_cluster_id` bigint NOT NULL COMMENT '集群id',
                                       `cluster_name` varchar(100) NOT NULL DEFAULT '' COMMENT '集群名称',
                                       `db_name` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库名',
                                       `vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VIP 和端口',
                                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                       PRIMARY KEY (`id`),
                                       KEY `idx_order_uuid` (`order_uuid`),
                                       KEY `idx_user_id` (`user_id`),
                                       KEY `idx_username` (`username`),
                                       KEY `idx_meta_cluster_db_name` (`meta_cluster_id`,`db_name`),
                                       KEY `idx_db_name` (`db_name`),
                                       KEY `idx_user_db_name` (`user_id`,`db_name`),
                                       KEY `idx_created_at` (`created_at`),
                                       KEY `idx_updated_at` (`updated_at`)
) COMMENT='MySQL数据库查询权限申请明细';

DROP TABLE IF EXISTS mysql_db_priv_trash;
CREATE TABLE `mysql_db_priv_trash` (
                                       `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                                       `mysql_db_priv_id` bigint NOT NULL COMMENT '源申请id',
                                       `order_uuid` VARCHAR(20) NOT NULL DEFAULT '订单编号',
                                       `user_id` bigint NOT NULL COMMENT '用户id',
                                       `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                                       `name_zh` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
                                       `meta_cluster_id` bigint NOT NULL COMMENT '集群id',
                                       `cluster_name` varchar(100) NOT NULL DEFAULT '' COMMENT '集群名称',
                                       `db_name` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库名',
                                       `vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VIP 和端口',
                                       `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                       `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                       PRIMARY KEY (`id`),
                                       KEY `idx_mysql_db_priv_id` (`mysql_db_priv_id`),
                                       KEY `idx_order_uuid` (`order_uuid`),
                                       KEY `idx_user_id` (`user_id`),
                                       KEY `idx_username` (`username`),
                                       KEY `idx_meta_cluster_db_name` (`meta_cluster_id`,`db_name`),
                                       KEY `idx_db_name` (`db_name`),
                                       KEY `idx_user_db_name` (`user_id`,`db_name`),
                                       KEY `idx_created_at` (`created_at`),
                                       KEY `idx_updated_at` (`updated_at`)
) COMMENT='MySQL数据库查询权限垃圾箱';
