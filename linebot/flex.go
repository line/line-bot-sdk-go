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
	ButtonStyleTypeLink			ButtonStyleType = "link"
	ButtonStyleTypePrimary		ButtonStyleType = "primary"
	ButtonStyleTypeSecondary	ButtonStyleType = "secondary"
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
	BackgroundColor 	string		`json:"backgroundColor,omitempty"`
	Separator 			bool		`json:"separator,omitempty"`
	SeparatorColor 		string		`json:"separatorColor,omitempty"`
}

type FlexStylesBlock struct {
	Header	*FlexStyle	`json:"header,omitempty"`
	Hero	*FlexStyle	`json:"hero,omitempty"`
	Body	*FlexStyle	`json:"body,omitempty"`
	Footer	*FlexStyle	`json:"footer,omitempty"`
}

type Flex struct {
	AltText 	string			`json:"altText"`
	Contents	FlexContainer	`json:"contents"`
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
	Header		FlexComponent		`json:"header,omitempty"`
	Hero		FlexComponent		`json:"hero,omitempty"`
	Body		FlexComponent		`json:"body,omitempty"`
	Footer		FlexComponent		`json:"footer,omitempty"`
	Styles		*FlexStylesBlock	`json:"styles,omitempty"`
}

// MarshalJSON method of ButtonsTemplate
func (b *BubbleFlex) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ContainerType			`json:"type"`
		Header				FlexComponent			`json:"header,omitempty"`
		Hero				FlexComponent			`json:"hero,omitempty"`
		Body				FlexComponent			`json:"body,omitempty"`
		Footer				FlexComponent			`json:"footer,omitempty"`
		Styles				*FlexStylesBlock		`json:"styles,omitempty"`
	}{
		Type:               ContainerTypeBubble,
		Header:				b.Header,
		Hero:				b.Hero,
		Body:				b.Body,
		Footer:				b.Footer,
		Styles:				b.Styles,
	})
}

type CarouselFlex struct {
	Contents []*BubbleFlex	`json:"contents"`
}

// MarshalJSON method of ButtonsTemplate
func (c *CarouselFlex) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ContainerType		`json:"type"`
		Contents 			[]*BubbleFlex		`json:"contents"`
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
	Layout		string				`json:"layout"`
	Contents	[]FlexComponent		`json:"contents"`
	Flex		*int				`json:"flex,omitempty"`
	Spacing		*SizeType			`json:"spacing,omitempty"`
	Margin		*SizeType			`json:"margin,omitempty"`
}

// MarshalJSON method of BoxComponent
func (b *BoxComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Layout				string				`json:"layout"`
		Contents 			[]FlexComponent		`json:"contents"`
		Flex				*int				`json:"flex,omitempty"`
		Spacing				*SizeType			`json:"spacing,omitempty"`
		Margin				*SizeType			`json:"margin,omitempty"`
	}{
		Type:               ComponentTypeBox,
		Layout:				b.Layout,
		Contents:			b.Contents,
		Flex:				b.Flex,
		Spacing:			b.Spacing,
		Margin:				b.Margin,
	})
}

func (b *BoxComponent) WithBoxComponentSizeOptions(flex *int, spacing *SizeType, margin *SizeType) *BoxComponent {
	b.Flex = flex
	b.Spacing = spacing
	b.Margin = margin
	return b
}

type ButtonComponent struct {
	Action		TemplateAction		`json:"action"`
	Flex		*int				`json:"flex,omitempty"`
	Margin		*SizeType			`json:"margin,omitempty"`
	Height		*SizeHeightType		`json:"height,omitempty"`
	Style		*ButtonStyleType	`json:"style,omitempty"`
	Color		*string				`json:"color,omitempty"`
	Gravity		*GravityType		`json:"gravity,omitempty"`
}

// MarshalJSON method of ButtonComponent
func (b *ButtonComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Action				TemplateAction		`json:"action"`
		Flex				*int				`json:"flex,omitempty"`
		Margin				*SizeType			`json:"margin,omitempty"`
		Height				*SizeHeightType		`json:"height,omitempty"`
		Style				*ButtonStyleType	`json:"style,omitempty"`
		Color				*string				`json:"color,omitempty"`
		Gravity				*GravityType		`json:"gravity,omitempty"`
	}{
		Type:               ComponentTypeButton,
		Action:				b.Action,
		Flex:				b.Flex,
		Margin:				b.Margin,
		Height:				b.Height,
		Style:				b.Style,
		Color:				b.Color,
		Gravity:			b.Gravity,
	})
}

