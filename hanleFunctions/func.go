package hanleFunctions

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/looplab/fsm"
	tele "gopkg.in/telebot.v3"
	"image"
	"image/jpeg"
	"io"
	"leomineTg/repository"
	"log"
	"os"
	"strings"
)

var PreviousState = "default"

var id = ""

type Chat struct {
	To  string
	FSM *fsm.FSM
}

const (
	AddDevice                 = "addDevice"
	AddDeviceName             = "addDeviceName"
	AddDeviceCost             = "addDeviceCost"
	AddDeviceImage            = "addDeviceImage"
	AddDeviceSize             = "addDeviceSize"
	AddDevicePower            = "addDevicePower"
	AddDeviceHashrate         = "addDeviceHashrate"
	AddDeviceAlgorithm        = "addDeviceAlgorithm"
	AddDeviceUid              = "addDeviceUid"
	AddDeviceVideoUrl         = "addDeviceVideoUrl"
	RemoveDevice              = "removeDevice"
	RemoveDeviceID            = "removeDeviceArticleRemoveArticle"
	RemoveCase                = "removeCase"
	AddReview                 = "addReview"
	RemoveReview              = "removeReview"
	AddDeviceReview           = "addDeviceReview"
	RemoveDeviceReview        = "removeDeviceReview"
	EditDevice                = "editDevice"
	editDeviceID              = "editDeviceId"
	EditDeviceName            = "editDeviceName"
	EditDeviceCost            = "editDeviceCost"
	EditDevicePopularity      = "editDevicePopularity"
	EditDeviceImage           = "editDeviceImage"
	EditDeviceSize            = "editDeviceSize"
	EditDevicePower           = "editDevicePower"
	EditDeviceHashrate        = "editDeviceHashrate"
	EditDeviceAlgorithm       = "editDeviceAlgorithm"
	EditDeviceUid             = "editDeviceUid"
	EditDeviceVideoUrl        = "editDeviceVideoUrl"
	EditArticle               = "editArticle"
	EditArticleID             = "editArticleID"
	EditArticleName           = "editArticleName"
	EditArticleText           = "editArticleText"
	EditArticleVideoUrl       = "editArticleVideoURl"
	EditCase                  = "editCase"
	EditCaseID                = "editCaseID"
	EditCaseName              = "editCaseName"
	EditCaseText              = "editCaseText"
	EditCaseVideoUrl          = "editCaseVideoURl"
	AddCase                   = "addCase"
	AddCaseID                 = "addCaseID"
	AddCaseName               = "addCaseName"
	AddCaseText               = "addCaseText"
	AddCaseVideoUrl           = "addCaseVideoURl"
	RemoveCaseID              = "removeCaseID"
	AddArticle                = "addArticle"
	RemoveArticle             = "removeArticle"
	AddArticleID              = "addArticleID"
	AddArticleName            = "addArticleName"
	AddArticleText            = "addArticleText"
	AddArticleVideoUrl        = "addArticleVideoURl"
	RemoveArticleID           = "removeArticleID"
	EditCompanyReviews        = "editCompanyReviews"
	EditCompanyReviewsID      = "editCompanyReviewsID"
	EditCompanyReviewsText    = "editCompanyReviewsText"
	EditCompanyReviewsName    = "editCompanyReviewsName"
	EditCompanyReviewsEmail   = "editCompanyReviewsEmail"
	EditCompanyReviewsStars   = "editCompanyReviewsStars"
	RemoveCompanyReviews      = "RemoveCompanyReviews"
	RemoveCompanyReviewsID    = "RemoveCompanyReviewsID"
	AddCompanyReviews         = "AddCompanyReviews"
	AddCompanyReviewsID       = "AddCompanyReviewsID"
	AddCompanyReviewsText     = "AddCompanyReviewsText"
	AddCompanyReviewsName     = "AddCompanyReviewsName"
	AddCompanyReviewsEmail    = "AddCompanyReviewsEmail"
	AddCompanyReviewsStars    = "AddCompanyReviewsStars"
	EditDeviceReviews         = "editDeviceReviews"
	EditDeviceReviewsID       = "editDeviceReviewsID"
	EditDeviceReviewsText     = "editDeviceReviewsText"
	EditDeviceReviewsName     = "editDeviceReviewsName"
	EditDeviceReviewsEmail    = "editDeviceReviewsEmail"
	EditDeviceReviewsStars    = "editDeviceReviewsStars"
	EditDeviceReviewsAmount   = "editDeviceReviewsAmount"
	EditDeviceReviewsDate     = "editDeviceReviewsDate"
	EditDeviceReviewsDeviceID = "editDeviceReviewsDeviceID"
	AddDeviceReviews          = "addDeviceReviews"
	AddDeviceReviewsID        = "addDeviceReviewsID"
	RemoveDeviceReviews       = "removeDeviceReviews"
	RemoveDeviceReviewsID     = "removeDeviceReviewsID"
	AddDeviceReviewsText      = "addDeviceReviewsText"
	AddDeviceReviewsName      = "addDeviceReviewsName"
	AddDeviceReviewsEmail     = "addDeviceReviewsEmail"
	AddDeviceReviewsStars     = "addDeviceReviewsStars"
	AddDeviceReviewsAmount    = "addDeviceReviewsAmount"
	AddDeviceReviewsDate      = "addDeviceReviewsDate"
	AddDeviceReviewsDeviceID  = "addDeviceReviewsDeviceID"
)

