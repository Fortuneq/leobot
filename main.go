package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v3"
	"leomineTg/hanleFunctions"
	"leomineTg/repository"
	"log"
	"time"
)

func main() {
	pref := tele.Settings{
		Token:  "6016884724:AAGci95V692aspTxZzvejeDiWWFJdWbYZkg",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	db, err := sqlx.Open("mysql", "root:dCmd5e5A6hUN8Yv@(193.109.84.90:3306)/leomine_schema")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repository.NewIndexRepository(db)

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	hanleFunctions.Handle(b, repo)

	b.Start()
}
