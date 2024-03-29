DROP TABLE comment;

CREATE TABLE comments
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT REFERENCES users (id),
    inventory_id BIGINT REFERENCES inventory (id),
    rating       BIGINT    NOT NULL,
    text         TEXT      NOT NULL,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO comments (user_id, inventory_id, rating, text)
VALUES (1, 1, 4, 'Хороший товар, рекомендую!'),
       (2, 2, 5, 'Отличное качество, быстрая доставка!'),
       (3, 3, 3, 'Среднего качества, но за такую цену нормально.'),
       (1, 4, 2, 'Не очень доволен, товар пришел с дефектом.'),
       (2, 5, 5, 'Супер! Буду заказывать еще.'),
       (3, 6, 1, 'Ужасное качество, не рекомендую.'),
       (1, 7, 4, 'Все хорошо, но могло быть и лучше.'),
       (2, 8, 3, 'За такую цену неплохо, но не впечатлило.'),
       (3, 9, 5, 'Покупкой доволен, спасибо!'),
       (1, 10, 4, 'Хороший товар, но доставка могла быть быстрее.'),
       (2, 11, 2, 'Товар пришел с опозданием, качество среднее.'),
       (3, 12, 5, 'Отличный товар, всем советую!'),
       (1, 13, 3, 'Неплохо, но есть недочеты.'),
       (2, 14, 4, 'В целом доволен, но есть мелкие недостатки.'),
       (3, 15, 1, 'Плохой товар, не стоит своих денег.'),
       (1, 16, 5, 'Прекрасное качество, всем рекомендую!'),
       (2, 17, 3, 'Сойдет, но не более.'),
       (3, 18, 4, 'Хороший товар, но могло быть и лучше.'),
       (1, 19, 2, 'Ожидал большего, разочарован.'),
       (2, 20, 5, 'Супер товар, обязательно закажу еще!'),
       (3, 21, 3, 'Среднего качества, но за такую цену нормально.'),
       (1, 22, 4, 'Доволен покупкой, спасибо!'),
       (2, 23, 2, 'Не оправдал ожиданий, к сожалению.'),
       (3, 24, 5, 'Очень хороший товар, рекомендую всем!'),
       (1, 25, 3, 'Нормально, но могло быть лучше.'),
       (2, 26, 4, 'Все хорошо, спасибо!'),
       (3, 27, 1, 'Качество оставляет желать лучшего.'),
       (1, 28, 5, 'Отличный товар, буду заказывать еще.'),
       (2, 29, 3, 'Среднее качество, но подойдет.'),
       (3, 30, 4, 'Хороший товар, но могло быть и лучше.'),
       (1, 31, 2, 'Не очень доволен, товар пришел с дефектом.'),
       (2, 32, 5, 'Супер! Буду заказывать еще.'),
       (3, 33, 1, 'Ужасное качество, не рекомендую.'),
       (1, 34, 4, 'Все хорошо, но могло быть и лучше.'),
       (2, 35, 3, 'За такую цену неплохо, но не впечатлило.'),
       (3, 36, 5, 'Покупкой доволен, спасибо!'),
       (1, 37, 4, 'Хороший товар, но доставка могла быть быстрее.'),
       (2, 38, 2, 'Товар пришел с опозданием, качество среднее.'),
       (3, 39, 5, 'Отличный товар, всем советую!'),
       (1, 40, 3, 'Неплохо, но есть недочеты.'),
       (2, 41, 4, 'В целом доволен, но есть мелкие недостатки.'),
       (3, 42, 1, 'Плохой товар, не стоит своих денег.'),
       (1, 43, 5, 'Прекрасное качество, всем рекомендую!'),
       (2, 44, 3, 'Сойдет, но не более.'),
       (3, 45, 4, 'Хороший товар, но могло быть и лучше.'),
       (1, 46, 2, 'Ожидал большего, разочарован.'),
       (2, 47, 5, 'Супер товар, обязательно закажу еще!'),
       (3, 48, 3, 'Среднего качества, но за такую цену нормально.'),
       (1, 49, 4, 'Доволен покупкой, спасибо!'),
       (2, 50, 2, 'Не оправдал ожиданий, к сожалению.'),
       (3, 51, 5, 'Очень хороший товар, рекомендую всем!'),
       (1, 52, 3, 'Нормально, но могло быть лучше.'),
       (2, 53, 4, 'Все хорошо, спасибо!'),
       (3, 54, 1, 'Качество оставляет желать лучшего.'),
       (1, 55, 5, 'Отличный товар, буду заказывать еще.'),
       (2, 56, 3, 'Среднее качество, но подойдет.');
