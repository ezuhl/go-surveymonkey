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

const (
	templatePath = "survey_templates"
)

func (sm *SMApi) GetSurveyTemplate(params *request.TemplatesRequest) ([]*response.Template, error) {

	req, err := sm.smHttp.NewRequest("GET", fmt.Sprintf("%s?%s", templatePath, params), nil)

	if err != nil {
		return nil, err
	}

	models := make([]*response.Template, 0)
	resp, err := sm.smHttp.Do(req, &models)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return models, fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return models, nil

}
