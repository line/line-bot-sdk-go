package linebot

import (
	"encoding/json"
)

// FlexType type
type FlexType string

// FlexType constants
const (
	FlexTypeFlex       FlexType = "flex"
)

// ContainerType type
type ContainerType string

// ContainerType constants
const (
	ContainerTypeBubble       ContainerType = "bubble"
	ContainerTypeCarousel     ContainerType = "carousel"
)

// ComponentType type
type ComponentType string

// ContainerType constants
const (
	ComponentTypeBox			ComponentType = "box"
	ComponentTypeButton     	ComponentType = "button"
	ComponentTypeFiller			ComponentType = "filler"
	ComponentTypeIcon			ComponentType = "icon"
	ComponentTypeImage			ComponentType = "image"
	ComponentTypeSeparator		ComponentType = "separator"
	ComponentTypeSpacer			ComponentType = "spacer"
	ComponentTypeText			ComponentType = "text"
)

// LayoutType type
type LayoutType string

// ContainerType constants
const (
	LayoutTypeHorizontal   LayoutType = "horizontal"
	LayoutTypeVertical     LayoutType = "vertical"
	LayoutTypeBaseline     LayoutType = "baseline"
)

// SizeType type
type SizeType string

// SizeType constants
const (
	SizeTypeNone		SizeType = "none"
	SizeTypeXxs			SizeType = "xxs"
	SizeTypeXs			SizeType = "xs"
	SizeTypeSm			SizeType = "sm"
	SizeTypeMd			SizeType = "md"
	SizeTypeLg			SizeType = "lg"
	SizeTypeXl			SizeType = "xl"
	SizeTypeXxl			SizeType = "xxl"
	SizeType3xl			SizeType = "3xl"
	SizeType4xl			SizeType = "4xl"
)

// AlignType type
type AlignType string

// AlignType constants
const (
	AlignTypeStart		AlignType = "start"
	AlignTypeEnd		AlignType = "end"
	AlignTypeCenter		AlignType = "center"
)

// AspectModeType type
type AspectModeType string

//AspectModeType constants
const (
	AspectModeTypeCover AspectModeType = "cover"
	AspectModeTypeFit	AspectModeType = "fit"
)

//	AspectRatioType type
type AspectRatioType string

// AspectRatioType constants
const (
	AspectRatioType1To1		AspectRatioType	= "1:1"
	AspectRatioType151To1	AspectRatioType	= "1.51:1"
	AspectRatioType1911To1	AspectRatioType	= "1.911:1"
	AspectRatioType4To3		AspectRatioType	= "4:3"
	AspectRatioType16To9	AspectRatioType	= "16:9"
	AspectRatioType20To13	AspectRatioType	= "20:13"
	AspectRatioType2To1		AspectRatioType	= "2:1"
	AspectRatioType3To1		AspectRatioType	= "3:1"
	AspectRatioType3To4		AspectRatioType	= "3:4"
	AspectRatioType9To16	AspectRatioType	= "9:16"
	AspectRatioType1To2		AspectRatioType	= "1:2"
	AspectRatioType1To3		AspectRatioType	= "1:3"
)

// SizeHeightType type
type SizeHeightType string

// SizeHeightType constants
const (
	SizeHeightTypeSm		SizeHeightType = "Sm"
	SizeHeightTypeMd		SizeHeightType = "Md"
)

// SizeWeightType type
type SizeWeightType string

// SizeWeightType constants
const (
	SizeWeightTypeRegular		SizeHeightType = "regular"
	SizeWeightTypeBold			SizeHeightType = "bold"
)

// ButtonStyleType type
type ButtonStyleType string

// SizeHeightType constants
const (
	SizeHeightTypeLink		ButtonStyleType = "link"
	SizeHeightTypePrimary	ButtonStyleType = "primary"
	SizeHeightTypeSecondary	ButtonStyleType = "secondary"
)

// ButtonStyleType type
type GravityType string

