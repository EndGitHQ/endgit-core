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

// ModerationLog holds the schema definition for the ModerationLog entity.
type ModerationLog struct {
	ent.Schema
}

func (ModerationLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "moderation_logs"},
	}
}

// Fields of the ModerationLog.
func (ModerationLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("action"),
		field.String("targetType"),
		field.String("targetId"),
		field.String("oldStatus"),
		field.String("newStatus"),
		field.Text("reason").
			Optional().
			Nillable(),
		field.String("actorId"),
		field.Time("createdAt").
			Default(time.Now),

		field.String("pluginId").
			Optional().
			Nillable(),
	}
}

// Edges of the ModerationLog.
func (ModerationLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plugin", Plugin.Type).
			Field("pluginId").
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.SetNull}),
	}
}

func (ModerationLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("pluginId"),
		index.Fields("actorId"),
		index.Fields("createdAt"),
	}
}
