package linebot

import (
	"encoding/json"
	"errors"
)

// UnmarshalTemplateJSON function
// Deprecated: Use OpenAPI based classes instead.
func UnmarshalTemplateJSON(data []byte) (Template, error) {
	raw := rawTemplate{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	return raw.Template, nil
}

// Deprecated: Use OpenAPI based classes instead.
type rawTemplate struct {
	Type     TemplateType `json:"type"`
	Template Template     `json:"-"`
}

func (t *rawTemplate) UnmarshalJSON(data []byte) error {
	type alias rawTemplate
	raw := alias{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var template Template
	switch raw.Type {
	case TemplateTypeButtons:
		template = &ButtonsTemplate{}
	case TemplateTypeConfirm:
		template = &ConfirmTemplate{}
	case TemplateTypeCarousel:
		template = &CarouselTemplate{}
	case TemplateTypeImageCarousel:
		template = &ImageCarouselTemplate{}
	default:
		return errors.New("invalid template type")
	}
	if err := json.Unmarshal(data, template); err != nil {
		return err
	}
	t.Type = raw.Type
	t.Template = template
	return nil
}

// UnmarshalJSON method for ButtonsTemplate
func (t *ButtonsTemplate) UnmarshalJSON(data []byte) error {
	type alias ButtonsTemplate
	raw := struct {
		Actions []rawAction `json:"actions"`
		*alias
	}{
		alias: (*alias)(t),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	actions := make([]TemplateAction, len(raw.Actions))
	for i, action := range raw.Actions {
		actions[i] = action.Action
	}
	t.Actions = actions
	return nil
}

// UnmarshalJSON method for ConfirmTemplate
func (t *ConfirmTemplate) UnmarshalJSON(data []byte) error {
	type alias ConfirmTemplate
	raw := struct {
		Actions []rawAction `json:"actions"`
		*alias
	}{
		alias: (*alias)(t),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	actions := make([]TemplateAction, len(raw.Actions))
	for i, action := range raw.Actions {
		actions[i] = action.Action
	}
	t.Actions = actions
	return nil
}

// Deprecated: Use OpenAPI based classes instead.
type rawColumn struct {
	Column CarouselColumn `json:"-"`
}

func (c *rawColumn) UnmarshalJSON(data []byte) error {
	type alias CarouselColumn
	raw := struct {
		DefaultAction rawAction   `json:"defaultAction"`
		Actions       []rawAction `json:"actions"`
		*alias
	}{
		alias: (*alias)(&c.Column),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.Column.DefaultAction = raw.DefaultAction.Action
	actions := make([]TemplateAction, len(raw.Actions))
	for i, action := range raw.Actions {
		actions[i] = action.Action
	}
	c.Column.Actions = actions
	return nil
}

// UnmarshalJSON method for CarouselTemplate
func (t *CarouselTemplate) UnmarshalJSON(data []byte) error {
	type alias CarouselTemplate
	raw := struct {
		Columns []rawColumn `json:"columns"`
		*alias
	}{
		alias: (*alias)(t),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	columns := make([]*CarouselColumn, len(raw.Columns))
	for i, column := range raw.Columns {
		c := column.Column
		columns[i] = &c
	}
	t.Columns = columns
	return nil
}

// Deprecated: Use OpenAPI based classes instead.
type rawImageColumn struct {
	Column ImageCarouselColumn `json:"-"`
}

func (c *rawImageColumn) UnmarshalJSON(data []byte) error {
	type alias ImageCarouselColumn
	raw := struct {
		Action rawAction `json:"action"`
		*alias
	}{
		alias: (*alias)(&c.Column),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.Column.Action = raw.Action.Action
	return nil
}

// UnmarshalJSON method for ImageCarouselTemplate
func (t *ImageCarouselTemplate) UnmarshalJSON(data []byte) error {
	type alias ImageCarouselTemplate
	raw := struct {
		Columns []rawImageColumn `json:"columns"`
		*alias
	}{
		alias: (*alias)(t),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	columns := make([]*ImageCarouselColumn, len(raw.Columns))
	for i, column := range raw.Columns {
		c := column.Column
		columns[i] = &c
	}
	t.Columns = columns
	return nil
}
