package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Entry struct {
	ID          int
	Name        string
	Description string
}

func main() {
	db, err := sql.Open("postgres", "user=user password=Ulyasar10389# dbname=user sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS entries (id SERIAL PRIMARY KEY, name TEXT, description TEXT)")
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		entries := []Entry{}
		rows, err := db.Query("SELECT * FROM entries")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var entry Entry
			err := rows.Scan(&entry.ID, &entry.Name, &entry.Description)
			if err != nil {
				panic(err)
			}
			entries = append(entries, entry)
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"entries": entries,
		})
	})

	router.POST("/add", func(c *gin.Context) {
		name := c.PostForm("name")
		description := c.PostForm("description")

		_, err := db.Exec("INSERT INTO entries (name, description) VALUES ($1, $2)", name, description)
		if err != nil {
			panic(err)
		}

		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/edit/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		name := c.PostForm("name")
		description := c.PostForm("description")

		_, err = db.Exec("UPDATE entries SET name=$1, description=$2 WHERE id=$3", name, description, id)
		if err != nil {
			panic(err)
		}

		c.Redirect(http.StatusFound, "/")
	})

	router.POST("/delete/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("DELETE FROM entries WHERE id=$1", id)
		if err != nil {
			panic(err)
		}

		c.Redirect(http.StatusFound, "/")
	})

	router.Run(":8080")
}
