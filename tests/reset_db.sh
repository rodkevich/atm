#!/bin/sh
echo "Reseting database..."
psql "${DATABASE_URL}" <<EOF
DROP TABLE IF EXISTS persons;
CREATE TABLE persons (id SERIAL NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, name VARCHAR, privileged BOOL NOT NULL DEFAULT false, CONSTRAINT id_pk PRIMARY KEY (id));
INSERT INTO persons (description) VALUES ('John Constantine');
EOF
