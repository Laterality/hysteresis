CREATE TABLE BAR
(
    id      bigint auto_increment not null,
    username varchar(22)          not null,
    foo_id bigint not null,
    primary key (id),
    foreign key(foo_id) REFERENCES foo(id)
);
