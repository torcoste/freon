package dao

import (
	"time"

	"github.com/AlekSi/pointer"
)

// go:generate reform
// reform:translations
type Translation struct {
	ID             int64      `reform:"id,pk"`
	LocalizationID int64      `reform:"localization_id"`
	IdentifierID   int64      `reform:"identifier_id"`
	CreatorID      int64      `reform:"creator_id"`
	Text           string     `reform:"text"`
	Status         int64      `reform:"status"`
	CreatedAt      time.Time  `reform:"created_at"`
	UpdatedAt      *time.Time `reform:"updated_at"`

	Localization *Localization
	Identifier   *Identifier
	Creator      *User
}

func (u *Translation) BeforeInsert() error {
	if u.UpdatedAt != nil {
		u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	}
	return nil
}

func (u *Translation) BeforeUpdate() error {
	now := time.Now().UTC()
	u.UpdatedAt = &now
	u.UpdatedAt = pointer.ToTime(u.UpdatedAt.UTC().Truncate(time.Second).AddDate(0, 0, 0))
	return nil
}
