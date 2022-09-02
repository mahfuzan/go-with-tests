package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ImageRepository interface {
	Store(image Image) error
	GetList() ([]Image, error)
	Get(id int64) (Image, error)
}

type MysqlRepository struct {
	sqlDb *sqlx.DB
}

func NewMysqlRepository(sqlDb *sqlx.DB) ImageRepository {
	return MysqlRepository{sqlDb: sqlDb}
}

func (sql MysqlRepository) Store(image Image) error {
	stmt, err := sql.sqlDb.Prepare("INSERT INTO images (file_name, file_path, source_url) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(image.FileName, image.FilePath, image.SourceUrl)
	if err != nil {
		return err
	}

	return nil
}

func (sql MysqlRepository) GetList() ([]Image, error) {
	imageList := []Image{}
	err := sql.sqlDb.Select(&imageList, "SELECT * FROM images")
	if err != nil {
		return nil, err
	}

	return imageList, nil
}

func (sql MysqlRepository) Get(id int64) (Image, error) {
	image := Image{}
	err := sql.sqlDb.Get(&image, "SELECT * FROM images WHERE id = ?", id)
	if err != nil {
		return Image{}, err
	}

	return image, nil
}

type Service struct {
	imageRepository ImageRepository
}

func NewService(imageRepository ImageRepository) Service {
	return Service{imageRepository: imageRepository}
}

func (service Service) Insert(filename, filepath, sourceUrl string) error {
	err := service.imageRepository.Store(Image{
		FileName:  filename,
		FilePath:  filepath,
		SourceUrl: sourceUrl,
	})

	if err != nil {
		return err
	}
	return nil
}

func (service Service) GetImageList() ([]Image, error) {
	imageList, err := service.imageRepository.GetList()
	if err != nil {
		return nil, err
	}

	return imageList, nil
}

func (service Service) GetById(id int64) (Image, error) {
	image, err := service.imageRepository.Get(id)
	if err != nil {
		return Image{}, err
	}

	return image, nil
}

var (
	mysqlHost = "127.0.0.1"
	mysqlUser = "root"
	mysqlPass = "root"
)

type Image struct {
	Id        int    `db:"id"`
	FileName  string `db:"file_name"`
	FilePath  string `db:"file_path"`
	SourceUrl string `db:"source_url"`
}

func newSqlConn(mysqlHost, mysqlUser, mysqlPass string) (*sqlx.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/training?parseTime=true", mysqlUser, mysqlPass, mysqlHost)
	sql, err := sqlx.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if err := sql.Ping(); err != nil {
		return nil, err
	}

	return sql, err
}

func main() {
	sqlDb, err := newSqlConn(mysqlHost, mysqlUser, mysqlPass)
	if err != nil {
		log.Fatal(err)
	}

	mysqlRepository := NewMysqlRepository(sqlDb)
	imageService := NewService(mysqlRepository)
	if err := imageService.Insert("testing.png", "testing.png", "testing.png"); err != nil {
		log.Fatal(err)
	}

	imageList, err := imageService.GetImageList()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(imageList)
	}

	imageById, err := imageService.GetById(1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(imageById)
	}

	log.Println("Done")
}