// SizeHeightType constants
const (
	GravityTypeTop		GravityType = "top"
	GravityTypeBottom	GravityType = "bottom"
	GravityTypeCenter	GravityType = "center"
)


// DirectionType type
type DirectionType string

// FlexType constants
const (
	DirectionTypeLtr       DirectionType = "ltr"
	DirectionTypeRtl       DirectionType = "rtl"
)


type FlexStyle struct {
	BackgroundColor 	string		`json:"backgroundColor, omitempty"`
	Separator 			bool		`json:"separator, omitempty"`
	SeparatorColor 		string		`json:"separatorColor, omitempty"`
}

type FlexStylesBlock struct {
	Header	FlexStyle	`json:"header, omitempty"`
	Hero	FlexStyle	`json:"hero, omitempty"`
	Body	FlexStyle	`json:"body, omitempty"`
	Footer	FlexStyle	`json:"footer, omitempty"`
}

type Flex struct {
	AltText 	string
	Contents	FlexContainer
}

// MarshalJSON method of ButtonsTemplate
func (f *Flex) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type                FlexType        `json:"type"`
		AltText				string			`json:"altText"`
		Contents			FlexContainer	`json:"contents"`
	}{
		Type:               FlexTypeFlex,
		AltText:    		f.AltText,
		Contents:     		f.Contents,
	})
}

// FlexContainer interface
type FlexContainer interface {
	json.Marshaler
	FlexContainer()
}

type BubbleFlex struct {
	Direction 	string
	Header		FlexComponent
	Hero		FlexComponent
	Body		FlexComponent
	Footer		FlexComponent
	Styles		FlexStylesBlock
}

// MarshalJSON method of ButtonsTemplate
func (b *BubbleFlex) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ContainerType		`json:"type"`
		Direction           string       		`json:"direction, omitempty"`
		Header				FlexComponent		`json:"header, omitempty"`
		Hero				FlexComponent		`json:"hero, omitempty"`
		Body				FlexComponent		`json:"body, omitempty"`
		Footer				FlexComponent		`json:"footer, omitempty"`
		Styles				FlexStylesBlock		`json:"styles, omitempty"`
	}{
		Type:               ContainerTypeBubble,
	})
}



type CarouselFlex struct {
	Contents []BubbleFlex
}

// MarshalJSON method of ButtonsTemplate
func (c *CarouselFlex) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ContainerType		`json:"type"`
		Contents 			[]BubbleFlex		`json:"contents"`
	}{
		Type:               ContainerTypeCarousel,
		Contents:			c.Contents,
	})
}

// implements TemplateAction interface
func (*BubbleFlex) FlexContainer() {}
func (*CarouselFlex) FlexContainer() {}


// FlexContainer interface
type FlexComponent interface {
	json.Marshaler
	FlexComponent()
}

type BoxComponent struct {
	Layout		string
	Contents	[]FlexComponent
	Flex		int
	Spacing		string
	Margin		string
}

// MarshalJSON method of BoxComponent
func (b *BoxComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Layout				string				`json:"layout"`
		Contents 			[]FlexComponent		`json:"contents"`
		Flex				int					`json:"flex, omitempty"`
		Spacing				string				`json:"spacing, omitempty"`
		Margin				string				`json:"margin, omitempty"`
	}{
		Type:               ComponentTypeBox,
		Layout:				b.Layout,
		Contents:			b.Contents,
		Flex:				b.Flex,
		Spacing:			b.Spacing,
		Margin:				b.Margin,
	})
}

type ButtonComponent struct {
	Action		TemplateAction
	Flex		int
	Margin		string
	Height		string
	Style		string
	Color		string
	Gravity		string
}

// MarshalJSON method of ButtonComponent
func (b *ButtonComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Action				TemplateAction		`json:"action"`
		Flex				int					`json:"flex, omitempty"`
		Margin				string				`json:"margin, omitempty"`
		Height				string				`json:"height, omitempty"`
		Style				string				`json:"style, omitempty"`
		Color				string				`json:"color, omitempty"`
		Gravity				string				`json:"gravity, omitempty"`
	}{
		Type:               ComponentTypeBox,
		Action:				b.Action,
		Flex:				b.Flex,
		Margin:				b.Margin,
		Height:				b.Height,
		Style:				b.Style,
		Color:				b.Color,
		Gravity:			b.Gravity,
	})
}

