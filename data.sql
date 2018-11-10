INSERT INTO pokemon_type (id, name) VALUES (1, 'Grass');
INSERT INTO pokemon_type (id, name) VALUES (2, 'Fire');
INSERT INTO pokemon_type (id, name) VALUES (3, 'Poison');
INSERT INTO pokemon_type (id, name) VALUES (4, 'Water');
INSERT INTO pokemon_type (id, name) VALUES (5, 'Flying');
INSERT INTO pokemon_type (id, name) VALUES (6, 'Normal');
INSERT INTO pokemon_type (id, name) VALUES (7, 'Bug');

INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (1, 'Bulbasaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (2, 'Ivysaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (3, 'Venasaur', 1, 3);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (4, 'Charmander', 2, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (5, 'Charmeleon', 2, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (6, 'Charizard', 2, 5);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (7, 'Squirtle', 4, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (8, 'Wartortle', 4, null);
INSERT INTO pokemon (id, name, type_one_id, type_two_id) VALUES (9, 'Blastoise', 4, null);

INSERT INTO junidex.evolution_type (id, name) VALUES (1, 'By Level');

INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (1, 1, 2, 1, '{"level": "16"}');
INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (2, 2, 3, 1, '{"level": "32"}');
INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (3, 4, 5, 1, '{"level": "16"}');
INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (4, 5, 6, 1, '{"level": "36"}');
INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (5, 7, 8, 1, '{"level": "16"}');
INSERT INTO junidex.chain_evolution (id, current_id, evolution_id, type_id, details) VALUES (6, 8, 9, 1, '{"level": "36"}');
