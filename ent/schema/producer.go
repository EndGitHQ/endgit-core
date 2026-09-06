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

// Producer holds the schema definition for the Producer entity.
type Producer struct {
	ent.Schema
}

func (Producer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "producers"},
	}
}

// Fields of the Producer.
func (Producer) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("githubUser"),
		field.Enum("role").
			Values("COLLABORATOR", "CONTRIBUTOR", "TRANSLATOR", "REQUESTER"),
		field.String("versionId"),
	}
}

// Edges of the Producer.
func (Producer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("version", Version.Type).
			Field("versionId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Producer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("versionId"),
	}
}
