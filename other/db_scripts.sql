drop table if exists order_dish_item_statuses;
drop table if exists order_dish_items;
drop table if exists order_items;
drop table if exists orders;

create table orders (
    id serial not null,
    client_id varchar(255) not null,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);

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

CREATE TABLE IF NOT EXISTS order_dish_items
(
    id serial NOT NULL,
    client_id varchar(255) not null,
    order_id integer not null,
    dish_name character varying(255) NOT NULL,
    time_created timestamp with time zone,
    status varchar(100) not null,
    CONSTRAINT order_dish_items_pkey PRIMARY KEY (id)
    );

create table if not exists order_dish_item_statuses(
    order_dish_item_id int not null,
    timestamp timestamp with time zone,
    status varchar(100) not null,
    constraint order_dish_item_statuses_fk1 FOREIGN KEY (order_dish_item_id) REFERENCES order_dish_items (id)
    );


select * from orders;
