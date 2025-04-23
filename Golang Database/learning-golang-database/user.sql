create table user(
username varchar(100) not null primary key,
password varchar(100) not null
);

describe user;

insert into user values ("admin","admin");

select * from user;
