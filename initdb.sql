create table gp
(
    id          bigserial not null
        constraint gp_pkey
            primary key,
    text        varchar,
    size        varchar,
    date        varchar,
    measurement varchar,
    model       varchar,
    picture     varchar,
    foto        varchar,
    hash        varchar
);
