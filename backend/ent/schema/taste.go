package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Taste holds the schema definition for the Taste entity.
type Taste struct {
	ent.Schema
}

// Fields of the Taste.
func (Taste) Fields() []ent.Field {
	return []ent.Field{
		field.String("taste_name").NotEmpty(),
	}
}

// Edges of the Taste.
func (Taste) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("eatinghistory", Eatinghistory.Type).StorageKey(edge.Column("taste_id")),
	}
}
