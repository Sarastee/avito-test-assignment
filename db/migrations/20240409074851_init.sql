-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS banners (
    id BIGSERIAL PRIMARY KEY,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    selected_revision BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS banner_revision_tags (
    id BIGSERIAL PRIMARY KEY,
    banner_id BIGINT NOT NULL,
    feature_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    CONSTRAINT fk_banner_revision_tags_banner_id FOREIGN KEY (banner_id) REFERENCES banners (id) ON DELETE CASCADE,
    CONSTRAINT banner_revision_tags_unique UNIQUE (feature_id, tag_id)
);

CREATE TABLE IF NOT EXISTS banner_revisions (
    id BIGSERIAL PRIMARY KEY,
    revision_id BIGINT NOT NULL DEFAULT 1,
    banner_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_banner_revisions_banner_id FOREIGN KEY (banner_id) REFERENCES banners (id) ON DELETE CASCADE,
    CONSTRAINT banner_revisions_unique UNIQUE (revision_id, banner_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS banner_revision_tags;
DROP TABLE IF EXISTS banner_revisions;
-- +goose StatementEnd
