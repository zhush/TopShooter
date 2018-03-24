drop database if exists topshooter;
create database if not exists topshooter default character set utf8 collate utf8_general_ci;
use topshooter;

create table t_account(accid bigint not null auto_increment,
                    accountName varchar(128) not null,
                    password varchar(128) not null,
                    createtm datetime not null,
                    lastLoginTm datetime not null,
                    gameid bigint,
                    primary key(accid));

insert into t_account(accountName, password, createtm, lastLoginTm, gameid) values("admin", md5("admin"), now(), now(), 0);
insert into t_account(accountName, password, createtm, lastLoginTm, gameid) values("test", md5("test"), now(), now(), 0);


create table t_role(roleid bigint not null auto_increment,
                    accid bigint not null,
                    nickName varchar(128) not null,
                    sex tinyint,
                    templateId int,
                    createtm datetime not null,
                    lastsceneid int default 0,
                    lastposX int,
                    lastposY int,
                    handWeapon int,
                    bulletCount int,
                    weaponList text,
                    level int,
                    gold  bigint,
                    primary key(roleid));

