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

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Annotations(entgql.OrderField("USERNAME")),
		field.String("password").Sensitive().Annotations(entgql.Skip(entgql.SkipType, entgql.SkipWhereInput)),
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").Default(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type).Annotations(entgql.RelayConnection()),
	}
}

// Annotations of the .User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entgql.QueryField("users"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{auth.MutationPrivacy("User")},
		Query:    privacy.QueryPolicy{auth.QueryPrivacy("User")},
	}
}
