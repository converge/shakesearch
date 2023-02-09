BEGIN;

CREATE TABLE IF NOT EXISTS shakesearch (
    id INTEGER NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    chapter VARCHAR,
    created TIMESTAMP
);

-- setweight is a postgres function that allow us to set the weight of a word in a tsvector
-- A is the highest weight, B is the second highest
-- this is used in the search query to give more weight to the title than the chapter, in that queries that matches the
-- title will be ranked higher than those that matches the chapter
UPDATE shakesearch SET title = (setweight(to_tsvector(title), 'A') || setweight(to_tsvector(chapter), 'B'));

END;
