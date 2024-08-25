DROP TABLE IF EXISTS Services_Provided; -- Сначала удаляем таблицу, зависящую от users и cars

DROP TABLE IF EXISTS Fines; -- Затем удаляем таблицу, зависящую от users и cars

DROP TABLE IF EXISTS Services; -- Затем таблицу services, которая не имеет внешних ключей

DROP TABLE IF EXISTS users; -- Затем таблицу users, которая ссылается на driver_licenses

DROP TABLE IF EXISTS cars; -- Затем таблицу cars

DROP TABLE IF EXISTS driver_licenses; -- Наконец, таблицу driver_licenses, от которой зависят другие таблицы
