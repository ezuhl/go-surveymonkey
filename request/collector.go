package request

import "github.com/google/go-querystring/query"

/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.txt', which is part of this source code package.
 */

type CollectorRequest struct {
	Page      int    `url:"page,omitempty"`
	PerPage   int    `url:"per_page,omitempty"`
	SortBy    string `url:"sort_by,omitempty"`
	SortOrder string `url:"sort_order,omitempty"`
	Name      string `url:"name,omitempty"`
	StartDate int    `url:"start_date,omitempty"`
	EndDate   int    `url:"end_date,omitempty"`
	Include   int    `url:"include,omitempty"`
}

func (u CollectorRequest) String() string {

	values, err := query.Values(u)
	if err != nil {
		return ""
	}

	return values.Encode()
}
