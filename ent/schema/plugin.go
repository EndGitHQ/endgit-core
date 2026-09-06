package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Plugin holds the schema definition for the Plugin entity.
type Plugin struct {
	ent.Schema
}

func (Plugin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plugins"},
	}
}

// Fields of the Plugin.
func (Plugin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("name").
			Unique(),
		field.String("slug").
			Unique(),
		field.String("displayName"),
		field.String("description"),
		field.Text("longDescription").
			Optional().
			Nillable(),
		field.String("iconUrl").
			Optional().
			Nillable(),
		field.String("repoUrl").
			Optional().
			Nillable(),
		field.String("license").
			Optional().
			Nillable(),
		field.Strings("tags").
			Default([]string{}),
		field.Strings("keywords").
			Default([]string{}),
		field.Enum("pluginType").
			Values("PYTHON", "CPP").
			Default("PYTHON"),
		field.Int("downloads").
			Default(0),
		field.Int("stars").
			Default(0),
		field.Int("commentCount").
			Default(0),
		field.Enum("status").
			Values("DRAFT", "BUILDING", "BUILD_FAILED", "PENDING_REVIEW", "APPROVED", "REJECTED", "FLAGGED", "SUSPENDED").
			Default("DRAFT"),
		field.Text("statusReason").
			Optional().
			Nillable(),
		field.Float("trustScore").
			Default(0),
		field.Enum("qualityBadge").
			Values("NONE", "BRONZE", "SILVER", "GOLD").
			Default("NONE"),
		field.Bool("isVerified").
			Default(false),
		field.Bool("isFeatured").
			Default(false),
		field.String("reviewBuildId").
			Optional().
			Nillable(),
		field.String("webhookId").
			Optional().
			Nillable(),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),

		// Foreign key to User.
		field.String("authorId"),
	}
}

// Edges of the Plugin.
func (Plugin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Field("authorId").
			Unique().
			Required(),
		edge.From("versions", Version.Type).
			Ref("plugin"),
		edge.From("builds", Build.Type).
			Ref("plugin"),
		edge.From("reviews", Review.Type).
			Ref("plugin"),
		edge.From("reports", Report.Type).
			Ref("plugin"),
		edge.From("ratings", Rating.Type).
			Ref("plugin"),
		edge.From("comments", PluginComment.Type).
			Ref("plugin"),
		edge.From("autoChecks", AutoCheck.Type).
			Ref("plugin"),
		edge.From("moderationLogs", ModerationLog.Type).
			Ref("plugin"),
		edge.From("analytics", PluginAnalytics.Type).
			Ref("plugin"),
	}
}
