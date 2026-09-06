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

// Build holds the schema definition for the Build entity.
type Build struct {
	ent.Schema
}

func (Build) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "builds"},
	}
}

// Fields of the Build.
func (Build) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Int("buildNumber"),
		field.String("commitHash").
			Optional().
			Nillable(),
		field.String("branch").
			Default("main"),
		field.Enum("status").
			Values("QUEUED", "RUNNING", "SUCCESS", "FAILED", "CANCELLED").
			Default("QUEUED"),
		field.Bool("isRelease").
			Default(false),
		field.String("triggerType").
			Default("MANUAL"),
		field.Text("commitMessage").
			Optional().
			Nillable(),
		field.Text("logs").
			Optional().
			Nillable(),
		field.String("artifactUrl").
			Optional().
			Nillable(),
		field.Int("artifactSize").
			Optional().
			Nillable(),

		field.String("artifactUrlLinux").
			Optional().
			Nillable(),
		field.Int("artifactSizeLinux").
			Optional().
			Nillable(),
		field.String("artifactUrlWin").
			Optional().
			Nillable(),
		field.Int("artifactSizeWin").
			Optional().
			Nillable(),

		field.String("winBuildStatus").
			Optional().
			Nillable(),
		field.String("linuxBuildStatus").
			Optional().
			Nillable(),
		field.String("ghActionsRunId").
			Optional().
			Nillable(),
		field.Int("duration").
			Optional().
			Nillable(),
		field.Int("safeScore").
			Optional().
			Nillable(),
		field.Text("scanResults").
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

// Edges of the Build.
func (Build) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Build) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId"),
		index.Fields("createdAt"),
		index.Fields("status"),
		index.Fields("pluginId", "createdAt"),
		index.Fields("pluginId", "buildNumber"),
	}
}
