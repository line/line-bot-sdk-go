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

type IUploadAudienceGroupOption interface {
	Apply(call *UploadAudienceGroupCall)
}

func WithUploadAudienceGroupCallIsIfaAudience(isIfaAudience bool) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallIsIfaAudience{
		isIfaAudience: isIfaAudience,
	}
}

type withUploadAudienceGroupCallIsIfaAudience struct {
	isIfaAudience bool
}

func (w *withUploadAudienceGroupCallIsIfaAudience) Apply(call *UploadAudienceGroupCall) {
	call.IsIfaAudience = w.isIfaAudience
}

func WithUploadAudienceGroupCallUploadDescription(uploadDescription string) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

type withUploadAudienceGroupCallUploadDescription struct {
	uploadDescription string
}

func (w *withUploadAudienceGroupCallUploadDescription) Apply(call *UploadAudienceGroupCall) {
	call.UploadDescription = w.uploadDescription
}

func WithUploadAudienceGroupCallAudiences(audiences ...string) IUploadAudienceGroupOption {
	return &withUploadAudienceGroupCallAudiences{
		audiences: audiences,
	}
}

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

type audience struct {
	ID string `json:"id,omitempty"`
}

type UploadAudienceGroupCall struct {
	c                 *Client
	ctx               context.Context
	Description       string     `json:"description,omitempty" validate:"required,max=120"`
	IsIfaAudience     bool       `json:"isIfaAudience,omitempty"`
	UploadDescription string     `json:"uploadDescription,omitempty"`
	Audiences         []audience `json:"audiences,omitempty" validate:"max=10000"`
}

// WithContext method
func (call *UploadAudienceGroupCall) WithContext(ctx context.Context) *UploadAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *UploadAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
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

type IUploadAudienceGroupByFileOption interface {
	Apply(call *UploadAudienceGroupByFileCall)
}

func WithUploadAudienceGroupByFileCallIsIfaAudience(isIfaAudience bool) IUploadAudienceGroupByFileOption {
	return &withUploadAudienceGroupByFileCallIsIfaAudience{
		isIfaAudience: isIfaAudience,
	}
}

type withUploadAudienceGroupByFileCallIsIfaAudience struct {
	isIfaAudience bool
}

func (w *withUploadAudienceGroupByFileCallIsIfaAudience) Apply(call *UploadAudienceGroupByFileCall) {
	call.IsIfaAudience = w.isIfaAudience
}

