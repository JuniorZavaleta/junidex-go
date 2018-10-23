INSERT INTO pokemon_type (id, name) VALUES (1, 'Grass');
INSERT INTO pokemon_type (id, name) VALUES (2, 'Fire');
INSERT INTO pokemon_type (id, name) VALUES (3, 'Poison');
INSERT INTO pokemon_type (id, name) VALUES (4, 'Water');
INSERT INTO pokemon_type (id, name) VALUES (5, 'Flying');

INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (1, 'Bulbasaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (2, 'Ivysaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (3, 'Venasaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (4, 'Charmander', 2, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (5, 'Charmeleon', 2, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (6, 'Charizard', 2, 5);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (7, 'Squirtle', 4, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (8, 'Wartortle', 4, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (9, 'Blastoise', 4, null);
