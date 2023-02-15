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

//
// Automattic customization for Atomic APM
//

package utility_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/utility"
)

func TestAtomicSite(t *testing.T) {
	req := &http.Request{
		Header: make(http.Header),
	}

	// No site id header
	assert.Equal(t, "", utility.AtomicSiteHeader(req.Header))

	req.Header.Set("x-atomic-site", "123")
	assert.Equal(t, "123", utility.AtomicSiteHeader(req.Header))
}
