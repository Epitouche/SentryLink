package schemas

import "gorm.io/gorm"

// @Ignore
type Database struct {
	Connection *gorm.DB
}
