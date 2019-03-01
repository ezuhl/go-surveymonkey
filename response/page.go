package response

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

type SurveyPage struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Data    []struct {
		ID          string `json:"id"`
		Position    int    `json:"position"`
		Href        string `json:"href"`
		Description string `json:"description"`
		Title       string `json:"title"`
	} `json:"data"`
	Page  int `json:"page"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
