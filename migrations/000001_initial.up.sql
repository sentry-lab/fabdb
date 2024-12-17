CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    email TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS workflows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS block_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    is_preset BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS block_type_properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    value_type TEXT CHECK (value_type IN ('str', 'bool', 'num')) NOT NULL,
    default_value TEXT,
    block_type_id UUID NOT NULL,

    CONSTRAINT fk_type_id
        FOREIGN KEY(block_type_id)
        REFERENCES block_types(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS blocks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    position INT NOT NULL,
    workflow_id UUID NOT NULL,
    block_type_id UUID NOT NULL,

    CONSTRAINT fk_workflow_id
        FOREIGN KEY(workflow_id)
        REFERENCES workflows(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_type_id
        FOREIGN KEY(block_type_id)
        REFERENCES block_types(id)
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block_properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    value TEXT,
    block_id UUID NOT NULL,

    CONSTRAINT fk_block_id
        FOREIGN KEY(block_id)
        REFERENCES blocks(id)
        ON DELETE CASCADE
);
