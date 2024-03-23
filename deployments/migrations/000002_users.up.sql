CREATE TYPE wendover.role AS ENUM('DIRECTOR', 'ADMIN', 'STAFF');

CREATE TABLE wendover.users (
    id UUID PRIMARY KEY DEFAULT wendover.uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE wendover.user_roles (
    user_id UUID NOT NULL REFERENCES wendover.users(id) ON DELETE CASCADE,
    activity_id UUID NOT NULL REFERENCES wendover.activities(id) ON DELETE CASCADE,
    role wendover.role NOT NULL DEFAULT 'STAFF',
    PRIMARY KEY (user_id, activity_id)
);