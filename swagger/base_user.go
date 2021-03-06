/* 
 * Project Limitless Builtin API
 *
 * This specification defines the builtin API available for Project Limitless installations.
 *
 * OpenAPI spec version: 0.0.0.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"time"
)

type BaseUser struct {

	ID int32 `json:"iD,omitempty"`

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	DateCreated time.Time `json:"dateCreated,omitempty"`

	IsDeleted bool `json:"isDeleted,omitempty"`
}
