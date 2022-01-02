// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dao

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type userTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *userTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("users").
func (v *userTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *userTableType) Columns() []string {
	return []string{
		"id",
		"uuid_id",
		"email",
		"password",
		"first_name",
		"second_name",
		"status",
		"role",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *userTableType) NewStruct() reform.Struct {
	return new(User)
}

// NewRecord makes a new record for that table.
func (v *userTableType) NewRecord() reform.Record {
	return new(User)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UserTable represents users view or table in SQL database.
var UserTable = &userTableType{
	s: parse.StructInfo{
		Type:    "User",
		SQLName: "users",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "UUIDID", Type: "string", Column: "uuid_id"},
			{Name: "Email", Type: "string", Column: "email"},
			{Name: "Password", Type: "string", Column: "password"},
			{Name: "FirstName", Type: "sql.NullString", Column: "first_name"},
			{Name: "SecondName", Type: "sql.NullString", Column: "second_name"},
			{Name: "Status", Type: "int64", Column: "status"},
			{Name: "Role", Type: "int64", Column: "role"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "*time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(User).Values(),
}

// String returns a string representation of this struct or record.
func (s User) String() string {
	res := make([]string, 10)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "UUIDID: " + reform.Inspect(s.UUIDID, true)
	res[2] = "Email: " + reform.Inspect(s.Email, true)
	res[3] = "Password: " + reform.Inspect(s.Password, true)
	res[4] = "FirstName: " + reform.Inspect(s.FirstName, true)
	res[5] = "SecondName: " + reform.Inspect(s.SecondName, true)
	res[6] = "Status: " + reform.Inspect(s.Status, true)
	res[7] = "Role: " + reform.Inspect(s.Role, true)
	res[8] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[9] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *User) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.UUIDID,
		s.Email,
		s.Password,
		s.FirstName,
		s.SecondName,
		s.Status,
		s.Role,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *User) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.UUIDID,
		&s.Email,
		&s.Password,
		&s.FirstName,
		&s.SecondName,
		&s.Status,
		&s.Role,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *User) View() reform.View {
	return UserTable
}

// Table returns Table object for that record.
func (s *User) Table() reform.Table {
	return UserTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *User) HasPK() bool {
	return s.ID != UserTable.z[UserTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *User) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = UserTable
	_ reform.Struct = (*User)(nil)
	_ reform.Table  = UserTable
	_ reform.Record = (*User)(nil)
	_ fmt.Stringer  = (*User)(nil)
)

func init() {
	parse.AssertUpToDate(&UserTable.s, new(User))
}
