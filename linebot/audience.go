package linebot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// AudienceStatusType type
type AudienceStatusType string

// String method
func (a AudienceStatusType) String() string {
	return string(a)
}

const (
	// INPROGRESS const
	INPROGRESS AudienceStatusType = "IN_PROGRESS"
	// READY const
	READY AudienceStatusType = "READY"
	// FAILED const
	FAILED AudienceStatusType = "FAILED"
	// EXPIRED const
	EXPIRED AudienceStatusType = "EXPIRED"
	// INACTIVE const
	INACTIVE AudienceStatusType = "INACTIVE"
	// ACTIVATING const
	ACTIVATING AudienceStatusType = "ACTIVATING"
)

// AudienceAuthorityLevelType type
type AudienceAuthorityLevelType string

// String method
func (a AudienceAuthorityLevelType) String() string {
	return string(a)
}

const (
	// PUBLIC const
	PUBLIC AudienceAuthorityLevelType = "PUBLIC"
	// PRIVATE const
	PRIVATE AudienceAuthorityLevelType = "PRIVATE"
)

// IUploadAudienceGroupOption type
type IUploadAudienceGroupOption interface {
	Apply(call *UploadAudienceGroupCall)
}

// WithUploadAudienceGroupCallIsIfaAudience func
// Deprecated: Use OpenAPI based classes instead.
func WithUploadAudienceGroupCallIsIfaAudience(isIfaAudience bool) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallIsIfaAudience{
		isIfaAudience: isIfaAudience,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withUploadAudienceGroupCallIsIfaAudience struct {
	isIfaAudience bool
}

func (w *withUploadAudienceGroupCallIsIfaAudience) Apply(call *UploadAudienceGroupCall) {
	call.IsIfaAudience = w.isIfaAudience
}

// WithUploadAudienceGroupCallUploadDescription func
// Deprecated: Use OpenAPI based classes instead.
func WithUploadAudienceGroupCallUploadDescription(uploadDescription string) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withUploadAudienceGroupCallUploadDescription struct {
	uploadDescription string
}

func (w *withUploadAudienceGroupCallUploadDescription) Apply(call *UploadAudienceGroupCall) {
	call.UploadDescription = w.uploadDescription
}

// WithUploadAudienceGroupCallAudiences func
// Deprecated: Use OpenAPI based classes instead.
func WithUploadAudienceGroupCallAudiences(audiences ...string) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallAudiences{
		audiences: audiences,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withUploadAudienceGroupCallAudiences struct {
	audiences []string
}

func (w *withUploadAudienceGroupCallAudiences) Apply(call *UploadAudienceGroupCall) {
	for _, item := range w.audiences {
		call.Audiences = append(call.Audiences, audience{ID: item})
	}
}

// UploadAudienceGroup method
func (client *Client) UploadAudienceGroup(description string, options ...IUploadAudienceGroupOption) *UploadAudienceGroupCall {
	call := &UploadAudienceGroupCall{
		c:           client,
		Description: description,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// Deprecated: Use OpenAPI based classes instead.
type audience struct {
	ID string `json:"id,omitempty"`
}

// UploadAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type UploadAudienceGroupCall struct {
	c                 *Client
	ctx               context.Context
	Description       string `validate:"required,max=120"`
	IsIfaAudience     bool
	UploadDescription string
	Audiences         []audience `validate:"max=10000"`
}

// WithContext method
func (call *UploadAudienceGroupCall) WithContext(ctx context.Context) *UploadAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *UploadAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		Description       string     `json:"description,omitempty"`
		IsIfaAudience     bool       `json:"isIfaAudience,omitempty"`
		UploadDescription string     `json:"uploadDescription,omitempty"`
		Audiences         []audience `json:"audiences,omitempty"`
	}{
		Description:       call.Description,
		IsIfaAudience:     call.IsIfaAudience,
		UploadDescription: call.UploadDescription,
		Audiences:         call.Audiences,
	})
}

// Do method
func (call *UploadAudienceGroupCall) Do() (*UploadAudienceGroupResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIAudienceGroupUpload, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToAudienceGroupResponse(res)
}

