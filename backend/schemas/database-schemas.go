package schemas

import "gorm.io/gorm"

type Database struct {
	Connection *gorm.DB
}
