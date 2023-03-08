package component

import "github.com/jmoiron/sqlx"

type AppContext interface {
	GetDBConn() *sqlx.DB
	SecretKey() string
}

type appContext struct {
	db        *sqlx.DB
	secretKey string
}

func NewAppContext(db *sqlx.DB, secretKey string) *appContext {
	return &appContext{
		db:        db,
		secretKey: secretKey,
	}
}

func (appCtx *appContext) GetDBConn() *sqlx.DB {
	return appCtx.db
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}
