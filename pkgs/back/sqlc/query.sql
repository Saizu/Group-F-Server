-- name: GetPlayer :one
SELECT * FROM players
WHERE id = $1 LIMIT 1;

-- name: ListPlayers :many
SELECT * FROM players
ORDER BY id;
