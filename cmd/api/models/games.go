package models

import (
	"database/sql"

	"github.com/gustafer/go-games/cmd/api/database"
)

type Game struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func InsertGame(game *Game) (createdGameId int, err error) {
	db, err := database.OpenConn()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	query := `INSERT INTO games (title, description) VALUES ($1, $2) RETURNING id`

	err = db.QueryRow(query, game.Title, game.Description).Scan(&game.Id)

	if err != nil {
		return game.Id, err
	}

	return game.Id, nil
}

func DeleteGame(id string) (int64, error) {
	db, err := database.OpenConn()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := `DELETE FROM games WHERE id = $1`

	res, err := db.Exec(query, id)

	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func PatchGame(id string, game *Game) (updatedGameId int, err error) {
	db, err := database.OpenConn()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := `UPDATE games SET title = $2, description= $3 WHERE id = $1 RETURNING id`

	if err := db.QueryRow(query, id, game.Title, game.Description).Scan(&game.Id); err != nil {
		return game.Id, err
	}

	return game.Id, nil
}

func QueryAllGames() ([]Game, error) {
	db, err := database.OpenConn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM games ORDER BY id"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var games []Game

	for rows.Next() {
		var game Game

		if err := rows.Scan(&game.Id, &game.Title, &game.Description); err != nil {
			return games, err
		}
		games = append(games, game)
	}

	if err = rows.Err(); err != nil {
		return games, err
	}
	return games, nil
}

func GameById(id string) (Game, error) {
	var game Game

	db, err := database.OpenConn()
	if err != nil {
		return game, err
	}
	defer db.Close()

	query := `SELECT * FROM games WHERE id = $1`

	if err := db.QueryRow(query, id).Scan(&game.Id, &game.Title, &game.Description); err != nil {
		if err == sql.ErrNoRows {
			return game, err
		}
		return game, err
	}
	return game, nil
}
