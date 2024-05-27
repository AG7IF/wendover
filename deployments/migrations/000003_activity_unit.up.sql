CREATE TABLE wendover.activity_units (
    id UUID PRIMARY KEY DEFAULT wendover.uuid_generate_v4(),
    activity_id UUID NOT NULL REFERENCES wendover.activities(id) ON DELETE CASCADE,
    commander_id UUID,
    superior_unit_id UUID REFERENCES wendover.activity_units(id) ON DELETE CASCADE,
    unit_name VARCHAR(255) NOT NULL
);
