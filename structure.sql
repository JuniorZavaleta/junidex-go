DROP database if exists junidex;

CREATE DATABASE junidex;

USE junidex;

CREATE TABLE pokemon
(
  id smallint unsigned PRIMARY KEY AUTO_INCREMENT,
  name varchar(30) NOT NULL,
  type_one_id tinyint unsigned NOT NULL,
  type_two_id tinyint unsigned
);

CREATE TABLE pokemon_type
(
  id tinyint unsigned PRIMARY KEY AUTO_INCREMENT,
  name varchar(12) NOT NULL
);

CREATE TABLE evolution_type
(
  id tinyint unsigned PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(12) NOT NULL
);

DROP TABLE chain_evolution;

CREATE TABLE chain_evolution
(
  id tinyint unsigned PRIMARY KEY AUTO_INCREMENT,
  current_id smallint unsigned NOT NULL,
  evolution_id smallint unsigned NOT NULL,
  type_id tinyint unsigned NOT NULL,
  details json
);

ALTER TABLE pokemon ADD CONSTRAINT pkmn_type_one_fk FOREIGN KEY (type_one_id) references pokemon_type(id);
ALTER TABLE pokemon ADD CONSTRAINT pkmn_type_two_fk FOREIGN KEY (type_two_id) references pokemon_type(id);
ALTER TABLE chain_evolution ADD CONSTRAINT ch_evol_current_fk FOREIGN KEY (current_id) references pokemon(id);
ALTER TABLE chain_evolution ADD CONSTRAINT ch_evol_evolution_fk FOREIGN KEY (evolution_id) references pokemon(id);
ALTER TABLE chain_evolution ADD CONSTRAINT ch_evol_type_fk FOREIGN KEY (type_id) references evolution_type(id);
