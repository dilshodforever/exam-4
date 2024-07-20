package postgres

// import (
// 	"database/sql"
// 	"fmt"

// 	"github.com/dilshodforever/4-oyimtixon-game-service/config"
// 	st "github.com/dilshodforever/4-oyimtixon-game-service/storage"

// 	_ "github.com/lib/pq"
// )

// type Storage struct {
// 	Db    *sql.DB
// 	Games st.GameStorage
// }

// func NewPostgresStorage() (st.InitRoot, error) {
// 	config := config.Load()
// 	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
// 		config.PostgresUser, config.PostgresPassword,
// 		config.PostgresHost, config.PostgresPort,
// 		config.PostgresDatabase)
// 	db, err := sql.Open("postgres", con)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Storage{Db: db, Games: &GameStorage{db}}, nil

// }

// func (s *Storage) Game() st.GameStorage {
// 	if s.Games == nil {
// 		s.Games = &GameStorage{s.Db}
// 	}
// 	return s.Games
// }
