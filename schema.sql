create database manager_inventory;

use manager_inventory;

create table
    USERS(
        id int not null auto_increment,
        email varchar(200) not null,
        name varchar(200) not null,
        password varchar(200) not null,
        primary key (id)
    );

create table
    PRODUCTS(
        id int not null auto_increment,
        name varchar(200) not null,
        description varchar(200) not null,
        price float not null,
        created_by int not null,
        primary key (id),
        foreign Key (created_by) references USERS(id)
    );

create table
    ROLES(
        id int not null auto_increment,
        name varchar(200) not null,
        primary key (id)
    );

create table
    USER_ROLES(
        id int not null auto_increment,
        user_id int not null,
        role_id int not null,
        primary key (id),
        foreign Key (user_id) references USERS(id),
        foreign Key (role_id) references ROLES(id)
    );

insert into ROLES (id, name)
values (1, 'admin')
insert into ROLES (id, name)
values (2, 'seller')
insert into ROLES (id, name)
values (3, 'customer')