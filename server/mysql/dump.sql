CREATE TABLE currency_conversions (
    id INT NOT NULL AUTO_INCREMENT,
    bid DECIMAL(10,2) NOT NULL,
    quote_time VARCHAR(30) NOT NULL,
    PRIMARY KEY (id)
);