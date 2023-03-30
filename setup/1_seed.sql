\c chargepoints

insert into groups (max_current) values (25), (12.5);

insert into connectors (group_id, status) values
(2, 'Available'),
(2, 'Charging'),
(1, 'Charging'),
(1, 'Available'),
(1, 'Charging'),
(2, 'Available'),
(1, 'Available'),
(1, 'Charging'),
(2, 'Available'),
(1, 'Charging');

