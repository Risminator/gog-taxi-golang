insert into gog_demo.dock (name, active, latitude, longitude) values 
('Dock1', true, 59.92549, 30.38959),
('Dock2', true, 59.9425, 30.40474),
('Dock3', false, 59.95356, 30.34985);

insert into gog_demo.vessel (model, seats, is_approved, latitude, longitude) values 
('Duck', 4, true, 55.123456, 37.1234567),
('Swan', 3, true, 55.123456, 37.1234567),
('Fish', 4, true, 55.123456, 37.1234567);

insert into gog_demo.driver (first_name, last_name, vessel_id, status, balance, cert_first_aid, cert_driving) values
('George', 'Boole', 1, 'waiting', 512, 1, 1),
('Mary', 'Sin', 2, 'waiting', 1024, 1, 1),
('Alan', 'Remedy', 3, 'waiting', 1024, 1, 1);

insert into gog_demo.customer (phone, first_name, last_name) values
('12345', 'Yan', 'Oreshko')