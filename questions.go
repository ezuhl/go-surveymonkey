package surveymonkey

import (
	"fmt"
	"github.com/ezuhl/go-surveymonkey/response"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

func (sm *SMApi) GetSurveyQuestions(surveyId, pageId int) (*response.SurveyListQuestion, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d/pages/%d/questions", surveyPath, surveyId, pageId), nil)

	if err != nil {
		return nil, err
	}

	models := &response.SurveyListQuestion{}
	resp, err := sm.smHttp.Do(req, &models)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return models, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return models, nil

}
