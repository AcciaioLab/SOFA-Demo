-- SCHEMA: NOT USED
-- Table: trailer_fea
-- Output from FEA
CREATE TABLE IF NOT EXISTS trailer_fea (
    NodeNumber INT NOT NULL,
    Damage REAL NOT NULL,
    UtilizationRatio REAL NOT NULL,
    EqStressRange REAL NOT NULL,
    MinStress REAL NOT NULL,
    MaxStress REAL NOT NULL,
    MaxStressRrange REAL NOT NULL,
    VIN TEXT REFERENCES trailers (VIN)
);

