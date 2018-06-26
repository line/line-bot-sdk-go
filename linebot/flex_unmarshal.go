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

var flexContainerTypeMapping = map[FlexContainerType]FlexContainer{
	FlexContainerTypeBubble:   &BubbleContainer{},
	FlexContainerTypeCarousel: &CarouselContainer{},
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

var templateActionTypeMapping = map[TemplateActionType]TemplateAction{
	TemplateActionTypeURI:            &URITemplateAction{},
	TemplateActionTypeMessage:        &MessageTemplateAction{},
	TemplateActionTypePostback:       &PostbackTemplateAction{},
	TemplateActionTypeDatetimePicker: &DatetimePickerTemplateAction{},
}

// UnmarshalFlexMessageJSON function
func UnmarshalFlexMessageJSON(data []byte) (FlexContainer, error) {
	o := struct {
		Type FlexContainerType `json:"type"`
	}{}
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	if v, ok := flexContainerTypeMapping[o.Type]; ok {
		i := reflect.New(reflect.TypeOf(v)).Interface()
		if err := json.Unmarshal(data, i); err != nil {
			return nil, err
		}
		if container, ok := reflect.Indirect(reflect.ValueOf(i)).Interface().(FlexContainer); ok {
			return container, nil
		}
	}
	return nil, errors.New("invalid container type")
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
	for _, contentData := range a.Contents {
		o := struct {
			Type FlexComponentType `json:"type"`
		}{}
		if err := json.Unmarshal(contentData, &o); err != nil {
			return err
		}
		if v, ok := flexComponentTypeMapping[o.Type]; ok {
			i := reflect.New(reflect.TypeOf(v)).Interface()
			if err := json.Unmarshal(contentData, i); err != nil {
				return err
			}
			if content, ok := reflect.Indirect(reflect.ValueOf(i)).Interface().(FlexComponent); ok {
				c.Contents = append(c.Contents, content)
			}
		} else {
			return errors.New("invalid component type")
		}
	}
	return nil
}

// UnmarshalJSON method for ButtonComponent
func (c *ButtonComponent) UnmarshalJSON(data []byte) error {
	type alias ButtonComponent
	a := struct {
		Action json.RawMessage `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	{
		o := struct {
			Type TemplateActionType `json:"type"`
		}{}
		if err := json.Unmarshal(a.Action, &o); err != nil {
			return err
		}
		if v, ok := templateActionTypeMapping[o.Type]; ok {
			i := reflect.New(reflect.TypeOf(v)).Interface()
			if err := json.Unmarshal(a.Action, i); err != nil {
				return err
			}
			if action, ok := reflect.Indirect(reflect.ValueOf(i)).Interface().(TemplateAction); ok {
				c.Action = action
			}
		} else {
			return errors.New("invalid action type")
		}
	}
	return nil
}

// UnmarshalJSON method for FillerComponent
func (c *FillerComponent) UnmarshalJSON(data []byte) error {
	return nil
}

// UnmarshalJSON method for IconComponent
func (c *IconComponent) UnmarshalJSON(data []byte) error {
	type alias IconComponent
	a := struct {
		*alias
	}{
		alias: (*alias)(c),
	}
	return json.Unmarshal(data, &a)
}

// UnmarshalJSON method for ImageComponent
func (c *ImageComponent) UnmarshalJSON(data []byte) error {
	type alias ImageComponent
	a := struct {
		Action json.RawMessage `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	if len(a.Action) > 0 {
		o := struct {
			Type TemplateActionType `json:"type"`
		}{}
		if err := json.Unmarshal(a.Action, &o); err != nil {
			return err
		}
		if v, ok := templateActionTypeMapping[o.Type]; ok {
			i := reflect.New(reflect.TypeOf(v)).Interface()
			if err := json.Unmarshal(a.Action, i); err != nil {
				return err
			}
			if action, ok := reflect.Indirect(reflect.ValueOf(i)).Interface().(TemplateAction); ok {
				c.Action = action
			}
		} else {
			return errors.New("invalid action type")
		}
	}
	return nil
}

// UnmarshalJSON method for SeparatorComponent
func (c *SeparatorComponent) UnmarshalJSON(data []byte) error {
	type alias SeparatorComponent
	a := struct {
		*alias
	}{
		alias: (*alias)(c),
	}
	return json.Unmarshal(data, &a)
}

// UnmarshalJSON method for SpacerComponent
func (c *SpacerComponent) UnmarshalJSON(data []byte) error {
	type alias SpacerComponent
	a := struct {
		*alias
	}{
		alias: (*alias)(c),
	}
	return json.Unmarshal(data, &a)
}

// UnmarshalJSON method for TextComponent
func (c *TextComponent) UnmarshalJSON(data []byte) error {
	type alias TextComponent
	a := struct {
		Action json.RawMessage `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	if len(a.Action) > 0 {
		o := struct {
			Type TemplateActionType `json:"type"`
		}{}
		if err := json.Unmarshal(a.Action, &o); err != nil {
			return err
		}
		if v, ok := templateActionTypeMapping[o.Type]; ok {
			i := reflect.New(reflect.TypeOf(v)).Interface()
			if err := json.Unmarshal(a.Action, i); err != nil {
				return err
			}
			if action, ok := reflect.Indirect(reflect.ValueOf(i)).Interface().(TemplateAction); ok {
				c.Action = action
			}
		} else {
			return errors.New("invalid action type")
		}
	}
	return nil
}
