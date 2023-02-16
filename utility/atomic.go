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

package utility

import (
	"net/http"
)

func AtomicSiteHeader(header http.Header) string {
	key := "x-atomic-site"

	if v := header.Get(key); v != "" {
		return v
	}

	// header.Get() internally canonicalizes key names, but metadata.Pairs uses
	// lowercase keys and returns arrays of values. Using the lowercase key
	// names and only returning the first item allows this function to be used
	// for gRPC metadata.
	if v, ok := header[key]; ok && len(v) > 0 {
		return v[0]
	}
	return ""
}
