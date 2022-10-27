create table if not exists customers (
id int primary key not null,
customers_name char(50) not null
);

create table if not exists products (
id int primary key not null,
product_name char(50) not null
);

create table if not exists orders (
id int primary key not null,
customer_id int not null,
products_id int not null,
order_date date not null,
total float not null
);
--insert table customers
insert into customers values('chintya');
insert into customers values('suci');

--insert table products
insert into products values(1,'gitar');
insert into products  values(2,'biola');

--insert table orders
insert into orders  values(1,1,1,'26/10/2022',1.0);
insert into orders values(2,2,2,'27/10/2022',2.0);


--delete table customer
delete from customers where customers_name = 'suci';

--delete table products
delete from products where id =1;

--delete table orders
delete from orders where id =1;


--update table customers 
update customers set customers_name = 'wardani'
where customers_name ='chintya';

--update table products 
update products set product_name  = 'bass'
where id=2;

--update table orders 
update orders set total = subquery.total + 2.0 
from (select id, total from orders where id = 2) as subquery where orders.id = 2;

