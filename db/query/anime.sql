
-- name: CheckID :one
SELECT EXISTS(SELECT 1 FROM animes WHERE id=$1) as "exists";

-- name: ListAnimes :many
SELECT * FROM ( SELECT * FROM animes
WHERE name LIKE ('%' + $3::text + '%')  ) AS "animes"
ORDER BY $1::text DESC
OFFSET ($2::int-1)*10 LIMIT 10;

-- name: CreateAnime :one
INSERT INTO animes (
  name, 
  type,
  summary,
  num_of_episodes, 
  other_names,
  status, 
  genre,
  released
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetAnime :one
SELECT * FROM animes
WHERE id = $1 LIMIT 1;

-- name: UpdateAnime :exec
UPDATE animes SET 
  name = $2, 
  type = $3,
  summary = $4,
  num_of_episodes = $5, 
  other_names = $6,
  status = $7, 
  genre = $8,
  released = $9
WHERE id = $1;

-- name: DeleteAnime :exec
DELETE FROM animes
WHERE id = $1;