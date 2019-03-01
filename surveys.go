package surveymonkey

import (
	"github.com/ezuhl/go-surveymonkey/response"
	"fmt"
	"github.com/ezuhl/go-surveymonkey/request"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */



 const (
 	surveyPath = "surveys"
 )



func (sm *SMApi) GetSurveys(params request.SurveyRequest) (*response.SurveyList,error) {

	req, err := sm.smHttp.NewRequest("GET",fmt.Sprintf("%s?%s",surveyPath,params),nil)

	if err != nil {
		return nil, err
	}

	models := &response.SurveyList{}
	resp,err := sm.smHttp.Do(req,&models)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return models,fmt.Errorf("request failure with code %d", resp.StatusCode)
	}

	return models, nil

}

