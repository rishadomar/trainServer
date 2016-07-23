CREATE DATABASE train;
CREATE USER 'train'@'localhost' IDENTIFIED BY 'kaluma';
GRANT ALL PRIVILEGES ON * . * TO 'train'@'localhost';
FLUSH PRIVILEGES;

USE train;

CREATE TABLE `location` (
  `ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `latitude` decimal(16,13) NOT NULL DEFAULT 0,
  `longitude` decimal(16,13) NOT NULL DEFAULT 0,
  `locationType` char(50) NOT NULL DEFAULT 'Point',
  `timeCreated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `route` (
  `ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL DEFAULT '',
  `code` char(12) NOT NULL DEFAULT '',
  `timeCreated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `route_location_link` (
  `ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `routeId` int(10) NOT NULL DEFAULT 0,
  `locationId` int(10) NOT NULL DEFAULT 0,
  `timeCreated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `passenger` (
    `ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `routeId` int(10) NOT NULL DEFAULT 0,
    `locationId` int(10) NOT NULL DEFAULT 0,
    `direction` int(1) NOT NULL DEFAULT 0,
    `timeCreated` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

# Test data
INSERT INTO `route` VALUES (null, 'Cape Flats', '05', null);

INSERT INTO `location` VALUES (null, 'Crawford', -33.9762856,18.4986613, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9738922,18.5013328, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.971592,18.5015636, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9682952,18.5023601, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9637712,18.5024891, 'Point', null);
INSERT INTO `location` VALUES (null, 'Athlone', -33.9626014,18.4993885, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9597699,18.5023776, 'Point', null);
INSERT INTO `location` VALUES (null, 'Hazendal', -33.9556684,18.5037623, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9521823,18.50311, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9508734,18.4997265, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.950093,18.4983154, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9484065,18.4963126, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9465185,18.4936725, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9443283,18.4914725, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9420123,18.4913511, 'Point', null);
INSERT INTO `location` VALUES (null, 'Pinelands', -33.9420123,18.4913511, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9378835,18.4913662, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9353533,18.4921249, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9332258,18.4938698, 'Point', null);
INSERT INTO `location` VALUES (null, 'Ndabeni', -33.929084,18.4961305, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9277747,18.4972533, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.925861,18.4967981, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9243754,18.4964643, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9242243,18.4942339, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9240733,18.4914876, 'Point', null);
INSERT INTO `location` VALUES (null, 'Maitland', -33.9242369,18.487391, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.924879,18.483355, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9254204,18.4784238, 'Point', null);
INSERT INTO `location` VALUES (null, 'Koeberg', -33.9254204,18.4784238, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9268364,18.4727998, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9269966,18.4686155, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9269966,18.4686155, 'Point', null);
INSERT INTO `location` VALUES (null, 'Salt River', -33.9274239,18.467135, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9268186,18.4631867, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9266227,18.4584017, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9262132,18.4528012, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9251984,18.4494109, 'Point', null);
INSERT INTO `location` VALUES (null, 'Woodstock', -33.9255189,18.4459992, 'Station', null);
INSERT INTO `location` VALUES (null, '', -33.9248601,18.4406133, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9253052,18.4398408, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9252874,18.4355493, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9242903,18.4316869, 'Point', null);
INSERT INTO `location` VALUES (null, '', -33.9236849,18.4294767, 'Point', null);
INSERT INTO `location` VALUES (null, 'Cape Town', -33.9236849,18.4294767, 'Station', null);

