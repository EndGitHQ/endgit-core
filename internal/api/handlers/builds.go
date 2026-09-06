package handlers

import (
	"net/http"
	"time"

	"github.com/EndGitHQ/endgit-core/ent"
	"github.com/EndGitHQ/endgit-core/ent/build"
	"github.com/labstack/echo/v5"
)

type BuildDetailsResponse struct {
	ID                string     `json:"id"`
	BuildNumber       int        `json:"buildNumber"`
	CommitHash        *string    `json:"commitHash,omitempty"`
	Branch            string     `json:"branch"`
	Status            string     `json:"status"`
	IsRelease         bool       `json:"isRelease"`
	TriggerType       string     `json:"triggerType"`
	CommitMessage     *string    `json:"commitMessage,omitempty"`
	Logs              *string    `json:"logs,omitempty"`
	ArtifactURL       *string    `json:"artifactUrl,omitempty"`
	ArtifactSize      *int       `json:"artifactSize,omitempty"`
	ArtifactURLLinux  *string    `json:"artifactUrlLinux,omitempty"`
	ArtifactSizeLinux *int       `json:"artifactSizeLinux,omitempty"`
	ArtifactURLWin    *string    `json:"artifactUrlWin,omitempty"`
	ArtifactSizeWin   *int       `json:"artifactSizeWin,omitempty"`
	WinBuildStatus    *string    `json:"winBuildStatus,omitempty"`
	LinuxBuildStatus  *string    `json:"linuxBuildStatus,omitempty"`
	GhActionsRunID    *string    `json:"ghActionsRunId,omitempty"`
	Duration          *int       `json:"duration,omitempty"`
	SafeScore         *int       `json:"safeScore,omitempty"`
	ScanResults       *string    `json:"scanResults,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	FinishedAt        *time.Time `json:"finishedAt,omitempty"`
	PluginID          string     `json:"pluginId"`
}

// GetBuildDetails godoc
// @Summary Build details
// @Description Returns details for a build by ID.
// @Tags builds
// @Produce json
// @Param id path string true "Build ID"
// @Success 200 {object} BuildDetailsResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /builds/{id} [get]
func GetBuildDetails(c *echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "build not found"})
	}

	client := ent.FromContext(c.Request().Context())
	if client == nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "database client not available"})
	}

	b, err := client.Build.
		Query().
		Where(build.IDEQ(id)).
		Only(c.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, ErrorResponse{Error: "build not found"})
		}
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to get build"})
	}

	return c.JSON(http.StatusOK, BuildDetailsResponse{
		ID:                b.ID,
		BuildNumber:       b.BuildNumber,
		CommitHash:        b.CommitHash,
		Branch:            b.Branch,
		Status:            b.Status.String(),
		IsRelease:         b.IsRelease,
		TriggerType:       b.TriggerType,
		CommitMessage:     b.CommitMessage,
		Logs:              b.Logs,
		ArtifactURL:       b.ArtifactUrl,
		ArtifactSize:      b.ArtifactSize,
		ArtifactURLLinux:  b.ArtifactUrlLinux,
		ArtifactSizeLinux: b.ArtifactSizeLinux,
		ArtifactURLWin:    b.ArtifactUrlWin,
		ArtifactSizeWin:   b.ArtifactSizeWin,
		WinBuildStatus:    b.WinBuildStatus,
		LinuxBuildStatus:  b.LinuxBuildStatus,
		GhActionsRunID:    b.GhActionsRunId,
		Duration:          b.Duration,
		SafeScore:         b.SafeScore,
		ScanResults:       b.ScanResults,
		CreatedAt:         b.CreatedAt,
		FinishedAt:        b.FinishedAt,
		PluginID:          b.PluginId,
	})
}
