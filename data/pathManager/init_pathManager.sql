CREATE USER pathmanager@localhost IDENTIFIED BY 'password';

CREATE DATABASE IF NOT EXISTS PathManager;

USE PathManager;

GRANT ALL PRIVILEGES ON *.* TO pathmanager;

CREATE TABLE IF NOT EXISTS Path
(
    name varchar(255) NOT NULL,
    altitude integer NOT NULL,
    length integer NOT NULL,
    level varchar(255) NOT NULL,
    cyclable boolean NOT NULL,
    family boolean NOT NULL,
    historical boolean NOT NULL,
    region varchar(255) NOT NULL,
    province varchar(255) NOT NULL,
    city varchar(255) NOT NULL,
    PRIMARY KEY (name)
);