// Code generated by construct, DO NOT EDIT.
package repository

import (
	"database/sql"
	"encoding/json"
	uuid "github.com/gofrs/uuid"
	construct "github.com/networkteam/construct"
	cjson "github.com/networkteam/construct/json"
	domain "myvendor.mytld/myproject/backend/domain"
	"time"
)

const (
	account_id             = "accounts.account_id"
	account_emailAddress   = "accounts.email_address"
	account_secret         = "accounts.secret"
	account_passwordHash   = "accounts.password_hash"
	account_role           = "accounts.role_identifier"
	account_lastLogin      = "accounts.last_login"
	account_organisationID = "accounts.organisation_id"
	account_createdAt      = "accounts.created_at"
	account_updatedAt      = "accounts.updated_at"
)

var accountSortFields = map[string]string{
	"createdat":    account_createdAt,
	"emailaddress": account_emailAddress,
	"lastlogin":    account_lastLogin,
	"role":         account_role,
	"updatedat":    account_updatedAt,
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

func AccountToChangeSet(r domain.Account) (c AccountChangeSet) {
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

var accountDefaultSelectJson = cjson.JsonBuildObject().
	Set("ID", cjson.Exp("accounts.account_id")).
	Set("EmailAddress", cjson.Exp("accounts.email_address")).
	Set("Secret", cjson.Exp("ENCODE(accounts.secret,'BASE64')")).
	Set("PasswordHash", cjson.Exp("ENCODE(accounts.password_hash,'BASE64')")).
	Set("Role", cjson.Exp("accounts.role_identifier")).
	Set("LastLogin", cjson.Exp("accounts.last_login")).
	Set("OrganisationID", cjson.Exp("accounts.organisation_id")).
	Set("CreatedAt", cjson.Exp("accounts.created_at")).
	Set("UpdatedAt", cjson.Exp("accounts.updated_at"))

func accountScanJsonRow(row construct.RowScanner) (result domain.Account, err error) {
	var data []byte
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return result, construct.ErrNotFound
		}
		return result, err
	}
	return result, json.Unmarshal(data, &result)
}
