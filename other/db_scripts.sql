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
    dish_name varchar(255) NOT NULL,
    comment varchar(255),
    price numeric(5,2) NOT NULL,
    CONSTRAINT client_order_items_pkey PRIMARY KEY (id),
    constraint client_order_items_order_fk FOREIGN KEY (client_order_id) REFERENCES client_orders (id)
);

alter table client_order_items
add constraint client_order_items_comment_not_err CHECK ( comment != 'error db' ) NOT VALID;

select * from client_orders;
