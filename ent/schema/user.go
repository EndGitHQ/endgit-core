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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(cuid.New).
			Unique(),
		field.String("githubId").
			Unique(),
		field.String("username").
			Unique(),
		field.String("displayName").
			Optional().
			Nillable(),
		field.String("email").
			Optional().
			Nillable(),
		field.String("avatarUrl").
			Optional().
			Nillable(),
		field.String("bio").
			Optional().
			Nillable(),
		field.Enum("trustLevel").
			Values("NEW", "TRUSTED", "MAINTAINER", "ADMIN").
			Default("NEW"),
		field.Int("trustScore").
			Default(0),
		field.Int("weeklyBuildQuota").
			Default(50),
		field.Int("weeklyBuildCount").
			Default(0),
		field.Time("quotaResetAt").
			Default(time.Now),
		field.Time("createdAt").
			Default(time.Now),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plugins", Plugin.Type).
			Ref("author"),
		edge.From("accounts", Account.Type).
			Ref("user"),
		edge.From("sessions", Session.Type).
			Ref("user"),
		edge.From("reviews", Review.Type).
			Ref("reviewer"),
		edge.From("reports", Report.Type).
			Ref("reporter"),
		edge.From("ratings", Rating.Type).
			Ref("user"),
		edge.From("pluginComments", PluginComment.Type).
			Ref("user"),
	}
}
