// Copyright 2020 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

func TestSelector(t *testing.T) {
	var testCases = []struct {
		Op   *Operator
		Want []byte
	}{
		{
			Op: OpAnd(
				NewAudienceObject(5614991017776),
				OpNot(
					NewAudienceObject(4389303728991),
				),
			),
			Want: []byte(`{"type":"operator","and":[{"type":"audience","audienceGroupId":5614991017776},{"type":"operator","not":{"type":"audience","audienceGroupId":4389303728991}}]}`),
		},
		{
			Op: OpAnd(
				NewGenderFilter(GenderMale, GenderFemale), // one of male/female
				NewAgeFilter(Age20, Age25),                // >= 20 && < 25
				NewAppTypeFilter(AppTypeAndroid, AppTypeIOS),
				NewAreaFilter(AreaJPAichi, AreaJPAkita),
				NewSubscriptionPeriodFilter(PeriodDay7, PeriodDay30),
			),
			Want: []byte(`{"type":"operator","and":[{"type":"gender","oneOf":["male","female"]},{"type":"age","gte":"age_20","lt":"age_25"},{"type":"appType","oneOf":["android","ios"]},{"type":"area","oneOf":["jp_23","jp_05"]},{"type":"subscriptionPeriod","gte":"day_7","lt":"day_30"}]}`),
		},
		{
			Op: OpOr(
				OpAnd(
					NewGenderFilter(GenderMale, GenderFemale), // one of male/female
					NewAgeFilter(Age20, Age25),                // >= 20 && < 25
					NewAppTypeFilter(AppTypeAndroid, AppTypeIOS),
					NewAreaFilter(AreaJPAichi, AreaJPAkita),
					NewSubscriptionPeriodFilter(PeriodDay7, PeriodDay30),
				),
				OpAnd(
					NewAgeFilter(Age35, Age40),         // >= 35 && < 40
					OpNot(NewGenderFilter(GenderMale)), // not male
				),
			),
			Want: []byte(`{"type":"operator","or":[{"type":"operator","and":[{"type":"gender","oneOf":["male","female"]},{"type":"age","gte":"age_20","lt":"age_25"},{"type":"appType","oneOf":["android","ios"]},{"type":"area","oneOf":["jp_23","jp_05"]},{"type":"subscriptionPeriod","gte":"day_7","lt":"day_30"}]},{"type":"operator","and":[{"type":"age","gte":"age_35","lt":"age_40"},{"type":"operator","not":{"type":"gender","oneOf":["male"]}}]}]}`),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotJSON, err := json.Marshal(tc.Op)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(gotJSON, tc.Want) {
				t.Errorf("json \n%s, want \n%s", gotJSON, tc.Want)
			}
		})
	}
}
