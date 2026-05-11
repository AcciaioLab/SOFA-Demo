-- Testing data input
-- Trailer information
COPY trailers 
FROM '/tmp/data/Trailer_Parameters.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);

-- FEA data
-- Per trailer
COPY trailer_fea 
FROM '/tmp/data/FMC4_FEA.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);

COPY trailer_fea 
FROM '/tmp/data/FMC5_FEA.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);

COPY trailer_fea 
FROM '/tmp/data/FMC6_FEA.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);

COPY trailer_fea 
FROM '/tmp/data/FMC7_FEA.csv'
WITH (
    FORMAT CSV,
    HEADER TRUE,
    DELIMITER ','
);