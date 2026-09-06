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

// Version holds the schema definition for the Version entity.
type Version struct {
	ent.Schema
}

func (Version) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "versions"},
	}
}

// Fields of the Version.
func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("version"),
		field.Text("changelog").
			Optional().
			Nillable(),
		field.Text("longDescription").
			Optional().
			Nillable(),
		field.String("fileUrl"),
		field.String("fileName"),
		field.Int("fileSize"),
		field.String("fileHash"),
		field.String("minApiVersion").
			Optional().
			Nillable(),
		field.Strings("supportedApis"),
		field.Int("downloads").
			Default(0),
		field.Bool("isLatest").
			Default(false),
		field.Bool("isPreRelease").
			Default(false),
		field.Enum("status").
			Values("DRAFT", "PENDING", "AUTO_PASSED", "SANDBOX_PASSED", "APPROVED", "REJECTED").
			Default("PENDING"),
		field.Text("statusReason").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Default(time.Now),

		field.String("vtScanId").
			Optional().
			Nillable(),
		field.String("vtStatus").
			Optional().
			Nillable(),
		field.Int("vtMalicious").
			Optional().
			Nillable(),
		field.Int("vtSuspicious").
			Optional().
			Nillable(),
		field.Int("vtUndetected").
			Optional().
			Nillable(),
		field.Int("vtTotal").
			Optional().
			Nillable(),
		field.String("vtPermalink").
			Optional().
			Nillable(),
		field.Time("vtScanDate").
			Optional().
			Nillable(),

		field.String("pluginId"),
	}
}

// Edges of the Version.
func (Version) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("dependencies", Dependency.Type).
			Ref("parentVersion"),
		edge.From("producers", Producer.Type).
			Ref("version"),
	}
}

func (Version) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId", "version").
			Unique(),
		index.Fields("pluginId"),
	}
}
