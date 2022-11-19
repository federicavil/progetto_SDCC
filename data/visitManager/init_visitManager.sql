CREATE USER IF NOT EXISTS visitmanager@localhost IDENTIFIED BY 'password';

CREATE DATABASE IF NOT EXISTS VisitManager;

USE VisitManager;

GRANT ALL PRIVILEGES ON *.* TO visitmanager@localhost;

CREATE TABLE IF NOT EXISTS Visit
(
    ID integer AUTO_INCREMENT PRIMARY KEY NOT NULL,
    ID_Path varchar(255) NOT NULL,
    Timestamp DOUBLE NOT NULL,
    Creator varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS User_to_Visit
(
    ID_Visit integer NOT NULL,
    ID_User varchar(255) NOT NULL,
    Accepted boolean,
    CONSTRAINT User_to_Visit_pkey PRIMARY KEY (ID_User, ID_Visit),
    FOREIGN KEY (ID_Visit) REFERENCES Visit(ID)
);