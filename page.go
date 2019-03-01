package surveymonkey

import (
	"fmt"
	"github.com/ezuhl/go-surveymonkey/request"
	"github.com/ezuhl/go-surveymonkey/response"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

func (sm *SMApi) CreateSurveyPage(id string, params *request.SurveyPage) (*response.SurveyPage, error) {

	req, err := sm.smHttp.NewRequest("POST", fmt.Sprintf("%s/%s/pages", surveyPath, id), params)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyPage{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) GetSurveyPage(surveyId, pageId int) (*response.SurveyPage, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d/pages/%d", surveyPath, surveyId, pageId), nil)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyPage{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ModifySurveyPage(surveyId, pageId int, params *request.SurveyTitle) (*response.SurveyPage, error) {

	req, err := sm.smHttp.NewRequest("PATCH", fmt.Sprintf("%s/%d/pages/%d", surveyPath, surveyId, pageId), params)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyPage{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ReplaceSurveyPage(surveyId, pageId int, params *request.SurveyTitle) (*response.SurveyPage, error) {

	req, err := sm.smHttp.NewRequest("PUT", fmt.Sprintf("%s/%d/pages/%d", surveyPath, surveyId, pageId), params)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyPage{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) DeleteSurveyPage(surveyId, pageId int) error {

	req, err := sm.smHttp.NewRequest("DELETE", fmt.Sprintf("%s/%d/pages/%d", surveyPath, surveyId, pageId), nil)

	if err != nil {
		return err
	}

	resp, err := sm.smHttp.Do(req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return nil

}
