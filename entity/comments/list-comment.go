package entity

//User represents users table in database
type ListComment struct {
	ID              uint64 `json:"id"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Comment         string `json:"comment"`
	Post_id         uint64 `json:"post_id"`
	Page_id         uint64 `json:"page_id"`
	Image           string `json:"image"`
	Gender          uint64 `json:"gender"`
	Public          uint64 `json:"public"`
	Created         string `json:"created"`
	Type            uint64 `json:"type"`
	Parent_id       uint64 `json:"parent_id"`
	Url             string `json:"url"`
	Post_title      string `json:"post_title"`
	Collapse_status uint64 `json:"collapse_status"`
	Satisfied       uint64 `json:"satisfied"`
	Unsatisfied     uint64 `json:"unsatisfied"`
	Useful          uint64 `json:"useful"`
	Old_id          uint64 `json:"old_id"`
	Star            uint64 `json:"star"`
	Web_page_id     uint64 `json:"web_page_id"`
	Post_id_on_web  uint64 `json:"post_id_on_web"`
	Title           string `json:"title"`
	Page_id_on_web  uint64 `json:"page_id_on_web"`
}
