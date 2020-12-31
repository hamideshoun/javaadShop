-- Database: books_database

-- DROP DATABASE books_database;

CREATE DATABASE books_database
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       CONNECTION LIMIT = -1;

-- Table: books

-- DROP TABLE books;

CREATE TABLE Customers
(
 ID serial NOT NULL,
 Name character varying NOT NULL,
 Address character varying,
 RegisterDate date,
 Tel integer
)
WITH (
 OIDS=FALSE
);
ALTER TABLE Customers
 OWNER TO hamid;
