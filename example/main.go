package main

import (
	_ "github.com/lib/pq"
	"github.com/westernmonster/xorm"
	"time"
)

type User struct {
	Id      int64     `xorm:"pk BIGINT"`
	Name    string    `xorm:"pk VARCHAR(64)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (p *UserRepo) CreateUser(ses xorm.Session, user *User) error {
	var session = &ses
	if e := session.Begin(); e != nil {
		return e
	}

	if _, e := session.Insert(user); e != nil {
		session.Rollback()
		return e
	}

	if e := session.Commit(); e != nil {
		return e
	}

	return nil
}

type Tag struct {
	Id      int64     `xorm:"pk BIGINT"`
	Name    string    `xorm:"pk VARCHAR(64)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type TagRepo struct{}

func NewTagRepo(ses xorm.Session) *TagRepo {
	return &TagRepo{}
}

func (p *TagRepo) CreateTag(ses xorm.Session, tag *Tag) error {
	var session = &ses
	if e := session.Begin(); e != nil {
		return e
	}

	if _, e := session.Insert(tag); e != nil {
		session.Rollback()
		return e
	}

	if e := session.Commit(); e != nil {
		return e
	}

	return nil
}

func main() {
	engine, err := xorm.NewEngine("postgres", "dbname=test user=test password=test host=localhost port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

	session := engine.NewSession()

	user := &User{
		Id:   3,
		Name: "test2",
	}

	tag := &Tag{
		Id:   3,
		Name: "test2",
	}

	repoUser := new(UserRepo)
	repoTag := new(TagRepo)

	if e := session.Begin(); e != nil {
		panic(e)
	}

	if e := repoUser.CreateUser(*session, user); e != nil {
		session.Rollback()
		return
	}

	if e := repoTag.CreateTag(*session, tag); e != nil {
		session.Rollback()
		return
	}

	if e := session.Commit(); e != nil {
		panic(e)
	}
}