func (b *ButtonComponent) WithButtonComponentStyleOptions(flex *int, margin *SizeType, height *SizeHeightType, style *ButtonStyleType, color *string, gravity *GravityType) *ButtonComponent {
	b.Flex = flex
	b.Margin = margin
	b.Height = height
	b.Style = style
	b.Color = color
	b.Gravity = gravity
	return b
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
	Url			string					`json:"url"`
	Margin		*SizeType				`json:"margin,omitempty"`
	Size		*SizeType				`json:"size,omitempty"`
	AspectRatio	*AspectRatioType		`json:"aspectRatio,omitempty"`
}

// MarshalJSON method of IconComponent
func (i *IconComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Url					string				`json:"url"`
		Margin				*SizeType			`json:"margin,omitempty"`
		Size				*SizeType			`json:"size,omitempty"`
		AspectRatio			*AspectRatioType	`json:"aspectRatio,omitempty"`
	}{
		Type:               ComponentTypeIcon,
		Url:				i.Url,
		Margin:				i.Margin,
		Size:				i.Size,
		AspectRatio:		i.AspectRatio,
	})
}

func (i *IconComponent) WithIconComponentStyleOptions(margin *SizeType, size *SizeType, aspectRatio *AspectRatioType) *IconComponent {
	i.Margin = margin
	i.Size = size
	i.AspectRatio = aspectRatio
	return i
}

type ImageComponent struct {
	Url					string				`json:"url"`
	Flex				*int				`json:"flex,omitempty"`
	Align				*AlignType			`json:"align,omitempty"`
	Size				*SizeType			`json:"size,omitempty"`
	AspectRatio			*AspectRatioType	`json:"aspectRatio,omitempty"`
	AspectMode			*AspectModeType		`json:"aspectMode,omitempty"`
	BackgroundColor		*string				`json:"backgroundColor,omitempty"`
	Margin				*SizeType			`json:"margin,omitempty"`
	Gravity				*GravityType		`json:"gravity,omitempty"`
	Action				TemplateAction		`json:"action"`
}

// MarshalJSON method of ImageComponent
func (i *ImageComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Url					string				`json:"url"`
		Flex				*int				`json:"flex,omitempty"`
		Align				*AlignType			`json:"align,omitempty"`
		Size				*SizeType			`json:"size,omitempty"`
		AspectRatio			*AspectRatioType	`json:"aspectRatio,omitempty"`
		AspectMode			*AspectModeType		`json:"aspectMode,omitempty"`
		BackgroundColor		*string				`json:"backgroundColor,omitempty"`
		Margin				*SizeType			`json:"margin,omitempty"`
		Gravity				*GravityType		`json:"gravity,omitempty"`
		Action				TemplateAction		`json:"action"`
	}{
		Type:               ComponentTypeImage,
		Url:				i.Url,
		Flex:				i.Flex,
		Align:				i.Align,
		Size:				i.Size,
		AspectRatio:		i.AspectRatio,
		AspectMode:			i.AspectMode,
		BackgroundColor:	i.BackgroundColor,
		Margin:				i.Margin,
		Gravity:			i.Gravity,
		Action:				i.Action,
	})
}

func (i *ImageComponent) WithImageComponentStyleOption(flex *int, align *AlignType, size *SizeType, aspectRatio *AspectRatioType, aspectMode *AspectModeType, backgroundColor *string, margin *SizeType, gravity *GravityType) *ImageComponent {
	i.Flex = flex
	i.Align = align
	i.Size = size
	i.AspectRatio = aspectRatio
	i.AspectMode = aspectMode
	i.BackgroundColor = backgroundColor
	i.Margin = margin
	i.Gravity = gravity
	return i
}

func (i *ImageComponent) WithImageComponentActionOption(action TemplateAction) *ImageComponent {
	i.Action = action
	return i
}

type SeparatorComponent struct {
	Margin				*SizeType		`json:"margin,omitempty"`
	Color				*string			`json:"color,omitempty"`
}

