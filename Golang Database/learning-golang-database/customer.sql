create table customer(
id varchar(100) not null primary key,
name varchar(100) not null
);

describe customer;

select * from customer;

delete from customer;

alter table customer
add column email varchar(100),
add column balance int default 0,
add column rating double default 0.0,
add column created_at timestamp default current_timestamp,
add column birth_date date,
add column married boolean default false;

update customer set email = null, birth_date = null where id = "P003";