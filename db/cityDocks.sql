insert into gog_demo.dock (name, active, longitude, latitude) values 
('Point 1', true, 30.302148, 59.937396),
('Point 2', true, 30.305966, 59.938856),
('Point 3', true, 30.307799, 59.948990),
('Point 4', true, 30.333690, 59.947466),
('Point 5', true, 30.332535, 59.952730),
('Point 6', true, 30.354926, 59.952989),
('Point 7', true, 30.302276, 59.940157),
('Point 8', true, 30.259806, 59.975065),
('Point 9', true, 30.297678, 59.929261),
('Point 10', true, 30.325440, 59.942201),
('Point 11', true, 30.315886, 59.921489),
('Point 12', true, 30.276684, 59.950122),
('Point 13', true, 30.224572, 59.968673),
('Point 14', true, 30.335371, 59.959295),
('Point 15', true, 30.405681, 59.948474),
('Point 16', true, 30.389835, 59.929912),
('Point 17', true, 30.462815, 59.869128),
('Point 18', true, 30.457290, 59.871251),
('Point 19', true, 30.399514, 59.951630);

insert into gog_demo.vessel (model, seats, is_approved, latitude, longitude) values 
('Duck', 4, true, 30.3551579, 59.9529771),
('Swan', 5, true, 30.3075201, 59.9443755),
('Fish', 4, true, 30.2931345, 59.9341729);

insert into gog_demo.driver (first_name, last_name, vessel_id, status, balance, cert_first_aid, cert_driving) values
('George', 'Boole', 1, 'waiting', 512, 1, 1),
('Mary', 'Sin', 2, 'waiting', 1024, 1, 1),
('Alan', 'Remedy', 3, 'waiting', 1024, 1, 1);

insert into gog_demo.customer (phone, first_name, last_name) values
('12345', 'Yan', 'Oreshko');