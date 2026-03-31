-- name: GetFeedByID :one
SELECT id, name, desc FROM feeds
WHERE id = ?
LIMIT 1;

-- name: ListAllFeeds :many
SELECT id, name, desc FROM feeds
ORDER BY name;

-- name: UpdateFeed :exec
UPDATE feeds
SET name = ?,
    desc = ?
WHERE id = ?;

-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = ?;


-- Sightings CRUD

-- name: CreateSighting :one
INSERT INTO sightings (
  created_at, title, summary, lat, long
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetSighting :one
SELECT id, feed_id, created_at, title, summary, lat, long FROM sightings
WHERE id = ?
LIMIT 1;

-- name: ListSightings :many
SELECT id, feed_id, created_at, title, summary, lat, long FROM sightings
WHERE
  lat >= ?
  AND lat <= ?
  AND long >= ?
  AND lat <= ?
  AND created_at >= ?
  AND created_at <= ?
  AND feed_id = ?
ORDER BY created_at DESC
LIMIT ?;


-- name: UpdateSighting :exec
UPDATE sightings
SET created_at = ?,
    title = ?,
    summary = ?,
    lat = ?,
    long = ?
WHERE id = ?;

-- name: DeleteSighting :exec
DELETE FROM sightings
WHERE id = ?;