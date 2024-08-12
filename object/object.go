package object

import "fmt"

const (
	INTEGER_OBJ = "INTEGER"
	BOOL_OBJ    = "BOOL"
	NULL_OBJ    = "NULL"
)

// generic object type
type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

// integer object
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// bool object
type Bool struct {
	Value bool
}

func (b *Bool) Type() ObjectType { return BOOL_OBJ }

func (b *Bool) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// null object
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) Inspect() string { return "null" }
