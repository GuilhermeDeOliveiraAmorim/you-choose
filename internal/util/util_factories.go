package util

import "gorm.io/gorm"

type ImputFactory struct {
	DB         *gorm.DB
	BucketName string
}
