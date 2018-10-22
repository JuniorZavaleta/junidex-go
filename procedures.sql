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
           c.name as current_pokemon,
           e.name as evolution_pokemon,
           t.name as ev_type_name,
           chain_evolution.details
    FROM chain_evolution
           JOIN pokemon c ON current_id = c.id
           JOIN pokemon e ON evolution_id = e.id
           JOIN evolution_type t ON type_id = t.id
    WHERE JSON_CONTAINS(Ids, JSON_ARRAY(current_id), '$');
  END;