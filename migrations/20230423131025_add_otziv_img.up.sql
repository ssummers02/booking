CREATE TABLE images
(
    id   BIGSERIAL PRIMARY KEY,
    inventory_id BIGINT  UNIQUE REFERENCES inventory (id),
    name TEXT NOT NULL,
    data BYTEA
);


CREATE TABLE comment
(
    id      BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users (id),
    inventory_id BIGINT REFERENCES inventory (id),
    text    TEXT NOT NULL

)