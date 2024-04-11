package banner

import "context"

// GetBanner function which get banner instance
func (r Repo) GetBanner(_ context.Context, _ int64, _ int64) (string, error) {
	panic("implement me")
}
