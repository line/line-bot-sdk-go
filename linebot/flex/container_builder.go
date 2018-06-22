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

package flex

import "github.com/line/line-bot-sdk-go/linebot"

// BubbleContainerBuilder type
type BubbleContainerBuilder struct {
	container *linebot.BubbleContainer
}

// NewBubbleContainerBuilder function
func NewBubbleContainerBuilder() *BubbleContainerBuilder {
	return &BubbleContainerBuilder{
		container: &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
		},
	}
}

// Body method
func (b *BubbleContainerBuilder) Body(body *linebot.BoxComponent) *BubbleContainerBuilder {
	b.container.Body = body
	return b
}

// Build method
func (b *BubbleContainerBuilder) Build() *linebot.BubbleContainer {
	return b.container
}

// CarouselContainerBuilder type
type CarouselContainerBuilder struct {
	container *linebot.CarouselContainer
}

// NewCarouselContainerBuilder function
func NewCarouselContainerBuilder() *CarouselContainerBuilder {
	return &CarouselContainerBuilder{
		container: &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
		},
	}
}

// Contents method
func (b *CarouselContainerBuilder) Contents(contents []*linebot.BubbleContainer) *CarouselContainerBuilder {
	b.container.Contents = contents
	return b
}

// Build method
func (b *CarouselContainerBuilder) Build() *linebot.CarouselContainer {
	return b.container
}
