CREATE TABLE roles
(
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT      NOT NULL
);

INSERT INTO roles (id, name)
VALUES (1, 'admin'),
       (2, 'user'),
       (3, 'owner');


CREATE TABLE users
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    first_name  TEXT        NOT NULL,
    surname     TEXT        NOT NULL,
    middle_name TEXT        NOT NULL,
    email       TEXT        NOT NULL,
    password    TEXT        NOT NULL,
    phone       TEXT        NOT NULL,
    role_id     BIGINT      NOT NULL REFERENCES roles (id),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ

);

INSERT INTO users (first_name, surname, middle_name, email, password, phone, role_id)
VALUES ('admin', 'admin', 'admin', 'admin@test.ru', '$2a$14$KbQYSlu2/0DRUoCCRmrEYe/d2ySZNQ8ofeBFGcLYKQR/yXP.W6TP6',
        '123', 1),
       ('user', 'user', 'user', 'user@test.ru', '$2a$14$KbQYSlu2/0DRUoCCRmrEYe/d2ySZNQ8ofeBFGcLYKQR/yXP.W6TP6', '123',
        2),
       ('owner', 'owner', 'owner', 'owner@test.ru', '$2a$14$KbQYSlu2/0DRUoCCRmrEYe/d2ySZNQ8ofeBFGcLYKQR/yXP.W6TP6',
        '123', 3);


CREATE TABLE cities
(
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT      NOT NULL
);

INSERT INTO cities (name)
VALUES ('Москва'),
       ('Санкт-Петербург'),
       ('Казань'),
       ('Новосибирск'),
       ('Екатеринбург'),
       ('Нижний Новгород'),
       ('Красноярск'),
       ('Самара'),
       ('Омск'),
       ('Челябинск'),
       ('Ростов-на-Дону'),
       ('Уфа'),
       ('Краснодар'),
       ('Пермь'),
       ('Воронеж'),
       ('Волгоград'),
       ('Саратов'),
       ('Киров'),
       ('Тюмень'),
       ('Тольятти'),
       ('Ижевск'),
       ('Барнаул'),
       ('Ульяновск'),
       ('Иркутск'),
       ('Хабаровск'),
       ('Ярославль'),
       ('Владивосток'),
       ('Махачкала'),
       ('Томск'),
       ('Оренбург'),
       ('Кемерово'),
       ('Новокузнецк'),
       ('Рязань'),
       ('Астрахань'),
       ('Пенза'),
       ('Липецк'),
       ('Калининград'),
       ('Тула'),
       ('Курск'),
       ('Брянск'),
       ('Сочи'),
       ('Чебоксары'),
       ('Курган'),
       ('Ставрополь'),
       ('Тверь'),
       ('Иваново'),
       ('Белгород'),
       ('Сургут'),
       ('Симферополь'),
       ('Саранск'),
       ('Петрозаводск'),
       ('Владикавказ'),
       ('Магнитогорск'),
       ('Набережные Челны'),
       ('Архангельск'),
       ('Нижний Тагил'),
       ('Балашиха'),
       ('Владимир'),
       ('Севастополь'),
       ('Смоленск'),
       ('Улан-Удэ');

