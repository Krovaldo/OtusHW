create database my_db;

create table users
(
	id serial primary key,
	username varchar(20) not null,
	email varchar(40) not null unique,
	password varchar(20) not null
);

create table orders
(
	id serial primary key,
	user_id int not null,
	order_date date not null,
	total_amount numeric not null,
	foreign key (user_id) references users(id) on delete cascade
);

create table products
(
	id serial primary key,
	product_name varchar(80) not null,
	price numeric not null
);

create table orderproducts
(
	order_id int not null,
	product_id int not null,
	quantity int not null,
	primary key (order_id, user_id, product_id),
	foreign key (order_id) references orders (id) on delete cascade,
	foreign key (product_id) references products (id) on delete cascade
);

insert into users (username, email, password)
values
('Vitalik', 'vitalya@yandex.ru', '12345!'),
('Dmitry', 'dima5@gmail.com', '11111'),
('Andrey', 'andr@gmail.com', 'qwerty');

update users
set username = 'Ilya', email = 'ilya123@gmail.com'
where username = 'Vitalik';

delete from users
where id = 3;

insert into products (product_name, price)
values
('Laptop Thinkpad', 78000),
('Wireless keyboard', 5000),
('Bluetooth headphones', 7000);

update products
set price = price + 5000
where id = 1;

delete from products
where price < 5500;

insert into orders (user_id, order_date, total_amount)
values
(1, '2024-11-01', 90000);

insert into orderproducts (order_id, product_id, quantity)
values
(1, 1, 1),
(1, 3, 1);

delete from orders
where id = 1;

select id, username, email from users;

select id, username, email from users
where username = 'Ilya';

select * from products;

select product_name from products
where price > 20000;

insert into orders (user_id, order_date, total_amount)
values
(2, '2024-09-01', 90000);

select o.id, o.order_date, o.total_amount, u.username, u.email 
from orders as o
join users as u on o.user_id = u.id
where o.user_id = 1;

select u.id, u.username, u.email, sum(o.total_amount), 
	avg(p.price) as average_product_price
from users as u
join orders as o on u.id = o.user_id
join orderproducts as op on o.id=op.order_id
join products as p on op.product_id = p.id
where u.username='Dmitry'
group by u.id;

create index idx_orders_all on orders(user_id, order_date, total_amount);

create index idx_products_insert on products(product_name, price);