// IUploadAudienceGroupByFileOption type
type IUploadAudienceGroupByFileOption interface {
	Apply(call *UploadAudienceGroupByFileCall)
}

// WithUploadAudienceGroupByFileCallIsIfaAudience func
// Deprecated: Use OpenAPI based classes instead.
func WithUploadAudienceGroupByFileCallIsIfaAudience(isIfaAudience bool) IUploadAudienceGroupByFileOption {
	return &withUploadAudienceGroupByFileCallIsIfaAudience{
		isIfaAudience: isIfaAudience,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withUploadAudienceGroupByFileCallIsIfaAudience struct {
	isIfaAudience bool
}

func (w *withUploadAudienceGroupByFileCallIsIfaAudience) Apply(call *UploadAudienceGroupByFileCall) {
	call.IsIfaAudience = w.isIfaAudience
}

// WithUploadAudienceGroupByFileCallUploadDescription func
// Deprecated: Use OpenAPI based classes instead.
func WithUploadAudienceGroupByFileCallUploadDescription(uploadDescription string) IUploadAudienceGroupByFileOption {
	return &withUploadAudienceGroupByFileCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withUploadAudienceGroupByFileCallUploadDescription struct {
	uploadDescription string
}

func (w *withUploadAudienceGroupByFileCallUploadDescription) Apply(call *UploadAudienceGroupByFileCall) {
	call.UploadDescription = w.uploadDescription
}

// UploadAudienceGroupByFile method
func (client *Client) UploadAudienceGroupByFile(description string, audiences []string, options ...IUploadAudienceGroupByFileOption) *UploadAudienceGroupByFileCall {
	call := &UploadAudienceGroupByFileCall{
		c:           client,
		Description: description,
		Audiences:   audiences,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// UploadAudienceGroupByFileCall type
// Deprecated: Use OpenAPI based classes instead.
type UploadAudienceGroupByFileCall struct {
	c                 *Client
	ctx               context.Context
	Description       string   `json:"description,omitempty" validate:"required,max=120"`
	IsIfaAudience     bool     `json:"isIfaAudience,omitempty"`
	UploadDescription string   `json:"uploadDescription,omitempty"`
	Audiences         []string `json:"audiences,omitempty" validate:"max=1500000"`
}

// WithContext method
func (call *UploadAudienceGroupByFileCall) WithContext(ctx context.Context) *UploadAudienceGroupByFileCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *UploadAudienceGroupByFileCall) Do() (*UploadAudienceGroupResponse, error) {
	buf := bytes.NewBuffer([]byte{})
	defer buf.Reset()
	_, errWriteString := buf.WriteString(strings.Join(call.Audiences, "\n"))
	if errWriteString != nil {
		return nil, errWriteString
	}

	form := map[string]io.Reader{
		"description": strings.NewReader(call.Description),
		"file":        buf,
	}
	if call.IsIfaAudience {
		form["isIfaAudience"] = strings.NewReader(strconv.FormatBool(call.IsIfaAudience))
	}
	if call.UploadDescription != "" {
		form["uploadDescription"] = strings.NewReader(call.UploadDescription)
	}
	res, err := call.c.postFormFile(call.ctx, APIAudienceGroupUploadByFile, form)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToAudienceGroupResponse(res)
}

// IAddAudiencesOption type
type IAddAudiencesOption interface {
	Apply(call *AddAudiencesCall)
}

// WithAddAudiencesCallUploadDescription type
// Deprecated: Use OpenAPI based classes instead.
func WithAddAudiencesCallUploadDescription(uploadDescription string) IAddAudiencesOption {
	return &withAddAudiencesCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withAddAudiencesCallUploadDescription struct {
	uploadDescription string
}

func (w *withAddAudiencesCallUploadDescription) Apply(call *AddAudiencesCall) {
	call.UploadDescription = w.uploadDescription
}

// AddAudiences method
func (client *Client) AddAudiences(audienceGroupID int, audiences []string, options ...IAddAudiencesOption) *AddAudiencesCall {
	call := &AddAudiencesCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
	}
	for _, item := range audiences {
		call.Audiences = append(call.Audiences, audience{ID: item})
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// AddAudiencesCall type
// Deprecated: Use OpenAPI based classes instead.
type AddAudiencesCall struct {
	c                 *Client
	ctx               context.Context
	AudienceGroupID   int `validate:"required"`
	UploadDescription string
	Audiences         []audience `validate:"required,max=10000"`
}

// WithContext method
func (call *AddAudiencesCall) WithContext(ctx context.Context) *AddAudiencesCall {
	call.ctx = ctx
	return call
}

func (call *AddAudiencesCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		AudienceGroupID   int        `json:"audienceGroupId,omitempty"`
		UploadDescription string     `json:"uploadDescription,omitempty"`
		Audiences         []audience `json:"audiences,omitempty"`
	}{
		AudienceGroupID:   call.AudienceGroupID,
		UploadDescription: call.UploadDescription,
		Audiences:         call.Audiences,
	})
}

// Do method
func (call *AddAudiencesCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.put(call.ctx, APIAudienceGroupUpload, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// IAddAudiencesByFileOption type
type IAddAudiencesByFileOption interface {
	Apply(call *AddAudiencesByFileCall)
}

// WithAddAudiencesByFileCallUploadDescription func
// Deprecated: Use OpenAPI based classes instead.
func WithAddAudiencesByFileCallUploadDescription(uploadDescription string) IAddAudiencesByFileOption {
	return &withAddAudiencesByFileCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withAddAudiencesByFileCallUploadDescription struct {
	uploadDescription string
}

func (w *withAddAudiencesByFileCallUploadDescription) Apply(call *AddAudiencesByFileCall) {
	call.UploadDescription = w.uploadDescription
}

// AddAudiencesByFile method
func (client *Client) AddAudiencesByFile(audienceGroupID int, audiences []string, options ...IAddAudiencesByFileOption) *AddAudiencesByFileCall {
	call := &AddAudiencesByFileCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
		Audiences:       audiences,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// AddAudiencesByFileCall type
// Deprecated: Use OpenAPI based classes instead.
type AddAudiencesByFileCall struct {
	c                 *Client
	ctx               context.Context
	AudienceGroupID   int      `json:"audienceGroupId,omitempty" validate:"required"`
	UploadDescription string   `json:"uploadDescription,omitempty"`
	Audiences         []string `json:"audiences,omitempty" validate:"required,max=1500000"`
}

// WithContext method
func (call *AddAudiencesByFileCall) WithContext(ctx context.Context) *AddAudiencesByFileCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *AddAudiencesByFileCall) Do() (*BasicResponse, error) {
	buf := bytes.NewBuffer([]byte{})
	defer buf.Reset()
	_, errWriteString := buf.WriteString(strings.Join(call.Audiences, "\n"))
	if errWriteString != nil {
		return nil, errWriteString
	}

	form := map[string]io.Reader{
		"audienceGroupId": strings.NewReader(strconv.FormatInt(int64(call.AudienceGroupID), 10)),
		"file":            buf,
	}
	if call.UploadDescription != "" {
		form["uploadDescription"] = strings.NewReader(call.UploadDescription)
	}
	res, err := call.c.putFormFile(call.ctx, APIAudienceGroupUploadByFile, form)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// IClickAudienceGroupOption type
type IClickAudienceGroupOption interface {
	Apply(call *ClickAudienceGroupCall)
}

// WithClickAudienceGroupCallClickURL func
// Deprecated: Use OpenAPI based classes instead.
func WithClickAudienceGroupCallClickURL(clickURL string) IClickAudienceGroupOption {
	return &withClickAudienceGroupCallClickURL{
		clickURL: clickURL,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withClickAudienceGroupCallClickURL struct {
	clickURL string
}

func (w *withClickAudienceGroupCallClickURL) Apply(call *ClickAudienceGroupCall) {
	call.ClickURL = w.clickURL
}

// ClickAudienceGroup method
func (client *Client) ClickAudienceGroup(description, requestID string, options ...IClickAudienceGroupOption) *ClickAudienceGroupCall {
	call := &ClickAudienceGroupCall{
		c:           client,
		Description: description,
		RequestID:   requestID,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// ClickAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type ClickAudienceGroupCall struct {
	c           *Client
	ctx         context.Context
	Description string `validate:"required,max=120"`
	RequestID   string `validate:"required"`
	ClickURL    string `validate:"max=2000"`
}

// WithContext method
func (call *ClickAudienceGroupCall) WithContext(ctx context.Context) *ClickAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *ClickAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		Description string `json:"description,omitempty"`
		RequestID   string `json:"requestId,omitempty"`
		ClickURL    string `json:"clickUrl,omitempty"`
	}{
		Description: call.Description,
		RequestID:   call.RequestID,
		ClickURL:    call.ClickURL,
	})
}

// Do method
func (call *ClickAudienceGroupCall) Do() (*ClickAudienceGroupResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIAudienceGroupClick, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToClickAudienceGroupResponse(res)
}

// IMPAudienceGroup method
func (client *Client) IMPAudienceGroup(description, requestID string) *IMPAudienceGroupCall {
	call := &IMPAudienceGroupCall{
		c:           client,
		Description: description,
		RequestID:   requestID,
	}
	return call
}

// IMPAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type IMPAudienceGroupCall struct {
	c           *Client
	ctx         context.Context
	Description string `validate:"required,max=120"`
	RequestID   string `validate:"required"`
}

// WithContext method
func (call *IMPAudienceGroupCall) WithContext(ctx context.Context) *IMPAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *IMPAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		Description string `json:"description,omitempty"`
		RequestID   string `json:"requestId,omitempty"`
	}{
		Description: call.Description,
		RequestID:   call.RequestID,
	})
}

// Do method
func (call *IMPAudienceGroupCall) Do() (*IMPAudienceGroupResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIAudienceGroupIMP, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToIMPAudienceGroupResponse(res)
}

// UpdateAudienceGroupDescription method
func (client *Client) UpdateAudienceGroupDescription(audienceGroupID int, description string) *UpdateAudienceGroupDescriptionCall {
	call := &UpdateAudienceGroupDescriptionCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
		Description:     description,
	}
	return call
}

// UpdateAudienceGroupDescriptionCall type
// Deprecated: Use OpenAPI based classes instead.
type UpdateAudienceGroupDescriptionCall struct {
	c               *Client
	ctx             context.Context
	AudienceGroupID int    `json:"-" validate:"required"`
	Description     string `validate:"required,max=120"`
}

// WithContext method
func (call *UpdateAudienceGroupDescriptionCall) WithContext(ctx context.Context) *UpdateAudienceGroupDescriptionCall {
	call.ctx = ctx
	return call
}

func (call *UpdateAudienceGroupDescriptionCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		Description string `json:"description,omitempty"`
	}{
		Description: call.Description,
	})
}

// Do method
func (call *UpdateAudienceGroupDescriptionCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.put(call.ctx, fmt.Sprintf(APIAudienceGroupUpdateDescription, call.AudienceGroupID), &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// ActivateAudienceGroup method
func (client *Client) ActivateAudienceGroup(audienceGroupID int) *ActivateAudienceGroupCall {
	call := &ActivateAudienceGroupCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
	}
	return call
}

// ActivateAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type ActivateAudienceGroupCall struct {
	c               *Client
	ctx             context.Context
	AudienceGroupID int `json:"-" validate:"required"`
}

// WithContext method
func (call *ActivateAudienceGroupCall) WithContext(ctx context.Context) *ActivateAudienceGroupCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *ActivateAudienceGroupCall) Do() (*BasicResponse, error) {
	res, err := call.c.put(call.ctx, fmt.Sprintf(APIAudienceGroupActivate, call.AudienceGroupID), nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// DeleteAudienceGroup method
func (client *Client) DeleteAudienceGroup(audienceGroupID int) *DeleteAudienceGroupCall {
	call := &DeleteAudienceGroupCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
	}
	return call
}

// DeleteAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type DeleteAudienceGroupCall struct {
	c               *Client
	ctx             context.Context
	AudienceGroupID int `json:"-" validate:"required"`
}

