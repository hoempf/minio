/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lifecycle

import "encoding/xml"

// AbortIncompleteMultipartUpload is unsupported, but some clients want to see
// the policy when setting it.
type AbortIncompleteMultipartUpload struct {
	XMLName             xml.Name `xml:"AbortIncompleteMultipartUpload"`
	DaysAfterInitiation int      `xml:"DaysAfterInitiation,omitempty"`
}

// MarshalXML so Pure Storage is happy.
func (a AbortIncompleteMultipartUpload) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if a.DaysAfterInitiation == 0 {
		a.DaysAfterInitiation = 1
	}
	type abortIncompleteMultipartUpload AbortIncompleteMultipartUpload
	return e.EncodeElement(abortIncompleteMultipartUpload(a), start)
}
