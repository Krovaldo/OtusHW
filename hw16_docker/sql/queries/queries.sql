-- name: InsertUser :exec
insert into users (username, email, password)
values ($1, $2, $3);

-- name: DeleteUserByID :exec
delete from users
where id = $1;

-- name: InsertProduct :exec
insert into products (product_name, price)
values ($1, $2);

-- name: DeleteProduct :exec
delete from products
where price = $1;

-- name: InsertOrder :exec
insert into orders (user_id, order_date, total_amount)
values ($1, $2, $3);

-- name: InsertOrderProduct :exec
insert into orderproducts (order_id, product_id, quantity)
values ($1, $2, $3);

-- name: DeleteOrder :exec
delete from orders
where id = $1;

-- name: GetAllUsers :many
select id, username, email from users;

-- name: GetUserByUsername :one
select id, username, email from users
where username = $1;

-- name: GetAllProducts :many
select * from products;

-- name: GetOrdersByUser :many
select o.id, o.order_date, o.total_amount, u.username, u.email
from orders as o
         join users as u on o.user_id = u.id
where o.user_id = $1;