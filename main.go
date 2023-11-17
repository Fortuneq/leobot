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

	db, err := sqlx.Open("mysql", "u2333338_root:Y3S4G9taZvwsYtU2@(31.31.196.165:3306)/u2333338_some")
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
