\c distribution

CREATE TABLE groups (
    id serial primary key,
    max_current float(10) check (max_current >= 0)
);

CREATE TABLE charge_points (
    id serial primary key,
    group_id integer not null,
    priority integer not null
);

CREATE TABLE charge_point_connectors (
    id serial primary key,
    charge_point_id integer not null,
    status varchar(10) check (status in ('Available', 'Charging'))
);