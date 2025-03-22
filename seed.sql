-- Insert data into Users
INSERT INTO
    "users" (
        "id",
        "name",
        "username",
        "password",
        "role",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e5',
        'John Doe',
        'johndoe',
        '$2a$10$Kf8OfW1QdeSonf3iAj4fiuOSBAOsy4UZb1JXdxoUSg.igJaW2omWm',
        2,
        NOW(),
        NOW(),
        NULL
    ),
    (
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e6',
        'Jane Doe',
        'janedoe',
        '$2a$10$Kf8OfW1QdeSonf3iAj4fiuOSBAOsy4UZb1JXdxoUSg.igJaW2omWm',
        1,
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into Classes
INSERT INTO
    "classes" (
        "id",
        "name",
        "description",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'a54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'Warrior',
        'A strong and resilient class specialized in close combat.',
        NOW(),
        NOW(),
        NULL
    ),
    (
        'a54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'Mage',
        'A class focused on casting powerful spells from a distance.',
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into Races
INSERT INTO
    "races" (
        "id",
        "name",
        "description",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'c54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'Human',
        'A versatile race with balanced attributes.',
        NOW(),
        NOW(),
        NULL
    ),
    (
        'c54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'Elf',
        'An agile and intelligent race known for their grace and magic.',
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into Difficulty Levels
INSERT INTO
    "difficulty_levels" (
        "id",
        "name",
        "description",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'd54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'Easy',
        'For beginners and casual players.',
        NOW(),
        NOW(),
        NULL
    ),
    (
        'd54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'Medium',
        'A balanced challenge for players with some experience.',
        NOW(),
        NOW(),
        NULL
    ),
    (
        'd54455e3-6b13-4c77-b74b-8b6c8f8d72e3',
        'Hard',
        'For experienced players looking for a tough challenge.',
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into Characters
INSERT INTO
    "characters" (
        "id",
        "user_id",
        "name",
        "description",
        "class_id",
        "race_id",
        "images",
        "private",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'e54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e5',
        'John Warrior',
        'A powerful warrior skilled in melee combat.',
        'a54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'c54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        '{}',
        FALSE,
        NOW(),
        NOW(),
        NULL
    ),
    (
        'e54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e6',
        'Jane Mage',
        'A mage capable of casting powerful elemental spells.',
        'a54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'c54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        '{}',
        FALSE,
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into Quests
INSERT INTO
    "quests" (
        "id",
        "name",
        "description",
        "difficulty_level_id",
        "images",
        "private",
        "created_at",
        "updated_at",
        "deleted_at"
    )
VALUES
    (
        'f54455e3-6b13-4c77-b74b-8b6c8f8d72e1',
        'Defeat the Goblin King',
        'A quest to defeat the Goblin King and bring peace to the village.',
        'd54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        '{}',
        FALSE,
        NOW(),
        NOW(),
        NULL
    ),
    (
        'f54455e3-6b13-4c77-b74b-8b6c8f8d72e2',
        'Rescue the Princess',
        'A quest to rescue the princess from the evil dragon.',
        'd54455e3-6b13-4c77-b74b-8b6c8f8d72e3',
        '{}',
        TRUE,
        NOW(),
        NOW(),
        NULL
    );

-- Insert data into User Quests
INSERT INTO
    "user_quests" ("user_id", "quest_id")
VALUES
    (
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e5',
        'f54455e3-6b13-4c77-b74b-8b6c8f8d72e1'
    ),
    (
        'b54455e3-6b13-4c77-b74b-8b6c8f8d72e6',
        'f54455e3-6b13-4c77-b74b-8b6c8f8d72e2'
    );