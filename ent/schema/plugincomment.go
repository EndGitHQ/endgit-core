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

// PluginComment holds the schema definition for the PluginComment entity.
type PluginComment struct {
	ent.Schema
}

func (PluginComment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plugin_comments"},
	}
}

// Fields of the PluginComment.
func (PluginComment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.Text("body"),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.String("userId"),
		field.String("pluginId"),
		field.String("parentId").
			Optional().
			Nillable(),
	}
}

// Edges of the PluginComment.
func (PluginComment) Edges() []ent.Edge {
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
		edge.To("parent", PluginComment.Type).
			Unique().
			Field("parentId").
			From("replies").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (PluginComment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId", "createdAt"),
		index.Fields("parentId"),
	}
}
