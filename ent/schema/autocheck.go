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

// AutoCheck holds the schema definition for the AutoCheck entity.
type AutoCheck struct {
	ent.Schema
}

func (AutoCheck) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auto_checks"},
	}
}

// Fields of the AutoCheck.
func (AutoCheck) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Enum("tier").
			Values("TIER_1_AUTO", "TIER_2_SANDBOX"),
		field.Enum("status").
			Values("RUNNING", "PASSED", "FAILED", "SKIPPED").
			Default("RUNNING"),
		field.Bool("structureOk").
			Optional().
			Nillable(),
		field.Bool("depsOk").
			Optional().
			Nillable(),
		field.Bool("semverOk").
			Optional().
			Nillable(),
		field.Bool("fileSizeOk").
			Optional().
			Nillable(),
		field.Bool("securityScanOk").
			Optional().
			Nillable(),
		field.Bool("sandboxLoadOk").
			Optional().
			Nillable(),
		field.Bool("sandboxCrashFree").
			Optional().
			Nillable(),
		field.Text("logs").
			Optional().
			Nillable(),
		field.Float("score").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("finishedAt").
			Optional().
			Nillable(),

		field.String("pluginId"),
	}
}

// Edges of the AutoCheck.
func (AutoCheck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (AutoCheck) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId"),
	}
}
