insert into gog_demo.dock (name, lat, lon, location, working) values 
('Dock1', 55.7538337, 37.6211812, st_setsrid(st_makepoint(37.6211812, 55.7538337), 4326), true),
('Dock2', 55.750934, 37.6235352, st_setsrid(st_makepoint(37.6235352, 55.750934), 4326), true),
('Dock3', 55.7533247, 37.6233312, st_setsrid(st_makepoint(37.6233312, 55.7533247), 4326), false);

insert into gog_demo.vessel (model, seats, is_approved, lon, lat, location) values 
('Duck', 4, true, 55.123456, 37.1234567, st_setsrid(st_makepoint(37.123456, 55.123456), 4326)),
('Swan', 3, true, 55.123456, 37.1234567, st_setsrid(st_makepoint(37.123456, 55.123456), 4326)),
('Fish', 4, true, 55.123456, 37.1234567, st_setsrid(st_makepoint(37.123456, 55.123456), 4326));

insert into gog_demo.driver (first_name, last_name, vessel_id, status, balance, cert_first_aid, cert_driving) values
('George', 'Boole', 1, 'waiting', 512, 1, 1),
('Mary', 'Sin', 2, 'waiting', 1024, 1, 1),
('Alan', 'Remedy', 3, 'waiting', 1024, 1, 1);