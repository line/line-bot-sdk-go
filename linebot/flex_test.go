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
	"reflect"
	"strconv"
	"testing"
)

func TestUnmarshalFlexMessageJSON(t *testing.T) {
	var testCases = []struct {
		JSON []byte
		Want FlexContainer
	}{
		{
			JSON: []byte(`{
  "type": "bubble",
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "hello"
      },
      {
        "type": "text",
        "text": "world"
      }
    ]
  }
}`),
			Want: &BubbleContainer{
				Type: FlexContainerTypeBubble,
				Body: &BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeVertical,
					Contents: []FlexComponent{
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "hello",
						},
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "world",
						},
					},
				},
			},
		},
		{
			JSON: []byte(`{
  "type": "carousel",
  "contents": [
    {
      "type": "bubble",
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": "First bubble"
          }
        ]
      }
    },
    {
      "type": "bubble",
      "body": {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": "Second bubble"
          }
        ],
        "action": {
          "type": "message",
          "text": "Second bubble"
        }
      }
    }
  ]
}`),
			Want: &CarouselContainer{
				Type: FlexContainerTypeCarousel,
				Contents: []*BubbleContainer{
					{
						Type: FlexContainerTypeBubble,
						Body: &BoxComponent{
							Type:   FlexComponentTypeBox,
							Layout: FlexBoxLayoutTypeVertical,
							Contents: []FlexComponent{
								&TextComponent{
									Type: FlexComponentTypeText,
									Text: "First bubble",
								},
							},
						},
					},
					{
						Type: FlexContainerTypeBubble,
						Body: &BoxComponent{
							Type:   FlexComponentTypeBox,
							Layout: FlexBoxLayoutTypeVertical,
							Contents: []FlexComponent{
								&TextComponent{
									Type: FlexComponentTypeText,
									Text: "Second bubble",
								},
							},
							Action: &MessageAction{Text: "Second bubble"},
						},
					},
				},
			},
		},
		{
			JSON: []byte(`{
  "type": "bubble",
  "size": "nano",
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "hello"
      },
      {
        "type": "text",
        "text": "world"
      }
    ]
  }
}`),
			Want: &BubbleContainer{
				Type: FlexContainerTypeBubble,
				Size: FlexBubbleSizeTypeNano,
				Body: &BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeVertical,
					Contents: []FlexComponent{
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "hello",
						},
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "world",
						},
					},
				},
			},
		},
		{
			JSON: []byte(`{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "https://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "Brown Cafe",
        "weight": "bold",
        "size": "xl"
      },
      {
        "type": "box",
        "layout": "baseline",
        "margin": "md",
        "contents": [
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
          },
          {
            "type": "text",
            "text": "4.0",
            "size": "sm",
            "color": "#999999",
            "margin": "md",
            "flex": 0
          }
        ]
      },
      {
        "type": "box",
        "layout": "vertical",
        "margin": "lg",
        "spacing": "sm",
        "contents": [
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Place",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1,
                "maxLines": 0
              },
              {
                "type": "text",
                "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5,
                "maxLines": 1
              }
            ]
          },
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Time",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "10:00 - 23:00",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          }
        ]
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "CALL",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "WEBSITE",
          "uri": "https://linecorp.com",
          "altUri": {
            "desktop": "https://line.me/ja/download"
          }
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      },
      {
        "type": "spacer"
      },
      {
        "type": "spacer",
        "size": "xs"
      }
    ],
    "flex": 0
  }
}`),
			Want: &BubbleContainer{
				Type: FlexContainerTypeBubble,
				Hero: &ImageComponent{
					Type:        FlexComponentTypeImage,
					URL:         "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_1_cafe.png",
					Size:        FlexImageSizeTypeFull,
					AspectRatio: FlexImageAspectRatioType20to13,
					AspectMode:  FlexImageAspectModeTypeCover,
					Action:      &URIAction{URI: "https://linecorp.com/"},
				},
				Body: &BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeVertical,
					Contents: []FlexComponent{
						&TextComponent{
							Type:   FlexComponentTypeText,
							Text:   "Brown Cafe",
							Size:   FlexTextSizeTypeXl,
							Weight: FlexTextWeightTypeBold,
						},
						&BoxComponent{
							Type:   FlexComponentTypeBox,
							Layout: FlexBoxLayoutTypeBaseline,
							Contents: []FlexComponent{
								&IconComponent{
									Type: FlexComponentTypeIcon,
									URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
									Size: FlexIconSizeTypeSm,
								},
								&IconComponent{
									Type: FlexComponentTypeIcon,
									URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
									Size: FlexIconSizeTypeSm,
								},
								&IconComponent{
									Type: FlexComponentTypeIcon,
									URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
									Size: FlexIconSizeTypeSm,
								},
								&IconComponent{
									Type: FlexComponentTypeIcon,
									URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gold_star_28.png",
									Size: FlexIconSizeTypeSm,
								},
								&IconComponent{
									Type: FlexComponentTypeIcon,
									URL:  "https://scdn.line-apps.com/n/channel_devcenter/img/fx/review_gray_star_28.png",
									Size: FlexIconSizeTypeSm,
								},
								&TextComponent{
									Type:   FlexComponentTypeText,
									Text:   "4.0",
									Flex:   IntPtr(0),
									Margin: FlexComponentMarginTypeMd,
									Size:   FlexTextSizeTypeSm,
									Color:  "#999999",
								},
							},
							Margin: FlexComponentMarginTypeMd,
						},
						&BoxComponent{
							Type:   FlexComponentTypeBox,
							Layout: FlexBoxLayoutTypeVertical,
							Contents: []FlexComponent{
								&BoxComponent{
									Type:   FlexComponentTypeBox,
									Layout: FlexBoxLayoutTypeBaseline,
									Contents: []FlexComponent{
										&TextComponent{
											Type:     FlexComponentTypeText,
											Text:     "Place",
											Flex:     IntPtr(1),
											Size:     FlexTextSizeTypeSm,
											Color:    "#aaaaaa",
											MaxLines: IntPtr(0),
										},
										&TextComponent{
											Type:     FlexComponentTypeText,
											Text:     "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
											Flex:     IntPtr(5),
											Size:     FlexTextSizeTypeSm,
											Wrap:     true,
											Color:    "#666666",
											MaxLines: IntPtr(1),
										},
									},
									Spacing: FlexComponentSpacingTypeSm,
								},
								&BoxComponent{
									Type:   FlexComponentTypeBox,
									Layout: FlexBoxLayoutTypeBaseline,
									Contents: []FlexComponent{
										&TextComponent{
											Type:  FlexComponentTypeText,
											Text:  "Time",
											Flex:  IntPtr(1),
											Size:  FlexTextSizeTypeSm,
											Color: "#aaaaaa",
										},
										&TextComponent{
											Type:  FlexComponentTypeText,
											Text:  "10:00 - 23:00",
											Flex:  IntPtr(5),
											Size:  FlexTextSizeTypeSm,
											Wrap:  true,
											Color: "#666666",
										},
									},
									Spacing: FlexComponentSpacingTypeSm,
								},
							},
							Spacing: FlexComponentSpacingTypeSm,
							Margin:  FlexComponentMarginTypeLg,
						},
					},
				},
				Footer: &BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeVertical,
					Contents: []FlexComponent{
						&ButtonComponent{
							Type: FlexComponentTypeButton,
							Action: &URIAction{
								Label: "CALL",
								URI:   "https://linecorp.com",
							},
							Height: FlexButtonHeightTypeSm,
							Style:  FlexButtonStyleTypeLink,
						},
						&ButtonComponent{
							Type: FlexComponentTypeButton,
							Action: &URIAction{
								Label: "WEBSITE",
								URI:   "https://linecorp.com",
								AltURI: &URIActionAltURI{
									Desktop: "https://line.me/ja/download",
								},
							},
							Height: FlexButtonHeightTypeSm,
							Style:  FlexButtonStyleTypeLink,
						},
						&SpacerComponent{
							Type: FlexComponentTypeSpacer,
							Size: FlexSpacerSizeTypeSm,
						},
						&SpacerComponent{
							Type: FlexComponentTypeSpacer,
						},
						&SpacerComponent{
							Type: FlexComponentTypeSpacer,
							Size: FlexSpacerSizeTypeXs,
						},
					},
					Spacing: FlexComponentSpacingTypeSm,
					Flex:    IntPtr(0),
				},
			},
		},
		{
			JSON: []byte(`{
  "type": "bubble",
  "body": {
    "type": "box",
    "layout": "horizontal",
    "contents": [
      {
        "type": "text",
        "text": "hello",
        "flex": 0
      },
      {
        "type": "filler",
        "flex": 4
      },
      {
        "type": "text",
        "text": "world",
        "flex": 2
	  },
	  {
        "type": "text",
        "contents": [
		  {
			"type": "span",
			"text": "hi"
		  },
		  {
			"type": "span",
			"text": "span",
			"size": "xl",
			"color": "#29cf5b"
		  }
		]
      }
    ]
  }
}`),
			Want: &BubbleContainer{
				Type: FlexContainerTypeBubble,
				Body: &BoxComponent{
					Type:   FlexComponentTypeBox,
					Layout: FlexBoxLayoutTypeHorizontal,
					Contents: []FlexComponent{
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "hello",
							Flex: IntPtr(0),
						},
						&FillerComponent{
							Type: FlexComponentTypeFiller,
							Flex: IntPtr(4),
						},
						&TextComponent{
							Type: FlexComponentTypeText,
							Text: "world",
							Flex: IntPtr(2),
						},
						&TextComponent{
							Type: FlexComponentTypeText,
							Contents: []*SpanComponent{
								{
									Type: FlexComponentTypeSpan,
									Text: "hi",
								},
								{
									Type:  FlexComponentTypeSpan,
									Text:  "span",
									Size:  FlexTextSizeTypeXl,
									Color: "#29cf5b",
								},
							},
						},
					},
				},
			},
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			container, err := UnmarshalFlexMessageJSON([]byte(tc.JSON))
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(container, tc.Want) {
				t.Errorf("Container %v, want %v", container, tc.Want)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	testCases := []struct {
		component FlexComponent
		want      []byte
	}{
		{
			&FillerComponent{
				Type: FlexComponentTypeFiller,
				Flex: nil,
			},
			[]byte(`{"type":"filler"}`),
		},
		{
			&FillerComponent{
				Type: FlexComponentTypeFiller,
				Flex: IntPtr(4),
			},
			[]byte(`{"type":"filler","flex":4}`),
		},
		{
			&SpanComponent{
				Type:       FlexComponentTypeSpan,
				Text:       "span",
				Size:       FlexTextSizeTypeMd,
				Weight:     FlexTextWeightTypeRegular,
				Color:      "#0000ff",
				Style:      FlexTextStyleTypeNormal,
				Decoration: FlexTextDecorationTypeNone,
			},
			[]byte(`{"type":"span","text":"span","size":"md","weight":"regular","color":"#0000ff","style":"normal","decoration":"none"}`),
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := json.Marshal(tc.component)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %s, want %s", string(got), string(tc.want))
			}
		})
	}
}

