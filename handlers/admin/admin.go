package AdminHandler

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	FCTemplateAdmin "github.com/snowheat/falconism/templates/admin"
	FCTypes "github.com/snowheat/falconism/types"
	"github.com/valyala/fasthttp"
)

func New(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")

	db, err := sql.Open("sqlite3", "./falconism.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := db.Query("SELECT rowid,post_title,post_content FROM posts")

	var blogposts = []FCTypes.BlogPost{}
	var rowid int
	var postTitle string
	var postContent string

	for rows.Next() {
		err = rows.Scan(&rowid, &postTitle, &postContent)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		blogposts = append(blogposts, FCTypes.BlogPost{ID: rowid, Title: postTitle, Content: postContent})

	}

	rows.Close() //good habit to close

	db.Close()

	FCTemplateAdmin.WritePageTemplate(ctx, blogposts)
}

func Post(ctx *fasthttp.RequestCtx) {
	//ctx.SetContentType("text/html; charset=utf-8")

	db, err := sql.Open("sqlite3", "./falconism.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stmt, err := db.Prepare("INSERT INTO posts(post_title, post_content) values(?,?)")
	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(string(ctx.PostArgs().Peek("title")), string(ctx.PostArgs().Peek("content")))
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(ctx.PostArgs().Peek("title")), string(ctx.PostArgs().Peek("content")))

	db.Close()
	ctx.Redirect("/admin", 302)
}
