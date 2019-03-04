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

func (sm *SMApi) GetSurveyPages(id int, param request.PageRequest) (*response.SurveyPage, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s/%d/pages?%s", surveyPath, id, param), nil)

	if err != nil {
		return nil, err
	}

	models := &response.SurveyPage{}
	resp, err := sm.smHttp.Do(req, &models)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return models, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return models, nil

}
