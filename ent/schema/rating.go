package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/lucsky/cuid"
)

// Rating holds the schema definition for the Rating entity.
type Rating struct {
	ent.Schema
}

func (Rating) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ratings"},
	}
}

// Fields of the Rating.
func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Int("score"),
		field.Text("comment").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Default(time.Now),

		field.String("userId"),
		field.String("pluginId"),
	}
}

// Edges of the Rating.
func (Rating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("userId").
			Unique().
			Required(),
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Rating) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("userId", "pluginId").
			Unique(),
		index.Fields("pluginId"),
	}
}
