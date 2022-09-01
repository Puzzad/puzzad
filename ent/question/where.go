// Code generated by ent, DO NOT EDIT.

package question

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/greboid/puzzad/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Answer applies equality check predicate on the "answer" field. It's identical to AnswerEQ.
func Answer(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswer), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Question {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Question {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AnswerEQ applies the EQ predicate on the "answer" field.
func AnswerEQ(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnswer), v))
	})
}

// AnswerNEQ applies the NEQ predicate on the "answer" field.
func AnswerNEQ(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnswer), v))
	})
}

// AnswerIn applies the In predicate on the "answer" field.
func AnswerIn(vs ...string) predicate.Question {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAnswer), v...))
	})
}

// AnswerNotIn applies the NotIn predicate on the "answer" field.
func AnswerNotIn(vs ...string) predicate.Question {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAnswer), v...))
	})
}

// AnswerGT applies the GT predicate on the "answer" field.
func AnswerGT(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnswer), v))
	})
}

// AnswerGTE applies the GTE predicate on the "answer" field.
func AnswerGTE(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnswer), v))
	})
}

// AnswerLT applies the LT predicate on the "answer" field.
func AnswerLT(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnswer), v))
	})
}

// AnswerLTE applies the LTE predicate on the "answer" field.
func AnswerLTE(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnswer), v))
	})
}

// AnswerContains applies the Contains predicate on the "answer" field.
func AnswerContains(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAnswer), v))
	})
}

// AnswerHasPrefix applies the HasPrefix predicate on the "answer" field.
func AnswerHasPrefix(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAnswer), v))
	})
}

// AnswerHasSuffix applies the HasSuffix predicate on the "answer" field.
func AnswerHasSuffix(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAnswer), v))
	})
}

// AnswerEqualFold applies the EqualFold predicate on the "answer" field.
func AnswerEqualFold(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAnswer), v))
	})
}

// AnswerContainsFold applies the ContainsFold predicate on the "answer" field.
func AnswerContainsFold(v string) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAnswer), v))
	})
}

// HasAdventure applies the HasEdge predicate on the "adventure" edge.
func HasAdventure() predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdventureTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdventureTable, AdventureColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdventureWith applies the HasEdge predicate on the "adventure" edge with a given conditions (other predicates).
func HasAdventureWith(preds ...predicate.Adventure) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdventureInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AdventureTable, AdventureColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Question) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
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
func Not(p predicate.Question) predicate.Question {
	return predicate.Question(func(s *sql.Selector) {
		p(s.Not())
	})
}
