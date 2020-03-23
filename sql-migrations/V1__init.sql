CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
);

INSERT INTO person (first_name, last_name, email) VALUES ('Jason', 'Moiron', 'jmoiron@jmoiron.net');
INSERT INTO person (first_name, last_name, email) VALUES ('John', 'Doe', 'johndoeDNE@gmail.net');
INSERT INTO person (first_name, last_name, email) VALUES ('Jane', 'Citizen', 'jane.citzen@example.com');
INSERT INTO place (country, city, telcode) VALUES ('United States', 'New York', '1');
INSERT INTO place (country, telcode) VALUES ('Hong Kong', '852');
INSERT INTO place (country, telcode) VALUES ('Singapore', '65');
