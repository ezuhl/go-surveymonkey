package request

import (
	"github.com/google/go-querystring/query"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */
type SurveyTitle struct {
	Title string `json:"title,omitempty"`
}

type SurveyRequest struct {
	Title           string `url:"title,omitempty"`
	Page            int    `url:"page,omitempty"`
	PerPage         int    `url:"per_page,omitempty"`
	SortBy          string `url:"sort_by,omitempty"`
	SortOrder       string `url:"sort_order,omitempty"`
	Include         string `url:"include,omitempty"`
	StartModifiedAt string `url:"start_modified_at,omitempty"`
	EndModifiedAt   string `url:"end_modified_at,omitempty"`
	FolderId        string `url:"folder_id,omitempty"`
}

func (u SurveyRequest) String() string {

	values, err := query.Values(u)
	if err != nil {
		return ""
	}

	return values.Encode()
}

type CreateSurvey struct {
	Page            int    `json:"page,omitempty"`
	PerPage         int    `json:"per_page,omitempty"`
	SortBy          string `json:"sort_by,omitempty"`    //title, date_modified, or num_responses
	SortOrder       string `json:"sort_order,omitempty"` //asc or desc
	Include         int    `string:"include,omitempty"`
	Title           string `json:"title,omitempty"`
	StartModifiedAt string `json:"start_modified_at,omitempty"` //YYYY-MM-DDTHH:MM:S
	EndModifiedAt   string `json:"end_modified_at,omitempty"`   //YYYY-MM-DDTHH:MM:S
	FolderId        string `json:"folder_id,omitempty"`
}

type ModifySurvey struct {
	Title           string            `json:"title,omitempty"`
	NickName        string            `json:"nickname,omitempty"`
	Language        string            `json:"language,omitempty"`
	ButtonsText     *ButtonsText      `json:"buttons_text,omitempty"`
	CustomVariables map[string]string `json:"custom_variables,omitempty"`
	Footer          bool              `json:"footer"`
}

type ButtonsText struct {
	NextButton string `json:"next_button,omitempty"`
	PrevButton string `json:"prev_button,omitempty"`
	ExitButton string `json:"exit_button,omitempty"`
	DoneButton string `json:"done_button,omitempty"`
}

type SurveyPage struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Position      int    `json:"position"`
	QuestionCount int    `json:"question_count"`
	ID            string `json:"id"`
	Href          string `json:"href"`
}
