DROP TABLE images;

ALTER TABLE inventory
    ADD COLUMN name TEXT,
    ADD COLUMN data BYTEA