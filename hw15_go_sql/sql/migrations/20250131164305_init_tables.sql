-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id serial primary key,
    username varchar(20) not null,
    email varchar(40) not null unique,
    password varchar(20) not null
    );

create table if not exists orders
(
    id serial primary key,
    user_id int not null,
    order_date date not null,
    total_amount numeric not null,
    foreign key (user_id) references users(id) on delete cascade
    );

create table if not exists products
(
    id serial primary key,
    product_name varchar(80) not null,
    price numeric not null
    );

create table if not exists orderProducts
(
    order_id int not null,
    product_id int not null,
    quantity int not null,
    primary key (order_id, product_id),
    foreign key (order_id) references orders (id) on delete cascade,
    foreign key (product_id) references products (id) on delete cascade
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists orderProducts;
drop table if exists products;
drop table if exists orders;
drop table if exists users;
-- +goose StatementEnd
