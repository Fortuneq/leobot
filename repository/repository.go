package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type IndexRepository struct {
	db *sqlx.DB
}

type Reviews struct {
	Email string `db:"email"`
	Text  string `db:"text"`
	Name  string `db:"name"`
	Stars string `db:"stars"`
	Id    string `db:"id"`
}

type DeviceReviews struct {
	Email    string `db:"email"`
	Text     string `db:"review_text"`
	Name     string `db:"name"`
	Stars    string `db:"stars"`
	Id       string `db:"id"`
	DeviceId string `db:"device_id"`
	Amount   string `db:"amount"`
	Date     string `db:"date"`
}

type Devices struct {
	Name      string `db:"name"`
	Id        string `db:"id"`
	Cost      string `db:"cost"`
	Image     string `db:"image"`
	Size      string `db:"size"`
	Power     string `db:"power"`
	Hashrate  string `db:"hashrate"`
	UID       string `db:"uid"`
	Algorithm string `db:"algorithm"`
	VideoUrl  string `db:"video_url"`
}

type Articles struct {
	Text     string `db:"text"`
	Name     string `db:"name"`
	Id       string `db:"id"`
	VideoUrl string `db:"video_url"`
}

type Cases struct {
	Text     string `db:"text"`
	Name     string `db:"name"`
	Id       string `db:"id"`
	VideoUrl string `db:"video_url"`
}

