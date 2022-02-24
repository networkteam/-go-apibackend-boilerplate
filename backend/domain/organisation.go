package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type Organisation struct {
	ID   uuid.UUID `read_col:"organisations.organisation_id" write_col:"organisation_id"`
	Name string    `read_col:"organisations.name,sortable" write_col:"name"`

	CreatedAt time.Time `read_col:"organisations.created_at,sortable"`
	UpdatedAt time.Time `read_col:"organisations.updated_at,sortable"`
}

type OrganisationsQuery struct {
	IDs []uuid.UUID
	Q   *string
}

func (f *OrganisationsQuery) SetOrganisationID(organisationID *uuid.UUID) {
	if organisationID != nil {
		f.IDs = []uuid.UUID{*organisationID}
	} else {
		f.IDs = nil
	}
}
