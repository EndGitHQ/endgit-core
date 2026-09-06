package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Dependency holds the schema definition for the Dependency entity.
type Dependency struct {
	ent.Schema
}

func (Dependency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "dependencies"},
	}
}

// Fields of the Dependency.
func (Dependency) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("name"),
		field.String("version"),
		field.String("versionId"),
	}
}

// Edges of the Dependency.
func (Dependency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parentVersion", Version.Type).
			Field("versionId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
