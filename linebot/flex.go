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

// FlexContainerType type
type FlexContainerType string

// FlexContainerType constants
const (
	FlexContainerTypeBubble   FlexContainerType = "bubble"
	FlexContainerTypeCarousel FlexContainerType = "carousel"
)

// FlexComponentType type
type FlexComponentType string

// FlexComponentType constants
const (
	FlexComponentTypeBox       FlexComponentType = "box"
	FlexComponentTypeButton    FlexComponentType = "button"
	FlexComponentTypeFiller    FlexComponentType = "filler"
	FlexComponentTypeIcon      FlexComponentType = "icon"
	FlexComponentTypeImage     FlexComponentType = "image"
	FlexComponentTypeSeparator FlexComponentType = "separator"
	FlexComponentTypeSpacer    FlexComponentType = "spacer"
	FlexComponentTypeText      FlexComponentType = "text"
)

// FlexBubbleDirectionType type
type FlexBubbleDirectionType string

// FlexBubbleDirectionType constants
const (
	FlexBubbleDirectionTypeLTR FlexBubbleDirectionType = "ltr"
	FlexBubbleDirectionTypeRTL FlexBubbleDirectionType = "rtl"
)

// FlexButtonStyleType type
type FlexButtonStyleType string

// FlexButtonStyleType constants
const (
	FlexButtonStyleTypeLink      FlexButtonStyleType = "link"
	FlexButtonStyleTypePrimary   FlexButtonStyleType = "primary"
	FlexButtonStyleTypeSecondary FlexButtonStyleType = "secondary"
)

// FlexButtonHeightType type
type FlexButtonHeightType string

// FlexButtonHeightType constants
const (
	FlexButtonHeightTypeMd FlexButtonHeightType = "md"
	FlexButtonHeightTypeSm FlexButtonHeightType = "sm"
)

// FlexIconAspectRatioType type
type FlexIconAspectRatioType string

// FlexIconAspectRatioType constants
const (
	FlexIconAspectRatioType1to1 FlexIconAspectRatioType = "1:1"
	FlexIconAspectRatioType2to1 FlexIconAspectRatioType = "2:1"
	FlexIconAspectRatioType3to1 FlexIconAspectRatioType = "3:1"
)

// FlexImageSizeType type
type FlexImageSizeType string

// FlexImageSizeType constants
const (
	FlexImageSizeTypeXxs FlexImageSizeType = "xxs"
	FlexImageSizeTypeXs  FlexImageSizeType = "xs"
	FlexImageSizeTypeSm  FlexImageSizeType = "sm"
	FlexImageSizeTypeMd  FlexImageSizeType = "md"
	FlexImageSizeTypeLg  FlexImageSizeType = "lg"
	FlexImageSizeTypeXl  FlexImageSizeType = "xl"
	FlexImageSizeTypeXxl FlexImageSizeType = "xxl"
	FlexImageSizeType3xl FlexImageSizeType = "3xl"
	FlexImageSizeType4xl FlexImageSizeType = "4xl"
	FlexImageSizeType5xl FlexImageSizeType = "5xl"
)

// FlexImageAspectRatioType type
type FlexImageAspectRatioType string

// FlexImageAspectRatioType constants
const (
	FlexImageAspectRatioType1to1    FlexImageAspectRatioType = "1:1"
	FlexImageAspectRatioType1_51to1 FlexImageAspectRatioType = "1.51:1"
	FlexImageAspectRatioType1_91to1 FlexImageAspectRatioType = "1.91:1"
	FlexImageAspectRatioType4to3    FlexImageAspectRatioType = "4:3"
	FlexImageAspectRatioType16to9   FlexImageAspectRatioType = "16:9"
	FlexImageAspectRatioType20to13  FlexImageAspectRatioType = "20:13"
	FlexImageAspectRatioType2to1    FlexImageAspectRatioType = "2:1"
	FlexImageAspectRatioType3to1    FlexImageAspectRatioType = "3:1"
	FlexImageAspectRatioType3to4    FlexImageAspectRatioType = "3:4"
	FlexImageAspectRatioType9to16   FlexImageAspectRatioType = "9:16"
	FlexImageAspectRatioType1to2    FlexImageAspectRatioType = "1:2"
	FlexImageAspectRatioType1to3    FlexImageAspectRatioType = "1:3"
)

// FlexImageAspectModeType type
type FlexImageAspectModeType string

