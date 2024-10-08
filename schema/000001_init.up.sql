CREATE TABLE users (
    id              serial       not null unique,
    username        varchar(255)    not null    unique,
    name            varchar(255)   not null,
    password_hash   varchar(255)    not null
);

/* username" are of incompatible types: integer and character varying */

CREATE TABLE posts
  (
    id serial not null unique,
    title text not null,
    body text not null,
    username varchar(255) references users (username) on delete cascade not null
  );


CREATE TABLE comments
  (
  id serial not null unique,
  body text not null,
  username varchar(255) references users (username) on delete cascade not null,
  post_id int references posts (id) on delete cascade not null
  );
