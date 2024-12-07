-- name: GetAnnounces :many
SELECT * FROM announces
ORDER BY time DESC;

-- name: PostAnnounce :one
INSERT INTO announces (
    title, body
) VALUES (
    $1, $2
) RETURNING *;



-- name: GetUsers :many
SELECT * FROM users
ORDER BY id ASC;

-- name: PostUser :one
INSERT INTO users (
    name
) VALUES (
    $1
) RETURNING *;

-- name: BanOrUnbanUser :one
UPDATE users SET banned = $2
WHERE id = $1
RETURNING *;