// MarshalJSON method of SeparatorComponent
func (s *SeparatorComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Margin				*SizeType			`json:"margin,omitempty"`
		Color				*string				`json:"color,omitempty"`
	}{
		Type:               ComponentTypeSeparator,
		Margin:				s.Margin,
		Color:				s.Color,
	})
}

func (s *SeparatorComponent) WithSeparatorComponentStyleOption(margin *SizeType, color *string) *SeparatorComponent {
	s.Margin = margin
	s.Color = color
	return s
}

type SpacerComponent struct {
	Size				string		`json:"size"`
}

// MarshalJSON method of SpacerComponent
func (s *SpacerComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType	`json:"type"`
		Size				string			`json:"size"`
	}{
		Type:               ComponentTypeSpacer,
		Size:				s.Size,
	})
}

type TextComponent struct {
	Text				string				`json:"text"`
	Flex				*int				`json:"flex,omitempty"`
	Margin				*SizeType			`json:"margin,omitempty"`
	Size				*SizeType			`json:"size,omitempty"`
	Align				*AlignType			`json:"align,omitempty"`
	Weight				*SizeWeightType		`json:"weight,omitempty"`
	Color				*string				`json:"color,omitempty"`
	Wrap				bool				`json:"wrap,omitempty"`
	Gravity				*GravityType		`json:"gravity,omitempty"`
}

// MarshalJSON method of TextComponent
func (t *TextComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type				ComponentType		`json:"type"`
		Text				string				`json:"text"`
		Flex				*int				`json:"flex,omitempty"`
		Margin				*SizeType			`json:"margin,omitempty"`
		Size				*SizeType			`json:"size,omitempty"`
		Align				*AlignType			`json:"align,omitempty"`
		Weight				*SizeWeightType		`json:"weight,omitempty"`
		Color				*string				`json:"color,omitempty"`
		Wrap				bool				`json:"wrap,omitempty"`
		Gravity				*GravityType		`json:"gravity,omitempty"`
	}{
		Type:               ComponentTypeText,
		Text:				t.Text,
		Flex:				t.Flex,
		Margin:				t.Margin,
		Size:				t.Size,
		Weight:				t.Weight,
		Color:				t.Color,
		Wrap:				t.Wrap,
		Align:				t.Align,
		Gravity:			t.Gravity,
	})
}

func (t *TextComponent) WithTextComponentStyleOption(flex *int, margin *SizeType, size *SizeType, weight *SizeWeightType, color *string, wrap bool, align *AlignType, gravity *GravityType) *TextComponent {
	t.Flex = flex
	t.Margin = margin
	t.Size = size
	t.Weight = weight
	t.Color = color
	t.Wrap = wrap
	t.Align = align
	t.Gravity = gravity
	return t
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
func NewFlex() *Flex {
	return &Flex{}
}

// NewCarouselFlex function
func NewCarouselFlex() *CarouselFlex {
	return &CarouselFlex{}
}

// NewBubbleFlex function
func NewBubbleFlex() *BubbleFlex {
	return &BubbleFlex{}
}

// NewFlexStylesBlock function
func NewFlexStylesBlock() *FlexStylesBlock {
	return &FlexStylesBlock{}
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
func NewBoxComponent(layout string, contents ...FlexComponent) *BoxComponent {
	return &BoxComponent{
		Layout:	layout,
		Contents: contents,
	}
}

// NewButtonComponent
func NewButtonComponent(action TemplateAction) *ButtonComponent {
	return &ButtonComponent{
		Action: action,
	}
}

// NewFillerComponent
func NewFillerComponent() *FillerComponent {
	return &FillerComponent{}
}

// NewIconComponent
func NewIconComponent(url string) *IconComponent {
	return &IconComponent{
		Url:url,
	}
}

//NewImageComponent
func NewImageComponent(url string) *ImageComponent {
	return &ImageComponent{
		Url: url,
	}
}

// NewSeparatorComponent
func NewSeparatorComponent() *SeparatorComponent {
	return &SeparatorComponent{}
}

// NewSpacerComponent
func NewSpacerComponent(size string) *SpacerComponent {
	return &SpacerComponent{
		Size: size,
	}
}

// NewTextComponent
func NewTextComponent(text string) *TextComponent {
	return &TextComponent{
		Text: text,
	}
}



