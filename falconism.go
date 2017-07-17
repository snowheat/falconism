package main

//go:generate qtc -dir=./templates

import (
	"github.com/buaazp/fasthttprouter"
	FCRouter "github.com/snowheat/falconism/router"
	FCDb "github.com/snowheat/falconism/system/db"
	FCServer "github.com/snowheat/falconism/system/server"
	"github.com/valyala/fasthttp"
)

func main() {
	FCDb.Init()

	router := fasthttprouter.New()
	server := fasthttp.Server{}

	FCRouter.Set(router)
	FCServer.Set(&server, router)
	FCServer.Run(&server)
}

/*
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
	createDatabaseIfNotExist()
	displayFalconismBanner()

	app := siris.New()
	siris.WithoutBanner(app)

	app.AttachView(view.HTML("./views", ".html") .Reload(true))

	app.OnErrorCode(siris.StatusNotFound, notFoundHandler)

	app.Get("/", homeHandler)
	adminRoutes := app.Party("/admin")
	adminRoutes.Get("/", adminIndexHandler)
	adminRoutes.Get("/post", adminPostHandler)

	app.Get("/blog/{title:string regexp(^[a-z-0-9]+)}", blogHandler)
	app.Get("/blog/topic/{title:string regexp(^[a-z-0-9]+)}", topicHandler)
	app.Get("/blog/archive/{year:string regexp(^[a-z-0-9]+)}", archiveHandler)

	if err := app.Run(siris.Addr(":1212"), siris.WithCharset("UTF-8")); err != nil {
		panic(err)
	}
}

func createDatabaseIfNotExist() {
	if _, err := os.Stat("./falconism.db"); os.IsNotExist(err) {
		os.Create("./falconism.db")
	}

	db, err := sql.Open("sqlite3", "./falconism.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
	   "post_datetime"  REAL,
	   "post_title"  TEXT,
	   "post_slug"  TEXT,
	   "post_content"  TEXT,
	   "ping_status"  INTEGER,
	   "to_ping"  INTEGER DEFAULT 1,
	   "sitemap_status"  INTEGER,
	   "to_sitemap"  INTEGER DEFAULT 1,
	   "meta_description"  TEXT
	);`)

	_, err = db.Exec(`
		CREATE TABLE "config" (
		"parameter"  TEXT,
		"value"  TEXT
		);
	`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()
}

func displayFalconismBanner() {
	fmt.Println(`
     ______      __                 _
    / ____/___ _/ /________  ____  (_)________ ___
   / /_  / __ '/ / ___/ __ \/ __ \/ / ___/ __ '__ \
  / __/ / /_/ / / /__/ /_/ / / / / (__  ) / / / / /
 /_/    \__,_/_/\___/\____/_/ /_/_/____/_/ /_/ /_/

 *** Ultra-minimalist blog engine optimized for cheap VPS ***

 Powered by Go & Go-Siris

 Now listening on: http://localhost:1212
 Blog admin: http://localhost:1212/admin

 Application started. Press CTRL+C to shut down.
	`)
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

func adminIndexHandler(ctx context.Context) {
	ctx.View("admin/index.html")
}

func adminPostHandler(ctx context.Context) {
	ctx.View("admin/post.html")
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

*/
