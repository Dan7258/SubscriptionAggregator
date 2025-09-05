package models

func (db PostgresDatabase) Migrate() error {
	return db.Conn.AutoMigrate(&Subscription{})
}
