package db

import (
	"time"
)

/*
Describe table columns like the following:
	ID           int `gorm:"primary_key"`
	Subscribed   bool
	Birthday     time.Time
    Age          int
    Name         string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Num          int     `sql:"AUTO_INCREMENT"`
    UserID       int     `sql:"index"` // tag `index` will create index for this field when using AutoMigrate
    Email        string  `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
    Name         string `sql:"index:idx_name_code"` // Will create combined index if find other fields defined same name
    Code         string `sql:"index:idx_name_code"` // `unique_index` also works
*/

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
