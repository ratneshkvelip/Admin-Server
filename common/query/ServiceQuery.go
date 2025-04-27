package query

func CreateAccount() (serviceQuery string) {
	serviceQuery = `INSERT INTO
				public.user_data
				(user_id,user_name,email_id,gender,mobile_no,country,login_token,user_password,status,created_by,updated_by)
				VALUES
				($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
	`
	return serviceQuery
}

func GetPassword() (serviceQuery string) {
	serviceQuery = `Select user_password from public.user_data where user_id=$1 OR email_id=$1 or mobile_no=$1
	`
	return serviceQuery
}

func GetUserName() (serviceQuery string) {
	serviceQuery = `Select user_id from public.user_data where user_id=$1
	`
	return serviceQuery
}

func GetUserMobile() (serviceQuery string) {
	serviceQuery = `Select mobile_no from public.user_data where mobile_no=$1
	`
	return serviceQuery
}

func GetUserEmailId() (serviceQuery string) {
	serviceQuery = `Select email_id from public.user_data where email_id=$1
	`
	return serviceQuery
}

func GetName() (serviceQuery string) {
	serviceQuery = `Select user_name from data1 where id=1
	`
	return serviceQuery
}