// FlexImageAspectModeType constants
const (
	FlexImageAspectModeTypeCover FlexImageAspectModeType = "cover"
	FlexImageAspectModeTypeFit   FlexImageAspectModeType = "fit"
)

// FlexBoxLayoutType type
type FlexBoxLayoutType string

// FlexBoxLayoutType constants
const (
	FlexBoxLayoutTypeHorizontal FlexBoxLayoutType = "horizontal"
	FlexBoxLayoutTypeVertical   FlexBoxLayoutType = "vertical"
	FlexBoxLayoutTypeBaseline   FlexBoxLayoutType = "baseline"
)

// FlexComponentSpacingType type
type FlexComponentSpacingType string

// FlexComponentSpacingType constants
const (
	FlexComponentSpacingTypeNone FlexComponentSpacingType = "none"
	FlexComponentSpacingTypeXs   FlexComponentSpacingType = "xs"
	FlexComponentSpacingTypeSm   FlexComponentSpacingType = "sm"
	FlexComponentSpacingTypeMd   FlexComponentSpacingType = "md"
	FlexComponentSpacingTypeLg   FlexComponentSpacingType = "lg"
	FlexComponentSpacingTypeXl   FlexComponentSpacingType = "xl"
	FlexComponentSpacingTypeXxl  FlexComponentSpacingType = "xxl"
)

// FlexComponentMarginType type
type FlexComponentMarginType string

// FlexComponentMarginType constants
const (
	FlexComponentMarginTypeNone FlexComponentMarginType = "none"
	FlexComponentMarginTypeXs   FlexComponentMarginType = "xs"
	FlexComponentMarginTypeSm   FlexComponentMarginType = "sm"
	FlexComponentMarginTypeMd   FlexComponentMarginType = "md"
	FlexComponentMarginTypeLg   FlexComponentMarginType = "lg"
	FlexComponentMarginTypeXl   FlexComponentMarginType = "xl"
	FlexComponentMarginTypeXxl  FlexComponentMarginType = "xxl"
)

// FlexComponentGravityType type
type FlexComponentGravityType string

// FlexComponentGravityType constants
const (
	FlexComponentGravityTypeTop    FlexComponentGravityType = "top"
	FlexComponentGravityTypeBottom FlexComponentGravityType = "bottom"
	FlexComponentGravityTypeCenter FlexComponentGravityType = "center"
)

// FlexComponentAlignType type
type FlexComponentAlignType string

// FlexComponentAlignType constants
const (
	FlexComponentAlignTypeStart  FlexComponentAlignType = "start"
	FlexComponentAlignTypeEnd    FlexComponentAlignType = "end"
	FlexComponentAlignTypeCenter FlexComponentAlignType = "center"
)

// FlexIconSizeType type
type FlexIconSizeType string

// FlexIconSizeType constants
const (
	FlexIconSizeTypeXxs FlexIconSizeType = "xxs"
	FlexIconSizeTypeXs  FlexIconSizeType = "xs"
	FlexIconSizeTypeSm  FlexIconSizeType = "sm"
	FlexIconSizeTypeMd  FlexIconSizeType = "md"
	FlexIconSizeTypeLg  FlexIconSizeType = "lg"
	FlexIconSizeTypeXl  FlexIconSizeType = "xl"
	FlexIconSizeTypeXxl FlexIconSizeType = "xxl"
	FlexIconSizeType3xl FlexIconSizeType = "3xl"
	FlexIconSizeType4xl FlexIconSizeType = "4xl"
	FlexIconSizeType5xl FlexIconSizeType = "5xl"
)

// FlexSpacerSizeType type
type FlexSpacerSizeType string

// FlexSpacerSizeType constants
const (
	FlexSpacerSizeTypeXs  FlexSpacerSizeType = "xs"
	FlexSpacerSizeTypeSm  FlexSpacerSizeType = "sm"
	FlexSpacerSizeTypeMd  FlexSpacerSizeType = "md"
	FlexSpacerSizeTypeLg  FlexSpacerSizeType = "lg"
	FlexSpacerSizeTypeXl  FlexSpacerSizeType = "xl"
	FlexSpacerSizeTypeXxl FlexSpacerSizeType = "xxl"
)

// FlexTextWeightType type
type FlexTextWeightType string

// FlexTextWeightType constants
const (
	FlexTextWeightTypeRegular FlexTextWeightType = "regular"
	FlexTextWeightTypeBold    FlexTextWeightType = "bold"
)

