
-- BlackBox  новый сервис настроек.


-- Создание таблиц и базы данных -- 
create database blackbox with encoding 'UTF-8';
\c blackbox;
\i /docker-entrypoint-initdb.d/db_schema.sql;