type FillerComponent struct {

}

// MarshalJSON method of FillerComponent
func (f *FillerComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
	}{
		Type:               ComponentTypeFiller,
	})
}

type IconComponent struct {
	Url			string
	Margin		string
	Size		string
	AspectRatio	string
}

// MarshalJSON method of IconComponent
func (i *IconComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Url					string				`json:"url"`
		Margin				string				`json:"margin, omitempty"`
		Size				string				`json:"size, omitempty"`
		AspectRatio			string				`json:"aspectRatio, omitempty"`
	}{
		Type:               ComponentTypeIcon,
		Url:				i.Url,
		Margin:				i.Margin,
		Size:				i.Size,
		AspectRatio:		i.AspectRatio,
	})
}

type ImageComponent struct {
	Url					string
	Flex				int
	Margin				string
	Align				string
	Gravity				string
	Size				string
	AspectRatio			string
	AspectMode			string
	BackgroundColor		string
	Action				TemplateAction
}

// MarshalJSON method of ImageComponent
func (i *ImageComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Url					string				`json:"url"`
		Flex				int					`json:"flex, omitempty"`
		Margin				string				`json:"margin, omitempty"`
		Align				string				`json:"align, omitempty"`
		Gravity				string				`json:"gravity, omitempty"`
		Size				string				`json:"size, omitempty"`
		AspectRatio			string				`json:"aspectRatio, omitempty"`
		AspectMode			string				`json:"aspectMode, omitempty"`
		BackgroundColor		string				`json:"backgroundColor, omitempty"`
		Action				TemplateAction		`json:"action"`
	}{
		Type:               ComponentTypeImage,
		Url:				i.Url,
		Flex:				i.Flex,
		Margin:				i.Margin,
		Align:				i.Align,
		Gravity:			i.Gravity,
		Size:				i.Size,
		AspectRatio:		i.AspectRatio,
		AspectMode:			i.AspectMode,
		BackgroundColor:	i.BackgroundColor,
		Action:				i.Action,
	})
}

type SeparatorComponent struct {
	Margin				string
	Color				string
}

// MarshalJSON method of SeparatorComponent
func (s *SeparatorComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Margin				string				`json:"margin, omitempty"`
		Color				string				`json:"color, omitempty"`
	}{
		Type:               ComponentTypeSeparator,
		Margin:				s.Margin,
		Color:				s.Color,
	})
}

type SpacerComponent struct {
	Size				string
}

// MarshalJSON method of SpacerComponent
func (s *SpacerComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Size				string				`json:"size, omitempty"`
	}{
		Type:               ComponentTypeSpacer,
		Size:				s.Size,
	})
}

type TextComponent struct {
	Text				string
	Flex				int
	Margin				string
	Size				string
	Align				string
	Gravity				string
	Wrap				bool
	Weight				string
	Color				string
}

// MarshalJSON method of TextComponent
func (t *TextComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Text				string				`json:"text"`
		Flex				int					`json:"flex, omitempty"`
		Margin				string				`json:"margin, omitempty"`
		Size				string				`json:"size, omitempty"`
		Align				string				`json:"align, omitempty"`
		Gravity				string				`json:"gravity, omitempty"`
		Wrap				bool				`json:"wrap, omitempty"`
		Weight				string				`json:"weight, omitempty"`
		Color				string				`json:"color, omitempty"`
	}{
		Type:               ComponentTypeText,
		Text:				t.Text,
		Flex:				t.Flex,
		Margin:				t.Margin,
		Size:				t.Size,
		Align:				t.Align,
		Gravity:			t.Gravity,
		Wrap:				t.Wrap,
		Weight:				t.Weight,
		Color:				t.Color,
	})
}

