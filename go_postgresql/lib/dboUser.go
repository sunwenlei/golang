package lib

//GetPersons get all persons data
func GetPersons() *[]MstPersonInfo {
	Connect()
	defer Disconnect()

	var strSQL string = "select * from mst_person_info;"

	rows, err := Db.Query(strSQL)

	defer rows.Close()

	if err == nil {
		var rs []MstPersonInfo

		for rows.Next() {
			var r MstPersonInfo
			err = rows.Scan(&r.PersonCd,
				&r.PersonName,
				&r.PersonNameKana,
				&r.Birthday,
				&r.Sex,
				&r.Zip,
				&r.Address1,
				&r.Address2,
				&r.Address3,
				&r.Address4,
				&r.Tel,
				&r.Mobile,
				&r.MailAddress,
				&r.AuthenticationDate,
				&r.DeleteFlag,
				&r.CreateUser,
				&r.CreateDate,
				&r.UpdateUser,
				&r.UpdateDate)

			rs = append(rs, r)
		}
		return &rs
	} else {
		return nil
	}

}
