package database

import (
	"fmt"
	"strings"
	"sync"

	"github.com/enorith/http/contracts"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
)

type OpFunc func(db *gorm.DB, col string, val interface{}) *gorm.DB

var (
	customFilterOps = make(map[string]OpFunc)
	mu              sync.RWMutex
)

func WithCustomFilter(op string, opFunc OpFunc) {
	mu.Lock()
	defer mu.Unlock()
	customFilterOps[op] = opFunc
}

func GetCustomFilter(op string) (OpFunc, bool) {
	mu.RLock()
	defer mu.RUnlock()
	fn, ok := customFilterOps[op]
	return fn, ok
}

func WithFilters(db *gorm.DB, req contracts.RequestContract, key ...string) *gorm.DB {
	k := "filters"
	if len(key) > 0 {
		k = key[0]
	}
	raw := req.Get(k)
	var filters map[string]interface{}
	if raw != nil {
		jsoniter.Unmarshal(raw, &filters)
		db = ApplyFilters(db, filters)
	}

	return db
}

func ApplyFilters(db *gorm.DB, filters map[string]interface{}, defOp ...string) *gorm.DB {
	for col, val := range filters {
		parts := strings.Split(col, ",")
		c := parts[0]
		op := "like"
		if len(defOp) > 0 {
			op = defOp[0]
		}

		if len(parts) > 1 {
			op = parts[1]
		}
		if op == "like" {
			val = fmt.Sprintf("%%%s%%", val)
		}
		if fn, ok := GetCustomFilter(op); ok {
			db = fn(db, c, val)
		} else {
			db = db.Where(fmt.Sprintf("%s %s ?", c, op), val)
		}
	}

	return db
}

func WithEnabledVersion(db *gorm.DB, table ...string) *gorm.DB {
	var t string
	if len(table) > 0 {
		t = table[0] + "."
	}
	return db.Where(fmt.Sprintf("(%senabled = 1 OR %sversion = 0)", t, t))
}

func Exists(tx *gorm.DB) (exists bool) {
	db := tx.Session(&gorm.Session{
		NewDB: true,
	})

	db.Raw("SELECT EXISTS(?)", tx).Scan(&exists)

	return
}

func WithSorts(db *gorm.DB, req contracts.RequestContract, key ...string) *gorm.DB {
	k := "sort"
	if len(key) > 0 {
		k = key[0]
	}
	raw := req.Get(k)
	var sorts map[string]string
	if raw != nil {
		jsoniter.Unmarshal(raw, &sorts)
		db = ApplySorts(db, sorts, "id desc")
	}

	return db
}

func ApplySorts(db *gorm.DB, sorts map[string]string, defSort ...string) *gorm.DB {

	if len(defSort) > 0 && len(sorts) == 0 {
		db = db.Order(defSort[0])
	}

	for col, val := range sorts {
		db = db.Order(fmt.Sprintf("%s %s", col, val))
	}
	return db
}
