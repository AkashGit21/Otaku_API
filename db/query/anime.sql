-- name: ListAnimes :many
SELECT * FROM animes 
ORDER BY NAME 
LIMIT 10 OFFSET ($1-1)*10;


-- name: CreateAnime :one
INSERT INTO animes (
  name, 
  description,
  num_of_episodes, 
  "cast",
  status, 
  genre
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetAnime :one
SELECT * FROM animes
WHERE id = $1 LIMIT 1;

-- name: UpdateAnime :exec
UPDATE animes SET 
  name = $2,
  description = $3,
  num_of_episodes = $4,
  "cast" = $5,
  status = $6,
  genre = $7
WHERE id = $1;

-- name: DeleteAnime :exec
DELETE FROM animes
WHERE id = $1;