package schema

import (
	"entgo.io/ent"
)

// AuthRole holds the schema definition for the AuthRole entity.
type AuthRole struct {
	ent.Schema
}

// Fields of the AuthRole.
func (AuthRole) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the AuthRole.
func (AuthRole) Edges() []ent.Edge {
	return nil
}
func (AuthRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
