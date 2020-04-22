package main

import (
	"database/sql"
	"fmt"
	"go_postgresql/lib"
	"time"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("This is main")

	//do get all persons
	fmt.Println("do get all persons")
	getPersons()

	//get person by code
	fmt.Println("do get person that person code is \"1\"")
	getPersonByCode("1")

	//create a new person
	fmt.Println("create a new person")
	newPersonCD := createNewPerson()

	//update the new person data
	fmt.Println("update the new person")
	updatePerson(newPersonCD)

	//delete the new person data
	fmt.Println("delete the new person")
	deletePerson(newPersonCD)

}

func getPersons() {
	persons := lib.GetPersons()

	fmt.Println(*persons)
}

func getPersonByCode(personCD string) {
	person, err := lib.GetPerson(personCD)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("no records matching with", personCD)
		return
	case err != nil:
		fmt.Println("an error occourd")
		fmt.Println(err)
		return

	}

	fmt.Println(*person)
}

func createNewPerson() string {
	var person lib.MstPersonInfo
	loc, _ := time.LoadLocation("Asia/Tokyo")

	// person.PersonCd = ""    create by sequence
	person.PersonName = "テスト　一郎"
	person.PersonNameKana = "ﾃｽﾄ ｲﾁﾛｳ"
	person.Birthday = time.Date(1990, 10, 10, 0, 0, 0, 0, loc)
	person.Sex = "0"
	person.Zip = "123-1234"
	person.Address1 = "東京都"
	person.Address2 = "千代田区"
	person.Address3 = "神田淡路町１－２－３－４"
	person.Address4 = "メゾンM　１２２３室"
	person.Tel = "0312345678"
	person.Mobile = "07012345678"
	person.MailAddress = "112233@go.com"
	person.AuthenticationDate = time.Now()
	person.DeleteFlag = "0"
	person.CreateUser = "test"
	person.UpdateUser = "test"

	newPersonCD, err := lib.CreatePerson(&person)

	if err != nil {
		fmt.Println("an error occourd")
		fmt.Println(err)
		return ""
	}

	fmt.Println("create person successed. the new personcd is ", newPersonCD)

	//get the person data benn created
	getPersonByCode(newPersonCD)

	return newPersonCD
}

func updatePerson(personCD string) {
	var person lib.MstPersonInfo
	var err error
	person1, err1 := lib.GetPerson(personCD)

	if err1 == sql.ErrNoRows {
		fmt.Println("no records matched")
		return
	}

	person = *person1

	person.PersonName = person.PersonName + "更新"
	person.PersonNameKana = person.PersonNameKana + "ｺｳｼﾝ"

	err = lib.UpdatePerson(&person, personCD)
	if err != nil {
		fmt.Println("an error occourd")
		fmt.Println(err)
		return
	}

	//get the person data benn created
	fmt.Println("show person data updated")
	getPersonByCode(personCD)
}

func deletePerson(personCD string) {
	result, err := lib.DeletePerson(personCD)

	if err == nil && result > 0 {
		fmt.Println("person data has been deleted")

		//get the person data benn deleted
		fmt.Println("show person data deleted")
		getPersonByCode(personCD)
	} else {
		if result == 0 {
			fmt.Println("no data matched")
		} else {
			fmt.Println("an error occourd")
			fmt.Println(err)
		}

	}
}
