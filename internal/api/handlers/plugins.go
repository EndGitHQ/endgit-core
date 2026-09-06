package handlers

import (
	"net/http"
	"strconv"

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
	Tags        []string `json:"tags"`
	Keywords    []string `json:"keywords"`
	PluginType  string   `json:"pluginType"`
	Downloads   int      `json:"downloads"`
	Stars       int      `json:"stars"`
	Status      string   `json:"status"`
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
			Tags:        p.Tags,
			Keywords:    p.Keywords,
			PluginType:  p.PluginType.String(),
			Downloads:   p.Downloads,
			Stars:       p.Stars,
			Status:      p.Status.String(),
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
