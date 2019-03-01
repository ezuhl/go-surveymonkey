package tests

import (
	"testing"
	"github.com/ezuhl/go-surveymonkey"
	"os"
	"github.com/ezuhl/go-surveymonkey/request"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

var access = os.Getenv("SM_API_KEY")
var url = os.Getenv("SM_URL")
//
func TestGetSurveys(t *testing.T) {

	//id 164963984

	api := surveymonkey.NewSMAPI(url,access,"dev")
	resp, err := api.GetSurveys(request.SurveyRequest{})

	if err  != nil {
		t.Log(err)
		return
	}

	for _, item := range resp.Data {
		t.Logf("%s- %s - %s", item.Title,item.ID, item.Href)
	}




}

 //
 func TestCategories(t *testing.T) {
	//
	//Access Token:
	//

	api := surveymonkey.NewSMAPI(url,access,"dev")
	resp, err := api.GetSurveyCategory(request.CategoryRequest{Page:1})

	if err  != nil {
		t.Log(err)
		return
	}


	t.Logf("items %v\n", resp)


 }




func TestSurveyPage(t *testing.T) {
	//
	//Access Token:
	//

	api := surveymonkey.NewSMAPI(url,access,"dev")



	resp, err := api.GetSurveyPages(164965608,request.PageRequest{Page:1})

	if err  != nil {
		t.Log(err)
		return
	}


	t.Log("item: " , resp)


}

func TestSurveyQuestion(t *testing.T) {
	//
	//Access Token:
	//

	api := surveymonkey.NewSMAPI(url,access,"dev")



	resp, err := api.GetSurveyQuestions(164965608,56269018)

	if err  != nil {
		t.Log(err)
		return
	}


	t.Log("item: " , resp)


}

