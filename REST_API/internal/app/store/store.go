package store

import (
	"database/sql"
	
	_ "go get github.com/lib/pq" // импортируем без методов
)
//Store
type Store struct {
	config *Config
	db     *sql.DB
	userRepository *UserRepository
}

// New Store
func New(config *Config) *Store {
	return &Config{
		config: config,
	}
}

/*make; ./apiserver

createdb restapi_dev

migrate create -ext sql -dir migrations create_users
migrate -path migrations - database "postgres://localhost/restapi_dev?sslmode=disable" up 
check table
psql -d restapi_dev
	\dt
	\d users;
	\q	
*/

//Open DB
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

//Close DB
func (s *Store) Close() {
	s.db.Close()
}

//User Repository
// Connet API store.User().Create( )
func (s *Store) User() *UserRepository{
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}