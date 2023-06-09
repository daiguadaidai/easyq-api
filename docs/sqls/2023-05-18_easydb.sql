DROP TABLE IF EXISTS `meta_cluster`;
CREATE TABLE `meta_cluster` (
                                `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增',
                                `name` varchar(100) NOT NULL COMMENT '集群名称',
                                `cluster_id` varchar(50) NOT NULL DEFAULT '' COMMENT '集群id',
                                `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0:否, 1:是',
                                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                `business_line` varchar(100) NOT NULL DEFAULT '' COMMENT '业务线名称',
                                `owner` varchar(50) NOT NULL DEFAULT '' COMMENT '负责人',
                                `domain_name` varchar(500) NOT NULL DEFAULT '' COMMENT '域名: 多个以逗号隔开',
                                `vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VIP 和端口',
                                `vpcgw_vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VPC GateWay VIP和端口',
                                `is_shard` tinyint NOT NULL DEFAULT '0' COMMENT '是否分片: 0:否, 1:是',
                                `category` tinyint NOT NULL DEFAULT '0' COMMENT '集群类别: 0:未知, 1.TCE自建MySQL, 2.公有云自建MySQL, 3.TCE-TDSQL, 4.公有云TDSQL-C, 5.公有云MySQL, 6.福州-MySQL(自建)',
                                `set_name` varchar(30) NOT NULL DEFAULT '' COMMENT 'set 名称',
                                `shard_type` varchar(25) NOT NULL DEFAULT '' COMMENT '分片类型: mycat, tdsql, zebra, sharding-jdbc, sharding-jdbc-only-db',
                                `read_host_port` varchar(30) NOT NULL DEFAULT '' COMMENT '只读地址',
                                `status` tinyint NOT NULL DEFAULT '2' COMMENT '-1: 已隔离, -2: 已删除, 0: 创建中, 1: 处理中, 2: 运行中, 3:未初始化, 4:初始化中, 5:删除中, 6: 重启中, 7: 迁移中',
                                PRIMARY KEY (`id`),
                                KEY `idx_name` (`name`),
                                KEY `idx_vip_port` (`vip_port`),
                                KEY `idx_vpcgw_vip_port` (`vpcgw_vip_port`),
                                KEY `idx_set_name` (`set_name`),
                                KEY `idx_cluster_id` (`cluster_id`)
) COMMENT='集群信息';

DROP TABLE IF EXISTS `instance`;
CREATE TABLE `instance` (
                            `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增',
                            `instance_id` varchar(50) NOT NULL DEFAULT '' COMMENT '实例ID',
                            `instance_name` varchar(50) NOT NULL DEFAULT '' COMMENT '实例名称',
                            `meta_cluster_id` bigint NOT NULL DEFAULT '0' COMMENT '集群id',
                            `machine_host` varchar(15) NOT NULL DEFAULT '' COMMENT '机器host',
                            `vpcgw_rip` varchar(15) NOT NULL DEFAULT '' COMMENT '机器host',
                            `port` int NOT NULL COMMENT '端口',
                            `role` varchar(15) NOT NULL DEFAULT '' COMMENT '实例角色: master, slave, backup',
                            `cpu` int NOT NULL DEFAULT '0' COMMENT '核数',
                            `mem` int NOT NULL DEFAULT '0' COMMENT '内存大小, 单位G',
                            `disk` int NOT NULL DEFAULT '0' COMMENT '磁盘容量, 单位G',
                            `master_host` varchar(15) NOT NULL DEFAULT '' COMMENT '主库host',
                            `master_port` int NOT NULL DEFAULT '0' COMMENT '主库端口',
                            `version` varchar(50) NOT NULL DEFAULT '' COMMENT '实例版本',
                            `is_maintenance` tinyint NOT NULL DEFAULT '0' COMMENT '是否在维护中: 0:否, 1:是',
                            `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0:否, 1:是',
                            `descript` text COMMENT '描述',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VIP 和端口',
                            `vpcgw_vip_port` varchar(190) NOT NULL DEFAULT '' COMMENT 'VPC GateWay VIP和端口',
                            `set_name` varchar(30) NOT NULL DEFAULT '' COMMENT 'set 名称',
                            PRIMARY KEY (`id`),
                            KEY `idx_instance_id` (`instance_id`),
                            KEY `idx_instance_name` (`instance_name`),
                            KEY `idx_created_at` (`created_at`),
                            KEY `idx_updated_at` (`updated_at`),
                            KEY `idx_machine_host_port` (`machine_host`,`port`),
                            KEY `idx_vpcgw_rip_port` (`vpcgw_rip`,`port`),
                            KEY `idx_master_host_port` (`master_host`,`master_port`),
                            KEY `idx_meta_cluster_id` (`meta_cluster_id`),
                            KEY `idx_set_name` (`set_name`)
) COMMENT='实例信息';