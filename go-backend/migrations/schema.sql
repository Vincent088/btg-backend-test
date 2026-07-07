CREATE TABLE nationality (
    nationality_id SERIAL PRIMARY KEY,
    nationality_name VARCHAR(50) NOT NULL,
    nationality_code CHAR(2) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS customer (
    cst_id SERIAL PRIMARY KEY,
    nationality_id INT NOT NULL REFERENCES nationality(nationality_id),
    cst_name CHAR(50) NOT NULL,
    cst_dob DATE NOT NULL,
    cst_phoneNum VARCHAR(20) NOT NULL,
    cst_email VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS family_list (
    fl_id SERIAL PRIMARY KEY,
    cst_id INT NOT NULL REFERENCES customer(cst_id) ON DELETE CASCADE,
    fl_relation VARCHAR(50) NOT NULL,
    fl_name VARCHAR(50) NOT NULL,
    fl_dob VARCHAR(50) NOT NULL
);