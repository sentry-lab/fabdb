package db

import "github.com/google/uuid"

type User struct {
	id    uuid.UUID
	name  string
	email string
}

type Workflow struct {
	id   uuid.UUID
	name string
}

type BlockType struct {
	id        uuid.UUID
	name      string
	is_preset bool
}

type BlockTypeProperty struct {
	id            uuid.UUID
	name          string
	value_type    string
	default_value string
	block_type_id uuid.UUID
	block_type    *BlockType
}

type Block struct {
	id            uuid.UUID
	name          string
	position      int
	workflow_id   uuid.UUID
	workflow      *Workflow
	block_type_id uuid.UUID
	block_type    *BlockType
}

type BlockProperty struct {
	id       uuid.UUID
	name     string
	value    string
	block_id uuid.UUID
	block    *Block
}
