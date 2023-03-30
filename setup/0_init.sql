\c chargepoints

CREATE TABLE groups (
    id serial primary key,
    max_current float(10) check (max_current >= 0)
);

CREATE TABLE connectors (
    id serial primary key,
    group_id integer not null,
    status varchar(10) check (status IN ('Available', 'Charging'))
);

create index group_id_idx on connectors (group_id);