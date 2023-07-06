ALTER TABLE mysql_db_priv
    ADD `is_shard` tinyint NOT NULL DEFAULT '0' COMMENT '是否分片: 0:否, 1:是',
    ADD `shard_type` varchar(25) NOT NULL DEFAULT '' COMMENT '分片类型: mycat, tdsql, zebra, sharding-jdbc, sharding-jdbc-only-db';

ALTER TABLE mysql_db_priv_apply
    ADD `is_shard` tinyint NOT NULL DEFAULT '0' COMMENT '是否分片: 0:否, 1:是',
    ADD `shard_type` varchar(25) NOT NULL DEFAULT '' COMMENT '分片类型: mycat, tdsql, zebra, sharding-jdbc, sharding-jdbc-only-db';

ALTER TABLE mysql_db_priv_trash
    ADD `is_shard` tinyint NOT NULL DEFAULT '0' COMMENT '是否分片: 0:否, 1:是',
    ADD `shard_type` varchar(25) NOT NULL DEFAULT '' COMMENT '分片类型: mycat, tdsql, zebra, sharding-jdbc, sharding-jdbc-only-db';
