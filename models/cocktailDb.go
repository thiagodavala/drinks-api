package models

import (
	"database/sql"
	"fmt"
	"os"
)

func AddCocktailDB(cocktail Cocktail) {

	db, err := sql.Open("mysql", getDbConnection())

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO cocktails (name,ingredients,method,garnish) VALUES (?,?,?,?)",
		cocktail.Name, cocktail.Ingredients, cocktail.Method, cocktail.Garnish)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func DeleteCocktail(id string) *Cocktail {

	db, err := sql.Open("mysql", getDbConnection())

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	_, err = db.Query("DELETE FROM cocktails where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	return &Cocktail{}
}

func GetCocktail(id string) *Cocktail {

	db, err := sql.Open("mysql", getDbConnection())

	ck := &Cocktail{}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM cocktails where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&ck.ID, &ck.Name, &ck.Garnish, &ck.Ingredients, &ck.Method)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return ck
}

func GetAllCocktails() []Cocktail {
	db, err := sql.Open("mysql", getDbConnection())

	// if there is an error opening the connection, handle it
	if err != nil {
		print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name, ingredients, method, garnish FROM cocktails")
	if err != nil {
		panic(err.Error())
	}

	var list []Cocktail

	for results.Next() {
		item := Cocktail{}
		err = results.Scan(&item.ID, &item.Name, &item.Ingredients, &item.Method, &item.Garnish)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, item)
	}
	return list
}

func getDbConnection() string {
	hostname := os.Getenv("DB_HOSTNAME")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	database := os.Getenv("DB_DATABASE")

	return user + ":" + password + "@tcp(" + hostname + ":3306)/" + database
}