// WithContext method
func (call *DeleteAudienceGroupCall) WithContext(ctx context.Context) *DeleteAudienceGroupCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *DeleteAudienceGroupCall) Do() (*BasicResponse, error) {
	res, err := call.c.delete(call.ctx, fmt.Sprintf(APIAudienceGroup, call.AudienceGroupID))
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

// GetAudienceGroup method
func (client *Client) GetAudienceGroup(audienceGroupID int) *GetAudienceGroupCall {
	call := &GetAudienceGroupCall{
		c:               client,
		AudienceGroupID: audienceGroupID,
	}
	return call
}

// GetAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type GetAudienceGroupCall struct {
	c               *Client
	ctx             context.Context
	AudienceGroupID int `json:"-" validate:"required"`
}

// WithContext method
func (call *GetAudienceGroupCall) WithContext(ctx context.Context) *GetAudienceGroupCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetAudienceGroupCall) Do() (*GetAudienceGroupResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, fmt.Sprintf(APIAudienceGroup, call.AudienceGroupID), nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToGetAudienceGroupResponse(res)
}

// IListAudienceGroupOption type
type IListAudienceGroupOption interface {
	Apply(call *ListAudienceGroupCall)
}

// WithListAudienceGroupCallDescription func
// Deprecated: Use OpenAPI based classes instead.
func WithListAudienceGroupCallDescription(description string) IListAudienceGroupOption {
	return &withListAudienceGroupCallDescription{
		description: description,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withListAudienceGroupCallDescription struct {
	description string
}

func (w *withListAudienceGroupCallDescription) Apply(call *ListAudienceGroupCall) {
	call.Description = w.description
}

// WithListAudienceGroupCallStatus func
// Deprecated: Use OpenAPI based classes instead.
func WithListAudienceGroupCallStatus(status AudienceStatusType) IListAudienceGroupOption {
	return &withListAudienceGroupCallStatus{
		status: status,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withListAudienceGroupCallStatus struct {
	status AudienceStatusType
}

func (w *withListAudienceGroupCallStatus) Apply(call *ListAudienceGroupCall) {
	call.Status = w.status
}

// WithListAudienceGroupCallSize func
// Deprecated: Use OpenAPI based classes instead.
func WithListAudienceGroupCallSize(size int) IListAudienceGroupOption {
	return &withListAudienceGroupCallSize{
		size: size,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withListAudienceGroupCallSize struct {
	size int
}

func (w *withListAudienceGroupCallSize) Apply(call *ListAudienceGroupCall) {
	call.Size = w.size
}

// WithListAudienceGroupCallIncludesExternalPublicGroups func
// Deprecated: Use OpenAPI based classes instead.
func WithListAudienceGroupCallIncludesExternalPublicGroups(includesExternalPublicGroups bool) IListAudienceGroupOption {
	return &withListAudienceGroupCallIncludesExternalPublicGroups{
		includesExternalPublicGroups: includesExternalPublicGroups,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withListAudienceGroupCallIncludesExternalPublicGroups struct {
	includesExternalPublicGroups bool
}

func (w *withListAudienceGroupCallIncludesExternalPublicGroups) Apply(call *ListAudienceGroupCall) {
	call.IncludesExternalPublicGroups = w.includesExternalPublicGroups
}

// WithListAudienceGroupCallCreateRoute func
// Deprecated: Use OpenAPI based classes instead.
func WithListAudienceGroupCallCreateRoute(createRoute string) IListAudienceGroupOption {
	return &withListAudienceGroupCallCreateRoute{
		createRoute: createRoute,
	}
}

// Deprecated: Use OpenAPI based classes instead.
type withListAudienceGroupCallCreateRoute struct {
	createRoute string
}

func (w *withListAudienceGroupCallCreateRoute) Apply(call *ListAudienceGroupCall) {
	call.CreateRoute = w.createRoute
}

// ListAudienceGroup method
func (client *Client) ListAudienceGroup(page int, options ...IListAudienceGroupOption) *ListAudienceGroupCall {
	call := &ListAudienceGroupCall{
		c:                            client,
		Page:                         page,
		IncludesExternalPublicGroups: true,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

// ListAudienceGroupCall type
// Deprecated: Use OpenAPI based classes instead.
type ListAudienceGroupCall struct {
	c                            *Client
	ctx                          context.Context
	Page                         int                `json:"-" validate:"required,min=1"`
	Description                  string             `json:"-"`
	Status                       AudienceStatusType `json:"-"`
	Size                         int                `json:"-" validate:"min=1,max=40"`
	IncludesExternalPublicGroups bool               `json:"-"`
	CreateRoute                  string             `json:"-"`
}

// WithContext method
func (call *ListAudienceGroupCall) WithContext(ctx context.Context) *ListAudienceGroupCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *ListAudienceGroupCall) Do() (*ListAudienceGroupResponse, error) {
	u := url.Values{}
	u.Set("page", strconv.FormatInt(int64(call.Page), 10))
	u.Set("size", strconv.FormatInt(int64(call.Size), 10))
	if call.Description != "" {
		u.Set("description", call.Description)
	}
	if call.Status != "" {
		u.Set("status", call.Status.String())
	}
	if !call.IncludesExternalPublicGroups {
		u.Set("includesExternalPublicGroups", strconv.FormatBool(call.IncludesExternalPublicGroups))
	}
	if call.CreateRoute != "" {
		u.Set("createRoute", call.CreateRoute)
	}
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIAudienceGroupList, u)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToListAudienceGroupResponse(res)
}

// GetAudienceGroupAuthorityLevel method
func (client *Client) GetAudienceGroupAuthorityLevel() *GetAudienceGroupAuthorityLevelCall {
	return &GetAudienceGroupAuthorityLevelCall{
		c: client,
	}
}

// GetAudienceGroupAuthorityLevelCall type
// Deprecated: Use OpenAPI based classes instead.
type GetAudienceGroupAuthorityLevelCall struct {
	c   *Client
	ctx context.Context
}

// WithContext method
func (call *GetAudienceGroupAuthorityLevelCall) WithContext(ctx context.Context) *GetAudienceGroupAuthorityLevelCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetAudienceGroupAuthorityLevelCall) Do() (*GetAudienceGroupAuthorityLevelResponse, error) {
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIAudienceGroupAuthorityLevel, nil)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToGetAudienceGroupAuthorityLevelResponse(res)
}

// ChangeAudienceGroupAuthorityLevel method
func (client *Client) ChangeAudienceGroupAuthorityLevel(authorityLevel AudienceAuthorityLevelType) *ChangeAudienceGroupAuthorityLevelCall {
	call := &ChangeAudienceGroupAuthorityLevelCall{
		c:              client,
		AuthorityLevel: authorityLevel,
	}
	return call
}

// ChangeAudienceGroupAuthorityLevelCall type
// Deprecated: Use OpenAPI based classes instead.
type ChangeAudienceGroupAuthorityLevelCall struct {
	c              *Client
	ctx            context.Context
	AuthorityLevel AudienceAuthorityLevelType `validate:"required"`
}

// WithContext method
func (call *ChangeAudienceGroupAuthorityLevelCall) WithContext(ctx context.Context) *ChangeAudienceGroupAuthorityLevelCall {
	call.ctx = ctx
	return call
}

func (call *ChangeAudienceGroupAuthorityLevelCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(struct {
		AuthorityLevel AudienceAuthorityLevelType `json:"authorityLevel,omitempty"`
	}{
		AuthorityLevel: call.AuthorityLevel,
	})
}

// Do method
func (call *ChangeAudienceGroupAuthorityLevelCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.put(call.ctx, APIAudienceGroupAuthorityLevel, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}