// FlexTextSizeType type
type FlexTextSizeType string

// FlexTextSizeType constants
const (
	FlexTextSizeTypeXxs FlexTextSizeType = "xxs"
	FlexTextSizeTypeXs  FlexTextSizeType = "xs"
	FlexTextSizeTypeSm  FlexTextSizeType = "sm"
	FlexTextSizeTypeMd  FlexTextSizeType = "md"
	FlexTextSizeTypeLg  FlexTextSizeType = "lg"
	FlexTextSizeTypeXl  FlexTextSizeType = "xl"
	FlexTextSizeTypeXxl FlexTextSizeType = "xxl"
	FlexTextSizeType3xl FlexTextSizeType = "3xl"
	FlexTextSizeType4xl FlexTextSizeType = "4xl"
	FlexTextSizeType5xl FlexTextSizeType = "5xl"
)

// FlexContainer interface
type FlexContainer interface {
	FlexContainer()
}

// BubbleContainer type
type BubbleContainer struct {
	Type      FlexContainerType       `json:"type"`
	Direction FlexBubbleDirectionType `json:"direction,omitempty"`
	Header    *BoxComponent           `json:"header,omitempty"`
	Hero      *ImageComponent         `json:"hero,omitempty"`
	Body      *BoxComponent           `json:"body,omitempty"`
	Footer    *BoxComponent           `json:"footer,omitempty"`
	Styles    *BubbleStyle            `json:"styles,omitempty"`
}

// FlexContainer method
func (*BubbleContainer) FlexContainer() {}

// CarouselContainer type
type CarouselContainer struct {
	Type     FlexContainerType  `json:"type"`
	Contents []*BubbleContainer `json:"contents"`
}

// FlexContainer method
func (*CarouselContainer) FlexContainer() {}

// BubbleStyle type
type BubbleStyle struct {
	Header BlockStyle
	Hero   BlockStyle
	Body   BlockStyle
	Footer BlockStyle
}

// BlockStyle type
type BlockStyle struct {
	BackgroundColor string
	Separator       bool
	SeparatorColor  string
}

// FlexComponent interface
type FlexComponent interface {
}

// BoxComponent type
type BoxComponent struct {
	Type     FlexComponentType        `json:"type"`
	Layout   FlexBoxLayoutType        `json:"layout"`
	Contents []FlexComponent          `json:"contents"`
	Flex     *int                     `json:"flex,omitempty"`
	Spacing  FlexComponentSpacingType `json:"spacing,omitempty"`
	Margin   FlexComponentMarginType  `json:"margin,omitempty"`
}

// ButtonComponent type
type ButtonComponent struct {
	Action  TemplateAction
	Flex    int
	Margin  FlexComponentMarginType
	Height  FlexButtonHeightType
	Style   FlexButtonStyleType
	Color   string
	Gravity FlexComponentGravityType
}

// FillerComponent type
type FillerComponent struct {
}

// IconComponent type
type IconComponent struct {
	URL         string
	Margin      FlexComponentMarginType
	Size        FlexIconSizeType
	AspectRatio FlexIconAspectRatioType
}

// ImageComponent type
type ImageComponent struct {
	URL             string
	Flex            int
	Margin          FlexComponentMarginType
	Align           FlexComponentAlignType
	Gravity         FlexComponentGravityType
	Size            FlexImageSizeType
	AspectRatio     FlexImageAspectRatioType
	AspectMode      FlexImageAspectModeType
	BackgroundColor string
	Action          TemplateAction
}

// SeparatorComponent type
type SeparatorComponent struct {
	Margin FlexComponentMarginType
	Color  string
}

// SpacerComponent type
type SpacerComponent struct {
	Size FlexSpacerSizeType
}

// TextComponent type
type TextComponent struct {
	Type    FlexComponentType         `json:"type"`
	Text    string                    `json:"text"`
	Flex    *int                      `json:"flex,omitempty"`
	Margin  *FlexComponentMarginType  `json:"margin,omitempty"`
	Size    *FlexTextSizeType         `json:"size,omitempty"`
	Align   *FlexComponentAlignType   `json:"align,omitempty"`
	Gravity *FlexComponentGravityType `json:"gravity,omitempty"`
	Wrap    *bool                     `json:"wrap,omitempty"`
	Weight  *FlexTextWeightType       `json:"weight,omitempty"`
	Color   *string                   `json:"color,omitempty"`
	Action  *TemplateAction           `json:"action,omitempty"`
}
