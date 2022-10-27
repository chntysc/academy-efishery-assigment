create table if not exists products (
id serial not null,
stock int not null
);

insert into products (stock) values (1);
insert into products (stock) values (2);
insert into products (stock) values (3);
insert into products (stock) values (4);
insert into products (stock) values (5);

update products set stock = subquery.stock - 2 
from (select id, stock from products where id = 1) as subquery where products.id = 1;

create function kurangi_stock (INT, INT) returns products as
'
update products set stock = subquery.stock - $2 from (select id, stock from products where id = $1) as subquery where products.id = $1;
select * from products where id = $1;
'
language 'sql';

select kurangi_stock(4,2);