CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment PRIMARY KEY,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    senha varchar(100) NOT NULL,
    criadoEm timestamp DEFAULT CURRENT_TIMESTAMP()
)   ENGINE=INNODB;