func BenchmarkUnmarshalFlexMessageJSON(b *testing.B) {
	var json = []byte(`{
		"type": "bubble",
		"header": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
				{
					"type": "text",
					"text": "NEWS DIGEST",
					"weight": "bold",
					"color": "#aaaaaa",
					"size": "sm"
				}
			]
		},
		"hero": {
			"type": "image",
			"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/01_4_news.png",
			"size": "full",
			"aspectRatio": "20:13",
			"aspectMode": "cover",
			"action": {
				"type": "uri",
				"uri": "https://linecorp.com/"
			}
		},
		"body": {
			"type": "box",
			"layout": "horizontal",
			"spacing": "md",
			"contents": [
				{
					"type": "box",
					"layout": "vertical",
					"flex": 1,
					"contents": [
						{
							"type": "image",
							"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/02_1_news_thumbnail_1.png",
							"aspectMode": "cover",
							"aspectRatio": "4:3",
							"size": "sm",
							"gravity": "bottom"
						},
						{
							"type": "image",
							"url": "https://scdn.line-apps.com/n/channel_devcenter/img/fx/02_1_news_thumbnail_2.png",
							"aspectMode": "cover",
							"aspectRatio": "4:3",
							"margin": "md",
							"size": "sm"
						}
					]
				},
				{
					"type": "box",
					"layout": "vertical",
					"flex": 2,
					"contents": [
						{
							"type": "text",
							"text": "7 Things to Know for Today",
							"gravity": "top",
							"size": "xs",
							"flex": 1
						},
						{
							"type": "separator"
						},
						{
							"type": "text",
							"text": "Hay fever goes wild",
							"gravity": "center",
							"size": "xs",
							"flex": 2
						},
						{
							"type": "separator"
						},
						{
							"type": "text",
							"text": "LINE Pay Begins Barcode Payment Service",
							"gravity": "center",
							"size": "xs",
							"flex": 2
						},
						{
							"type": "separator"
						},
						{
							"type": "text",
							"text": "LINE Adds LINE Wallet",
							"gravity": "bottom",
							"size": "xs",
							"flex": 1
						},
						{
							"type": "text",
							"contents": [
								{
									"type": "span",
									"text": "LINE",
									"size": "xxl",
									"weight": "bold",
									"style": "italic",
									"color": "#4f8f00"
								},
								{
									"type": "span",
									"text": "MUSIC",
									"size": "xxl",
									"weight": "bold",
									"style": "italic",
									"color": "#4f8f00"
								}
							]
						}
					]
				}
			]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
				{
					"type": "button",
					"action": {
						"type": "uri",
						"label": "More",
						"uri": "https://linecorp.com"
					}
				}
			]
		}
	}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := UnmarshalFlexMessageJSON(json)
		if err != nil {
			b.Fatal(err)
		}
	}
}
