\c distribution

insert into groups (max_current) values (25), (12.5);

insert into charge_points (group_id, priority) values
(2, 2),
(2, 3),
(1, 1),
(1, 6),
(1, 1),
(2, 4),
(1, 1),
(1, 3),
(2, 2),
(1, 5);

insert into charge_point_connectors (charge_point_id, status) values
(1, 'Available'),
(1, 'Available'),
(2, 'Charging'),
(2, 'Charging'),
(3, 'Available'),
(3, 'Charging'),
(4, 'Available'),
(4, 'Available'),
(5, 'Charging'),
(5, 'Available'),
(6, 'Charging'),
(6, 'Charging'),
(7, 'Charging'),
(7, 'Charging'),
(8, 'Charging'),
(8, 'Charging'),
(9, 'Charging'),
(9, 'Charging'),
(10, 'Charging'),
(10, 'Charging');
