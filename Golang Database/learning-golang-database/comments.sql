create table comments (
id int not null primary key auto_increment,
email varchar(100) not null,
comment text
);

describe comments;

select * from comments;