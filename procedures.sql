USE junidex;

create procedure get_all_pokemon()
  BEGIN
    SELECT p.id, p.name, p.type_one_id, p.type_one_id, p.type_two_id, p.has_preevolution, t1.name as type_one, t2.name as type_two
    FROM pokemon as p
           JOIN pokemon_type as t1 ON p.type_one_id = t1.id
           LEFT JOIN pokemon_type as t2 ON p.type_two_id = t2.id;
  END;

CREATE PROCEDURE filter_pokemon(TypeOne varchar(30), TypeTwo varchar(30))
  BEGIN
    SELECT p.id, p.name, p.type_one_id, p.type_one_id, p.type_two_id, p.has_preevolution, t1.name as type_one, t2.name as type_two
    FROM pokemon as p
           JOIN pokemon_type as t1 ON p.type_one_id = t1.id
           LEFT JOIN pokemon_type as t2 ON p.type_two_id = t2.id
    WHERE
        CASE
          WHEN TypeOne IS NULL OR TypeOne = '' THEN 1
          WHEN t1.name like concat('%',TypeOne,'%') THEN 1
          ELSE 0 END = 1

      AND

        CASE
          WHEN TypeTwo IS NULL OR TypeTwo = '' THEN 1
          WHEN t2.name like concat('%',TypeTwo,'%') THEN 1
          ELSE 0 END = 1
    ;
  END;

CREATE PROCEDURE get_complete_chain_evolution(PokemonId tinyint unsigned)
  BEGIN
    #   IF
    DECLARE HasPreevolution BOOLEAN DEFAULT TRUE;
    DECLARE IsEvolution BOOLEAN DEFAULT FALSE;
    DECLARE BaseId TINYINT UNSIGNED DEFAULT PokemonId;
    DECLARE Ids JSON DEFAULT "[]";

    SELECT count(*) > 0 INTO IsEvolution
    FROM chain_evolution
           JOIN pokemon ON current_id = pokemon.id
    WHERE evolution_id = PokemonId;

    IF IsEvolution THEN
      WHILE HasPreevolution = TRUE DO
        SELECT current_id, pokemon.has_preevolution INTO BaseId, HasPreevolution
        FROM chain_evolution
               JOIN pokemon ON current_id = pokemon.id
        WHERE evolution_id = BaseId;
      END WHILE;
    END IF;

    SET Ids = JSON_ARRAY(BaseId);

    -- Primera evolución
    SELECT JSON_MERGE_PRESERVE(json_arrayagg(evolution_id), Ids) INTO Ids
    FROM chain_evolution
    WHERE current_id = BaseId;

    -- ultima evolucion
    SELECT JSON_MERGE_PRESERVE(JSON_ARRAYAGG(evolution_id), Ids) INTO Ids
    FROM chain_evolution
    WHERE json_contains(Ids, json_array(current_id), '$');

    SELECT
           chain_evolution.id,
           c.id as current_id,
           e.id as evolution_id,
           t.name as ev_type_name,
           chain_evolution.details
    FROM chain_evolution
           JOIN pokemon c ON current_id = c.id
           JOIN pokemon e ON evolution_id = e.id
           JOIN evolution_type t ON type_id = t.id
    WHERE JSON_CONTAINS(Ids, JSON_ARRAY(current_id), '$');
  END;

CREATE PROCEDURE create_pokemon(IN PokemonData json)
  BEGIN
    DECLARE PokemonId SMALLINT UNSIGNED DEFAULT 0;
    DECLARE TypeTwo TINYINT UNSIGNED DEFAULT NULL;

    SELECT
           CASE WHEN JSON_EXTRACT(PokemonData, '$.TypeTwoId.Null') = FALSE
                     THEN JSON_EXTRACT(PokemonData, '$.TypeTwoId.Value')
                ELSE NULL
               END INTO TypeTwo;


    INSERT INTO pokemon (name, type_one_id, type_two_id, has_preevolution)
    values (JSON_UNQUOTE(JSON_EXTRACT(PokemonData, '$.Name')),
            JSON_EXTRACT(PokemonData, '$.TypeOneId'),
            TypeTwo,
            JSON_EXTRACT(PokemonData, '$.HasPreEvol'));
    SELECT LAST_INSERT_ID() INTO PokemonId;

    SELECT p.*, t1.name as type_one, t2.name as type_two
    FROM pokemon as p
           JOIN pokemon_type as t1 ON p.type_one_id = t1.id
           LEFT JOIN pokemon_type as t2 ON p.type_two_id = t2.id
    WHERE p.id = PokemonId;
  END;
