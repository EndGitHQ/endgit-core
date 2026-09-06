package handlers

import (
	"net/http"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/EndGitHQ/endgit-core/ent"
	"github.com/EndGitHQ/endgit-core/ent/plugin"
	"github.com/labstack/echo/v5"
)

const (
	defaultPluginsLimit = 20
	maxPluginsLimit     = 100
)

type PluginListItem struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	IconURL     *string  `json:"iconUrl,omitempty"`
	Keywords    []string `json:"keywords"`
	PluginType  string   `json:"pluginType"`
	Downloads   int      `json:"downloads"`
	Stars       int      `json:"stars"`
	IsVerified  bool     `json:"isVerified"`
	IsFeatured  bool     `json:"isFeatured"`
	AuthorID    string   `json:"authorId"`
}

type ListPluginsResponse struct {
	Items  []PluginListItem `json:"items"`
	Count  int              `json:"count"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
}

type PluginDetailsResponse struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Slug            string    `json:"slug"`
	DisplayName     string    `json:"displayName"`
	Description     string    `json:"description"`
	LongDescription *string   `json:"longDescription,omitempty"`
	IconURL         *string   `json:"iconUrl,omitempty"`
	RepoURL         *string   `json:"repoUrl,omitempty"`
	License         *string   `json:"license,omitempty"`
	Tags            []string  `json:"tags"`
	Keywords        []string  `json:"keywords"`
	PluginType      string    `json:"pluginType"`
	Downloads       int       `json:"downloads"`
	Stars           int       `json:"stars"`
	CommentCount    int       `json:"commentCount"`
	Status          string    `json:"status"`
	StatusReason    *string   `json:"statusReason,omitempty"`
	TrustScore      float64   `json:"trustScore"`
	QualityBadge    string    `json:"qualityBadge"`
	IsVerified      bool      `json:"isVerified"`
	IsFeatured      bool      `json:"isFeatured"`
	ReviewBuildID   *string   `json:"reviewBuildId,omitempty"`
	WebhookID       *string   `json:"webhookId,omitempty"`
	AuthorID        string    `json:"authorId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// ListPlugins godoc
// @Summary List plugins
// @Description Returns a paginated list of plugins.
// @Tags plugins
// @Produce json
// @Param limit query int false "Number of items (default 20, max 100)"
// @Param offset query int false "Pagination offset (default 0)"
// @Success 200 {object} ListPluginsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /plugins [get]
func ListPlugins(c *echo.Context) error {
	limit, err := parsePaginationInt(c.QueryParam("limit"), defaultPluginsLimit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid limit"})
	}
	if limit > maxPluginsLimit {
		limit = maxPluginsLimit
	}

	offset, err := parsePaginationInt(c.QueryParam("offset"), 0)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid offset"})
	}

	client := ent.FromContext(c.Request().Context())
	if client == nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "database client not available"})
	}

	items, err := client.Plugin.
		Query().
		Order(plugin.ByCreatedAt(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to list plugins"})
	}

	respItems := make([]PluginListItem, 0, len(items))
	for _, p := range items {
		respItems = append(respItems, PluginListItem{
			ID:          p.ID,
			Name:        p.Name,
			Slug:        p.Slug,
			DisplayName: p.DisplayName,
			Description: p.Description,
			IconURL:     p.IconUrl,
			Keywords:    p.Keywords,
			PluginType:  p.PluginType.String(),
			Downloads:   p.Downloads,
			Stars:       p.Stars,
			IsVerified:  p.IsVerified,
			IsFeatured:  p.IsFeatured,
			AuthorID:    p.AuthorId,
		})
	}

	return c.JSON(http.StatusOK, ListPluginsResponse{
		Items:  respItems,
		Count:  len(respItems),
		Limit:  limit,
		Offset: offset,
	})
}

// GetPluginDetails godoc
// @Summary Plugin details
// @Description Returns details for a plugin by slug.
// @Tags plugins
// @Produce json
// @Param slug path string true "Plugin slug"
// @Success 200 {object} PluginDetailsResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /plugins/{slug} [get]
func GetPluginDetails(c *echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "plugin not found"})
	}

	client := ent.FromContext(c.Request().Context())
	if client == nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "database client not available"})
	}

	p, err := client.Plugin.
		Query().
		Where(plugin.SlugEQ(slug)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: "plugin not found"})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get plugin"})
	}

	return c.JSON(http.StatusOK, PluginDetailsResponse{
		ID:              p.ID,
		Name:            p.Name,
		Slug:            p.Slug,
		DisplayName:     p.DisplayName,
		Description:     p.Description,
		LongDescription: p.LongDescription,
		IconURL:         p.IconUrl,
		RepoURL:         p.RepoUrl,
		License:         p.License,
		Tags:            p.Tags,
		Keywords:        p.Keywords,
		PluginType:      p.PluginType.String(),
		Downloads:       p.Downloads,
		Stars:           p.Stars,
		CommentCount:    p.CommentCount,
		Status:          p.Status.String(),
		StatusReason:    p.StatusReason,
		TrustScore:      p.TrustScore,
		QualityBadge:    p.QualityBadge.String(),
		IsVerified:      p.IsVerified,
		IsFeatured:      p.IsFeatured,
		ReviewBuildID:   p.ReviewBuildId,
		WebhookID:       p.WebhookId,
		AuthorID:        p.AuthorId,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	})
}

func parsePaginationInt(raw string, fallback int) (int, error) {
	if raw == "" {
		return fallback, nil
	}
	v, err := strconv.Atoi(raw)
	if err != nil || v < 0 {
		return 0, strconv.ErrSyntax
	}
	return v, nil
}
