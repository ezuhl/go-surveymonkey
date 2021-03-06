package go_surveymonkey

import (
	"github.com/ezuhl/go-surveymonkey/client"
	"time"
)

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

const (
	apiTimeOut = time.Second * 5
)

type SMApi struct {
	smHttp    *client.SMHttp
	cacheTime time.Duration
}

func NewSMAPI(apiUrl, apiKey, env string) *SMApi {

	s := &SMApi{}
	s.smHttp = client.NewSMHttp(apiUrl, apiKey, env, apiTimeOut)

	return s
}
