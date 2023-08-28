mysql 版本

```sql
mysql> select version();
+-----------+
| version() |
+-----------+
| 8.0.33    |
+-----------+
```



查看隔离级别

```sql
mysql> select @@GLOBAL.transaction_isolation, @@transaction_isolation;
+--------------------------------+-------------------------+
| @@GLOBAL.transaction_isolation | @@transaction_isolation |
+--------------------------------+-------------------------+
| REPEATABLE-READ                | REPEATABLE-READ         |
+--------------------------------+-------------------------+
1 row in set (0.01 sec)
```

mysql默认隔离级别是`可重复读`。



创建表

```sql
CREATE TABLE test_transaction_isolation(
    `id` INT NOT NULL AUTO_INCREMENT  COMMENT '用户id' ,
    `name` VARCHAR(255) NOT NULL   COMMENT '用户名称' ,
    PRIMARY KEY (id)
)  COMMENT = '测试transaction_isolation';
```



可重复读隔离级别幻读现象

| 事务A                                                        | 事务B                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| begin;                                                       | begin;                                                       |
|                                                              | SELECT * FROM test_transaction_isolation WHERE id = 3;       |
| INSERT INTO test_transaction_isolation (id, name) VALUES (3,'test_user'); |                                                              |
| commit;                                                      |                                                              |
|                                                              | SELECT * FROM test_transaction_isolation WHERE id = 3;       |
|                                                              | UPDATE test_transaction_isolation SET name = 'tree' WHERE id = 3; |
|                                                              | SELECT * FROM test_transaction_isolation WHERE id = 3;       |
|                                                              | commit;                                                      |

[参考](https://www.liaoxuefeng.com/wiki/1177760294764384/1245268672511968)



```sql
CREATE TABLE test_select_for_update(
    `id` INT NOT NULL AUTO_INCREMENT  COMMENT '用户id' ,
    `name` VARCHAR(255) NOT NULL   COMMENT '用户名称' ,
 	  `unuse_id` INT    COMMENT '用于测试select for update' ,
    PRIMARY KEY (id)
)  COMMENT = '测试select for update';


insert into test_select_for_update (`id`,`name`,`unuse_id`) values (0,'test_user',9000);
```

| 事务A                                                        | 事务B                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| begin;                                                       | begin;                                                       |
| select max(unuse_id) from test_select_for_update for update;-- 结果9000 |                                                              |
|                                                              | select max(unuse_id) from test_select_for_update for update;-- 卡住 |
| 执行sql 1                                                    |                                                              |
| commit;                                                      |                                                              |
|                                                              | select得到结果9001                                           |
|                                                              | 执行sql 2                                                    |
|                                                              | commit;                                                      |

```sql
-- sql 1
insert into test_select_for_update (`id`,`name`,`unuse_id`) values (0,'test_user_A',9001);


-- sql 2
insert into test_select_for_update (`id`,`name`,`unuse_id`) values (0,'test_user_B',9002);
```

 