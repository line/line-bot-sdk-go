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

// BoxComponentBuilder type
type BoxComponentBuilder struct {
	component *linebot.BoxComponent
}

// NewBoxComponentBuilder function
func NewBoxComponentBuilder() *BoxComponentBuilder {
	return &BoxComponentBuilder{
		component: &linebot.BoxComponent{
			Type: linebot.FlexComponentTypeBox,
		},
	}
}

// Layout method
func (b *BoxComponentBuilder) Layout(layout linebot.FlexBoxLayoutType) *BoxComponentBuilder {
	b.component.Layout = layout
	return b
}

// Contents method
func (b *BoxComponentBuilder) Contents(contents []linebot.FlexComponent) *BoxComponentBuilder {
	b.component.Contents = contents
	return b
}

// Build method
func (b *BoxComponentBuilder) Build() *linebot.BoxComponent {
	return b.component
}

// TextComponentBuilder type
type TextComponentBuilder struct {
	component *linebot.TextComponent
}

// NewTextComponentBuilder function
func NewTextComponentBuilder() *TextComponentBuilder {
	return &TextComponentBuilder{
		component: &linebot.TextComponent{
			Type: linebot.FlexComponentTypeText,
		},
	}
}

// Text method
func (b *TextComponentBuilder) Text(text string) *TextComponentBuilder {
	b.component.Text = text
	return b
}

// Build method
func (b *TextComponentBuilder) Build() *linebot.TextComponent {
	return b.component
}
