
CREATE TABLE books (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE category (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

create table category_book(
    id UUID primary key,
    book_id UUID not null references books(id),
    category_id UUID not null references category(id)
);

insert into category_book(id,book_id,category_id)values
('ee5167a6-90dc-11ed-a1eb-0242ac120002','ffbae55c-c85d-4772-ad32-23c82db39409','50d2c016-adc2-4ea1-9457-940adce688f0'),
('024ef18a-90cc-11ed-a1eb-0242ac120002','ffbae55c-c85d-4772-ad32-23c82db39409','4e1f3d27-1267-4e86-ba03-e62c508a7b69'),
('35009d36-90cc-11ed-a1eb-0242ac120002','9187bbf3-b601-44e7-851c-a7f3c7415439','50d2c016-adc2-4ea1-9457-940adce688f0'),
('5db04b14-90cc-11ed-a1eb-0242ac120002','9187bbf3-b601-44e7-851c-a7f3c7415439','f6fba637-09ab-45ff-9b22-ff1aba32666b'),
('cca3249e-90de-11ed-a1eb-0242ac120002','c1503f86-be40-4ea9-bd4a-51675f7d5610','ddd11b05-59a0-447c-b390-c4fda21219dd');


