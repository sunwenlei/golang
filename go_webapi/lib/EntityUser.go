package lib

import (
	"database/sql"
	"time"
)

//MstPersonInfo entiry of person
type MstPersonInfo struct {
	PersonCd           string
	PersonName         string
	PersonNameKana     string
	Birthday           time.Time
	Sex                string
	Zip                string
	Address1           string
	Address2           string
	Address3           string
	Address4           string
	Tel                string
	Mobile             string
	MailAddress        string
	AuthenticationDate time.Time
	DeleteFlag         string
	CreateUser         string
	CreateDate         time.Time
	UpdateUser         string
	UpdateDate         time.Time
}

//MstPersonInfoDB entity of db recores person because some field has null value
type MstPersonInfoDB struct {
	PersonCd           sql.NullString
	PersonName         sql.NullString
	PersonNameKana     sql.NullString
	Birthday           sql.NullTime
	Sex                sql.NullString
	Zip                sql.NullString
	Address1           sql.NullString
	Address2           sql.NullString
	Address3           sql.NullString
	Address4           sql.NullString
	Tel                sql.NullString
	Mobile             sql.NullString
	MailAddress        sql.NullString
	AuthenticationDate sql.NullTime
	DeleteFlag         sql.NullString
	CreateUser         sql.NullString
	CreateDate         sql.NullTime
	UpdateUser         sql.NullString
	UpdateDate         sql.NullTime
}