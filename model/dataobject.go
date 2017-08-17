package model

import (
	"log"
	"os"
	"strings"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"
)

type Follow struct {
	User   string `json:"user",gorethink:"user"`
	Follow string `json:"follow",gorethink:"follow"`
}

type Twitt struct {
	ID   string    `json:"id",gorethink:"id"`
	Text string    `json:"text",gorethink:"text"`
	User string    `json:"user",gorethink:"user"`
	Date time.Time `json:"date",gorethink:"date"`
}

var session *r.Session

func InitSesson() error {
	dbmicroblog := os.Getenv("RETHINKDB_HOST")
	if dbmicroblog == "" {
		dbmicroblog = "localhost"
	}

	log.Printf("RETHINKDB_HOST: %s\n", dbmicroblog)
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: dbmicroblog,
	})
	if err != nil {
		return err
	}

	err = CreateDBIfNotExist()
	if err != nil {
		return err
	}

	err = CreateTwittsTableIfNotExist()
	if err != nil {
		return err
	}

	err = CreateFollowTableIfNotExist()
	if err != nil {
		return err
	}

	return err
}

//проверяем существует ли бд. если нет, то создаем
func CreateDBIfNotExist() error {
	res, err := r.DBList().Run(session)
	if err != nil {
		return err
	}

	var dbList []string
	err = res.All(&dbList)
	if err != nil {
		return err
	}

	for _, item := range dbList {
		if item == "microblog" {
			return nil
		}
	}

	_, err = r.DBCreate("microblog").Run(session)
	if err != nil {
		return err
	}

	return nil
}

//проверяем существует ли таблица twitt. если нет, то создаем и заполняем
func CreateTwittsTableIfNotExist() error {

	res, err := r.DB("microblog").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "twitt" {
			return nil
		}
	}

	_, err = r.DB("microblog").TableCreate("twitt", r.TableCreateOpts{PrimaryKey: "ID"}).Run(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "djbelyak",
		"text": "У меня большой опыт: я всю жизнь работаю с идиотами.",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "djbelyak",
		"text": "Мои парни этим займутся.",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "MeXoS",
		"text": "Сейчас я людей соберу!",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "MeXoS",
		"text": "Где-то у меня была еще водка.",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "Cracktv",
		"text": "Кто тебя научил так драться?",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "Cracktv",
		"text": "Зачем мне тебе помогать?",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "syanaw",
		"text": "Сдавайся! Сбережешь мне время!",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}

	_, err = r.DB("microblog").Table("twitt").Insert(map[string]interface{}{
		"ID":   r.UUID(),
		"user": "syanaw",
		"text": "Меду! Несите еще меду!",
		"date": r.Now(),
	}).RunWrite(session)
	if err != nil {
		return err
	}
	return err
}

//проверяем существует ли таблица Follow. если нет, то создаем и заполняем
func CreateFollowTableIfNotExist() error {

	res, err := r.DB("microblog").TableList().Run(session)
	if err != nil {
		return err
	}

	var tableList []string
	err = res.All(&tableList)
	if err != nil {
		return err
	}

	for _, item := range tableList {
		if item == "follow" {
			return nil
		}
	}

	_, err = r.DB("microblog").TableCreate("follow", r.TableCreateOpts{PrimaryKey: "User"}).Run(session)
	if err != nil {
		return err
	}
	data := new(Follow)
	data.User = "MeXoS"
	data.Follow = "djbelyak Cracktv"
	_, err = r.DB("microblog").Table("follow").Insert(data).RunWrite(session)

	return err
}

//получаем наши новости
func GetNews() ([]Twitt, error) {
	//var userName string = "MeXos"
	res, err := r.DB("microblog").Table("follow").Get("MeXoS").Run(session)
	if err != nil {
		return nil, err
	}

	var response Follow
	err = res.One(&response)
	if err != nil {
		return nil, err
	}

	var following string = response.Follow
	var arr []string = strings.Split(following, " ")
	var answer []Twitt
	var temp []Twitt
	for i := 0; i < len(arr); i++ {
		log.Println("i: ", i)
		log.Println("user: ", arr[i])
		res, err := r.DB("microblog").Table("twitt").Filter(map[string]interface{}{"user": arr[i],}).Run(session)
		if err != nil {
			return nil, err
		}

		err = res.All(&temp)
		if err != nil {
			return nil, err
		}
		for j := 0; j<len(temp); j++{
			log.Println(temp[j].Text)
			answer = append(answer, temp[j])
		}
	}
	return answer, nil
}

//получение моих твиттов
func GetMyTwitts() ([]Twitt, error) {
	res, err := r.DB("microblog").Table("twitt").Filter(map[string]interface{}{"user": "MeXoS"}).Run(session)
	//log.Println()
	if err != nil {
		return nil, err
	}
	var response []Twitt
	err = res.All(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}