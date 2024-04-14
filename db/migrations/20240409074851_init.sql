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

CREATE TABLE IF NOT EXISTS banner_feature_tags (
    id BIGSERIAL PRIMARY KEY,
    banner_id BIGINT NOT NULL,
    feature_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    CONSTRAINT fk_banner_feature_tags_banner_id FOREIGN KEY (banner_id) REFERENCES banners (id) ON DELETE CASCADE,
    CONSTRAINT banner_feature_tags_unique UNIQUE (feature_id, tag_id)
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

CREATE OR REPLACE FUNCTION banner_update_trigger() RETURNS TRIGGER AS
$$
BEGIN
UPDATE banners SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER banner_update
    AFTER UPDATE
    ON banners
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE FUNCTION banner_update_trigger();


CREATE OR REPLACE FUNCTION banner_insert_revision_trigger() RETURNS TRIGGER AS
$$
DECLARE
    v_count      int    := 0;
    selected_revision bigint := 0;
BEGIN
SELECT count(*) INTO v_count FROM banner_revisions WHERE banner_id = NEW.banner_id;
IF v_count != 0 THEN
SELECT max(revision_id) INTO selected_revision FROM banner_revisions WHERE banner_id = NEW.banner_id;
NEW.revision_id = selected_revision + 1;
DELETE FROM banner_revisions WHERE banner_id = NEW.banner_id and revision_id < selected_revision - 1;
UPDATE banners SET selected_revision = NEW.revision_id WHERE id = NEW.banner_id;
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER banner_insert_revision
    BEFORE INSERT
    ON banner_revisions
    FOR EACH ROW
EXECUTE FUNCTION banner_insert_revision_trigger();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS banner_feature_tags;
DROP TABLE IF EXISTS banner_revisions;
-- +goose StatementEnd
