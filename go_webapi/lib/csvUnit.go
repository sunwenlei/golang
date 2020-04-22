package lib

import (
	"encoding/csv"
	"log"
	"os"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

//CreateCSV create a csv file
func CreateCSV(filename string, persons []MstPersonInfo) {
	// O_WRONLY:write only, O_CREATE:create it when not exists
	// create file in tmp folder
	file, err := os.OpenFile("./tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()

	err = file.Truncate(0)
	failOnError(err)

	writer := csv.NewWriter(file)
	writer.Write([]string{"person_cd", "person_name", "person_name_kana", "birthday", "sex", "zip", "address1", "address2", "address3", "address4", "tel", "mobile", "mail_address", "authentication_date", "delete_flag", "create_user", "create_date", "update_user", "update_date"})
	for _, person := range persons {
		writer.Write([]string{person.PersonCd,
			person.PersonName,
			person.PersonNameKana,
			person.Birthday.Format("2020-01-01"),
			person.Sex,
			person.Zip,
			person.Address1,
			person.Address2,
			person.Address3,
			person.Address4,
			person.Tel,
			person.Mobile,
			person.MailAddress,
			person.AuthenticationDate.Format("2020-01-01"),
			person.DeleteFlag,
			person.CreateUser,
			person.CreateDate.Format("2020-01-01 20:10:01"),
			person.UpdateUser,
			person.UpdateDate.Format("2020-01-01 20:10:01"),
		})
	}

	// also can use length index and will get same result
	// var lenthPerson int = len(persons)
	// for i := 0; i < lenthPerson; i++ {
	// 	var person = persons[i]
	// 	writer.Write([]string{person.PersonCd,
	// 		person.PersonName,
	// 		person.PersonNameKana,
	// 		person.Birthday.Format("2020-01-01"),
	// 		person.Sex,
	// 		person.Zip,
	// 		person.Address1,
	// 		person.Address2,
	// 		person.Address3,
	// 		person.Address4,
	// 		person.Tel,
	// 		person.Mobile,
	// 		person.MailAddress,
	// 		person.AuthenticationDate.Format("2020-01-01"),
	// 		person.DeleteFlag,
	// 		person.CreateUser,
	// 		person.CreateDate.Format("2020-01-01 20:10:01"),
	// 		person.UpdateUser,
	// 		person.UpdateDate.Format("2020-01-01 20:10:01"),
	// 	})
	// }

	writer.Flush()
}
