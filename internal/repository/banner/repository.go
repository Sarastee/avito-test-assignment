package banner

import (
	"github.com/rs/zerolog"
	"github.com/sarastee/avito-test-assignment/internal/repository"
	"github.com/sarastee/platform_common/pkg/db"
)

const (
	bannersTable           = "banners"
	idColumn               = "id"
	isActiveColumn         = "is_active"
	selectedRevisionColumn = "selected_revision"

	bannerFeatureTagsTable = "banner_feature_tags"
	bannerIDColumn         = "banner_id"
	featureIDColumn        = "feature_id"
	tagIDColumn            = "tag_id"

	bannerRevisionsTable = "banner_revisions"
	contentColumn        = "content"
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
