package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	bannersTable             = "banners"
	bannerIDColumn           = "banner_id"
	selectedRevisionIDColumn = "selected_revision_id"

	bannerRevisionsTable = "banner_revisions"
	revisionIDColumn     = "revision_id"
	featureIDColumn      = "feature_id"
	contentColumn        = "content"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
	isActiveColumn       = "is_active"

	bannerRevisionTagsTable = "banner_revision_tags"
	idColumn                = "id"
	tagIDColumn             = "tag_id"

	selectedRevisionsTable = "selected_revisions"
)

var _ repository.BannerRepository = (*Repo)(nil)

// Repo banner repository for CRUD operations
type Repo struct {
	logger *zerolog.Logger
	db     db.Client
}

// NewRepo function which get new repo instance
func NewRepo(logger *zerolog.Logger, dbClient db.Client) *Repo {
	return &Repo{
		logger: logger,
		db:     dbClient,
	}
}
