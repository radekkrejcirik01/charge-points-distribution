\c chargepoints

CREATE TABLE statuses (
    id serial primary key,
    connector_status varchar(10) not null
);


create index statuses_connector_status_idx on statuses (connector_status);
