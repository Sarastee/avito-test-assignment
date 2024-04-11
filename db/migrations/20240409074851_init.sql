-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS banners (
    banner_id BIGSERIAL PRIMARY KEY,
    selected_revision_id BIGINT NULL
);

CREATE TABLE IF NOT EXISTS banner_revisions (
    revision_id SERIAL PRIMARY KEY,
    banner_id BIGSERIAL NOT NULL,
    feature_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    CONSTRAINT fk_banner_id FOREIGN KEY (banner_id) REFERENCES banners (banner_id) ON DELETE CASCADE
);

ALTER TABLE banners
    ADD CONSTRAINT fk_selected_revision FOREIGN KEY (selected_revision_id) REFERENCES banner_revisions (revision_id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS banner_revision_tags (
    id BIGSERIAL PRIMARY KEY,
    revision_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    CONSTRAINT fk_banner_revision_tags_revision_id FOREIGN KEY (revision_id) REFERENCES banner_revisions(revision_id) ON DELETE CASCADE,
    UNIQUE (revision_id, tag_id)
);

CREATE TABLE IF NOT EXISTS selected_revisions (
    banner_id INT NOT NULL,
    revision_id BIGINT NOT NULL,
    feature_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    CONSTRAINT fk_selected_revisions_revision_id FOREIGN KEY (revision_id) REFERENCES banner_revisions(revision_id) ON DELETE CASCADE,
    PRIMARY KEY (feature_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
