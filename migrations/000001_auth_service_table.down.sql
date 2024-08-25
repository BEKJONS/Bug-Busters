-- Удаляем таблицы, если они существуют, в правильном порядке

-- Сначала удаляем таблицы, которые зависят от других
DROP TABLE IF EXISTS Services_Provided;
DROP TABLE IF EXISTS Fines;
DROP TABLE IF EXISTS Services;

-- Затем удаляем таблицу связей пользователей и машин
DROP TABLE IF EXISTS users_cars;

-- Теперь можно безопасно удалить таблицу пользователей
DROP TABLE IF EXISTS users;

-- Удаляем оставшиеся таблицы
DROP TABLE IF EXISTS cars;
DROP TABLE IF EXISTS driver_licenses;
