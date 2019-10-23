BEGIN;

CREATE TABLE accounts
(
    id              uuid  NOT NULL
        CONSTRAINT accounts_pkey PRIMARY KEY,
    type            text  NOT NULL,
    role_identifier text  NOT NULL,
    secret          bytea NOT NULL,
    email_address   text,
    password_hash   bytea,
    device_label    text,
    organisation_id uuid
        CONSTRAINT organisation_id_fkey REFERENCES organisations ON DELETE CASCADE,
    first_name      text,
    last_name       text,
    device_token    text,
    device_os       text
);

CREATE UNIQUE INDEX accounts_email_address_idx ON accounts (lower(email_address));
CREATE UNIQUE INDEX accounts_organisation_id_device_label_idx ON accounts (organisation_id, lower(device_label));

CREATE TABLE organisations
(
    id                uuid NOT NULL
        CONSTRAINT organisations_pkey PRIMARY KEY,
    organisation_name text NOT NULL
        CONSTRAINT organisation_name_key UNIQUE
);

CREATE UNIQUE INDEX organisations_organisation_name_unique_idx ON organisations (lower(organisation_name));

COMMIT;
