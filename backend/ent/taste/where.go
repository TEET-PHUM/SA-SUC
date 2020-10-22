// Code generated by entc, DO NOT EDIT.

package taste

import (
	"github.com/Teeth/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TasteName applies equality check predicate on the "taste_name" field. It's identical to TasteNameEQ.
func TasteName(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTasteName), v))
	})
}

// TasteNameEQ applies the EQ predicate on the "taste_name" field.
func TasteNameEQ(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTasteName), v))
	})
}

// TasteNameNEQ applies the NEQ predicate on the "taste_name" field.
func TasteNameNEQ(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTasteName), v))
	})
}

// TasteNameIn applies the In predicate on the "taste_name" field.
func TasteNameIn(vs ...string) predicate.Taste {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Taste(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTasteName), v...))
	})
}

// TasteNameNotIn applies the NotIn predicate on the "taste_name" field.
func TasteNameNotIn(vs ...string) predicate.Taste {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Taste(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTasteName), v...))
	})
}

// TasteNameGT applies the GT predicate on the "taste_name" field.
func TasteNameGT(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTasteName), v))
	})
}

// TasteNameGTE applies the GTE predicate on the "taste_name" field.
func TasteNameGTE(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTasteName), v))
	})
}

// TasteNameLT applies the LT predicate on the "taste_name" field.
func TasteNameLT(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTasteName), v))
	})
}

// TasteNameLTE applies the LTE predicate on the "taste_name" field.
func TasteNameLTE(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTasteName), v))
	})
}

// TasteNameContains applies the Contains predicate on the "taste_name" field.
func TasteNameContains(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTasteName), v))
	})
}

// TasteNameHasPrefix applies the HasPrefix predicate on the "taste_name" field.
func TasteNameHasPrefix(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTasteName), v))
	})
}

// TasteNameHasSuffix applies the HasSuffix predicate on the "taste_name" field.
func TasteNameHasSuffix(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTasteName), v))
	})
}

// TasteNameEqualFold applies the EqualFold predicate on the "taste_name" field.
func TasteNameEqualFold(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTasteName), v))
	})
}

// TasteNameContainsFold applies the ContainsFold predicate on the "taste_name" field.
func TasteNameContainsFold(v string) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTasteName), v))
	})
}

// HasEatinghistory applies the HasEdge predicate on the "eatinghistory" edge.
func HasEatinghistory() predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EatinghistoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EatinghistoryTable, EatinghistoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEatinghistoryWith applies the HasEdge predicate on the "eatinghistory" edge with a given conditions (other predicates).
func HasEatinghistoryWith(preds ...predicate.Eatinghistory) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EatinghistoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EatinghistoryTable, EatinghistoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Taste) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Taste) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Taste) predicate.Taste {
	return predicate.Taste(func(s *sql.Selector) {
		p(s.Not())
	})
}