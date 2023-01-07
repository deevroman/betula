// Command betula runs Betula, a personal link collection software.
package main

import (
	"fmt"
	"git.sr.ht/~bouncepaw/betula/auth"
	"git.sr.ht/~bouncepaw/betula/web"
	"log"
	"os"
	"path/filepath"

	"git.sr.ht/~bouncepaw/betula/db"
	_ "git.sr.ht/~bouncepaw/betula/web" // For init()

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello Betula!")

	if len(os.Args) < 2 {
		log.Fatalln("Pass a database file name!")
	}

	filename, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	db.Initialize(filename)
	defer db.Finalize()
	auth.Initialize()

	// TODO: make it configurable
	web.Start()
}
