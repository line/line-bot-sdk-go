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
)

// UnmarshalFlexMessageJSON function
func UnmarshalFlexMessageJSON(data []byte) (FlexContainer, error) {
	o := struct {
		Type FlexContainerType `json:"type"`
	}{}
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}
	switch o.Type {
	case FlexContainerTypeBubble:
		container := BubbleContainer{}
		if err := json.Unmarshal(data, &container); err != nil {
			return nil, err
		}
		return &container, nil
	case FlexContainerTypeCarousel:
		container := CarouselContainer{}
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

	c.Contents = make([]FlexComponent, len(a.Contents))
	for i, contentData := range a.Contents {
		o := struct {
			Type FlexComponentType `json:"type"`
		}{}
		if err := json.Unmarshal(contentData, &o); err != nil {
			return err
		}
		switch o.Type {
		case FlexComponentTypeBox:
			component := BoxComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeButton:
			component := ButtonComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeFiller:
			component := FillerComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeIcon:
			component := IconComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeImage:
			component := ImageComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeSeparator:
			component := SeparatorComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeSpacer:
			component := SpacerComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		case FlexComponentTypeText:
			component := TextComponent{}
			if err := json.Unmarshal(contentData, &component); err != nil {
				return err
			}
			c.Contents[i] = &component
		default:
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
		action, err := unmarshalTemplateAction(a.Action)
		if err != nil {
			return err
		}
		c.Action = action
	}
	return nil
}

// UnmarshalJSON method for FillerComponent
func (c *FillerComponent) UnmarshalJSON(data []byte) error {
	type alias FillerComponent
	a := struct {
		*alias
	}{
		alias: (*alias)(c),
	}
	return json.Unmarshal(data, &a)
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
		action, err := unmarshalTemplateAction(a.Action)
		if err != nil {
			return err
		}
		c.Action = action
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
		action, err := unmarshalTemplateAction(a.Action)
		if err != nil {
			return err
		}
		c.Action = action
	}
	return nil
}

func unmarshalTemplateAction(data []byte) (TemplateAction, error) {
	o := struct {
		Type TemplateActionType `json:"type"`
	}{}
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, err
	}
	switch o.Type {
	case TemplateActionTypeURI:
		action := URITemplateAction{}
		if err := json.Unmarshal(data, &action); err != nil {
			return nil, err
		}
		return &action, nil
	case TemplateActionTypeMessage:
		action := MessageTemplateAction{}
		if err := json.Unmarshal(data, &action); err != nil {
			return nil, err
		}
		return &action, nil
	case TemplateActionTypePostback:
		action := PostbackTemplateAction{}
		if err := json.Unmarshal(data, &action); err != nil {
			return nil, err
		}
		return &action, nil
	case TemplateActionTypeDatetimePicker:
		action := DatetimePickerTemplateAction{}
		if err := json.Unmarshal(data, &action); err != nil {
			return nil, err
		}
		return &action, nil
	default:
		return nil, errors.New("invalid action type")
	}
}
