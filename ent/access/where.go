// Code generated by ent, DO NOT EDIT.

package access

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/greboid/puzzad/ent/predicate"
)

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// AdventureID applies equality check predicate on the "adventure_id" field. It's identical to AdventureIDEQ.
func AdventureID(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdventureID), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCode), v))
	})
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCode), v))
	})
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCode), v))
	})
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCode), v))
	})
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCode), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// AdventureIDEQ applies the EQ predicate on the "adventure_id" field.
func AdventureIDEQ(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdventureID), v))
	})
}

// AdventureIDNEQ applies the NEQ predicate on the "adventure_id" field.
func AdventureIDNEQ(v int) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdventureID), v))
	})
}

// AdventureIDIn applies the In predicate on the "adventure_id" field.
func AdventureIDIn(vs ...int) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAdventureID), v...))
	})
}

// AdventureIDNotIn applies the NotIn predicate on the "adventure_id" field.
func AdventureIDNotIn(vs ...int) predicate.Access {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Access(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAdventureID), v...))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAdventures applies the HasEdge predicate on the "adventures" edge.
func HasAdventures() predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, AdventuresColumn),
			sqlgraph.To(AdventuresInverseTable, AdventureFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AdventuresTable, AdventuresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdventuresWith applies the HasEdge predicate on the "adventures" edge with a given conditions (other predicates).
func HasAdventuresWith(preds ...predicate.Adventure) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, AdventuresColumn),
			sqlgraph.To(AdventuresInverseTable, AdventureFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AdventuresTable, AdventuresColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
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
func Not(p predicate.Access) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		p(s.Not())
	})
}
