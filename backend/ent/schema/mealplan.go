package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Mealplan holds the schema definition for the Mealplan entity.
type Mealplan struct {
	ent.Schema
}

// Fields of the Mealplan.
func (Mealplan) Fields() []ent.Field {
	return []ent.Field{
		field.String("mealplan_name").NotEmpty(),
	}
}

// Edges of the Mealplan.
func (Mealplan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("mealplan").Unique(),
		edge.To("eatinghistory", Eatinghistory.Type).StorageKey(edge.Column("mealplan_id")),
	}
}
