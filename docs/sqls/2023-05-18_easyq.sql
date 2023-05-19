DROP TABLE IF EXISTS user;
CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                        `username` varchar(50) NOT NULL COMMENT '用户名',
                        `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
                        `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮件',
                        `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '电话',
                        `name_en` varchar(50) NOT NULL DEFAULT '' COMMENT '英文名',
                        `name_zh` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
                        `role` varchar(15) NOT NULL DEFAULT '' COMMENT '角色: dba, dev, guest',
                        `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0:否, 1:是',
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `udx_username` (`username`),
                        KEY `idx_created_at` (`created_at`),
                        KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB CHARSET=utf8mb4 COMMENT='用户表';

DROP TABLE IF EXISTS mysql_db_priv;
CREATE TABLE mysql_db_priv(
                              `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                              `user_id` bigint NOT NULL COMMENT '用户id',
                              `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
                              `meta_cluster_id` bigint NOT NULL COMMENT '集群id',
                              `cluster_name` varchar(100) NOT NULL DEFAULT '' COMMENT '集群名称',
                              `db_name` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库名',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_user_id` (`user_id`),
                              KEY `idx_username` (`username`),
                              KEY `idx_meta_cluster_db_name` (`meta_cluster_id`, `db_name`),
                              KEY `idx_db_name` (`db_name`),
                              KEY `idx_user_db_name` (`user_id`, `db_name`),
                              KEY `idx_created_at` (`created_at`),
                              KEY `idx_updated_at` (`updated_at`)
) COMMENT = 'MySQL数据库查询权限';