// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package config

import (
	"fmt"
	"net/url"
)

type urls []*url.URL

func (u *urls) Unpack(c interface{}) error {
	if c == nil {
		return nil
	}
	hosts, ok := c.([]interface{})
	if !ok {
		return fmt.Errorf("hosts must be a list, got: %#v", c)
	}

	if len(hosts) == 0 {
		return nil
	}

	nu := make(urls, len(hosts))
	for i, host := range hosts {
		h, ok := host.(string)
		if !ok {
			return fmt.Errorf("host must be a string, got: %#v", h)
		}
		url, err := url.Parse(h)
		if err != nil {
			return err
		}
		nu[i] = url
	}
	*u = nu

	return nil
}