func (r *IndexRepository) GetAllReviews(ctx context.Context) ([]Reviews, error) {

	//Абстрактный sql ,  с которого получаем данные

	q := "SELECT id,stars, name,text,email FROM company_reviews"

	place := []Reviews{}

	err := r.db.SelectContext(ctx, &place, q)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func (r *IndexRepository) GetAllDevices(ctx context.Context) ([]Devices, error) {

	//Абстрактный sql ,  с которого получаем данные

	q := "SELECT id,name, cost,image,size,power,hashrate,algorithm,uid,video_url FROM devices "

	place := []Devices{}

	err := r.db.SelectContext(ctx, &place, q)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func (r *IndexRepository) EditDeviceName(ctx context.Context, id, name string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET name = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, name, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceCost(ctx context.Context, id, cost string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET cost = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, cost, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDevicePopularity(ctx context.Context, id string) error {
	//Абстрактный sql ,  с которого получаем данные
	str := ""
	s := "SELECT recommended from devices where id = ?"
	err := r.db.Select(ctx, str, s, id)
	if err != nil {
		println(err)
		return err
	}
	q := "Update devices SET recommended = ? Where id = ?"
	i, err := strconv.ParseInt(str, 16, 32)
	if err != nil {
		println(err)
		return err
	}
	if i == 0 {
		i = 1
	} else {
		i = 0
	}
	_, err = r.db.ExecContext(ctx, q, id, i)
	if err != nil {
		println(err)
		return err
	}
	return nil
}
func (r *IndexRepository) EditDeviceSize(ctx context.Context, id, size string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET size = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, size, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceImage(ctx context.Context, id, image string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET image = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, image, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDevicePower(ctx context.Context, id, power string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET power = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, power, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceHashrate(ctx context.Context, id, hashrate string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET hashrate = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, hashrate, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceUid(ctx context.Context, id, uid string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET uid = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, uid, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceVideoUrl(ctx context.Context, id, videoUrl string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET video_url = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, videoUrl, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceAlgorithm(ctx context.Context, id, algorithm string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update devices SET algorithm = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, algorithm, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) GetDeviceReviews(ctx context.Context) ([]DeviceReviews, error) {

	//Абстрактный sql ,  с которого получаем данные

	q := "SELECT id,name,date,amount,device_id,stars,amount,email,review_text FROM device_reviews"

	place := []DeviceReviews{}

	err := r.db.SelectContext(ctx, &place, q)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func (r *IndexRepository) GetArticles(ctx context.Context) ([]Articles, error) {

	//Абстрактный sql ,  с которого получаем данные

	q := "SELECT id,name,text,video_url FROM articles"

	place := []Articles{}

	err := r.db.SelectContext(ctx, &place, q)
	if err != nil {
		return nil, err
	}
	return place, nil
}

func (r *IndexRepository) EditArticleVideoUrl(ctx context.Context, id, videoUrl string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update articles SET video_url = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, videoUrl, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditArticleName(ctx context.Context, id, name string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update articles SET name = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, name, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditArticleText(ctx context.Context, id, text string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update articles SET text = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, text, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCaseVideoUrl(ctx context.Context, id, videoUrl string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update cases SET video_url = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, videoUrl, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCaseName(ctx context.Context, id, name string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update cases SET name = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, name, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCaseText(ctx context.Context, id, text string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update cases SET text = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, text, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCompanyReviewsText(ctx context.Context, id, text string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update company_reviews SET text = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, text, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCompanyReviewsEmail(ctx context.Context, id, email string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update company_reviews SET email = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, email, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCompanyReviewsName(ctx context.Context, id, name string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update company_reviews SET  name = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, name, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditCompanyReviewsStars(ctx context.Context, id, stars string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update company_reviews SET  stars = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, stars, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsName(ctx context.Context, id, name string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  name = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, name, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsStars(ctx context.Context, id, stars string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  stars = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, stars, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsAmount(ctx context.Context, id, amount string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  amount = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, amount, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsText(ctx context.Context, id, text string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  review_text = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, text, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsDate(ctx context.Context, id, date string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  date = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, date, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsDeviceID(ctx context.Context, id, deviceID string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  device_id = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, deviceID, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) EditDeviceReviewsEmail(ctx context.Context, id, email string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "Update device_reviews SET  email = ? Where id = ?"

	_, err := r.db.ExecContext(ctx, q, email, id)
	if err != nil {
		println(err)
		return err
	}
	return nil
}

func (r *IndexRepository) DeviceInsertion(ctx context.Context, d Devices) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "INSERT  INTO devices(name,cost,image,size,power,hashrate,algorithm,uid,video_url)  VALUES (?,?,?,?,?,?,?,?,?)"

	_, err := r.db.ExecContext(ctx, q, d.Name, d.Cost, d.Image, d.Size, d.Power, d.Hashrate, d.Algorithm, d.UID, d.VideoUrl)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(d)
		return err
	}
	return nil
}

func (r *IndexRepository) DeviceRemove(ctx context.Context, id string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "DELETE FROM devices WHERE id = ?"

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *IndexRepository) CaseInsertion(ctx context.Context, d Cases) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "INSERT  INTO cases(name,text,video_url)  VALUES (?,?,?)"

	_, err := r.db.ExecContext(ctx, q, d.Name, d.Text, d.VideoUrl)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(d)
		return err
	}
	return nil
}

func (r *IndexRepository) ArticleInsertion(ctx context.Context, d Articles) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "INSERT  INTO articles(name,text,video_url)  VALUES (?,?,?)"

	_, err := r.db.ExecContext(ctx, q, d.Name, d.Text, d.VideoUrl)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(d)
		return err
	}
	return nil
}

func (r *IndexRepository) CompanyReviewsInsertion(ctx context.Context, d Reviews) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "INSERT  INTO company_reviews(name,text,email,stars)  VALUES (?,?,?,?)"

	_, err := r.db.ExecContext(ctx, q, d.Name, d.Text, d.Email, d.Stars)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(d)
		return err
	}
	return nil
}

func (r *IndexRepository) DeviceReviewsInsertion(ctx context.Context, d DeviceReviews) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "INSERT  INTO device_reviews(name,review_text,email,stars,device_id,date,amount)  VALUES (?,?,?,?,?,?,?)"

	_, err := r.db.ExecContext(ctx, q, d.Name, d.Text, d.Email, d.Stars, d.DeviceId, d.Date, d.Amount)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(d)
		return err
	}
	return nil
}
func (r *IndexRepository) CaseRemove(ctx context.Context, id string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "DELETE FROM cases WHERE id = ?"

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *IndexRepository) ArticleRemove(ctx context.Context, id string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "DELETE FROM articles WHERE id = ?"

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *IndexRepository) CompanyReviewRemove(ctx context.Context, id string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "DELETE FROM company_reviews WHERE id = ?"

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *IndexRepository) DeviceReviewRemove(ctx context.Context, id string) error {

	//Абстрактный sql ,  с которого получаем данные

	q := "DELETE FROM device_reviews WHERE id = ?"

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
func (r *IndexRepository) GetCases(ctx context.Context) ([]Cases, error) {

	//Абстрактный sql ,  с которого получаем данные

	q := "SELECT id,name,text,video_url FROM cases"

	place := []Cases{}

	err := r.db.SelectContext(ctx, &place, q)
	if err != nil {
		return nil, err
	}
	return place, nil
}
func NewIndexRepository(db *sqlx.DB) *IndexRepository {
	return &IndexRepository{db}
}
