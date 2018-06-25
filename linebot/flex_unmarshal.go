// Copyright 2018 LINE Corporation
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
	"errors"
	"reflect"
)

type typedObject struct {
	Type string `json:"type"`
}

var flexComponentTypeMapping = map[FlexComponentType]FlexComponent{
	FlexComponentTypeBox:       &BoxComponent{},
	FlexComponentTypeButton:    &ButtonComponent{},
	FlexComponentTypeFiller:    &FillerComponent{},
	FlexComponentTypeIcon:      &IconComponent{},
	FlexComponentTypeImage:     &ImageComponent{},
	FlexComponentTypeSeparator: &SeparatorComponent{},
	FlexComponentTypeSpacer:    &SpacerComponent{},
	FlexComponentTypeText:      &TextComponent{},
}

// UnmarshalFlexMessageJSON function
func UnmarshalFlexMessageJSON(data []byte) (FlexContainer, error) {
	var o typedObject
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	switch FlexContainerType(o.Type) {
	case FlexContainerTypeBubble:
		var container BubbleContainer
		if err := json.Unmarshal(data, &container); err != nil {
			return nil, err
		}
		return &container, nil
	case FlexContainerTypeCarousel:
		var container CarouselContainer
		if err := json.Unmarshal(data, &container); err != nil {
			return nil, err
		}
		return &container, nil
	default:
		return nil, errors.New("invalid container type")
	}
}

// UnmarshalJSON method for BoxComponent
func (c *BoxComponent) UnmarshalJSON(data []byte) error {
	type alias BoxComponent
	a := struct {
		Contents []json.RawMessage `json:"contents"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	c.Contents = make([]FlexComponent, 0, len(a.Contents))
	for _, content := range a.Contents {
		o := struct {
			Type FlexComponentType `json:"type"`
		}{}
		if err := json.Unmarshal(content, &o); err != nil {
			return err
		}

		if v, ok := flexComponentTypeMapping[o.Type]; ok {
			t := reflect.TypeOf(v)
			component := reflect.New(t).Interface()
			if err := json.Unmarshal(content, component); err != nil {
				return err
			}
			c.Contents = append(c.Contents, component)
		} else {
			return errors.New("invalid component type")
		}
	}
	return nil
}
