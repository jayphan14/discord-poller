CREATE TABLE companies (
  id          BIGSERIAL PRIMARY KEY,
  name        TEXT      NOT NULL,
  notified    BOOLEAN   NOT NULL,
  import_date DATE      NOT NULL
);