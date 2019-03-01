package response

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */




type SurveyCategoryModel struct {
	Page  int `json:"page"`
	PerPage int `json:"per_page"`
	Total int `json:"total"`
	Links map[string]string `json:"links"`
	Data  []map[string]string `json:"data"`

}



type Questions struct {
	Question []interface{} `json:"questions"`
}




type SurveyModel struct {
	Title string `json:"title,omitempty"`
	NickName string `json:"nickname,omitempty"`
	Category string `json:"category,omitempty"`
	Language string `json:"language"`
	QuestionCount int `json:"question_count"`
	PageCount string `json:"page_count"`
	DateCreated string `json:"date_created"`
	DateModified string `json:"date_modified"`
	Id string `json:"id"`
	FolderId string `json:"folder_id,omitempty"`
	Href string `json:"href"`
	Pages []Questions `json:"pages,omitempty"`
	ButtonsText map[string]string `json:"buttons_text"`
	CustomVariables map[string]string `json:"custom_variables"`
	Footer bool `json:"footer"`
	FooterId string `json:"footer_id"`
	Preview string `json:"preview"`
	EditUrl string `json:"edit_url"`
	CollectUrl string `json:"collect_url"`
	AnalyzeUrl string `json:"analyze_url"`
	SummaryUrl string `json:"summary_url"`
	ResponseCount int `json:"response_count,omitempty"`
}



type SurveyList struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Data    []struct {
		Name string `json:"nickname"`
		ID   string `json:"id"`
		Href string `json:"href"`
		Title string `json:"title"`
	} `json:"data"`
	Page  int `json:"page"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
