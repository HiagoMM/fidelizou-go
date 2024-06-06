CREATE TYPE user_role AS ENUM ('PROVIDER', 'CLIENT');

CREATE TYPE gender AS ENUM ('MALE', 'FEMALE', 'OTHER');

CREATE TYPE auth_provider AS ENUM ('GOOLE');

CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "email" TEXT NOT NULL UNIQUE,
    "name" TEXT NOT NULL,
    "photo" TEXT NOT NULL,
    "auth_provider" auth_provider NOT NULL,
    "credits" INT DEFAULT 0,
    "birth_date" DATE,
    "gender" gender,
    "user_role" user_role NOT NULL DEFAULT 'CLIENT',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "program" (
    "id" SERIAL PRIMARY KEY,
    "title" TEXT,
    "description" TEXT,
    "fidelity_card_front_image" TEXT,
    "fidelity_card_back_image" TEXT,
    "fidelity_card_point_image" TEXT,
    "fidelity_card_max_points" INT,
    "final_date" TIMESTAMP,
    "active_links" TEXT ,
    "owner_id" INT REFERENCES "user",
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "fidelity_card" (
    "id" SERIAL PRIMARY KEY,
    "points" INT,
    "dates_of_uses" TIMESTAMP ,
    "user_id" INT REFERENCES "user",
    "program_id" INT REFERENCES "program",
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
