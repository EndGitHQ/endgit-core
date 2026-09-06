package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/EndGitHQ/endgit-core/ent"
	"github.com/EndGitHQ/endgit-core/ent/plugin"
	"github.com/EndGitHQ/endgit-core/ent/user"
	"github.com/EndGitHQ/endgit-core/internal/db"
)

type seedPlugin struct {
	Name        string
	Slug        string
	DisplayName string
	Description string
	Tags        []string
	Keywords    []string
	PluginType  plugin.PluginType
	Downloads   int
	Stars       int
	IsVerified  bool
	IsFeatured  bool
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, cfg, err := db.OpenFromEnv(ctx)
	if err != nil {
		log.Fatalf("seed failed: open db: %v", err)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			log.Printf("warning: close db: %v", closeErr)
		}
	}()

	maintainer, createdUser, err := ensureSeedUser(ctx, client)
	if err != nil {
		log.Fatalf("seed failed: ensure user: %v", err)
	}

	plugins := []seedPlugin{
		{
			Name:        "syntax-shield",
			Slug:        "syntax-shield",
			DisplayName: "Syntax Shield",
			Description: "Static checks for repository plugin manifests.",
			Tags:        []string{"quality", "lint"},
			Keywords:    []string{"schema", "validation", "ci"},
			PluginType:  plugin.PluginTypePYTHON,
			Downloads:   1240,
			Stars:       74,
			IsVerified:  true,
			IsFeatured:  true,
		},
		{
			Name:        "release-notes-bot",
			Slug:        "release-notes-bot",
			DisplayName: "Release Notes Bot",
			Description: "Builds changelogs from merged pull requests.",
			Tags:        []string{"release", "automation"},
			Keywords:    []string{"changelog", "git", "github"},
			PluginType:  plugin.PluginTypePYTHON,
			Downloads:   860,
			Stars:       39,
			IsVerified:  true,
			IsFeatured:  false,
		},
		{
			Name:        "cpp-build-sentry",
			Slug:        "cpp-build-sentry",
			DisplayName: "C++ Build Sentry",
			Description: "Monitors C++ plugin build health and cache misses.",
			Tags:        []string{"build", "cpp"},
			Keywords:    []string{"cmake", "cache", "diagnostics"},
			PluginType:  plugin.PluginTypeCPP,
			Downloads:   410,
			Stars:       22,
			IsVerified:  false,
			IsFeatured:  false,
		},
	}

	createdPlugins, existingPlugins, err := seedPlugins(ctx, client, maintainer.ID, plugins)
	if err != nil {
		log.Fatalf("seed failed: seed plugins: %v", err)
	}

	fmt.Printf("seed complete (driver=%s db=%s): user(created=%t), plugins(created=%d existing=%d)\n",
		cfg.Driver,
		cfg.Name,
		createdUser,
		createdPlugins,
		existingPlugins,
	)
}

func ensureSeedUser(ctx context.Context, client *ent.Client) (*ent.User, bool, error) {
	u, err := client.User.Query().Where(user.GithubIdEQ("seed-maintainer-1")).Only(ctx)
	if err == nil {
		return u, false, nil
	}
	if !ent.IsNotFound(err) {
		return nil, false, err
	}

	u, err = client.User.Create().
		SetGithubId("seed-maintainer-1").
		SetUsername("seed-maintainer").
		SetDisplayName("Seed Maintainer").
		SetTrustLevel(user.TrustLevelMAINTAINER).
		Save(ctx)
	if err != nil {
		return nil, false, err
	}
	return u, true, nil
}

func seedPlugins(ctx context.Context, client *ent.Client, authorID string, candidates []seedPlugin) (int, int, error) {
	created := 0
	existing := 0

	for _, candidate := range candidates {
		exists, err := client.Plugin.Query().Where(plugin.SlugEQ(candidate.Slug)).Exist(ctx)
		if err != nil {
			return created, existing, err
		}
		if exists {
			existing++
			continue
		}

		_, err = client.Plugin.Create().
			SetName(candidate.Name).
			SetSlug(candidate.Slug).
			SetDisplayName(candidate.DisplayName).
			SetDescription(candidate.Description).
			SetTags(candidate.Tags).
			SetKeywords(candidate.Keywords).
			SetPluginType(candidate.PluginType).
			SetDownloads(candidate.Downloads).
			SetStars(candidate.Stars).
			SetIsVerified(candidate.IsVerified).
			SetIsFeatured(candidate.IsFeatured).
			SetStatus(plugin.StatusAPPROVED).
			SetQualityBadge(plugin.QualityBadgeBRONZE).
			SetAuthorID(authorID).
			Save(ctx)
		if err != nil {
			var constraintError *ent.ConstraintError
			if errors.As(err, &constraintError) {
				existing++
				continue
			}
			return created, existing, err
		}

		created++
	}

	return created, existing, nil
}
