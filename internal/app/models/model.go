package models

import "gorm.io/gorm"

type Model interface {
	GetID() int64
}

type AfterPost[T Model] interface {
	AfterPost(db *gorm.DB, origin T) error
}

type BeforePost[T Model] interface {
	BeforePost(db *gorm.DB, origin T) error
}
