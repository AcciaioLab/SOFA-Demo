-- Testing data input
-- Trailer information
COPY trailers 
FROM '/tmp/data/Trailer_Parameters.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);