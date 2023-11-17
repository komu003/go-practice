// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"time"
)

// APIMicropostsCountGetInternalServerError is response for APIMicropostsCountGet operation.
type APIMicropostsCountGetInternalServerError struct{}

func (*APIMicropostsCountGetInternalServerError) aPIMicropostsCountGetRes() {}

// APIMicropostsGetInternalServerError is response for APIMicropostsGet operation.
type APIMicropostsGetInternalServerError struct{}

func (*APIMicropostsGetInternalServerError) aPIMicropostsGetRes() {}

type APIMicropostsGetOKApplicationJSON []Micropost

func (*APIMicropostsGetOKApplicationJSON) aPIMicropostsGetRes() {}

// APIUsersCountGetInternalServerError is response for APIUsersCountGet operation.
type APIUsersCountGetInternalServerError struct{}

func (*APIUsersCountGetInternalServerError) aPIUsersCountGetRes() {}

// APIUsersGetInternalServerError is response for APIUsersGet operation.
type APIUsersGetInternalServerError struct{}

func (*APIUsersGetInternalServerError) aPIUsersGetRes() {}

type APIUsersGetOKApplicationJSON []User

func (*APIUsersGetOKApplicationJSON) aPIUsersGetRes() {}

// Ref: #/components/schemas/CountResponse
type CountResponse struct {
	// カウント数.
	Count OptInt `json:"count"`
}

// GetCount returns the value of Count.
func (s *CountResponse) GetCount() OptInt {
	return s.Count
}

// SetCount sets the value of Count.
func (s *CountResponse) SetCount(val OptInt) {
	s.Count = val
}

func (*CountResponse) aPIMicropostsCountGetRes() {}
func (*CountResponse) aPIUsersCountGetRes()      {}

// Ref: #/components/schemas/Micropost
type Micropost struct {
	// マイクロポストのID.
	ID OptInt `json:"id"`
	// マイクロポストの内容.
	Content OptString `json:"content"`
	// 作成したユーザーのID.
	UserId OptInt `json:"userId"`
	// マイクロポストの作成日時.
	CreatedAt OptDateTime `json:"createdAt"`
	// マイクロポストの更新日時.
	UpdatedAt OptDateTime `json:"updatedAt"`
	User      OptUser     `json:"user"`
}

// GetID returns the value of ID.
func (s *Micropost) GetID() OptInt {
	return s.ID
}

// GetContent returns the value of Content.
func (s *Micropost) GetContent() OptString {
	return s.Content
}

// GetUserId returns the value of UserId.
func (s *Micropost) GetUserId() OptInt {
	return s.UserId
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Micropost) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Micropost) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// GetUser returns the value of User.
func (s *Micropost) GetUser() OptUser {
	return s.User
}

// SetID sets the value of ID.
func (s *Micropost) SetID(val OptInt) {
	s.ID = val
}

// SetContent sets the value of Content.
func (s *Micropost) SetContent(val OptString) {
	s.Content = val
}

// SetUserId sets the value of UserId.
func (s *Micropost) SetUserId(val OptInt) {
	s.UserId = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Micropost) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Micropost) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// SetUser sets the value of User.
func (s *Micropost) SetUser(val OptUser) {
	s.User = val
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUser returns new OptUser with value set to v.
func NewOptUser(v User) OptUser {
	return OptUser{
		Value: v,
		Set:   true,
	}
}

// OptUser is optional User.
type OptUser struct {
	Value User
	Set   bool
}

// IsSet returns true if OptUser was set.
func (o OptUser) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUser) Reset() {
	var v User
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUser) SetTo(v User) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUser) Get() (v User, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUser) Or(d User) User {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/User
type User struct {
	// ユーザーのID.
	ID OptInt `json:"id"`
	// ユーザーの名前.
	Name OptString `json:"name"`
	// ユーザーのメールアドレス.
	Email OptString `json:"email"`
	// アカウントの作成日時.
	CreatedAt OptDateTime `json:"createdAt"`
	// アカウントの更新日時.
	UpdatedAt OptDateTime `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *User) GetID() OptInt {
	return s.ID
}

// GetName returns the value of Name.
func (s *User) GetName() OptString {
	return s.Name
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() OptString {
	return s.Email
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *User) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val OptInt) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *User) SetName(val OptString) {
	s.Name = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val OptString) {
	s.Email = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *User) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}
