drop table if exists orders;

create table orders (
                        id serial not null,
                        client_id varchar(255) not null,
                        CONSTRAINT orders_pkey PRIMARY KEY (id)
);

drop table if exists order_items;

CREATE TABLE IF NOT EXISTS order_items
(
    id serial NOT NULL,
    client_id varchar(255) not null,
    order_id integer not null,
    dish_name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price numeric(5,2) NOT NULL,
    CONSTRAINT order_items_pkey PRIMARY KEY (id)
    );

select * from orders;
