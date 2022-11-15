CREATE DATABASE "PathManager"
    WITH
    OWNER = pathmanager
    ENCODING = 'UTF8'
    LC_COLLATE = 'Italian_Italy.1252'
    LC_CTYPE = 'Italian_Italy.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public."Path"
(
    name text COLLATE pg_catalog."default" NOT NULL,
    altitude integer NOT NULL,
    length integer NOT NULL,
    level text COLLATE pg_catalog."default" NOT NULL,
    cyclable boolean NOT NULL,
    family boolean NOT NULL,
    historical boolean NOT NULL,
    region text COLLATE pg_catalog."default" NOT NULL,
    province text COLLATE pg_catalog."default" NOT NULL,
    city text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Path_pkey" PRIMARY KEY (name)
);