package dto

type GetListCommentDTO struct {
	Post_id     []int64 `json:"post_id" form:"post_id" binding:"required"`
	Page_id     string  `json:"page_id" form:"page_id" binding:"required"`
	Is_question uint64  `json:"is_question" form:"is_question"`
	Limit       string  `json:"limit" form:"limit"`
	Offset      string  `json:"offset" form:"offset"`
}
