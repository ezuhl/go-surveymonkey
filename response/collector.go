package response

import "time"

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

type CollectorsResponse struct {
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
	Data    []struct {
		Href     string `json:"href"`
		Nickname string `json:"nickname"`
		ID       string `json:"id"`
		Title    string `json:"title"`
	} `json:"data"`
	Page  int `json:"page"`
	Links struct {
		Self string `json:"self"`
		Last string `json:"last"`
		Next string `json:"next"`
	} `json:"links"`
}

type CollectorResponse struct {
	Status                  string      `json:"status"`
	ID                      string      `json:"id"`
	Type                    string      `json:"type"`
	Name                    string      `json:"name"`
	Hash                    string      `json:"hash"`
	ThankYouMessage         string      `json:"thank_you_message"`
	DisqualificationMessage string      `json:"disqualification_message"`
	CloseDate               time.Time   `json:"close_date"`
	ClosedPageMessage       string      `json:"closed_page_message"`
	RedirectURL             string      `json:"redirect_url"`
	DisplaySurveyResults    bool        `json:"display_survey_results"`
	EditResponseType        string      `json:"edit_response_type"`
	AnonymousType           string      `json:"anonymous_type"`
	AllowMultipleResponses  bool        `json:"allow_multiple_responses"`
	DateModified            time.Time   `json:"date_modified"`
	URL                     string      `json:"url"`
	DateCreated             time.Time   `json:"date_created"`
	SenderEmail             interface{} `json:"sender_email"`
	PasswordEnabled         bool        `json:"password_enabled"`
	ResponseLimit           int         `json:"response_limit"`
	RedirectType            string      `json:"redirect_type"`
	Href                    string      `json:"href"`
}
