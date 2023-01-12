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

CREATE TABLE book_category (
    category_id UUID references category(id), 
    books_id UUID references books(id)
);


SELECT
    b.id,
    b.name,
    b.price,
    b.description,
    b.created_at,
    b.updated_at,
    (
        SELECT
            ARRAY_AGG(category_id)
        FROM book_category AS bc 
        WHERE bc.books_id = '0a32eb14-a3ff-4a02-8445-f0f85bdc8eeb'
    ) AS category_ids
FROM
    books AS b
WHERE b.id = '0a32eb14-a3ff-4a02-8445-f0f85bdc8eeb'
