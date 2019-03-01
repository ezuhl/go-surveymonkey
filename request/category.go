package request

import "github.com/google/go-querystring/query"

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */






type CategoryRequest struct {
	Page int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
	Language string `url:"language,omitempty"`

}

func (u CategoryRequest) String() string {

	values, err := query.Values(u)
	if err != nil {
		return ""
	}

	return values.Encode()
}