CREATE TABLE resorts
(
    id          BIGSERIAL   NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    owner_id    BIGINT      NOT NULL REFERENCES users (id),
    city_id     BIGINT      NOT NULL REFERENCES cities (id),
    description TEXT        NOT NULL,
    address     TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
INSERT INTO resorts (name, city_id, description, address, owner_id)
VALUES ('Курорт в Москве 1', 1, 'Описание курорта в Москве 1', 'Адрес курорта в Москве', 3),
       ('Курорт в Москве 2', 1, 'Описание курорта в Москве 2', 'Адрес курорта в Москве', 3),
       ('Курорт в Москве 3', 1, 'Описание курорта в Москве 3', 'Адрес курорта в Москве', 3),
       ('Курорт в Москве 4', 1, 'Описание курорта в Москве 4', 'Адрес курорта в Москве', 3),
       ('Курорт в Санкт-Петербурге', 2, 'Описание курорта в Санкт-Петербурге', 'Адрес курорта в Санкт-Петербурге', 3),
       ('Курорт в Казани 1', 3, 'Описание курорта в Казани', 'Адрес курорта в Казани', 3),
       ('Курорт в Казани 2', 3, 'Описание курорта в Казани 2', 'Адрес курорта в Казани', 3);
CREATE TABLE inventory_type
(
    id   BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT      NOT NULL
);
INSERT INTO inventory_type (name)
VALUES ('Сноуборд'),
       ('Коньки'),
       ('Лыжи'),
       ('Пропуск'),
       ('Экипировка');

CREATE TABLE inventory
(
    id              BIGSERIAL   NOT NULL PRIMARY KEY,
    type_id         BIGINT      NOT NULL,
    resort_id       BIGINT      NOT NULL REFERENCES resorts (id),
    price           BIGINT      NOT NULL DEFAULT 0,
    photo           TEXT        NOT NULL DEFAULT '',
    available_start TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    available_end   TIMESTAMPTZ NULL     DEFAULT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
INSERT INTO inventory (type_id, resort_id, price)
VALUES (1, 1, 1000),
       (2, 1, 500),
       (3, 1, 100),
       (4, 1, 100),
       (5, 1, 100),
       (1, 2, 1000),
       (2, 2, 500),
       (3, 2, 100),
       (2, 1, 500),
       (3, 1, 100),
       (4, 1, 100),
       (5, 1, 100),
       (1, 2, 1000),
       (2, 2, 500),
       (3, 2, 100),
       (4, 2, 100),
       (5, 2, 100),
       (1, 3, 1000),
       (2, 3, 500),
       (3, 3, 100),
       (4, 3, 100),
       (5, 3, 100),
       (1, 4, 1000),
       (2, 4, 500),
       (3, 4, 100),
       (4, 4, 100),
       (5, 4, 100),
       (1, 5, 1000),
       (2, 5, 500),
       (3, 5, 100),
       (4, 5, 100),
       (5, 5, 100),
       (1, 6, 1000),
       (2, 6, 500),
       (3, 6, 100),
       (4, 6, 100),
       (5, 6, 100),
       (1, 4, 1000),
       (2, 4, 500),
       (3, 4, 100),
       (4, 4, 100),
       (5, 4, 100),
       (1, 5, 1000),
       (2, 5, 500),
       (3, 5, 100),
       (4, 5, 100),
       (5, 5, 100),
       (1, 6, 1000),
       (2, 6, 500),
       (3, 6, 100),
       (4, 6, 100),
       (5, 6, 100),
       (1, 7, 1000),
       (2, 7, 500),
       (3, 7, 100),
       (1, 6, 1000),
       (2, 6, 500),
       (3, 6, 100),
       (4, 6, 100),
       (5, 6, 100),
       (1, 7, 1000),
       (2, 7, 500),
       (3, 7, 100),
       (4, 7, 100),
       (5, 7, 100),

       (5, 4, 100),
       (1, 5, 1000),
       (2, 5, 500),
       (3, 5, 100),
       (4, 5, 100),
       (5, 5, 100),
       (1, 6, 1000),
       (2, 6, 500),
       (3, 6, 100),
       (4, 6, 100),
       (5, 6, 100),
       (1, 7, 1000),
       (2, 7, 500),
       (3, 7, 100),
       (1, 6, 1000),
       (2, 6, 500),
       (3, 6, 100),
       (4, 6, 100),
       (5, 6, 100);


CREATE TABLE bookings
(
    id           BIGSERIAL   NOT NULL PRIMARY KEY,
    user_id      BIGINT      NOT NULL REFERENCES users (id),
    inventory_id BIGINT      NOT NULL REFERENCES inventory (id),
    start_date   DATE        NOT NULL,
    end_date     DATE        NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);