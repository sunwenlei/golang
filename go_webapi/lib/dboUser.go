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
	}

	return nil
}

//GetPerson get person info by person cd
func GetPerson(personcd string) (*MstPersonInfo, error) {
	Connect()
	defer Disconnect()
	var err error
	err = nil
	var r MstPersonInfoDB
	var result MstPersonInfo

	err = Db.QueryRow("select * from mst_person_info where person_cd= $1;", personcd).Scan(&r.PersonCd,
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

	if err != nil {
		return nil, err
	}

	if r.PersonCd.Valid {
		result.PersonCd = r.PersonCd.String
	}
	if r.PersonName.Valid {
		result.PersonName = r.PersonName.String
	}
	if r.PersonNameKana.Valid {
		result.PersonNameKana = r.PersonNameKana.String
	}
	if r.Birthday.Valid {
		result.Birthday = r.Birthday.Time
	}
	if r.Sex.Valid {
		result.Sex = r.Sex.String
	}
	if r.Zip.Valid {
		result.Zip = r.Zip.String
	}
	if r.Address1.Valid {
		result.Address1 = r.Address1.String
	}
	if r.Address2.Valid {
		result.Address2 = r.Address2.String
	}
	if r.Address3.Valid {
		result.Address3 = r.Address3.String
	}
	if r.Address4.Valid {
		result.Address4 = r.Address4.String
	}
	if r.Tel.Valid {
		result.Tel = r.Tel.String
	}
	if r.Mobile.Valid {
		result.Mobile = r.Mobile.String
	}
	if r.MailAddress.Valid {
		result.MailAddress = r.MailAddress.String
	}
	if r.AuthenticationDate.Valid {
		result.AuthenticationDate = r.AuthenticationDate.Time
	}
	if r.DeleteFlag.Valid {
		result.DeleteFlag = r.DeleteFlag.String
	}
	if r.CreateUser.Valid {
		result.CreateUser = r.CreateUser.String
	}
	if r.CreateDate.Valid {
		result.CreateDate = r.CreateDate.Time
	}
	if r.UpdateUser.Valid {
		result.UpdateUser = r.UpdateUser.String
	}
	if r.UpdateDate.Valid {
		result.UpdateDate = r.UpdateDate.Time
	}

	return &result, nil
}

//CreatePerson create a new person data
func CreatePerson(person *MstPersonInfo) (string, error) {
	Connect()
	defer Disconnect()
	var err error
	err = nil

	var newPersonCD string = ""

	strSQL := `insert into mst_person_info (person_cd, person_name, person_name_kana, birthday, sex, zip, address1, address2, address3, address4, tel, mobile, mail_address, authentication_date, delete_flag, create_user, create_date, update_user, update_date) 
	values 
	( to_char(CURRENT_DATE,'YY') || LPAD(nextval('SEQ_PERSON_CD')::varchar, 8, '0'),$1,$2, $3, $4,$5,$6, $7,$8, $9, $10,$11,$12,$13,$14,$15,CURRENT_TIMESTAMP,$16,CURRENT_TIMESTAMP)
	returning person_cd;`

	stmt, err := Db.Prepare(strSQL)

	err = stmt.QueryRow(&person.PersonName,
		&person.PersonNameKana,
		&person.Birthday,
		&person.Sex,
		&person.Zip,
		&person.Address1,
		&person.Address2,
		&person.Address3,
		&person.Address4,
		&person.Tel,
		&person.Mobile,
		&person.MailAddress,
		&person.AuthenticationDate,
		&person.DeleteFlag,
		&person.CreateUser,
		&person.UpdateUser).Scan(&newPersonCD)

	return newPersonCD, err

}

//UpdatePerson update person data
func UpdatePerson(person *MstPersonInfo, personcd string) error {
	Connect()
	defer Disconnect()
	var err error
	err = nil

	strSQL := `update mst_person_info 
		set  person_name = $1 
		, person_name_kana = $2 
		, birthday = $3 
		, sex = $4 
		, zip = $5 
		, address1 = $6 
		, address2 = $7 
		, address3 = $8 
		, address4 = $9 
		, tel = $10 
		, mobile = $11 
		, mail_address = $12 
		, authentication_date = $13 
		, delete_flag = $14 
		, create_user = $15 
		, update_user = $16 
		, Update_date = current_timestamp
		where person_cd = $17;`

	_, err = Db.Exec(strSQL,
		&person.PersonName,
		&person.PersonNameKana,
		&person.Birthday,
		&person.Sex,
		&person.Zip,
		&person.Address1,
		&person.Address2,
		&person.Address3,
		&person.Address4,
		&person.Tel,
		&person.Mobile,
		&person.MailAddress,
		&person.AuthenticationDate,
		&person.DeleteFlag,
		&person.CreateUser,
		&person.UpdateUser,
		&person.PersonCd)

	return err
}

//DeletePerson delete person data by person code
func DeletePerson(personcd string) (int64, error) {
	Connect()
	defer Disconnect()

	result, err := Db.Exec("delete from mst_person_info where person_cd = $1;", personcd)

	rowCnt, err := result.RowsAffected()

	return rowCnt, err
}
