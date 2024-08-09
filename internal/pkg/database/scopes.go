package database

import (
	"github.com/enorith/supports/collection"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Scope func(*gorm.DB) *gorm.DB

type UpsertOpt func(clause.OnConflict) clause.OnConflict

func UpsertOptColumns(cols ...string) UpsertOpt {
	return func(on clause.OnConflict) clause.OnConflict {
		on.Columns = collection.Map(cols, func(col string) clause.Column {
			return clause.Column{Name: col}
		})
		return on
	}
}

func UpsertOptUpdateColumns(cols ...string) UpsertOpt {
	return func(on clause.OnConflict) clause.OnConflict {
		on.DoUpdates = clause.AssignmentColumns(cols)
		return on
	}
}

func UpsertOptUpdateAssign(assigns map[string]interface{}) UpsertOpt {
	return func(on clause.OnConflict) clause.OnConflict {
		on.DoUpdates = clause.Assignments(assigns)
		return on
	}
}

func WithUpsert(opts ...UpsertOpt) Scope {
	return func(db *gorm.DB) *gorm.DB {
		var on clause.OnConflict

		if len(opts) == 0 {
			on.DoNothing = true
		}
		for _, opt := range opts {
			on = opt(on)
		}

		return db.Clauses(on)
	}
}

func WithHas(relation string, fn Scope) Scope {
	return func(db *gorm.DB) *gorm.DB {
		asso := db.Session(&gorm.Session{}).Association(relation)
		if asso.Error == nil {
			refs := asso.Relationship.References
			newDB := db.Session(&gorm.Session{NewDB: true})
			switch asso.Relationship.Type {
			case schema.BelongsTo:
				if len(refs) > 0 {
					db.Where(refs[0].ForeignKey.DBName+" IN (?)", newDB.Table(asso.Relationship.FieldSchema.Table).Scopes(fn).Select(refs[0].PrimaryKey.DBName))
				}
			}
		}

		return db
	}
}
