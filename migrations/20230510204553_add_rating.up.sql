ALTER TABLE resorts
    ADD COLUMN avg_rating NUMERIC(5, 2) NOT NULL DEFAULT 0;

BEGIN;
CREATE TEMPORARY TABLE temp_avg_ratings
AS
SELECT r.id                     AS resort_id,
       SUM(c.rating) / COUNT(*) AS avg_rating -- Вычислите средний рейтинг
FROM comments c
         JOIN
     inventory i ON c.inventory_id = i.id
         JOIN
     resorts r ON i.resort_id = r.id
GROUP BY r.id;

-- Обновите таблицу resorts, установив avg_rating на основе значений из временной таблицы
UPDATE
    resorts
SET avg_rating = temp_avg_ratings.avg_rating
FROM temp_avg_ratings
WHERE resorts.id = temp_avg_ratings.resort_id;

-- Удалите временную таблицу
DROP TABLE temp_avg_ratings;

COMMIT;




-- Создаем функцию для обновления среднего рейтинга
CREATE OR REPLACE FUNCTION update_avg_rating() RETURNS TRIGGER AS $$
DECLARE
    resort_id_to_update BIGINT;
BEGIN
    -- Определите resort_id, который нужно обновить
    SELECT resort_id
    INTO resort_id_to_update
    FROM inventory
    WHERE inventory.id = (
        CASE
            WHEN TG_OP = 'DELETE' THEN OLD.inventory_id
            ELSE NEW.inventory_id
            END
        );

    -- Обновите только тот рейтинг, который изменили
    UPDATE resorts
    SET avg_rating = (
        SELECT SUM(comments.rating) / COUNT(*)
        FROM comments
                 JOIN inventory ON comments.inventory_id = inventory.id
        WHERE inventory.resort_id = resorts.id
    )
    WHERE id = resort_id_to_update;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Создаем триггер для обновления среднего рейтинга при добавлении, удалении или изменении комментариев
CREATE TRIGGER update_avg_rating_trigger
    AFTER INSERT OR UPDATE OR DELETE ON comments
    FOR EACH ROW
EXECUTE FUNCTION update_avg_rating();

