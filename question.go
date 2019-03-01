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

func (sm *SMApi) CreateQuestion(surveyId, pageId int, question *request.SurveyQuestion) (*response.SurveyQuestion, error) {

	req, err := sm.smHttp.NewRequest("POST", fmt.Sprintf("%s/%d/pages/%d/questions", surveyPath, surveyId, pageId), question)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyQuestion{}
	resp, err := sm.smHttp.Do(req, &model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ReplaceQuestion(surveyId, pageId int, question *request.SurveyQuestion) (*response.SurveyQuestion, error) {

	req, err := sm.smHttp.NewRequest("PUT", fmt.Sprintf("%s/%d/pages/%d/questions", surveyPath, surveyId, pageId), question)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyQuestion{}
	resp, err := sm.smHttp.Do(req, &model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) GetQuestion(surveyId, pageId int, questionId int) (*response.SurveyQuestion, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d/pages/%d/questions/%d", surveyPath, surveyId, pageId, questionId), nil)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyQuestion{}
	resp, err := sm.smHttp.Do(req, &model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) ModifyQuestion(surveyId, pageId, questionId int, questionItem interface{}) (*response.SurveyQuestion, error) {

	req, err := sm.smHttp.NewRequest("PATCH", fmt.Sprintf("%s/%d/pages/%d/questions/%d", surveyPath, surveyId, pageId, questionId), questionItem)

	if err != nil {
		return nil, err
	}

	model := &response.SurveyQuestion{}
	resp, err := sm.smHttp.Do(req, &model)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return model, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return model, nil

}

func (sm *SMApi) DeleteQuestion(surveyId, pageId int, questionId int) error {

	req, err := sm.smHttp.NewRequest("DELETE", fmt.Sprintf("%s/%d/pages/%d/questions/%d", surveyPath, surveyId, pageId, questionId), nil)

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
