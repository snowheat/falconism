package main

import (
	//"strconv"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-siris/siris"
	"github.com/go-siris/siris/context"
	"github.com/go-siris/siris/view"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if _, err := os.Stat("./falconism.db"); os.IsNotExist(err) {
		os.Create("./falconism.db")
	}

	db, err := sql.Open("sqlite3", "./falconism.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()

	/*
	   CREATE TABLE "posts" (
	   "post_datetime"  REAL,
	   "post_title"  TEXT,
	   "post_slug"  TEXT,
	   "post_content"  TEXT,
	   "ping_status"  INTEGER,
	   "to_ping"  INTEGER DEFAULT 1,
	   "sitemap_status"  INTEGER,
	   "to_sitemap"  INTEGER DEFAULT 1,
	   "meta_description"  TEXT
	   );
	*/

	app := siris.New()

	app.AttachView(view.HTML("./views", ".html").Reload(true))

	app.OnErrorCode(siris.StatusNotFound, notFoundHandler)

	app.Get("/", homeHandler)
	app.Get("/admin/{title:string regexp(^[a-z-0-9]+)}", adminHandler)
	app.Get("/app/{title:string regexp(^[a-z-0-9]+)}", appHandler)
	app.Get("/blog/{title:string regexp(^[a-z-0-9]+)}", blogHandler)
	app.Get("/blog/topic/{title:string regexp(^[a-z-0-9]+)}", topicHandler)
	app.Get("/blog/archive/{year:string regexp(^[a-z-0-9]+)}", archiveHandler)

	if err := app.Run(siris.Addr(":1212"), siris.WithCharset("UTF-8")); err != nil {
		panic(err)
	}
}

func homeHandler(ctx context.Context) {

	var kota = []string{"Broklin", "Jumhud", "Bahchinok", "China Daratan Islamabad"}
	var posts []Post

	posts = append(posts, Post{
		Title:   "Suratang",
		Content: "Suatu ketika",
		Date:    "2017-07-15",
	})

	posts = append(posts, Post{
		Title:   "Pergi ke pulau kemah",
		Content: "jaman dulu saya naik tiram",
		Date:    "2017-07-14",
	})

	ctx.ViewData("Post", posts)
	ctx.ViewData("Kota", kota)
	ctx.ViewData("Username", "insan")
	ctx.View("home.html")
}

func adminHandler(ctx context.Context) {
	ctx.HTML("Admin")
}

func blogHandler(ctx context.Context) {
	ctx.HTML("Blog")
}

func topicHandler(ctx context.Context) {
	ctx.HTML("Topic")
}

func archiveHandler(ctx context.Context) {
	ctx.HTML("Archive")
}

func appHandler(ctx context.Context) {
	ctx.HTML("App")
}

func notFoundHandler(ctx context.Context) {
	ctx.HTML("Page not found")
}

type Post struct {
	Title   string
	Content string
	Date    string
}
