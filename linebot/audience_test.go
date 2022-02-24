package linebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
	"time"
)

// TestUploadAudienceGroup tests UploadAudienceGroup
func TestUploadAudienceGroup(t *testing.T) {
	type RequestBody struct {
		Description       string     `json:"description,omitempty"`
		IsIfaAudience     bool       `json:"isIfaAudience,omitempty"`
		UploadDescription string     `json:"uploadDescription,omitempty"`
		Audiences         []audience `json:"audiences,omitempty"`
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *UploadAudienceGroupResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		Description       string
		IsIfaAudience     bool
		UploadDescription string
		Audiences         []string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:        "Create Audience Fail",
			Description:  "audienceGroupName_01",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:        "Create Audience not include userIDs",
			Description:  "audienceGroupName_01",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:       "Create Audience By UserID",
			Description: "audienceGroupName_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					Audiences: []audience{
						{
							ID: "U4af4980627",
						},
						{
							ID: "U4af4980628",
						},
						{
							ID: "U4af4980629",
						},
					},
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:         "Create Audience Is IFA",
			Description:   "audienceGroupName_01",
			IsIfaAudience: true,
			RequestID:     "12222",
			ResponseCode:  http.StatusOK,
			Response:      []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description:   "audienceGroupName_01",
					IsIfaAudience: true,
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:             "Create Audience have uploadDescription",
			Description:       "audienceGroupName_01",
			UploadDescription: "audienceGroupNameJob_01",
			RequestID:         "12222",
			ResponseCode:      http.StatusOK,
			Response:          []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description:       "audienceGroupName_01",
					UploadDescription: "audienceGroupNameJob_01",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:             "Create Audience",
			Description:       "audienceGroupName_01",
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					Description:       "audienceGroupName_01",
					UploadDescription: "audienceGroupNameJob_01",
					Audiences: []audience{
						{
							ID: "U4af4980627",
						},
						{
							ID: "U4af4980628",
						},
						{
							ID: "U4af4980629",
						},
					},
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *UploadAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IUploadAudienceGroupOption
			if len(tc.Audiences) > 0 {
				options = append(options, WithUploadAudienceGroupCallAudiences(tc.Audiences...))
			}
			if tc.IsIfaAudience {
				options = append(options, WithUploadAudienceGroupCallIsIfaAudience(true))
			}
			if tc.UploadDescription != "" {
				options = append(options, WithUploadAudienceGroupCallUploadDescription(tc.UploadDescription))
			}
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.UploadAudienceGroup(tc.Description, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestUploadAudienceGroupByFile tests UploadAudienceGroupByFile
func TestUploadAudienceGroupByFile(t *testing.T) {
	type RequestBody struct {
		Description       string
		IsIfaAudience     bool
		UploadDescription string
		Audiences         string
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *UploadAudienceGroupResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		Description       string
		IsIfaAudience     bool
		UploadDescription string
		Audiences         []string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:        "Create Audience By File Fail",
			Description:  "audienceGroupName_01",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:        "Create Audience By File not include userIDs",
			Description:  "audienceGroupName_01",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:       "Create Audience By File By UserID",
			Description: "audienceGroupName_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					Audiences:   "U4af4980627\nU4af4980628\nU4af4980629",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:         "Create Audience By File Is IFA",
			Description:   "audienceGroupName_01",
			IsIfaAudience: true,
			RequestID:     "12222",
			ResponseCode:  http.StatusOK,
			Response:      []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description:   "audienceGroupName_01",
					IsIfaAudience: true,
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:             "Create Audience By File have uploadDescription",
			Description:       "audienceGroupName_01",
			UploadDescription: "audienceGroupNameJob_01",
			RequestID:         "12222",
			ResponseCode:      http.StatusOK,
			Response:          []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description:       "audienceGroupName_01",
					UploadDescription: "audienceGroupNameJob_01",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
		{
			Label:             "Create Audience By File",
			Description:       "audienceGroupName_01",
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"UPLOAD","description":"audienceGroupName_01","created":1613698278,"permission":"READ_WRITE","expireTimestamp":1629250278}`),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					Description:       "audienceGroupName_01",
					UploadDescription: "audienceGroupNameJob_01",
					Audiences:         "U4af4980627\nU4af4980628\nU4af4980629",
				},
				Response: &UploadAudienceGroupResponse{
					RequestID:       "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "UPLOAD",
					Description:     "audienceGroupName_01",
					Created:         1613698278,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629250278,
					IsIfaAudience:   false,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			t.Fatal(err)
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		buf := bytes.NewBuffer(nil)
		_, errCopy := io.Copy(buf, file)
		if err != nil {
			t.Fatal(errCopy)
		}
		if buf.String() != tc.Want.RequestBody.Audiences {
			t.Errorf("Audiences %s; want %s", buf.String(), tc.Want.RequestBody.Audiences)
		}
		if v := r.FormValue("description"); v != "" {
			if v != tc.Want.RequestBody.Description {
				t.Errorf("Description %s; want %s", v, tc.Want.RequestBody.Description)
			}
		}
		if v := r.FormValue("uploadDescription"); v != "" {
			if v != tc.Want.RequestBody.UploadDescription {
				t.Errorf("UploadDescription %s; want %s", v, tc.Want.RequestBody.UploadDescription)
			}
		}
		if v := r.FormValue("isIfaAudience"); v != "" {
			isIfaAudience, err := strconv.ParseBool(v)
			if err != nil {
				t.Fatal(errCopy)
			}
			if isIfaAudience != tc.Want.RequestBody.IsIfaAudience {
				t.Errorf("IsIfaAudience %s; want %v", v, tc.Want.RequestBody.IsIfaAudience)
			}
		}

		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *UploadAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IUploadAudienceGroupByFileOption
			if tc.IsIfaAudience {
				options = append(options, WithUploadAudienceGroupByFileCallIsIfaAudience(true))
			}
			if tc.UploadDescription != "" {
				options = append(options, WithUploadAudienceGroupByFileCallUploadDescription(tc.UploadDescription))
			}
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.UploadAudienceGroupByFile(tc.Description, tc.Audiences, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestAddAudiences tests AddAudiences
func TestAddAudiences(t *testing.T) {
	type RequestBody struct {
		AudienceGroupID   int        `json:"audienceGroupId,omitempty"`
		UploadDescription string     `json:"uploadDescription,omitempty"`
		Audiences         []audience `json:"audiences,omitempty"`
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		UploadDescription string
		Audiences         []string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:             "Add Audience Fail",
			AudienceGroupID:   4389303728991,
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					AudienceGroupID:   4389303728991,
					UploadDescription: "audienceGroupNameJob_01",
					Audiences: []audience{
						{
							ID: "U4af4980627",
						},
						{
							ID: "U4af4980628",
						},
						{
							ID: "U4af4980629",
						},
					},
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:           "add Audience no uploadDescription",
			AudienceGroupID: 4389303728991,
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					AudienceGroupID: 4389303728991,
					Audiences: []audience{
						{
							ID: "U4af4980627",
						},
						{
							ID: "U4af4980628",
						},
						{
							ID: "U4af4980629",
						},
					},
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
		{
			Label:             "add Audience",
			AudienceGroupID:   4389303728991,
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUpload,
				RequestBody: RequestBody{
					AudienceGroupID:   4389303728991,
					UploadDescription: "audienceGroupNameJob_01",
					Audiences: []audience{
						{
							ID: "U4af4980627",
						},
						{
							ID: "U4af4980628",
						},
						{
							ID: "U4af4980629",
						},
					},
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IAddAudiencesOption
			if tc.UploadDescription != "" {
				options = append(options, WithAddAudiencesCallUploadDescription(tc.UploadDescription))
			}
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.AddAudiences(tc.AudienceGroupID, tc.Audiences, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				log.Println(err)
				log.Println(tc.Want.Error)
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestAddAudiencesByFile tests AddAudiencesByFile
func TestAddAudiencesByFile(t *testing.T) {
	type RequestBody struct {
		AudienceGroupID   int
		UploadDescription string
		Audiences         string
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		UploadDescription string
		Audiences         []string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:             "Add Audience By File Fail",
			AudienceGroupID:   4389303728991,
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					AudienceGroupID:   4389303728991,
					UploadDescription: "audienceGroupNameJob_01",
					Audiences:         "U4af4980627\nU4af4980628\nU4af4980629",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:           "Add Audience By File no uploadDescription",
			AudienceGroupID: 4389303728991,
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					AudienceGroupID: 4389303728991,
					Audiences:       "U4af4980627\nU4af4980628\nU4af4980629",
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
		{
			Label:             "Add Audience By File",
			AudienceGroupID:   4389303728991,
			UploadDescription: "audienceGroupNameJob_01",
			Audiences: []string{
				"U4af4980627", "U4af4980628", "U4af4980629",
			},
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupUploadByFile,
				RequestBody: RequestBody{
					AudienceGroupID:   4389303728991,
					UploadDescription: "audienceGroupNameJob_01",
					Audiences:         "U4af4980627\nU4af4980628\nU4af4980629",
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			t.Fatal(err)
		}
		file, _, err := r.FormFile("file")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		buf := bytes.NewBuffer(nil)
		_, errCopy := io.Copy(buf, file)
		if err != nil {
			t.Fatal(errCopy)
		}
		if buf.String() != tc.Want.RequestBody.Audiences {
			t.Errorf("Audiences %s; want %s", buf.String(), tc.Want.RequestBody.Audiences)
		}
		if v := r.FormValue("audienceGroupId"); v != "" {
			audienceGroupID, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			if int(audienceGroupID) != tc.Want.RequestBody.AudienceGroupID {
				t.Errorf("audienceGroupId %v; want %v", int(audienceGroupID), tc.Want.RequestBody.AudienceGroupID)
			}
		}
		if v := r.FormValue("uploadDescription"); v != "" {
			if v != tc.Want.RequestBody.UploadDescription {
				t.Errorf("UploadDescription %s; want %s", v, tc.Want.RequestBody.UploadDescription)
			}
		}

		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IAddAudiencesByFileOption
			if tc.UploadDescription != "" {
				options = append(options, WithAddAudiencesByFileCallUploadDescription(tc.UploadDescription))
			}
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.AddAudiencesByFile(tc.AudienceGroupID, tc.Audiences, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestClickAudienceGroup tests ClickAudienceGroup
func TestClickAudienceGroup(t *testing.T) {
	type RequestBody struct {
		Description string `json:"description,omitempty"`
		RequestID   string `json:"requestId,omitempty"`
		ClickURL    string `json:"clickUrl,omitempty"`
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *ClickAudienceGroupResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		Description       string
		XRequestID        string
		ClickURL          string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:        "Click Audience Fail",
			Description:  "audienceGroupName_01",
			XRequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
			ClickURL:     "https://developers.line.biz/",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupClick,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					RequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
					ClickURL:    "https://developers.line.biz/",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:        "Click Audience No ClickURL",
			Description:  "audienceGroupName_01",
			XRequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"CLICK","description":"audienceGroupName_01","created":1613705240,"permission":"READ_WRITE","expireTimestamp":1629257239,"requestId":"bb9744f9-47fa-4a29-941e-1234567890ab","clickUrl":"https://developers.line.biz/"}`),
			Want: want{
				URLPath: APIAudienceGroupClick,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					RequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
				},
				Response: &ClickAudienceGroupResponse{
					XRequestID:      "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "CLICK",
					Description:     "audienceGroupName_01",
					Created:         1613705240,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629257239,
					IsIfaAudience:   false,
					RequestID:       "bb9744f9-47fa-4a29-941e-1234567890ab",
					ClickURL:        "https://developers.line.biz/",
				},
			},
		},
		{
			Label:        "Click Audience",
			Description:  "audienceGroupName_01",
			XRequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
			ClickURL:     "https://developers.line.biz/",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"CLICK","description":"audienceGroupName_01","created":1613705240,"permission":"READ_WRITE","expireTimestamp":1629257239,"requestId":"bb9744f9-47fa-4a29-941e-1234567890ab","clickUrl":"https://developers.line.biz/"}`),
			Want: want{
				URLPath: APIAudienceGroupClick,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					RequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
					ClickURL:    "https://developers.line.biz/",
				},
				Response: &ClickAudienceGroupResponse{
					XRequestID:      "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "CLICK",
					Description:     "audienceGroupName_01",
					Created:         1613705240,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629257239,
					IsIfaAudience:   false,
					RequestID:       "bb9744f9-47fa-4a29-941e-1234567890ab",
					ClickURL:        "https://developers.line.biz/",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *ClickAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IClickAudienceGroupOption
			if tc.ClickURL != "" {
				options = append(options, WithClickAudienceGroupCallClickURL(tc.ClickURL))
			}
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.ClickAudienceGroup(tc.Description, tc.XRequestID, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				log.Println(err)
				log.Println(tc.Want.Error)
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestIMPAudienceGroup tests IMPAudienceGroup
func TestIMPAudienceGroup(t *testing.T) {
	type RequestBody struct {
		Description string `json:"description,omitempty"`
		RequestID   string `json:"requestId,omitempty"`
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *IMPAudienceGroupResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		Description       string
		XRequestID        string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:        "IMP Audience Fail",
			Description:  "audienceGroupName_01",
			XRequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath: APIAudienceGroupIMP,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					RequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:        "IMP Audience",
			Description:  "audienceGroupName_01",
			XRequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroupId":1234567890123,"createRoute":"MESSAGING_API","type":"IMP","description":"audienceGroupName_01","created":1613707097,"permission":"READ_WRITE","expireTimestamp":1629259095,"requestId":"bb9744f9-47fa-4a29-941e-1234567890ab"}`),
			Want: want{
				URLPath: APIAudienceGroupIMP,
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
					RequestID:   "bb9744f9-47fa-4a29-941e-1234567890ab",
				},
				Response: &IMPAudienceGroupResponse{
					XRequestID:      "12222",
					AudienceGroupID: 1234567890123,
					CreateRoute:     "MESSAGING_API",
					Type:            "IMP",
					Description:     "audienceGroupName_01",
					Created:         1613707097,
					Permission:      "READ_WRITE",
					ExpireTimestamp: 1629259095,
					IsIfaAudience:   false,
					RequestID:       "bb9744f9-47fa-4a29-941e-1234567890ab",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPost {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPost)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *IMPAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.IMPAudienceGroup(tc.Description, tc.XRequestID).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				log.Println(err)
				log.Println(tc.Want.Error)
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestUpdateAudienceGroupDescription tests UpdateAudienceGroupDescription
func TestUpdateAudienceGroupDescription(t *testing.T) {
	type RequestBody struct {
		Description string `json:"description,omitempty"`
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		Description       string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:           "Update Audience Description Fail",
			AudienceGroupID: 1234567890123,
			Description:     "audienceGroupName_01",
			RequestID:       "12222",
			ResponseCode:    http.StatusBadRequest,
			Response:        []byte(``),
			Want: want{
				URLPath: fmt.Sprintf(APIAudienceGroupUpdateDescription, 1234567890123),
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:           "Update Audience Description",
			AudienceGroupID: 1234567890123,
			Description:     "audienceGroupName_01",
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(``),
			Want: want{
				URLPath: fmt.Sprintf(APIAudienceGroupUpdateDescription, 1234567890123),
				RequestBody: RequestBody{
					Description: "audienceGroupName_01",
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.UpdateAudienceGroupDescription(tc.AudienceGroupID, tc.Description).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestActivateAudienceGroup tests ActivateAudienceGroup
func TestActivateAudienceGroup(t *testing.T) {
	type RequestBody struct {
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:           "Activate Audience Fail",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusBadRequest,
			Response:        []byte(``),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroupActivate, 1234567890123),
				RequestBody: RequestBody{},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:           "Activate Audience",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(``),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroupActivate, 1234567890123),
				RequestBody: RequestBody{},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.ActivateAudienceGroup(tc.AudienceGroupID).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestDeleteAudienceGroup tests DeleteAudienceGroup
func TestDeleteAudienceGroup(t *testing.T) {
	type RequestBody struct {
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:           "Delete Audience Fail",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusBadRequest,
			Response:        []byte(``),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 1234567890123),
				RequestBody: RequestBody{},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:           "Delete Audience",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(``),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 1234567890123),
				RequestBody: RequestBody{},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodDelete {
			t.Errorf("Method %s; want %s", r.Method, http.MethodDelete)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.DeleteAudienceGroup(tc.AudienceGroupID).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestGetAudienceGroup tests GetAudienceGroup
func TestGetAudienceGroup(t *testing.T) {
	type RequestBody struct {
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *GetAudienceGroupResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AudienceGroupID   int
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:           "Get Audience Fail",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusBadRequest,
			Response:        []byte(`{"message":"audience group not found","details":[{"message":"AUDIENCE_GROUP_NOT_FOUND","property":""}],"error":"","error_description":""}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 1234567890123),
				RequestBody: RequestBody{},
				Error: &APIError{
					Code: http.StatusBadRequest,
					Response: &ErrorResponse{
						Message: "audience group not found",
						Details: []errorResponseDetail{
							{
								Message: "AUDIENCE_GROUP_NOT_FOUND",
							},
						},
					},
				},
			},
		},
		{
			Label:           "Get UPLOAD Audience",
			AudienceGroupID: 1234567890123,
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(`{"audienceGroup":{"audienceGroupId":1234567890123,"createRoute":"OA_MANAGER","type":"UPLOAD","description":"audienceGroupName_01","status":"READY","audienceCount":1887,"created":1608617466,"permission":"READ","expireTimestamp":1624342266},"jobs":[{"audienceGroupJobId":12345678,"audienceGroupId":1234567890123,"description":"audience_list.txt","type":"DIFF_ADD","status":"FINISHED","created":1608617472,"jobStatus":"FINISHED"}]}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 1234567890123),
				RequestBody: RequestBody{},
				Response: &GetAudienceGroupResponse{
					RequestID: "12222",
					AudienceGroup: AudienceGroup{
						AudienceGroupID: 1234567890123,
						CreateRoute:     "OA_MANAGER",
						Type:            "UPLOAD",
						Description:     "audienceGroupName_01",
						Status:          "READY",
						AudienceCount:   1887,
						Created:         1608617466,
						Permission:      "READ",
						IsIfaAudience:   false,
						ExpireTimestamp: 1624342266,
					},
					Jobs: []Job{
						{
							AudienceGroupJobID: 12345678,
							AudienceGroupID:    1234567890123,
							Description:        "audience_list.txt",
							Type:               "DIFF_ADD",
							Status:             "FINISHED",
							AudienceCount:      0,
							Created:            1608617472,
							JobStatus:          "FINISHED",
						},
					},
				},
			},
		},
		{
			Label:           "Get CLICK Audience",
			AudienceGroupID: 1234567890987,
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(`{"audienceGroup":{"audienceGroupId":1234567890987,"createRoute":"OA_MANAGER","type":"CLICK","description":"audienceGroupName_02","status":"IN_PROGRESS","audienceCount":8619,"created":1611114828,"permission":"READ","requestId":"c10c3d86-f565-...","clickUrl":"https://example.com/","expireTimestamp":1624342266}}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 1234567890987),
				RequestBody: RequestBody{},
				Response: &GetAudienceGroupResponse{
					RequestID: "12222",
					AudienceGroup: AudienceGroup{
						AudienceGroupID: 1234567890987,
						CreateRoute:     "OA_MANAGER",
						Type:            "CLICK",
						Description:     "audienceGroupName_02",
						Status:          "IN_PROGRESS",
						AudienceCount:   8619,
						Created:         1611114828,
						Permission:      "READ",
						IsIfaAudience:   false,
						ExpireTimestamp: 1624342266,
						RequestID:       "c10c3d86-f565-...",
						ClickURL:        "https://example.com/",
					},
				},
			},
		},
		{
			Label:           "Get APP_EVENT Audience",
			AudienceGroupID: 2345678909876,
			RequestID:       "12222",
			ResponseCode:    http.StatusOK,
			Response:        []byte(`{"audienceGroup":{"audienceGroupId":2345678909876,"createRoute":"AD_MANAGER","type":"APP_EVENT","description":"audienceGroupName_03","status":"READY","audienceCount":8619,"created":1608619802,"permission":"READ","activated":1610068515,"inactivatedTimestamp":1625620516}}`),
			Want: want{
				URLPath:     fmt.Sprintf(APIAudienceGroup, 2345678909876),
				RequestBody: RequestBody{},
				Response: &GetAudienceGroupResponse{
					RequestID: "12222",
					AudienceGroup: AudienceGroup{
						AudienceGroupID:      2345678909876,
						CreateRoute:          "AD_MANAGER",
						Type:                 "APP_EVENT",
						Description:          "audienceGroupName_03",
						Status:               "READY",
						AudienceCount:        8619,
						Created:              1608619802,
						Permission:           "READ",
						Activated:            1610068515,
						InactivatedTimestamp: 1625620516,
						IsIfaAudience:        false,
					},
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *GetAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.GetAudienceGroup(tc.AudienceGroupID).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestListAudienceGroup tests ListAudienceGroup
func TestListAudienceGroup(t *testing.T) {
	type RequestBody struct {
	}
	type RequestParameter struct {
		Page                         int
		Description                  string
		Status                       AudienceStatusType
		Size                         int
		IncludesExternalPublicGroups bool
		CreateRoute                  string
	}
	type want struct {
		URLPath          string
		RequestBody      RequestBody
		RequestParameter RequestParameter
		Response         *ListAudienceGroupResponse
		Error            error
	}
	testCases := []struct {
		Label                        string
		Page                         int
		Description                  string
		Status                       AudienceStatusType
		Size                         int
		IncludesExternalPublicGroups bool
		CreateRoute                  string
		RequestID                    string
		AcceptedRequestID            string
		ResponseCode                 int
		Response                     []byte
		Want                         want
	}{
		{
			Label:        "List Audience Fail",
			Page:         1,
			Description:  "audienceGroupName",
			Status:       INPROGRESS,
			Size:         41,
			CreateRoute:  "OA_MANAGER",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(`{"message":"size: must be less than or equal to 40","details":[{"message":"TOO_HIGH","property":""}],"error":"","error_description":""}`),
			Want: want{
				URLPath:     APIAudienceGroupList,
				RequestBody: RequestBody{},
				RequestParameter: RequestParameter{
					Page:        1,
					Description: "audienceGroupName",
					Status:      INPROGRESS,
					Size:        41,
					CreateRoute: "OA_MANAGER",
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
					Response: &ErrorResponse{
						Message: "size: must be less than or equal to 40",
						Details: []errorResponseDetail{
							{
								Message: "TOO_HIGH",
							},
						},
					},
				},
			},
		},
		{
			Label:        "List Audience",
			Page:         1,
			Description:  "audienceGroupName",
			Size:         40,
			CreateRoute:  "OA_MANAGER",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"audienceGroups":[{"audienceGroupId":1234567890123,"createRoute":"OA_MANAGER","type":"CLICK","description":"audienceGroupName_01","status":"IN_PROGRESS","audienceCount":8619,"created":1611114828,"permission":"READ","requestId":"c10c3d86-f565-...","clickUrl":"https://example.com/","expireTimestamp":1626753228},{"audienceGroupId":2345678901234,"createRoute":"AD_MANAGER","type":"APP_EVENT","description":"audienceGroupName_02","status":"READY","audienceCount":3368,"created":1608619802,"permission":"READ","activated":1610068515,"inactivatedTimestamp":1625620516}],"totalCount":2,"page":40,"size":1}`),
			Want: want{
				URLPath:     APIAudienceGroupList,
				RequestBody: RequestBody{},
				RequestParameter: RequestParameter{
					Page:        1,
					Description: "audienceGroupName",
					Size:        40,
					CreateRoute: "OA_MANAGER",
				},
				Response: &ListAudienceGroupResponse{
					RequestID: "12222",
					AudienceGroups: []AudienceGroup{
						{
							AudienceGroupID: 1234567890123,
							CreateRoute:     "OA_MANAGER",
							Type:            "CLICK",
							Description:     "audienceGroupName_01",
							Status:          "IN_PROGRESS",
							AudienceCount:   8619,
							Created:         1611114828,
							Permission:      "READ",
							IsIfaAudience:   false,
							ExpireTimestamp: 1626753228,
							RequestID:       "c10c3d86-f565-...",
							ClickURL:        "https://example.com/",
						},
						{
							AudienceGroupID:      2345678901234,
							CreateRoute:          "AD_MANAGER",
							Type:                 "APP_EVENT",
							Description:          "audienceGroupName_02",
							Status:               "READY",
							AudienceCount:        3368,
							Created:              1608619802,
							Permission:           "READ",
							IsIfaAudience:        false,
							Activated:            1610068515,
							InactivatedTimestamp: 1625620516,
						},
					},
					HasNextPage:                      false,
					TotalCount:                       2,
					ReadWriteAudienceGroupTotalCount: 0,
					Page:                             40,
					Size:                             1,
				},
			},
		},
		{
			Label:                        "List Audience no data",
			Page:                         1,
			Description:                  "audienceGroupName",
			IncludesExternalPublicGroups: false,
			Size:                         40,
			CreateRoute:                  "OA_MANAGER",
			RequestID:                    "12222",
			ResponseCode:                 http.StatusOK,
			Response:                     []byte(`{"page":40,"size":1}`),
			Want: want{
				URLPath:     APIAudienceGroupList,
				RequestBody: RequestBody{},
				RequestParameter: RequestParameter{
					Page:                         1,
					Description:                  "audienceGroupName",
					IncludesExternalPublicGroups: false,
					Size:                         40,
					CreateRoute:                  "OA_MANAGER",
				},
				Response: &ListAudienceGroupResponse{
					RequestID:                        "12222",
					HasNextPage:                      false,
					TotalCount:                       0,
					ReadWriteAudienceGroupTotalCount: 0,
					Page:                             40,
					Size:                             1,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}

		if v := r.URL.Query().Get("page"); v != "" {
			page, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			if int(page) != tc.Want.RequestParameter.Page {
				t.Errorf("Request Page %v; want %v", page, tc.Want.RequestParameter.Page)
			}
		}
		if v := r.URL.Query().Get("size"); v != "" {
			size, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				t.Fatal(err)
			}
			if int(size) != tc.Want.RequestParameter.Size {
				t.Errorf("Request Size %v; want %v", size, tc.Want.RequestParameter.Size)
			}
		}
		if v := r.URL.Query().Get("description"); v != "" {
			if v != tc.Want.RequestParameter.Description {
				t.Errorf("Request Description %v; want %v", v, tc.Want.RequestParameter.Description)
			}
		}
		if v := r.URL.Query().Get("status"); v != "" {
			if v != tc.Want.RequestParameter.Status.String() {
				t.Errorf("Request Status %v; want %v", v, tc.Want.RequestParameter.Status)
			}
		}
		if v := r.URL.Query().Get("includesExternalPublicGroups"); v != "" {
			includesExternalPublicGroups, err := strconv.ParseBool(v)
			if err != nil {
				t.Fatal(err)
			}
			if includesExternalPublicGroups != tc.Want.RequestParameter.IncludesExternalPublicGroups {
				t.Errorf("Request IncludesExternalPublicGroups %v; want %v", v, tc.Want.RequestParameter.IncludesExternalPublicGroups)
			}
		}
		if v := r.URL.Query().Get("createRoute"); v != "" {
			if v != tc.Want.RequestParameter.CreateRoute {
				t.Errorf("Request CreateRoute %v; want %v", v, tc.Want.RequestParameter.CreateRoute)
			}
		}

		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *ListAudienceGroupResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			var options []IListAudienceGroupOption
			if tc.Description != "" {
				options = append(options, WithListAudienceGroupCallDescription(tc.Description))
			}
			if tc.Status != "" {
				options = append(options, WithListAudienceGroupCallStatus(tc.Status))
			}
			if !tc.IncludesExternalPublicGroups {
				options = append(options, WithListAudienceGroupCallIncludesExternalPublicGroups(tc.IncludesExternalPublicGroups))
			}
			if tc.Size > 0 {
				options = append(options, WithListAudienceGroupCallSize(tc.Size))
			}
			if tc.CreateRoute != "" {
				options = append(options, WithListAudienceGroupCallCreateRoute(tc.CreateRoute))
			}

			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.ListAudienceGroup(tc.Page, options...).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestGetAudienceGroupAuthorityLevel tests GetAudienceGroupAuthorityLevel
func TestGetAudienceGroupAuthorityLevel(t *testing.T) {
	type RequestBody struct {
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *GetAudienceGroupAuthorityLevelResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:        "Get Audience AuthorityLevel Fail",
			RequestID:    "12222",
			ResponseCode: http.StatusBadRequest,
			Response:     []byte(``),
			Want: want{
				URLPath:     APIAudienceGroupAuthorityLevel,
				RequestBody: RequestBody{},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:        "Get Audience AuthorityLevel",
			RequestID:    "12222",
			ResponseCode: http.StatusOK,
			Response:     []byte(`{"authorityLevel":"PUBLIC"}`),
			Want: want{
				URLPath:     APIAudienceGroupAuthorityLevel,
				RequestBody: RequestBody{},
				Response: &GetAudienceGroupAuthorityLevelResponse{
					RequestID:      "12222",
					AuthorityLevel: PUBLIC,
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodGet {
			t.Errorf("Method %s; want %s", r.Method, http.MethodGet)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}
		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *GetAudienceGroupAuthorityLevelResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.GetAudienceGroupAuthorityLevel().WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				log.Println(err)
				log.Println(tc.Want.Error)
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}

// TestChangeAudienceGroupAuthorityLevel tests ChangeAudienceGroupAuthorityLevel
func TestChangeAudienceGroupAuthorityLevel(t *testing.T) {
	type RequestBody struct {
		AuthorityLevel AudienceAuthorityLevelType
	}
	type want struct {
		URLPath     string
		RequestBody RequestBody
		Response    *BasicResponse
		Error       error
	}
	testCases := []struct {
		Label             string
		AuthorityLevel    AudienceAuthorityLevelType
		RequestID         string
		AcceptedRequestID string
		ResponseCode      int
		Response          []byte
		Want              want
	}{
		{
			Label:          "Change Audience AuthorityLevel Fail",
			AuthorityLevel: PUBLIC,
			RequestID:      "12222",
			ResponseCode:   http.StatusBadRequest,
			Response:       []byte(``),
			Want: want{
				URLPath: APIAudienceGroupAuthorityLevel,
				RequestBody: RequestBody{
					AuthorityLevel: PUBLIC,
				},
				Error: &APIError{
					Code: http.StatusBadRequest,
				},
			},
		},
		{
			Label:          "Change Audience AuthorityLevel ",
			AuthorityLevel: PUBLIC,
			RequestID:      "12222",
			ResponseCode:   http.StatusOK,
			Response:       []byte(``),
			Want: want{
				URLPath: APIAudienceGroupAuthorityLevel,
				RequestBody: RequestBody{
					AuthorityLevel: PUBLIC,
				},
				Response: &BasicResponse{
					RequestID: "12222",
				},
			},
		},
	}

	var currentTestIdx int
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		tc := testCases[currentTestIdx]
		if r.Method != http.MethodPut {
			t.Errorf("Method %s; want %s", r.Method, http.MethodPut)
		}
		if r.URL.Path != tc.Want.URLPath {
			t.Errorf("URLPath %s; want %s", r.URL.Path, tc.Want.URLPath)
		}

		var result RequestBody
		if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(result, tc.Want.RequestBody) {
			t.Errorf("Request %v; want %v", result, tc.Want.RequestBody)
		}

		w.Header().Set("X-Line-Request-Id", tc.RequestID)
		if tc.AcceptedRequestID != "" {
			w.Header().Set("X-Line-Accepted-Request-Id", tc.AcceptedRequestID)
		}
		w.WriteHeader(tc.ResponseCode)
		w.Write(tc.Response)
	}))
	defer server.Close()

	dataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		t.Error("Unexpected data API call")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"Not found"}`))
	}))
	defer dataServer.Close()

	client, err := mockClient(server, dataServer)
	if err != nil {
		t.Fatal(err)
	}

	var res *BasicResponse
	for i, tc := range testCases {
		currentTestIdx = i
		t.Run(strconv.Itoa(i)+"/"+tc.Label, func(t *testing.T) {
			timeoutCtx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancelFn()
			res, err = client.ChangeAudienceGroupAuthorityLevel(tc.AuthorityLevel).WithContext(timeoutCtx).Do()
			if tc.Want.Error != nil {
				log.Println(err)
				log.Println(tc.Want.Error)
				if !reflect.DeepEqual(err, tc.Want.Error) {
					t.Errorf("Error %v; want %v", err, tc.Want.Error)
				}
			} else {
				if err != nil {
					t.Error(err)
				}
			}
			if tc.Want.Response != nil {
				if !reflect.DeepEqual(res, tc.Want.Response) {
					t.Errorf("Response %v; want %v", res, tc.Want.Response)
				}
			}
		})
	}
}
