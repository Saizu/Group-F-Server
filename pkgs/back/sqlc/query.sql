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



-- name: GetItems :many
SELECT * FROM items
ORDER BY id ASC;

-- name: PostItem :one
INSERT INTO items (
    name
) VALUES (
    $1
) RETURNING *;



-- name: GetUsersItems :many
SELECT * FROM users_items
ORDER BY usrid ASC, itmid ASC;

-- name: PostItemToUser :one
INSERT INTO users_items ( usrid, itmid, amount )
VALUES ( $1, $2, $3 )
ON CONFLICT ( usrid, itmid )
DO UPDATE SET amount = users_items.amount + EXCLUDED.amount
RETURNING *;

-- name: PostItemToAllUsers :many
INSERT INTO users_items ( usrid, itmid, amount )
SELECT id, $1, $2 FROM users
ON CONFLICT ( usrid, itmid )
DO UPDATE SET amount = users_items.amount + EXCLUDED.amount
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;

-- name: GetItemsByUser :many
SELECT * FROM users_items
WHERE usrid = $1
ORDER BY itmid ASC;

-- name: GetUserIdByName :one
SELECT id FROM users WHERE name = $1;
