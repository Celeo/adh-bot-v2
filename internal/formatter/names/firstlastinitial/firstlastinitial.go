/*
 * Copyright Daniel Hawton
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package firstlastinitial

import (
	"fmt"

	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/internal/formatter/staff"
)

func GenerateNameFromUser(u *dto.UserResponse, staffformat, staffsep string) string {
	title := staff.GetTitle(u, staffformat, staffsep)
	if title != "" {
		title = fmt.Sprintf(" | %s", title)
	}

	if len(fmt.Sprintf("%s %s.%s", u.FirstName, u.LastName[:1], title)) > 32 {
		diff := len(fmt.Sprintf("%s %s.%s", u.FirstName, u.LastName[:1], title)) - 32 - 3
		return fmt.Sprintf("%s %s.%s", u.FirstName[:len(u.FirstName)-diff], u.LastName[:1], title)
	}

	return fmt.Sprintf("%s %s.%s", u.FirstName, u.LastName[:1], title)
}
