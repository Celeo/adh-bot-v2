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

package models

import "github.com/adh-partnership/api/pkg/database"

type Bot struct {
	ID    int64
	Key   string
	Value string
}

func Get(key string) (string, error) {
	var bot Bot
	err := database.DB.Where(&Bot{Key: key}).First(&bot).Error
	if err != nil {
		return "", err
	}

	return bot.Value, nil
}

func Set(key, value string) error {
	bot := Bot{
		Key:   key,
		Value: value,
	}

	err := database.DB.Save(&bot).Error
	if err != nil {
		return err
	}

	return nil
}