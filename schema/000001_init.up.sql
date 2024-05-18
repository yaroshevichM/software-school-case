CREATE TABLE IF NOT EXISTS "Subscriptions" 
(
    id serial not null unique,
    email varchar(255) not null
);
