package db

type User struct {
	Id    string
	Name  string
	Email string
}

type Workflow struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BlockType struct {
	id        string
	name      string
	is_preset bool
}

type ValueType string

var Text ValueType = "text"
var Boolean ValueType = "boolean"
var Number ValueType = "number"
var Date ValueType = "date"

type BlockTypeProperty struct {
	id            string
	name          string
	value_type    ValueType
	default_value string
	block_type_id string
	block_type    *BlockType
}

type Block struct {
	id            string
	position      int
	workflow_id   string
	workflow      *Workflow
	block_type_id string
	block_type    *BlockType
}

type BlockProperty struct {
	id       string
	name     string
	value    string
	block_id string
	block    *Block
}
