package go_surveymonkey

import (
	"fmt"
	"github.com/ezuhl/go-surveymonkey/request"
	"github.com/ezuhl/go-surveymonkey/response"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

func (sm *SMApi) CreateSurvey(options *request.CreateSurvey) (*response.SurveyModel, error) {

	req, err := sm.smHttp.NewRequest("POST", surveyPath, options)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyModel{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ReplaceSurvey(id int, options *request.CreateSurvey) (*response.SurveyModel, error) {

	req, err := sm.smHttp.NewRequest("PUT", fmt.Sprintf("%s/%d", surveyPath, id), options)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyModel{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) DeleteSurvey(id int) error {

	req, err := sm.smHttp.NewRequest("DELETE", fmt.Sprintf("%s/%d", surveyPath, id), nil)

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

func (sm *SMApi) GetSurveyById(id int) (*response.SurveyModel, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d", surveyPath, id), nil)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyModel{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ModifySurveyById(id int, mod *request.ModifySurvey) (*response.SurveyModel, error) {

	req, err := sm.smHttp.NewRequest("PATCH", fmt.Sprintf("%s/%d", surveyPath, id), mod)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyModel{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) GetSurveyDetail(id int) (*response.SurveyModel, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d/details", surveyPath, id), nil)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyModel{}
	resp, err := sm.smHttp.Do(req, model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}
