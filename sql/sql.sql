CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment,
  name varchar(50) not null,
  username varchar(50) not null UNIQUE,
  email varchar(50) not null UNIQUE,
  password varchar(100) not null,
  createdAt timestamp DEFAULT current_timestamp(),
  PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE followers(
  user_id int not null,
  follower_id int not null,

  PRIMARY KEY (user_id, follower_id),

  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE,

  FOREIGN KEY (follower_id)
  REFERENCES users(id)
  ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE posts(
  id int auto_increment,
  title varchar(50) not null,
  content varchar(300) not null,
  author_id int not null,
  likes int default 0,
  createdAt timestamp default current_timestamp(),

  PRIMARY KEY (id),

  FOREIGN KEY (author_id)
  REFERENCES users(id)
  ON DELETE CASCADE
) ENGINE=INNODB;