-- name: GetAnnounces :many
SELECT * FROM announces
ORDER BY time;

-- name: PostAnnounce :one
INSERT INTO announces (
    title, body
) VALUES (
    $1, $2
) RETURNING *;