// implements TemplateAction interface
func (*BoxComponent) FlexComponent() {}
func (*ButtonComponent) FlexComponent() {}
func (*FillerComponent) FlexComponent() {}
func (*IconComponent) FlexComponent() {}
func (*ImageComponent) FlexComponent() {}
func (*SeparatorComponent) FlexComponent() {}
func (*SpacerComponent) FlexComponent() {}
func (*TextComponent) FlexComponent() {}

// NewFlex function
func NewFlex(AltText string, Contents FlexContainer) *Flex {
	return &Flex{
		AltText: AltText,
		Contents: Contents,
	}
}

// NewCarouselFlex function
func NewCarouselFlex(bubbleFlexs ...BubbleFlex) *CarouselFlex {
	return &CarouselFlex{
		Contents: bubbleFlexs,
	}
}

// NewBubbleFlex function
func NewBubbleFlex(direction string, header FlexComponent, hero FlexComponent, body FlexComponent, footer FlexComponent, styles FlexStylesBlock) *BubbleFlex {
	return &BubbleFlex{
		Direction: 	direction,
		Header:		header,
		Hero:		hero,
		Body:		body,
		Footer:		footer,
		Styles:		styles,
	}
}

// NewFlexStylesBlock function
func NewFlexStylesBlock(header FlexStyle, hero FlexStyle, body FlexStyle, footer FlexStyle) *FlexStylesBlock {
	return &FlexStylesBlock{
		Header:	header,
		Hero:	hero,
		Body:	body,
		Footer:	footer,
	}
}

// NewFlexStyle function
func NewFlexStyle(backgroundColor string, separator bool, separatorColor string) *FlexStyle {
	return &FlexStyle{
		BackgroundColor: backgroundColor,
		Separator: separator,
		SeparatorColor: separatorColor,
	}

}
// NewBoxComponent function
func NewBoxComponent(layout string, flex int, spacing string, margin string, contents ...FlexComponent) *BoxComponent {
	return &BoxComponent{
		Layout:	layout,
		Contents: contents,
		Flex: flex,
		Spacing: spacing,
		Margin: margin,
	}
}

// NewButtonComponent
func NewButtonComponent(action TemplateAction, flex int, margin string, height string, style string, color string, gravity string) *ButtonComponent {
	return &ButtonComponent{
		Action: action,
		Flex: flex,
		Margin: margin,
		Height: height,
		Style: style,
		Color: color,
		Gravity: gravity,
	}
}

// NewFillerComponent
func NewFillerComponent() *FillerComponent {
	return &FillerComponent{}
}

// NewIconComponent
func NewIconComponent(url string, margin string, size string, aspectRatio string) *IconComponent {
	return &IconComponent{
		Url:url,
		Margin: margin,
		Size: size,
		AspectRatio: aspectRatio,
	}
}

//NewImageComponent
func NewImageComponent(url string, flex int, margin string, align string, gravity string, size string, aspectRatio string, aspectMode string, backgroundColor string, action TemplateAction) *ImageComponent {
	return &ImageComponent{
		Url: url,
		Flex: flex,
		Margin: margin,
		Align: align,
		Gravity: gravity,
		Size: size,
		AspectRatio: aspectRatio,
		AspectMode: aspectMode,
		BackgroundColor: backgroundColor,
		Action: action,
	}
}

// NewSeparatorComponent
func NewSeparatorComponent(margin string, color string) *SeparatorComponent {
	return &SeparatorComponent{
		Margin: margin,
		Color:color,
	}
}

// NewSpacerComponent
func NewSpacerComponent(size string) *SpacerComponent {
	return &SpacerComponent{
		Size: size,
	}
}

// NewTextComponent
func NewTextComponent(text string, flex int, margin string, size string, align string, gravity string, wrap bool, weight string, color string) *TextComponent {
	return &TextComponent{
		Text: text,
		Flex: flex,
		Margin: margin,
		Size: size,
		Align: align,
		Gravity: gravity,
		Wrap: wrap,
		Weight: weight,
		Color: color,
	}
}



