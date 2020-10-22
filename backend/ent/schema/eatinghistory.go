package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Eatinghistory holds the schema definition for the Eatinghistory entity.
type Eatinghistory struct {
	ent.Schema
}

// Fields of the Eatinghistory.
func (Eatinghistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("added_time"),
	}
}

// Edges of the Eatinghistory.
func (Eatinghistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("foodmenu", Foodmenu.Type).Ref("eatinghistory").Unique(),
		edge.From("mealplan", Mealplan.Type).Ref("eatinghistory").Unique(),
		edge.From("taste", Taste.Type).Ref("eatinghistory").Unique(),
		edge.From("user", User.Type).Ref("eatinghistory").Unique(),
	}
}
