package response

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */



type Template struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Data    []struct {
		Category     string `json:"category"`
		Name         string `json:"name"`
		Title        string `json:"title"`
		Available    bool   `json:"available"`
		ID           string `json:"id"`
		NumQuestions int    `json:"num_questions"`
		Description  string `json:"description"`
		PreviewLink  string `json:"preview_link"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}