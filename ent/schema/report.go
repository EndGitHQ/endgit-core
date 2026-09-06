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

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

func (Report) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reports"},
	}
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Enum("reason").
			Values("MALWARE", "SPAM", "COPIED", "BROKEN", "INAPPROPRIATE", "OTHER"),
		field.Text("details").
			Optional().
			Nillable(),
		field.Bool("resolved").
			Default(false),
		field.Time("createdAt").
			Default(time.Now),

		field.String("reporterId"),
		field.String("pluginId"),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reporter", User.Type).
			Field("reporterId").
			Unique().
			Required(),
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Report) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId"),
	}
}
