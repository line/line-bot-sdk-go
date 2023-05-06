package linebot

import (
	"reflect"
	"strconv"
	"testing"
)

func TestUnmarshalTemplateJSON(t *testing.T) {
	testCases := []struct {
		JSON []byte
		Want Template
	}{
		{
			JSON: []byte(`{
	"type": "buttons",
	"thumbnailImageUrl": "https://example.com/image.jpg",
	"title": "Menu",
	"text": "Please select",
	"actions": [
		{
			"type": "postback",
			"label": "postback",
			"data": "action=buy&itemid=1",
			"displayText": "postback text"
		},
		{
			"type": "message",
			"label": "message",
			"text": "message text"
		},
		{
			"type": "uri",
			"label": "uri",
			"uri": "http://example.com/",
			"altUri": {
				"desktop": "http://example.com/desktop"
			}
		}
	]
}`),
			Want: &ButtonsTemplate{
				ThumbnailImageURL: "https://example.com/image.jpg",
				Title:             "Menu",
				Text:              "Please select",
				Actions: []TemplateAction{
					&PostbackAction{
						Label:       "postback",
						DisplayText: "postback text",
						Data:        "action=buy&itemid=1",
					},
					&MessageAction{
						Label: "message",
						Text:  "message text",
					},
					&URIAction{
						Label: "uri",
						URI:   "http://example.com/",
						AltURI: &URIActionAltURI{
							Desktop: "http://example.com/desktop",
						},
					},
				},
			},
		},
		{
			JSON: []byte(`{
	"type": "confirm",
	"text": "Are you sure?",
	"actions": [
		{
			"type": "message",
			"label": "Yes",
			"text": "yes"
		},
		{
			"type": "message",
			"label": "No",
			"text": "no"
		}
	]
}`),
			Want: &ConfirmTemplate{
				Text: "Are you sure?",
				Actions: []TemplateAction{
					&MessageAction{
						Label: "Yes",
						Text:  "yes",
					},
					&MessageAction{
						Label: "No",
						Text:  "no",
					},
				},
			},
		},
		{
			JSON: []byte(`{
    "type": "carousel",
    "columns": [
    	{
			"thumbnailImageUrl": "https://example.com/bot/images/item1.jpg",
			"imageBackgroundColor": "#FFFFFF",
			"title": "this is menu",
			"text": "description",
			"defaultAction": {
			"type": "uri",
			"label": "View detail",
			"uri": "http://example.com/page/123"
			},
			"actions": [
			{
				"type": "postback",
				"label": "Buy",
				"data": "action=buy&itemid=111"
			},
			{
				"type": "postback",
				"label": "Add to cart",
				"data": "action=add&itemid=111"
			},
			{
				"type": "uri",
				"label": "View detail",
				"uri": "http://example.com/page/111"
			}
			]
		},
		{
			"thumbnailImageUrl": "https://example.com/bot/images/item2.jpg",
			"imageBackgroundColor": "#000000",
			"title": "this is menu",
			"text": "description",
			"defaultAction": {
			"type": "uri",
			"label": "View detail",
			"uri": "http://example.com/page/222"
			},
			"actions": [
			{
				"type": "postback",
				"label": "Buy",
				"data": "action=buy&itemid=222"
			},
			{
				"type": "postback",
				"label": "Add to cart",
				"data": "action=add&itemid=222"
			},
			{
				"type": "uri",
				"label": "View detail",
				"uri": "http://example.com/page/222"
			}
			]
		}
		],
		"imageAspectRatio": "rectangle",
		"imageSize": "cover"
}`),
			Want: &CarouselTemplate{
				Columns: []*CarouselColumn{
					{
						ThumbnailImageURL:    "https://example.com/bot/images/item1.jpg",
						ImageBackgroundColor: "#FFFFFF",
						Title:                "this is menu",
						Text:                 "description",
						DefaultAction: &URIAction{
							Label: "View detail",
							URI:   "http://example.com/page/123",
						},
						Actions: []TemplateAction{
							&PostbackAction{
								Label: "Buy",
								Data:  "action=buy&itemid=111",
							},
							&PostbackAction{
								Label: "Add to cart",
								Data:  "action=add&itemid=111",
							},
							&URIAction{
								Label: "View detail",
								URI:   "http://example.com/page/111",
							},
						},
					},
					{
						ThumbnailImageURL:    "https://example.com/bot/images/item2.jpg",
						ImageBackgroundColor: "#000000",
						Title:                "this is menu",
						Text:                 "description",
						DefaultAction: &URIAction{
							Label: "View detail",
							URI:   "http://example.com/page/222",
						},
						Actions: []TemplateAction{
							&PostbackAction{
								Label: "Buy",
								Data:  "action=buy&itemid=222",
							},
							&PostbackAction{
								Label: "Add to cart",
								Data:  "action=add&itemid=222",
							},
							&URIAction{
								Label: "View detail",
								URI:   "http://example.com/page/222",
							},
						},
					},
				},
				ImageAspectRatio: "rectangle",
				ImageSize:        "cover",
			},
		},
		{
			JSON: []byte(`{
	"type": "image_carousel",
	"columns": [
	{
		"imageUrl": "https://example.com/bot/images/item1.jpg",
		"action": {
		"type": "postback",
		"label": "Buy",
		"data": "action=buy&itemid=111"
		}
	},
	{
		"imageUrl": "https://example.com/bot/images/item2.jpg",
		"action": {
		"type": "message",
		"label": "Yes",
		"text": "yes"
		}
	},
	{
		"imageUrl": "https://example.com/bot/images/item3.jpg",
		"action": {
		"type": "uri",
		"label": "View detail",
		"uri": "http://example.com/page/222"
		}
	}
	]
}`),
			Want: &ImageCarouselTemplate{
				Columns: []*ImageCarouselColumn{
					{
						ImageURL: "https://example.com/bot/images/item1.jpg",
						Action: &PostbackAction{
							Label: "Buy",
							Data:  "action=buy&itemid=111",
						},
					},
					{
						ImageURL: "https://example.com/bot/images/item2.jpg",
						Action: &MessageAction{
							Label: "Yes",
							Text:  "yes",
						},
					},
					{
						ImageURL: "https://example.com/bot/images/item3.jpg",
						Action: &URIAction{
							Label: "View detail",
							URI:   "http://example.com/page/222",
						},
					},
				},
			},
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			template, err := UnmarshalTemplateJSON(tc.JSON)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(template, tc.Want) {
				t.Errorf("got %v, want %v", template, tc.Want)
			}
		})
	}
}

