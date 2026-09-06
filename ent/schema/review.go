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

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

func (Review) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reviews"},
	}
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Enum("decision").
			Values("APPROVED", "REJECTED", "REQUEST_CHANGES"),
		field.Text("comment").
			Optional().
			Nillable(),
		field.Bool("codeClean").
			Optional().
			Nillable(),
		field.Bool("noBackdoor").
			Optional().
			Nillable(),
		field.Bool("rulesOk").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Default(time.Now),

		field.String("reviewerId"),
		field.String("pluginId"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reviewer", User.Type).
			Field("reviewerId").
			Unique().
			Required(),
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Review) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId"),
		index.Fields("reviewerId"),
	}
}
