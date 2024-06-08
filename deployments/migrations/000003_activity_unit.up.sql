CREATE TABLE activity_units (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    activity_id UUID NOT NULL REFERENCES activities(id) ON DELETE CASCADE,
    commander_id UUID,
    superior_unit_id UUID REFERENCES activity_units(id) ON DELETE CASCADE,
    unit_name VARCHAR(255) NOT NULL
);

CREATE FUNCTION generate_activity_unit_id() RETURNS trigger AS $generate_activity_unit_id$
    BEGIN
        NEW.id := uuid_generate_v4();
        return NEW;
    END
$generate_activity_unit_id$ LANGUAGE plpgsql;

CREATE TRIGGER set_activity_units_id BEFORE INSERT
ON activity_units FOR EACH ROW
WHEN (NEW.id = uuid_nil())
EXECUTE FUNCTION generate_activity_unit_id();
