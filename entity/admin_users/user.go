package entity

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `admin_users`
func TableName() string {
	return "admin_users"
}

//User represents users table in database
type Admin_users struct {
	ID                      uint64 `json:"id"`
	Phone                   string `json:"phone"`
	Username                string `json:"username"`
	Password                string `json:"-"`
	Email                   string `json:"email"`
	Activation_code         string `json:"activation_code"`
	Forgotten_password_code string `json:"forgotten_password_code"`
	Forgotten_password_time uint64 `json:"forgotten_password_time"`
	Remember_code           string `json:"remember_code"`
	Date_start              string `json:"date_start"`
	Date_end                string `json:"date_end"`
	Created_on              uint64 `json:"created_on"`
	Last_login              uint64 `json:"last_login"`
	Online                  uint64 `json:"online"`
	Last_action             uint64 `json:"last_action"`
	Active                  uint64 `json:"active"`
	Contract                uint64 `json:"contract"`
	Verify                  uint64 `json:"verify"`
	First_name              string `json:"first_name"`
	Sex                     uint64 `json:"sex"`
	Address                 string `json:"address"`
	Address_contact         string `json:"address_contact"`
	Phone_contact           string `json:"phone_contact"`
	Id_card                 string `json:"id_card"`
	Id_date                 string `json:"id_date"`
	Birthday                string `json:"birthday"`
	Line                    string `json:"line "`
	Admin_user_shift_id     uint64 `json:"admin_user_shift_id"`
	Image_url               string `json:"image_url"`
	Main_store_id           uint64 `json:"main_store_id"`
	Can_work                uint64 `json:"can_work"`
	Wp_id                   string `json:"wp_id"`
	Id_cls                  uint64 `json:"id_cls"`
	Checkin_late_allow      uint64 `json:"checkin_late_allow"`
	Checkout_early_allow    uint64 `json:"checkout_early_allow"`
	Literacy_id             uint64 `json:"literacy_id"`
	Nation_id               uint64 `json:"nation_id"`
	Province_id             string `json:"province_id"`
	District_id             uint64 `json:"district_id"`
	Tag_name                string `json:"tag_name"`
	Reason_id               uint64 `json:"reason_id"`
	Contract_id             uint64 `json:"contract_id"`
	Main_group_id           uint64 `json:"main_group_id"`
	Parent_id               uint64 `json:"parent_id"`
	Status                  uint64 `json:"status"`
	Position                string `json:"position"`
	User_created_id         uint64 `json:"user_created_id"`
	Is_return               uint64 `json:"is_return"`
	Position_id             uint64 `json:"position_id"`
	Company_id              uint64 `json:"company_id"`
	User_permission         string `json:"user_permission"`
	Customer_id             uint64 `json:"customer_id"`
	Level_id                uint64 `json:"level_id"`
	Degree_id               uint64 `json:"degree_id"`
	Collabor_type           string `json:"collabor_type"`
	Bank_account            string `json:"bank_account"`
	Bank_number             string `json:"bank_number"`
	Is_deleted              uint64 `json:"is_deleted"`

	Token string `json:"token,omitempty"`
}
