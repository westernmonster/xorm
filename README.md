# xorm
xorm是一个简单而强大的Go语言ORM库. 通过它可以使数据库操作非常简便。

## 说明

```Go
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

type UserRepo struct {
	session *xorm.Session
}

func NewUserRepo(ses xorm.Session) *UserRepo {
	return &UserRepo{session: &ses}
}

func (p *UserRepo) CreateUser(user *User) error {
	if e := p.session.Begin(); e != nil {
		return e
	}

	if _, e := p.session.Insert(user); e != nil {
		p.session.Rollback()
		return e
	}

	if e := p.session.Commit(); e != nil {
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

type TagRepo struct {
	session *xorm.Session
}

func NewTagRepo(ses xorm.Session) *TagRepo {
	return &TagRepo{session: &ses}
}

func (p *TagRepo) CreateTag(tag *Tag) error {
	if e := p.session.Begin(); e != nil {
		return e
	}

	if _, e := p.session.Insert(tag); e != nil {
		p.session.Rollback()
		return e
	}

	if e := p.session.Commit(); e != nil {
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
		Id:   2,
		Name: "test2",
	}

	tag := &Tag{
		Id:   2,
		Name: "test2",
	}

	if e := session.Begin(); e != nil {
		panic(e)
	}

	// 需要在Begin后调用初始化
	repoUser := NewUserRepo(*session)
	repoTag := NewTagRepo(*session)

	if e := repoUser.CreateUser(user); e != nil {
		session.Rollback()
		return
	}

	if e := repoTag.CreateTag(tag); e != nil {
		session.Rollback()
		return
	}

	if e := session.Commit(); e != nil {
		panic(e)
	}
}
```

