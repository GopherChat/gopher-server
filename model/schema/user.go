package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			MaxLen(50).
			Optional(),
		field.String("first_name").
			MaxLen(50).
			Optional(),
		field.String("last_name").
			MaxLen(50).
			Optional(),
		field.String("email").
			MaxLen(100).
			Unique(),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