func NewChat(to string) *Chat {
	d := &Chat{
		To: to,
	}

	deviceArr := []string{
		EditDeviceAlgorithm, EditDeviceName, EditDevicePower, EditDeviceHashrate, EditDevicePopularity, EditDeviceSize, EditDeviceUid,
		EditDeviceCost, EditDeviceImage, EditDeviceVideoUrl, editDeviceID, EditDevice, EditArticle, EditArticleText, EditArticleName, EditArticleVideoUrl, EditArticleID, EditDevice, EditCaseText, EditCaseID, EditCase, EditCaseName, EditCaseVideoUrl,
		"default", EditCompanyReviews, EditCompanyReviewsID, EditCompanyReviewsText, EditCompanyReviewsName, EditCompanyReviewsEmail, EditCompanyReviewsStars,
		EditDeviceReviews, EditDeviceReviewsID, EditDeviceReviewsText, EditDeviceReviewsName, EditDeviceReviewsEmail, EditDeviceReviewsStars, EditDeviceReviewsAmount, EditDeviceReviewsDate, EditDeviceReviewsDeviceID,
		AddDevice,
		AddDeviceName,
		AddDeviceCost,
		AddDeviceImage,
		AddDeviceSize,
		AddDevicePower,
		AddDeviceHashrate,
		AddDeviceAlgorithm,
		AddDeviceUid,
		AddDeviceVideoUrl,
		RemoveDevice,
		AddArticleID,
		AddArticleName,
		AddArticleText,
		AddArticleVideoUrl,
		RemoveArticleID,
		AddArticle,
		RemoveArticle,
		AddCase,
		AddCaseID,
		AddCaseName,
		AddCaseText,
		AddCaseVideoUrl,
		RemoveCaseID,
		RemoveCase,
		AddReview,
		RemoveCompanyReviews,
		RemoveCompanyReviewsID,
		AddCompanyReviews,
		AddCompanyReviewsID,
		AddCompanyReviewsText,
		AddCompanyReviewsName,
		AddCompanyReviewsEmail,
		AddCompanyReviewsStars,
		RemoveReview,
		AddDeviceReview,
		AddDeviceReviews,
		AddDeviceReviewsID,
		RemoveDeviceReviews,
		RemoveDeviceReviewsID,
		AddDeviceReviewsText,
		AddDeviceReviewsName,
		AddDeviceReviewsEmail,
		AddDeviceReviewsStars,
		AddDeviceReviewsAmount,
		AddDeviceReviewsDate,
		AddDeviceReviewsDeviceID,
		RemoveDeviceReview}
	d.FSM = fsm.NewFSM(
		"default",
		fsm.Events{
			{Name: EditDevice, Src: deviceArr, Dst: EditDevice},
			{Name: editDeviceID, Src: deviceArr, Dst: editDeviceID},
			{Name: EditDeviceAlgorithm, Src: deviceArr, Dst: EditDeviceAlgorithm},
			{Name: EditDeviceName, Src: deviceArr, Dst: EditDeviceName},
			{Name: EditDevicePower, Src: deviceArr, Dst: EditDevicePower},
			{Name: EditDeviceHashrate, Src: deviceArr, Dst: EditDeviceHashrate},
			{Name: EditDevicePopularity, Src: deviceArr, Dst: EditDevicePopularity},
			{Name: EditDeviceSize, Src: deviceArr, Dst: EditDeviceSize},
			{Name: EditDeviceUid, Src: deviceArr, Dst: EditDeviceUid},
			{Name: EditDeviceCost, Src: deviceArr, Dst: EditDeviceCost},
			{Name: EditDeviceImage, Src: deviceArr, Dst: EditDeviceImage},
			{Name: EditDeviceVideoUrl, Src: deviceArr, Dst: EditDeviceVideoUrl},
			{Name: EditArticle, Src: deviceArr, Dst: EditArticle},
			{Name: EditArticleText, Src: deviceArr, Dst: EditArticleText},
			{Name: EditArticleName, Src: deviceArr, Dst: EditArticleName},
			{Name: EditArticleVideoUrl, Src: deviceArr, Dst: EditArticleVideoUrl},
			{Name: EditArticleID, Src: deviceArr, Dst: EditArticleID},
			{Name: EditCase, Src: deviceArr, Dst: EditCase},
			{Name: EditCaseText, Src: deviceArr, Dst: EditCaseText},
			{Name: EditCaseName, Src: deviceArr, Dst: EditCaseName},
			{Name: EditCaseVideoUrl, Src: deviceArr, Dst: EditCaseVideoUrl},
			{Name: EditCaseID, Src: deviceArr, Dst: EditCaseID},
			{Name: EditCompanyReviews, Src: deviceArr, Dst: EditCompanyReviews},
			{Name: EditCompanyReviewsID, Src: deviceArr, Dst: EditCompanyReviewsID},
			{Name: EditCompanyReviewsText, Src: deviceArr, Dst: EditCompanyReviewsText},
			{Name: EditCompanyReviewsEmail, Src: deviceArr, Dst: EditCompanyReviewsEmail},
			{Name: EditCompanyReviewsName, Src: deviceArr, Dst: EditCompanyReviewsName},
			{Name: EditCompanyReviewsStars, Src: deviceArr, Dst: EditCompanyReviewsStars},
			{Name: EditDeviceReviews, Src: deviceArr, Dst: EditDeviceReviews},
			{Name: EditDeviceReviewsID, Src: deviceArr, Dst: EditDeviceReviewsID},
			{Name: EditDeviceReviewsText, Src: deviceArr, Dst: EditDeviceReviewsText},
			{Name: EditDeviceReviewsEmail, Src: deviceArr, Dst: EditDeviceReviewsEmail},
			{Name: EditDeviceReviewsName, Src: deviceArr, Dst: EditDeviceReviewsName},
			{Name: EditDeviceReviewsStars, Src: deviceArr, Dst: EditDeviceReviewsStars},
			{Name: EditDeviceReviewsAmount, Src: deviceArr, Dst: EditDeviceReviewsAmount},
			{Name: EditDeviceReviewsDeviceID, Src: deviceArr, Dst: EditDeviceReviewsDeviceID},
			{Name: EditDeviceReviewsDate, Src: deviceArr, Dst: EditDeviceReviewsDate},
			{Name: AddDevice, Src: deviceArr, Dst: AddDevice},
			{Name: RemoveDevice, Src: deviceArr, Dst: RemoveDevice},
			{Name: AddArticle, Src: deviceArr, Dst: AddArticle},
			{Name: RemoveArticle, Src: deviceArr, Dst: RemoveArticle},
			{Name: AddCase, Src: deviceArr, Dst: AddCase},
			{Name: RemoveCase, Src: deviceArr, Dst: RemoveCase},
			{Name: AddReview, Src: deviceArr, Dst: AddReview},
			{Name: RemoveReview, Src: deviceArr, Dst: RemoveReview},
			{Name: AddDeviceReview, Src: deviceArr, Dst: AddDeviceReview},
			{Name: RemoveDeviceReview, Src: deviceArr, Dst: RemoveDeviceReview},
			{Name: AddDeviceName, Src: deviceArr, Dst: AddDeviceName},
			{Name: AddDeviceCost, Src: deviceArr, Dst: AddDeviceCost},
			{Name: AddDeviceImage, Src: deviceArr, Dst: AddDeviceImage},
			{Name: AddDeviceSize, Src: deviceArr, Dst: AddDeviceSize},
			{Name: AddDevicePower, Src: deviceArr, Dst: AddDevicePower},
			{Name: AddDeviceHashrate, Src: deviceArr, Dst: AddDeviceHashrate},
			{Name: AddDeviceAlgorithm, Src: deviceArr, Dst: AddDeviceAlgorithm},
			{Name: AddDeviceUid, Src: deviceArr, Dst: AddDeviceUid},
			{Name: AddDeviceVideoUrl, Src: deviceArr, Dst: AddDeviceVideoUrl},
			{Name: AddCaseID, Src: deviceArr, Dst: AddCaseID},
			{Name: AddCaseName, Src: deviceArr, Dst: AddCaseName},
			{Name: AddCaseText, Src: deviceArr, Dst: AddCaseText},
			{Name: AddCaseVideoUrl, Src: deviceArr, Dst: AddCaseVideoUrl},
			{Name: RemoveCaseID, Src: deviceArr, Dst: RemoveCaseID},
			{Name: AddArticleID, Src: deviceArr, Dst: AddArticleID},
			{Name: AddArticleName, Src: deviceArr, Dst: AddArticleName},
			{Name: AddArticleText, Src: deviceArr, Dst: AddArticleText},
			{Name: AddArticleVideoUrl, Src: deviceArr, Dst: AddArticleVideoUrl},
			{Name: RemoveArticleID, Src: deviceArr, Dst: RemoveArticleID},
			{Name: RemoveCompanyReviews, Src: deviceArr, Dst: RemoveCompanyReviews},
			{Name: RemoveCompanyReviewsID, Src: deviceArr, Dst: RemoveCompanyReviewsID},
			{Name: AddCompanyReviews, Src: deviceArr, Dst: AddCompanyReviews},
			{Name: AddCompanyReviewsID, Src: deviceArr, Dst: AddCompanyReviewsID},
			{Name: AddCompanyReviewsText, Src: deviceArr, Dst: AddCompanyReviewsText},
			{Name: AddCompanyReviewsName, Src: deviceArr, Dst: AddCompanyReviewsName},
			{Name: AddCompanyReviewsEmail, Src: deviceArr, Dst: AddCompanyReviewsEmail},
			{Name: AddCompanyReviewsStars, Src: deviceArr, Dst: AddCompanyReviewsStars},
			{Name: AddDeviceReviews, Src: deviceArr, Dst: AddDeviceReviews},
			{Name: AddDeviceReviewsID, Src: deviceArr, Dst: AddDeviceReviewsID},
			{Name: RemoveDeviceReviews, Src: deviceArr, Dst: RemoveDeviceReviews},
			{Name: RemoveDeviceReviewsID, Src: deviceArr, Dst: RemoveDeviceReviewsID},
			{Name: AddDeviceReviewsText, Src: deviceArr, Dst: AddDeviceReviewsText},
			{Name: AddDeviceReviewsName, Src: deviceArr, Dst: AddDeviceReviewsName},
			{Name: AddDeviceReviewsEmail, Src: deviceArr, Dst: AddDeviceReviewsEmail},
			{Name: AddDeviceReviewsStars, Src: deviceArr, Dst: AddDeviceReviewsStars},
			{Name: AddDeviceReviewsAmount, Src: deviceArr, Dst: AddDeviceReviewsAmount},
			{Name: AddDeviceReviewsDate, Src: deviceArr, Dst: AddDeviceReviewsDate},
			{Name: AddDeviceReviewsDeviceID, Src: deviceArr, Dst: AddDeviceReviewsDeviceID},
		},
		fsm.Callbacks{"enter_state": func(_ context.Context, e *fsm.Event) { d.enterState(e) }},
	)

	return d
}

