package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Foodmenu holds the schema definition for the Foodmenu entity.
type Foodmenu struct {
	ent.Schema
}

// Fields of the Foodmenu.
func (Foodmenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("foodmenu_name").NotEmpty(),
		field.String("foodmenu_type").NotEmpty(),
	}
}

// Edges of the Foodmenu.
func (Foodmenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("foodmenu").Unique(),
		edge.To("eatinghistory", Eatinghistory.Type).StorageKey(edge.Column("foodmenu_id")),
	}
}