func WithUploadAudienceGroupByFileCallUploadDescription(uploadDescription string) IUploadAudienceGroupByFileOption {
	return &withUploadAudienceGroupByFileCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

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

type IAddAudiencesOption interface {
	Apply(call *AddAudiencesCall)
}

func WithAddAudiencesCallUploadDescription(uploadDescription string) IAddAudiencesOption {
	return &withAddAudiencesCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

type withAddAudiencesCallUploadDescription struct {
	uploadDescription string
}

func (w *withAddAudiencesCallUploadDescription) Apply(call *AddAudiencesCall) {
	call.UploadDescription = w.uploadDescription
}

// AddAudiences method
func (client *Client) AddAudiences(audienceGroupId int, audiences []string, options ...IAddAudiencesOption) *AddAudiencesCall {
	call := &AddAudiencesCall{
		c:               client,
		AudienceGroupID: audienceGroupId,
	}
	for _, item := range audiences {
		call.Audiences = append(call.Audiences, audience{ID: item})
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

type AddAudiencesCall struct {
	c                 *Client
	ctx               context.Context
	AudienceGroupID   int        `json:"audienceGroupId,omitempty" validate:"required"`
	UploadDescription string     `json:"uploadDescription,omitempty"`
	Audiences         []audience `json:"audiences,omitempty" validate:"required,max=10000"`
}

// WithContext method
func (call *AddAudiencesCall) WithContext(ctx context.Context) *AddAudiencesCall {
	call.ctx = ctx
	return call
}

func (call *AddAudiencesCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
}

// Do method
func (call *AddAudiencesCall) Do() (*BasicResponse, error) {
	var buf bytes.Buffer
	if err := call.encodeJSON(&buf); err != nil {
		return nil, err
	}
	res, err := call.c.post(call.ctx, APIAudienceGroupUpload, &buf)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToBasicResponse(res)
}

type IAddAudiencesByFileOption interface {
	Apply(call *AddAudiencesByFileCall)
}

func WithAddAudiencesByFileCallUploadDescription(uploadDescription string) IAddAudiencesByFileOption {
	return &withAddAudiencesByFileCallUploadDescription{
		uploadDescription: uploadDescription,
	}
}

type withAddAudiencesByFileCallUploadDescription struct {
	uploadDescription string
}

func (w *withAddAudiencesByFileCallUploadDescription) Apply(call *AddAudiencesByFileCall) {
	call.UploadDescription = w.uploadDescription
}

// AddAudiencesByFile method
func (client *Client) AddAudiencesByFile(audienceGroupId int, audiences []string, options ...IAddAudiencesByFileOption) *AddAudiencesByFileCall {
	call := &AddAudiencesByFileCall{
		c:               client,
		AudienceGroupID: audienceGroupId,
		Audiences:       audiences,
	}
	for _, item := range options {
		item.Apply(call)
	}
	return call
}

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

type IClickAudienceGroupOption interface {
	Apply(call *ClickAudienceGroupCall)
}

func WithClickAudienceGroupCallClickURL(clickURL string) IClickAudienceGroupOption {
	return &withClickAudienceGroupCallClickURL{
		clickURL: clickURL,
	}
}

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

type ClickAudienceGroupCall struct {
	c           *Client
	ctx         context.Context
	Description string `json:"description,omitempty" validate:"required,max=120"`
	RequestID   string `json:"requestId,omitempty" validate:"required"`
	ClickURL    string `json:"clickUrl,omitempty" validate:"max=2000"`
}

// WithContext method
func (call *ClickAudienceGroupCall) WithContext(ctx context.Context) *ClickAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *ClickAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
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

type IMPAudienceGroupCall struct {
	c           *Client
	ctx         context.Context
	Description string `json:"description,omitempty" validate:"required,max=120"`
	RequestID   string `json:"requestId,omitempty" validate:"required"`
}

// WithContext method
func (call *IMPAudienceGroupCall) WithContext(ctx context.Context) *IMPAudienceGroupCall {
	call.ctx = ctx
	return call
}

func (call *IMPAudienceGroupCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
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
func (client *Client) UpdateAudienceGroupDescription(audienceGroupId int, description string) *UpdateAudienceGroupDescriptionCall {
	call := &UpdateAudienceGroupDescriptionCall{
		c:               client,
		AudienceGroupID: audienceGroupId,
		Description:     description,
	}
	return call
}

type UpdateAudienceGroupDescriptionCall struct {
	c               *Client
	ctx             context.Context
	AudienceGroupID int    `json:"-" validate:"required"`
	Description     string `json:"description,omitempty" validate:"required,max=120"`
}

// WithContext method
func (call *UpdateAudienceGroupDescriptionCall) WithContext(ctx context.Context) *UpdateAudienceGroupDescriptionCall {
	call.ctx = ctx
	return call
}

func (call *UpdateAudienceGroupDescriptionCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
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
func (client *Client) ActivateAudienceGroup(audienceGroupId int) *ActivateAudienceGroupCall {
	call := &ActivateAudienceGroupCall{
		c:               client,
		AudienceGroupID: audienceGroupId,
	}
	return call
}

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

// GetAudienceGroup method
func (client *Client) GetAudienceGroup(audienceGroupId int) *GetAudienceGroupCall {
	call := &GetAudienceGroupCall{
		c:               client,
		AudienceGroupID: audienceGroupId,
	}
	return call
}

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

type IListAudienceGroupOption interface {
	Apply(call *ListAudienceGroupCall)
}

func WithListAudienceGroupCallDescription(description string) IListAudienceGroupOption {
	return &withListAudienceGroupCallDescription{
		description: description,
	}
}

type withListAudienceGroupCallDescription struct {
	description string
}

func (w *withListAudienceGroupCallDescription) Apply(call *ListAudienceGroupCall) {
	call.Description = w.description
}

func WithListAudienceGroupCallStatus(status string) IListAudienceGroupOption {
	return &withListAudienceGroupCallStatus{
		status: status,
	}
}

type withListAudienceGroupCallStatus struct {
	status string
}

func (w *withListAudienceGroupCallStatus) Apply(call *ListAudienceGroupCall) {
	call.Status = w.status
}

func WithListAudienceGroupCallSize(size int) IListAudienceGroupOption {
	return &withListAudienceGroupCallSize{
		size: size,
	}
}

type withListAudienceGroupCallSize struct {
	size int
}

func (w *withListAudienceGroupCallSize) Apply(call *ListAudienceGroupCall) {
	call.Size = w.size
}

func WithListAudienceGroupCallIncludesExternalPublicGroups(includesExternalPublicGroups bool) IListAudienceGroupOption {
	return &withListAudienceGroupCallIncludesExternalPublicGroups{
		includesExternalPublicGroups: includesExternalPublicGroups,
	}
}

type withListAudienceGroupCallIncludesExternalPublicGroups struct {
	includesExternalPublicGroups bool
}

func (w *withListAudienceGroupCallIncludesExternalPublicGroups) Apply(call *ListAudienceGroupCall) {
	call.IncludesExternalPublicGroups = w.includesExternalPublicGroups
}

func WithListAudienceGroupCallCreateRoute(createRoute string) IListAudienceGroupOption {
	return &withListAudienceGroupCallCreateRoute{
		createRoute: createRoute,
	}
}

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

type ListAudienceGroupCall struct {
	c                            *Client
	ctx                          context.Context
	Page                         int    `json:"-" validate:"required,min=1"`
	Description                  string `json:"-"`
	Status                       string `json:"-"`
	Size                         int    `json:"-" validate:"min=1,max=40"`
	IncludesExternalPublicGroups bool   `json:"-"`
	CreateRoute                  string `json:"-"`
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
		u.Set("status", call.Status)
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

// ChangeAudienceGroupAuthorityLevel method
func (client *Client) ChangeAudienceGroupAuthorityLevel(authorityLevel string) *ChangeAudienceGroupAuthorityLevelCall {
	call := &ChangeAudienceGroupAuthorityLevelCall{
		c:              client,
		AuthorityLevel: authorityLevel,
	}
	return call
}

type ChangeAudienceGroupAuthorityLevelCall struct {
	c              *Client
	ctx            context.Context
	AuthorityLevel string `json:"authorityLevel,omitempty" validate:"required"`
}

// WithContext method
func (call *ChangeAudienceGroupAuthorityLevelCall) WithContext(ctx context.Context) *ChangeAudienceGroupAuthorityLevelCall {
	call.ctx = ctx
	return call
}

func (call *ChangeAudienceGroupAuthorityLevelCall) encodeJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(call)
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
