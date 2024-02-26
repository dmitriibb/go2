drop table if exists client_order_dish_item_statuses;
drop table if exists client_order_dish_items;
drop table if exists client_order_items;
drop table if exists client_orders;

create table client_orders (
    id serial not null,
    client_id varchar(255) not null,
    CONSTRAINT client_orders_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS client_order_items
(
    id serial NOT NULL,
    client_id varchar(255) not null,
    client_order_id integer not null,
    dish_name character varying(255) NOT NULL,
    quantity integer NOT NULL,
    price numeric(5,2) NOT NULL,
    CONSTRAINT client_order_items_pkey PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS client_order_dish_items
(
    id serial NOT NULL,
    client_id varchar(255) not null,
    client_order_id integer not null,
    dish_name character varying(255) NOT NULL,
    time_created timestamp with time zone,
    status varchar(100) not null,
    CONSTRAINT client_order_dish_items_pkey PRIMARY KEY (id)
    );

create table if not exists client_order_dish_item_statuses(
    client_order_dish_item_id int not null,
    timestamp timestamp with time zone,
    status varchar(100) not null,
    constraint client_order_dish_item_statuses_fk1 FOREIGN KEY (client_order_dish_item_id) REFERENCES client_order_dish_items (id)
    );


select * from client_orders;
