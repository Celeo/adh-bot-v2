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

package facility

import (
	"encoding/json"
	"fmt"

	"github.com/adh-partnership/api/pkg/database/dto"

	"github.com/vpaza/bot/pkg/cache"
	"github.com/vpaza/bot/pkg/network"
)

var ErrUserNotFound = fmt.Errorf("user not found")

func (f *Facility) GetRoster() []*dto.UserResponse {
	content, err := cache.Get(
		fmt.Sprintf("/%s/roster", f.Facility),
	)
	if err != nil {
		status, content, err := network.Call(
			"GET",
			fmt.Sprintf("%s/v1/user/all", f.API),
			"application/json",
			nil,
			nil,
		)
		if err != nil {
			log.Errorf("Failed to get roster: %s", err)
			return nil
		}

		if status != 200 {
			log.Errorf("Failed to get roster: %s", content)
			return nil
		}

		err = cache.Set(
			fmt.Sprintf("/%s/roster", f.Facility),
			content,
		)
		if err != nil {
			log.Errorf("Failed to cache %s roster: %s", f.Facility, err)
		}
	}

	var roster []*dto.UserResponse
	err = json.Unmarshal(content, &roster)
	if err != nil {
		log.Errorf("Failed to unmarshal roster: %s", err)
		return nil
	}

	return roster
}

func (f *Facility) FindUserByDiscordID(id string) (*dto.UserResponse, error) {
	roster := f.GetRoster()
	if roster == nil {
		return nil, fmt.Errorf("failed to get roster")
	}

	for _, user := range roster {
		if user.DiscordID == id {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}
