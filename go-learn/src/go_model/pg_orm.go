package go_model

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Profile struct {
	Id int
	Lang string
	UserId int
}

type User struct {
	Id int
	Name string
	Profile *Profile
}

func (u User)String() string {
	return fmt.Sprintf("User<%d> <%s>", u.Id, u.Name)
}

func CreateModel() {

	db := pg.Connect(&pg.Options{
		User: "chenjian",
		Database: "pg_orm",
	})

	defer db.Close()

	err :=createSchema(db)

	if err != nil {
		panic(err)
	}

	user1 := &User{
		Name:"chenjian",
	}

	err = db.Insert(user1)
	if err != nil {
		panic(err)
	}

	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(users)

}
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil)}{
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//func main() {
//	CreateModel()
//}





