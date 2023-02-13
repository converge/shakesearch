BEGIN;

CREATE INDEX chapter_idx ON shakesearch USING GIN (to_tsvector('english', title || ' ' || chapter));

END;