func (d *Chat) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func Handle(b *tele.Bot, repo *repository.IndexRepository) {
	review := repository.DeviceReviews{}
	device := repository.Devices{}
	cas := repository.Cases{}
	compReview := repository.Reviews{}
	article := repository.Articles{}

	state := NewChat("chat")

	b.Handle("/get_reviews", func(c tele.Context) error {
		reviews, err := repo.GetAllReviews(context.Background())
		if err != nil {
			return c.Send(err.Error())
		}
		fmt.Println(reviews)
		for _, v := range reviews {
			s := fmt.Sprintf("Айди: %s\nПочта: %s\nТекст: %s\nИмя: %s\nЗвезды: %s\n", v.Id, v.Email, v.Text, v.Name, v.Stars)
			c.Send(s)
		}
		return nil
	})

	b.Handle("/start", func(c tele.Context) error {

		return c.Send("Данный бот создан для админов leomine для редактирования элементов сайта , команды можно посмотреть слева")
	})

	b.Handle("/get_device_reviews", func(c tele.Context) error {
		reviews, err := repo.GetDeviceReviews(context.Background())
		if err != nil {
			return c.Send(err.Error())
		}
		fmt.Println(reviews)
		for _, v := range reviews {
			s := fmt.Sprintf("Айди: %s\nПочта: %s\nТекст: %s\nИмя: %s\nЗвезды: %s\nКоличество: %s\nДата: %s\nАйдиДевайса: %s\n", v.Id, v.Email, v.Text, v.Name, v.Stars, v.Amount, v.Date, v.DeviceId)
			c.Send(s)
		}
		return nil
	})

	b.Handle("/edit_device", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDevice)
		if err != nil {
			fmt.Println(err)
		}
		keyboardDevice(state, c, b, repo)
		return nil
	})

	b.Handle("/edit_article", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditArticle)
		if err != nil {
			fmt.Println(err)
		}
		keyboardArticle(state, c, b)
		return nil
	})

	b.Handle("/edit_case", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCase)
		if err != nil {
			fmt.Println(err)
		}
		keyboardCase(state, c, b)
		return nil
	})

	b.Handle("/edit_company_reviews", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCompanyReviews)
		if err != nil {
			fmt.Println(err)
		}
		keyboardCompanyReviews(state, c, b)
		return nil
	})

	b.Handle("/edit_device_reviews", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviews)
		if err != nil {
			fmt.Println(err)
		}
		keyboardDeviceReviews(state, c, b)
		return nil
	})

	b.Handle("/add_device", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), AddDevice)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите Имя")
		err = state.FSM.Event(context.Background(), AddDeviceName)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/remove_device", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), RemoveDevice)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите айди девайса")
		err = state.FSM.Event(context.Background(), RemoveDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/add_case", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), AddCase)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите Имя")
		err = state.FSM.Event(context.Background(), AddCaseName)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/remove_case", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), RemoveCase)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите айди кейса")
		err = state.FSM.Event(context.Background(), RemoveCaseID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	b.Handle("/add_article", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), AddArticle)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите Имя")
		err = state.FSM.Event(context.Background(), AddArticleName)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/remove_article", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), RemoveArticle)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите айди статьи")
		err = state.FSM.Event(context.Background(), RemoveArticleID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	b.Handle("/add_company_review", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), AddCompanyReviews)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите Имя")
		err = state.FSM.Event(context.Background(), AddCompanyReviewsName)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/remove_company_review", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), RemoveCompanyReviews)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите айди статьи")
		err = state.FSM.Event(context.Background(), RemoveCompanyReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	b.Handle("/add_device_review", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), AddDeviceReview)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите Имя")
		err = state.FSM.Event(context.Background(), AddDeviceReviewsName)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	b.Handle("/remove_device_review", func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), RemoveDeviceReviews)
		if err != nil {
			fmt.Println(err)
		}
		c.Send("Пришлите айди статьи")
		err = state.FSM.Event(context.Background(), RemoveDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	b.Handle("/get_cases", func(c tele.Context) error {
		reviews, err := repo.GetCases(context.Background())
		if err != nil {
			return c.Send(err.Error())
		}
		fmt.Println(reviews)
		for _, v := range reviews {
			s := fmt.Sprintf("Айди: %s\nТекст: %s\nИмя: %s\nВидео Ссылка: %s\n", v.Id, v.Text, v.Name, v.VideoUrl)
			c.Send(s)
		}
		return nil
	})

	b.Handle("/get_articles", func(c tele.Context) error {
		reviews, err := repo.GetArticles(context.Background())
		if err != nil {
			return c.Send(err.Error())
		}
		fmt.Println(reviews)
		for _, v := range reviews {
			s := fmt.Sprintf("Айди: %s\nТекст: %s\nИмя: %s\nВидео Ссылка: %s\n", v.Id, v.Text, v.Name, v.VideoUrl)
			c.Send(s)
		}
		return nil
	})

	b.Handle("/get_devices", func(c tele.Context) error {
		reviews, err := repo.GetAllDevices(context.Background())
		if err != nil {
			return c.Send(err.Error())
		}
		fmt.Println(reviews)
		for _, v := range reviews {
			s := fmt.Sprintf("Айди: %s\nИмя: %s\nВидео Ссылка: %s\nСтоимость: %s\nАлгоритм: %s\nРазмер: %s\nХэшрейт: %s\nУникальныйАйди: %s\n", v.Id, v.Name, v.VideoUrl, v.Cost, v.Algorithm, v.Size, v.Hashrate, v.UID)
			c.Send(s)
			_, err := base64toJpg(v.Image)
			if err != nil {
				c.Send(err)
			}
			g := &tele.Photo{File: tele.FromDisk("test.png")}
			b.Send(c.Sender(), g)

		}
		return nil
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		switch {
		case state.FSM.Current() == EditDeviceImage:
			photo := c.Message().Photo
			b.Download(&photo.File, "best.png")
			f, _ := os.Open("./best.png")

			// Read entire JPG into byte slice.
			reader := bufio.NewReader(f)
			content, _ := io.ReadAll(reader)

			encoded := base64.StdEncoding.EncodeToString(content)
			fmt.Println(encoded)
			err := repo.EditDeviceImage(context.Background(), id, encoded)
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == AddDeviceImage:
			photo := c.Message().Photo
			b.Download(&photo.File, "best.png")
			f, _ := os.Open("./best.png")

			// Read entire JPG into byte slice.
			reader := bufio.NewReader(f)
			content, _ := io.ReadAll(reader)

			encoded := base64.StdEncoding.EncodeToString(content)
			device.Image = encoded
			err := repo.DeviceInsertion(context.Background(), device)
			if err != nil {
				c.Send(err.Error())
			} else {
				c.Send("Успешно")
			}
		}
		return nil
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		switch {
		case state.FSM.Current() == EditDevice:
			keyboardDevice(state, c, b, repo)
		case state.FSM.Current() == editDeviceID:
			id = c.Text()
			err := state.FSM.Event(context.Background(), PreviousState)
			if err != nil {
				fmt.Println(err)
			}
			c.Send("Теперь введите значение, на что хотите поменять")
		case state.FSM.Current() == EditArticleID:
			id = c.Text()
			err := state.FSM.Event(context.Background(), PreviousState)
			if err != nil {
				fmt.Println(err)
			}
			c.Send("Теперь введите значение, на что хотите поменять")
		case state.FSM.Current() == EditCaseID:
			id = c.Text()
			err := state.FSM.Event(context.Background(), PreviousState)
			if err != nil {
				fmt.Println(err)
			}
			c.Send("Теперь введите значение, на что хотите поменять")
		case state.FSM.Current() == EditCompanyReviewsID:
			id = c.Text()
			err := state.FSM.Event(context.Background(), PreviousState)
			if err != nil {
				fmt.Println(err)
			}
			c.Send("Теперь введите значение, на что хотите поменять")
		case state.FSM.Current() == EditDeviceReviewsID:
			id = c.Text()
			err := state.FSM.Event(context.Background(), PreviousState)
			if err != nil {
				fmt.Println(err)
			}
			c.Send("Теперь введите значение, на что хотите поменять")
		case state.FSM.Current() == EditDeviceName:
			err := repo.EditDeviceName(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceAlgorithm:
			err := repo.EditDeviceAlgorithm(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceSize:
			err := repo.EditDeviceSize(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceCost:
			err := repo.EditDeviceCost(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDevicePower:
			err := repo.EditDevicePower(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceHashrate:
			err := repo.EditDeviceHashrate(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceUid:
			err := repo.EditDeviceUid(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceVideoUrl:
			err := repo.EditDeviceVideoUrl(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditArticleName:
			err := repo.EditArticleName(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditArticleVideoUrl:
			err := repo.EditArticleVideoUrl(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditArticleText:
			err := repo.EditArticleText(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditCaseName:
			err := repo.EditCaseName(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditCaseVideoUrl:
			err := repo.EditCaseVideoUrl(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditCaseText:
			err := repo.EditCaseText(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditCompanyReviewsName:
			err := repo.EditCompanyReviewsName(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditCompanyReviewsStars:
			err := repo.EditCompanyReviewsStars(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditCompanyReviewsText:
			err := repo.EditCompanyReviewsText(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditCompanyReviewsEmail:
			err := repo.EditCompanyReviewsEmail(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditDeviceReviewsName:
			err := repo.EditDeviceReviewsName(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditDeviceReviewsStars:
			err := repo.EditDeviceReviewsStars(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditDeviceReviewsText:
			err := repo.EditDeviceReviewsText(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceReviewsEmail:
			err := repo.EditDeviceReviewsEmail(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == EditDeviceReviewsAmount:
			err := repo.EditDeviceReviewsAmount(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditDeviceReviewsDeviceID:
			err := repo.EditDeviceReviewsDeviceID(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == EditDeviceReviewsDate:
			err := repo.EditDeviceReviewsDate(context.Background(), id, c.Text())
			if err != nil {
				c.Send(err.Error())
				return err
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == AddDeviceName:
			device.Name = c.Text()
			c.Send("Пришлите Стоимость")
			err := state.FSM.Event(context.Background(), AddDeviceCost)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceCost:
			device.Cost = c.Text()
			c.Send("Пришлите размер")
			err := state.FSM.Event(context.Background(), AddDeviceSize)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceSize:
			device.Size = c.Text()
			c.Send("Пришлите потребление")
			err := state.FSM.Event(context.Background(), AddDevicePower)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDevicePower:
			device.Power = c.Text()
			c.Send("Пришлите Хэшрейт")
			err := state.FSM.Event(context.Background(), AddDeviceHashrate)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceHashrate:
			device.Hashrate = c.Text()
			c.Send("Пришлите Алгоритм")
			err := state.FSM.Event(context.Background(), AddDeviceAlgorithm)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceAlgorithm:
			device.Algorithm = c.Text()
			c.Send("Пришлите Уид")
			err := state.FSM.Event(context.Background(), AddDeviceUid)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceUid:
			device.UID = c.Text()
			c.Send("Пришлите видео url")
			err := state.FSM.Event(context.Background(), AddDeviceVideoUrl)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceVideoUrl:
			device.VideoUrl = c.Text()
			c.Send("Пришлите изображение")
			err := state.FSM.Event(context.Background(), AddDeviceImage)
			if err != nil {
				fmt.Println(err)
			}

		case state.FSM.Current() == RemoveDeviceID:
			err := repo.DeviceRemove(context.Background(), c.Text())
			if err != nil {
				c.Send(err.Error())
				fmt.Println(err.Error())
			} else {
				c.Send("Успешно")
			}

		case state.FSM.Current() == AddCaseName:
			cas.Name = c.Text()
			c.Send("Пришлите текст")
			err := state.FSM.Event(context.Background(), AddCaseText)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddCaseText:
			cas.Text = c.Text()
			c.Send("Пришлите Url")
			err := state.FSM.Event(context.Background(), AddCaseVideoUrl)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddCaseVideoUrl:
			cas.VideoUrl = c.Text()
			err := repo.CaseInsertion(context.Background(), cas)
			if err != nil {
				c.Send(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == RemoveCaseID:
			err := repo.CaseRemove(context.Background(), c.Text())
			if err != nil {
				c.Send(err.Error())
				fmt.Println(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == AddArticleName:
			article.Name = c.Text()
			c.Send("Пришлите текст")
			err := state.FSM.Event(context.Background(), AddArticleText)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddArticleText:
			article.Text = c.Text()
			c.Send("Пришлите Url")
			err := state.FSM.Event(context.Background(), AddArticleVideoUrl)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddArticleVideoUrl:
			article.VideoUrl = c.Text()
			err := repo.ArticleInsertion(context.Background(), article)
			if err != nil {
				c.Send(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == RemoveArticleID:
			err := repo.ArticleRemove(context.Background(), c.Text())
			if err != nil {
				c.Send(err.Error())
				fmt.Println(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == RemoveCompanyReviewsID:
			err := repo.CompanyReviewRemove(context.Background(), c.Text())
			if err != nil {
				c.Send(err.Error())
				fmt.Println(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == RemoveDeviceReviewsID:
			err := repo.DeviceReviewRemove(context.Background(), c.Text())
			if err != nil {
				c.Send(err.Error())
				fmt.Println(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == AddCompanyReviewsName:
			compReview.Name = c.Text()
			c.Send("Пришлите текст")
			err := state.FSM.Event(context.Background(), AddCompanyReviewsText)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddCompanyReviewsText:
			compReview.Text = c.Text()
			c.Send("Пришлите Звезды")
			err := state.FSM.Event(context.Background(), AddCompanyReviewsStars)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddCompanyReviewsStars:
			compReview.Stars = c.Text()
			c.Send("Пришлите Url")
			err := state.FSM.Event(context.Background(), AddCompanyReviewsEmail)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddCompanyReviewsEmail:
			compReview.Email = c.Text()
			err := repo.CompanyReviewsInsertion(context.Background(), compReview)
			if err != nil {
				c.Send(err.Error())
			} else {
				c.Send("Успешно")
			}
		case state.FSM.Current() == AddDeviceReviewsName:
			review.Name = c.Text()
			c.Send("Пришлите текст")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsText)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsText:
			review.Text = c.Text()
			c.Send("Пришлите Звезды")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsStars)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsStars:
			review.Stars = c.Text()
			c.Send("Пришлите Дату")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsDate)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsDate:
			review.Date = c.Text()
			c.Send("Пришлите Количетсво")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsAmount)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsAmount:
			review.Amount = c.Text()
			c.Send("Пришлите DeviceID")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsDeviceID)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsDeviceID:
			review.DeviceId = c.Text()
			c.Send("Пришлите email")
			err := state.FSM.Event(context.Background(), AddDeviceReviewsEmail)
			if err != nil {
				fmt.Println(err)
			}
		case state.FSM.Current() == AddDeviceReviewsEmail:
			review.Email = c.Text()
			err := repo.DeviceReviewsInsertion(context.Background(), review)
			if err != nil {
				c.Send(err.Error())
			} else {
				c.Send("Успешно")
			}
		}
		return nil
	})
}

func keyboardArticle(state *Chat, c tele.Context, b *tele.Bot) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnName := menu.Text("Имя")
	btnText := menu.Text("Текст")
	btnVideoUrl := menu.Text("Ссылка на видео")
	menu.Reply(
		menu.Row(btnText),
		menu.Row(btnName),
		menu.Row(btnVideoUrl),
	)
	c.Send("Что вы хотите редактировать ? ", menu)

	b.Handle(&btnText, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditArticleID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditArticleText
		return c.Send("напиши мне айди статьи , текст которой меняешь")
	})

	b.Handle(&btnName, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditArticleID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditArticleName
		return c.Send("напиши мне айди статьи , имя которой меняешь")
	})
	b.Handle(&btnVideoUrl, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditArticleID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditArticleVideoUrl
		return c.Send("напиши мне айди статьи , видео которой меняешь")
	})
}

func keyboardCompanyReviews(state *Chat, c tele.Context, b *tele.Bot) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnName := menu.Text("Имя")
	btnEmail := menu.Text("Email")
	btnText := menu.Text("Текст")
	btnStars := menu.Text("Звезды")
	menu.Reply(
		menu.Row(btnText),
		menu.Row(btnName),
		menu.Row(btnStars),
		menu.Row(btnEmail),
	)
	c.Send("Что вы хотите редактировать ? ", menu)

	b.Handle(&btnText, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCompanyReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCompanyReviewsText
		return c.Send("напиши мне айди отзыва , текст которой меняешь")
	})
	b.Handle(&btnEmail, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCompanyReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCompanyReviewsEmail
		return c.Send("напиши мне айди отзыва , текст которой меняешь")
	})

	b.Handle(&btnName, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCompanyReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCompanyReviewsName
		return c.Send("напиши мне айди отзыва , имя которой меняешь")
	})
	b.Handle(&btnStars, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCompanyReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCompanyReviewsStars
		return c.Send("напиши мне айди отзыва , видео которой меняешь")
	})
}

func keyboardDeviceReviews(state *Chat, c tele.Context, b *tele.Bot) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnName := menu.Text("Имя")
	btnEmail := menu.Text("Email")
	btnText := menu.Text("Текст")
	btnStars := menu.Text("Звезды")
	btnDeviceID := menu.Text("АйдиДевайса")
	btnDate := menu.Text("Дата")
	btnAmount := menu.Text("Количество")

	menu.Reply(
		menu.Row(btnText),
		menu.Row(btnName),
		menu.Row(btnStars),
		menu.Row(btnEmail),
		menu.Row(btnDeviceID),
		menu.Row(btnDate),
		menu.Row(btnAmount),
	)
	c.Send("Что вы хотите редактировать ? ", menu)

	b.Handle(&btnText, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsText
		return c.Send("напиши мне айди отзыва , текст которого меняешь")
	})
	b.Handle(&btnEmail, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsEmail
		return c.Send("напиши мне айди отзыва , текст которого меняешь")
	})

	b.Handle(&btnName, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsName
		return c.Send("напиши мне айди отзыва , имя которого меняешь")
	})
	b.Handle(&btnStars, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsStars
		return c.Send("напиши мне айди отзыва , видео которой меняешь")
	})
	b.Handle(&btnDate, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsDate
		return c.Send("напиши мне айди отзыва , текст которого меняешь")
	})

	b.Handle(&btnAmount, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsAmount
		return c.Send("напиши мне айди отзыва , имя которого меняешь")
	})
	b.Handle(&btnDeviceID, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditDeviceReviewsID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceReviewsDeviceID
		return c.Send("напиши мне айди отзыва , видео которой меняешь")
	})
}

func keyboardCase(state *Chat, c tele.Context, b *tele.Bot) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnName := menu.Text("Имя")
	btnText := menu.Text("Текст")
	btnVideoUrl := menu.Text("Ссылка на видео")
	menu.Reply(
		menu.Row(btnText),
		menu.Row(btnName),
		menu.Row(btnVideoUrl),
	)
	c.Send("Что вы хотите редактировать ? ", menu)

	b.Handle(&btnText, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCaseID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCaseText
		return c.Send("напиши мне айди статьи , текст которой меняешь")
	})

	b.Handle(&btnName, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCaseID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCaseName
		return c.Send("напиши мне айди статьи , имя которой меняешь")
	})
	b.Handle(&btnVideoUrl, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), EditCaseID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditCaseVideoUrl
		return c.Send("напиши мне айди статьи , видео которой меняешь")
	})
}

func keyboardDevice(state *Chat, c tele.Context, b *tele.Bot, repo *repository.IndexRepository) {
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnAlgorithm := menu.Text("Алгоритм: 🧮")
	btnName := menu.Text("Имя: 📛")
	btnSize := menu.Text("Размер: 📏")
	btnCost := menu.Text("Стоимость: 💰")
	btnPoplura := menu.Text("Популярность: 👍")
	btnImage := menu.Text("Изображение: 🖼️")
	btnHashrate := menu.Text("Хэшрейт: ⛏️")
	btnVideoUrl := menu.Text("Ссылка на видео: 📺")
	btnPower := menu.Text("Потребление: ⚡")
	btnUid := menu.Text("UID: 🔑")

	menu.Reply(
		menu.Row(btnAlgorithm),
		menu.Row(btnName),
		menu.Row(btnSize),
		menu.Row(btnCost),
		menu.Row(btnImage),
		menu.Row(btnPoplura),
		menu.Row(btnHashrate),
		menu.Row(btnVideoUrl),
		menu.Row(btnPower),
		menu.Row(btnUid),
	)
	c.Send("Что вы хотите редактировать ? ", menu)

	b.Handle(&btnAlgorithm, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceAlgorithm
		return c.Send("напиши мне айди девайса , алгоритм которого меняешь")
	})

	b.Handle(&btnCost, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceCost

		return c.Send("напиши мне айди девайса , стоимость которого меняешь")
	})
	b.Handle(&btnPoplura, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		err = repo.EditDevicePopularity(context.Background(), id)
		if err != nil {
			return c.Send(err.Error())
		} else {
			return c.Send("Успешно")
		}
	})

	b.Handle(&btnHashrate, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceHashrate

		return c.Send("напиши мне айди девайса , хэшрейт которого меняешь")
	})
	b.Handle(&btnImage, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceImage

		return c.Send("напиши мне айди девайса , изображение которого меняешь")
	})
	b.Handle(&btnPower, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}

		PreviousState = EditDevicePower

		return c.Send("напиши мне айди девайса , потребление которого меняешь")
	})
	b.Handle(&btnUid, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceUid

		return c.Send("напиши мне айди девайса , uid которого меняешь")
	})
	b.Handle(&btnVideoUrl, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceVideoUrl

		return c.Send("напиши мне айди девайса , ссылку на видео которого меняешь")
	})
	b.Handle(&btnSize, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceSize

		return c.Send("напиши мне айди девайса , размер которого меняешь")
	})
	b.Handle(&btnName, func(c tele.Context) error {
		err := state.FSM.Event(context.Background(), editDeviceID)
		if err != nil {
			fmt.Println(err)
		}
		PreviousState = EditDeviceName

		return c.Send("напиши мне айди девайса , имя которого меняешь")
	})
}

func base64toJpg(data string) (f *os.File, err error) {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	pngFilename := "test.png"
	f, err = os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {

		return nil, err
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		return nil, err
	}
	fmt.Println("Jpg file", pngFilename, "created")
	return f, err
}
