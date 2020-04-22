package lib

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
