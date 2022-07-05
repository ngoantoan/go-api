package controller

import (
	"fmt"
	"net/http"
	"strings"

	dto "seoulspa_api/dto/comments"
	entity "seoulspa_api/entity/comments"
	"seoulspa_api/helper"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	GetListComment(ctx *gin.Context)
}

func GetListComment(ctx *gin.Context) {
	var request dto.GetListCommentDTO
	errDTO := ctx.ShouldBindJSON(&request)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	query := `SELECT
			pc.*,
			wpstr.web_page_id,
			wpstr.post_id_on_web,
			wpstr.title,wpr.page_id_on_web
		FROM post_comments AS pc
		JOIN web_post_resources AS wpstr ON pc.post_id = wpstr.id
		JOIN online_domains AS wpr ON wpstr.web_page_id = wpr.id
		WHERE wpstr.post_id_on_web IN(` + arrayToString(request.Post_id, ",") + `) AND pc.parent_id = 0 AND wpr.page_id_on_web = (` + request.Page_id + `)
	`

	if request.Is_question == 0 {
		query += " AND pc.star = 0"
	} else {
		query += " AND pc.star > 0"
	}

	if request.Limit == "" {
		query += " LIMIT 5"
	} else {
		query += " LIMIT " + request.Offset + "," + request.Limit
	}

	rows, err := db.Raw(query).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var c entity.ListComment
	var cs []entity.ListComment

	for rows.Next() {
		rows.Scan(
			&c.ID,
			&c.Name,
			&c.Phone,
			&c.Comment,
			&c.Post_id,
			&c.Page_id,
			&c.Image,
			&c.Gender,
			&c.Public,
			&c.Created,
			&c.Type,
			&c.Parent_id,
			&c.Url,
			&c.Post_title,
			&c.Collapse_status,
			&c.Satisfied,
			&c.Unsatisfied,
			&c.Useful,
			&c.Old_id,
			&c.Star,
			&c.Web_page_id,
			&c.Post_id_on_web,
			&c.Title,
			&c.Page_id_on_web,
		)

		cs = append(cs, c)
	}

	res := helper.BuildResponse(true, "OK", cs)
	ctx.JSON(http.StatusOK, res)
}

func arrayToString(a []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