func BenchmarkUnmarshalTemplateJSON(b *testing.B) {
	var jsonData = []byte(`{
    "type": "carousel",
    "columns": [
    	{
			"thumbnailImageUrl": "https://example.com/bot/images/item1.jpg",
			"imageBackgroundColor": "#FFFFFF",
			"title": "this is menu",
			"text": "description",
			"defaultAction": {
			"type": "uri",
			"label": "View detail",
			"uri": "http://example.com/page/123"
			},
			"actions": [
			{
				"type": "postback",
				"label": "Buy",
				"data": "action=buy&itemid=111"
			},
			{
				"type": "postback",
				"label": "Add to cart",
				"data": "action=add&itemid=111"
			},
			{
				"type": "uri",
				"label": "View detail",
				"uri": "http://example.com/page/111"
			}
			]
		},
		{
			"thumbnailImageUrl": "https://example.com/bot/images/item2.jpg",
			"imageBackgroundColor": "#000000",
			"title": "this is menu",
			"text": "description",
			"defaultAction": {
			"type": "uri",
			"label": "View detail",
			"uri": "http://example.com/page/222"
			},
			"actions": [
			{
				"type": "postback",
				"label": "Buy",
				"data": "action=buy&itemid=222"
			},
			{
				"type": "postback",
				"label": "Add to cart",
				"data": "action=add&itemid=222"
			},
			{
				"type": "uri",
				"label": "View detail",
				"uri": "http://example.com/page/222"
			}
			]
		}
		],
		"imageAspectRatio": "rectangle",
		"imageSize": "cover"
}`)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := UnmarshalTemplateJSON(jsonData)
		if err != nil {
			b.Fatal(err)
		}
	}
}
