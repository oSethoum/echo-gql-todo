package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
	"todo/auth"
	"todo/ent/privacy"
)

type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").Default(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput)),
		field.String("text").Annotations(entgql.OrderField("TEXT")),
		field.Bool("done").Default(false).Annotations(entgql.OrderField("DONE")),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Unique().Ref("todos"),
	}
}

// Annotations of the .Todo.
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "todos"},
		entgql.QueryField("todos"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Policy defines the privacy policy of the Todo.
func (Todo) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{auth.MutationPrivacy("Todo")},
		Query:    privacy.QueryPolicy{auth.QueryPrivacy("Todo")},
	}
}
