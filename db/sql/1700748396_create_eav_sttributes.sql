CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS eav_attributes
(
    code         VARCHAR(255) PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    value_type   SMALLINT     NOT NULL CHECK (min_length >= 0),
    description  VARCHAR(255) NOT NULL,
    field_format VARCHAR(255),
    regexp       VARCHAR(255),
    min_length   SMALLINT CHECK (min_length >= 0),
    max_length   SMALLINT CHECK (max_length >= 0),
    is_selection BOOLEAN      NOT NULL DEFAULT FALSE,
    is_required  BOOLEAN      NOT NULL DEFAULT FALSE,
    is_visible   BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS eav_attribute_options
(
    id                 BIGSERIAL PRIMARY KEY,
    eav_attribute_code VARCHAR(255)  NOT NULL,
    value              VARCHAR(1024) NOT NULL,
    ordering           SMALLINT      NOT NULL CHECK (ordering >= 0),
    visible            BOOLEAN       NOT NULL DEFAULT FALSE,
    created_at         TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT eav_attribute_options_fk_eav_attribute_code FOREIGN KEY (eav_attribute_code) REFERENCES eav_attributes (code)
);
CREATE INDEX eav_attribute_options_idx_eav_attribute_code
    ON eav_attribute_options (eav_attribute_code);

CREATE TABLE IF NOT EXISTS eav_attribute_values
(
    user_id            BIGINT       NOT NULL,
    eav_attribute_code VARCHAR(255) NOT NULL,
    value              TEXT         NOT NULL,
    created_at         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, eav_attribute_code),
    CONSTRAINT eav_attribute_values_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT eav_attribute_code FOREIGN KEY (eav_attribute_code) REFERENCES eav_attributes (code)
);
CREATE INDEX eav_attribute_values_idx_user_id
    ON eav_attribute_values (user_id);
CREATE INDEX eav_attribute_values_idx_eav_attribute_code
    ON eav_attribute_values (eav_attribute_code);

-- DROP TABLE IF EXISTS eav_attribute_values;
-- DROP TABLE IF EXISTS eav_attribute_options;
-- DROP TABLE IF EXISTS eav_attributes;
-- DROP TABLE IF EXISTS users;
