CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment,
  name varchar(50) NOT NULL,
  username varchar(50) NOT NULL UNIQUE,
  email varchar(50) NOT NULL UNIQUE,
  password varchar(100) NOT NULL,
  createdAt timestamp DEFAULT current_timestamp(),
  PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int NOT NULL,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;