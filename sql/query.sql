-- name: GetUserById :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1;

-- name: UpdateUserRole :exec
UPDATE "user" SET user_role = $1 WHERE id = $2;

-- name: GetMyFidelityCards :many
SELECT
    fidelity_card.id,
    fidelity_card.points,
    fidelity_card.dates_of_uses,
    fidelity_card.program_id,
    fidelity_card.created_at,
    program.title,
    program.description,
    program.fidelity_card_front_image,
    program.fidelity_card_back_image,
    program.fidelity_card_point_image,
    program.fidelity_card_max_points,
    program.final_date
FROM "fidelity_card"
INNER JOIN "program" ON fidelity_card.program_id = program.id
WHERE fidelity_card.user_id = $1;

-- name: GetMyPrograms :many
SELECT * FROM "program" WHERE owner_id = $1;