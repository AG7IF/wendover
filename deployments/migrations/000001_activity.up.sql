CREATE SCHEMA IF NOT EXISTS wendover;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA wendover;

CREATE TABLE wendover.activities (
    id UUID PRIMARY KEY DEFAULT wendover.uuid_generate_v4(),
    key VARCHAR(64) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL DEFAULT '',
    start_datetime timestamp,
    end_datetime timestamp,
    cadet_student_fee INTEGER NOT NULL CHECK (cadet_student_fee >= 0) DEFAULT 0,
    cadet_cadre_fee INTEGER NOT NULL CHECK (cadet_student_fee >= 0) DEFAULT 0,
    senior_student_fee INTEGER NOT NULL CHECK (cadet_student_fee >= 0) DEFAULT 0,
    senior_cadre_fee INTEGER NOT NULL CHECK (cadet_student_fee >= 0) DEFAULT 0
);