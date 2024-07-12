create table movies(
    id serial primary key,
    title varchar,
    year integer
);

INSERT INTO movies(title, year) VALUES
('The Shawshank Redemption', 1994),
('The Godfather', 1972),
('The Godfather: Part II', 1974),
('The Dark Knight', 2008),