-- 创建数据库
create database database_learning;
use database_learning;

-- 查看数据库
show databases;

-- 测试索引
DROP TABLE IF EXISTS index_learning;
CREATE TABLE index_learning(
                               `id` INT NOT NULL AUTO_INCREMENT  COMMENT '自增主键' ,
                               `name` VARCHAR(255)    COMMENT '唯一索引' ,
                               PRIMARY KEY (id),
                               unique index name_index(name)

)  COMMENT = '测试唯一索引和主键';


-- 查看表
show tables;
select * from index_learning;


-- 插入数据
insert into index_learning values (1,'user1'),(2,'user2'),(3,'user3');
insert into index_learning values (4,null); -- OK
insert into index_learning values (5,null); -- OK
insert into index_learning values (null,'user6'); -- OK 自增主键如果是null则从max(id)+1开始自动插入，如果还没有数据，则从1开始
insert into index_learning values (null,'user7');


-- 继续测试
DROP TABLE IF EXISTS index_learning;
CREATE TABLE index_learning(
                               `id` VARCHAR(255) ,
                               `name` VARCHAR(255)   NOT NULL   COMMENT '唯一索引' ,
                               PRIMARY KEY (id),
                               unique index name_index(name)

)  COMMENT = '测试主键是否可以为空';

insert into index_learning values ('userKey1','user1'),('userKey2','user2'),('userKey3','user3');
insert into index_learning values (null,'user4'); -- wrong :  Column 'id' cannot be null
