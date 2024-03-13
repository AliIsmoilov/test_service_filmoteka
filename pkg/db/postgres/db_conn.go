package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver

	"test_service_filmoteka/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance
func NewPsqlDB(c *config.Config) (*gorm.DB, error) {
	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	// 	c.Postgres.PostgresqlHost,
	// 	c.Postgres.PostgresqlPort,
	// 	c.Postgres.PostgresqlUser,
	// 	c.Postgres.PostgresqlDbname,
	// 	c.Postgres.PostgresqlPassword,
	// )

	// db, err := sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	// if err != nil {
	// 	return nil, err
	// }

	// db.SetMaxOpenConns(maxOpenConns)
	// db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	// db.SetMaxIdleConns(maxIdleConns)
	// db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	// -------------------------------------------------------------------
	// -----------------------------------------------------------------
	fmt.Println("init db")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlPassword,
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlDbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * 1, // Slow SQL threshold
			// LogLevel:                  logger.Info,                                               // For debugging you can set Info level
			LogLevel: logger.Error, // For debugging you can set Info level
			// IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      false,        // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get database", err)
	} else {
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetMaxIdleConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	return db, nil
}
