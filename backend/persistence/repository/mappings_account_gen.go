// Code generated by construct, DO NOT EDIT.
package repository

import (
	uuid "github.com/gofrs/uuid"
	qrb "github.com/networkteam/qrb"
	builder "github.com/networkteam/qrb/builder"
	fn "github.com/networkteam/qrb/fn"

	"myvendor.mytld/myproject/backend/domain/model"
	domain "myvendor.mytld/myproject/backend/domain/types"

	"time"
)

var account = struct {
	builder.Identer
	ID             builder.IdentExp
	EmailAddress   builder.IdentExp
	Secret         builder.IdentExp
	PasswordHash   builder.IdentExp
	Role           builder.IdentExp
	LastLogin      builder.IdentExp
	OrganisationID builder.IdentExp
	CreatedAt      builder.IdentExp
	UpdatedAt      builder.IdentExp
}{
	CreatedAt:      qrb.N("accounts.created_at"),
	EmailAddress:   qrb.N("accounts.email_address"),
	ID:             qrb.N("accounts.account_id"),
	Identer:        qrb.N("accounts"),
	LastLogin:      qrb.N("accounts.last_login"),
	OrganisationID: qrb.N("accounts.organisation_id"),
	PasswordHash:   qrb.N("accounts.password_hash"),
	Role:           qrb.N("accounts.role_identifier"),
	Secret:         qrb.N("accounts.secret"),
	UpdatedAt:      qrb.N("accounts.updated_at"),
}

var accountSortFields = map[string]builder.IdentExp{
	"createdat":    account.CreatedAt,
	"emailaddress": account.EmailAddress,
	"lastlogin":    account.LastLogin,
	"role":         account.Role,
	"updatedat":    account.UpdatedAt,
}

type AccountChangeSet struct {
	ID             *uuid.UUID
	EmailAddress   *string
	Secret         []byte
	PasswordHash   []byte
	Role           *domain.Role
	LastLogin      **time.Time
	OrganisationID *uuid.NullUUID
}

func (c AccountChangeSet) toMap() map[string]interface{} {
	m := make(map[string]interface{})
	if c.ID != nil {
		m["account_id"] = *c.ID
	}
	if c.EmailAddress != nil {
		m["email_address"] = *c.EmailAddress
	}
	if c.Secret != nil {
		m["secret"] = c.Secret
	}
	if c.PasswordHash != nil {
		m["password_hash"] = c.PasswordHash
	}
	if c.Role != nil {
		m["role_identifier"] = *c.Role
	}
	if c.LastLogin != nil {
		m["last_login"] = *c.LastLogin
	}
	if c.OrganisationID != nil {
		m["organisation_id"] = *c.OrganisationID
	}
	return m
}

func AccountToChangeSet(r model.Account) (c AccountChangeSet) {
	if r.ID != uuid.Nil {
		c.ID = &r.ID
	}
	c.EmailAddress = &r.EmailAddress
	c.Secret = r.Secret
	c.PasswordHash = r.PasswordHash
	c.Role = &r.Role
	c.LastLogin = &r.LastLogin
	c.OrganisationID = &r.OrganisationID
	return
}

var accountDefaultJson = fn.JsonBuildObject().
	Prop("ID", account.ID).
	Prop("EmailAddress", account.EmailAddress).
	Prop("Secret", qrb.Func("ENCODE", account.Secret, qrb.String("BASE64"))).
	Prop("PasswordHash", qrb.Func("ENCODE", account.PasswordHash, qrb.String("BASE64"))).
	Prop("Role", account.Role).
	Prop("LastLogin", account.LastLogin).
	Prop("OrganisationID", account.OrganisationID).
	Prop("CreatedAt", account.CreatedAt).
	Prop("UpdatedAt", account.UpdatedAt)
