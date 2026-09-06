package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/lucsky/cuid"
)

// PluginAnalytics holds the schema definition for the PluginAnalytics entity.
type PluginAnalytics struct {
	ent.Schema
}

func (PluginAnalytics) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plugin_analytics"},
	}
}

// Fields of the PluginAnalytics.
func (PluginAnalytics) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("pluginId"),
		// stored just the date part for daily aggregation
		field.Time("date").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Int("downloads").
			Default(0),
		field.Int("views").
			Default(0),
	}
}

// Edges of the PluginAnalytics.
func (PluginAnalytics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Required().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (PluginAnalytics) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId", "date").
			Unique(),
	}
}
