CREATE TYPE role AS ENUM('DIRECTOR', 'ADMIN', 'STAFF');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    activity_id UUID NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    role role NOT NULL DEFAULT 'STAFF',
    PRIMARY KEY (user_id, activity_id)
);