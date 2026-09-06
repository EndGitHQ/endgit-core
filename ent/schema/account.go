package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/lucsky/cuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "accounts"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("userId"),
		field.String("type"),
		field.String("provider"),
		field.String("providerAccountId"),
		field.Text("refresh_token").
			Optional().
			Nillable(),
		field.Text("access_token").
			Optional().
			Nillable(),
		field.Int("expires_at").
			Optional().
			Nillable(),
		field.String("token_type").
			Optional().
			Nillable(),
		field.String("scope").
			Optional().
			Nillable(),
		field.Text("id_token").
			Optional().
			Nillable(),
		field.String("session_state").
			Optional().
			Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("userId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "providerAccountId").
			Unique(),
	}
}
