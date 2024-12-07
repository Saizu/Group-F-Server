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



-- name: GetInquiries :many
SELECT * FROM inquiries
ORDER BY time DESC;

-- name: PostInquiry :one
INSERT INTO inquiries (
    usrid, title, body
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ReplyInquiry :one
UPDATE inquiries SET reply = $2
WHERE id = $1
RETURNING *;
