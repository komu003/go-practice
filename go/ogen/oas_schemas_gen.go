// Code generated by ogen, DO NOT EDIT.

package ogen

// APIMicropostsCountGetDef is default response for APIMicropostsCountGet operation.
type APIMicropostsCountGetDef struct {
	StatusCode int
}

// GetStatusCode returns the value of StatusCode.
func (s *APIMicropostsCountGetDef) GetStatusCode() int {
	return s.StatusCode
}

// SetStatusCode sets the value of StatusCode.
func (s *APIMicropostsCountGetDef) SetStatusCode(val int) {
	s.StatusCode = val
}

func (*APIMicropostsCountGetDef) aPIMicropostsCountGetRes() {}

// APIMicropostsCountGetInternalServerError is response for APIMicropostsCountGet operation.
type APIMicropostsCountGetInternalServerError struct{}

func (*APIMicropostsCountGetInternalServerError) aPIMicropostsCountGetRes() {}

// APIUsersCountGetDef is default response for APIUsersCountGet operation.
type APIUsersCountGetDef struct {
	StatusCode int
}

// GetStatusCode returns the value of StatusCode.
func (s *APIUsersCountGetDef) GetStatusCode() int {
	return s.StatusCode
}

// SetStatusCode sets the value of StatusCode.
func (s *APIUsersCountGetDef) SetStatusCode(val int) {
	s.StatusCode = val
}

func (*APIUsersCountGetDef) aPIUsersCountGetRes() {}

// APIUsersCountGetInternalServerError is response for APIUsersCountGet operation.
type APIUsersCountGetInternalServerError struct{}

func (*APIUsersCountGetInternalServerError) aPIUsersCountGetRes() {}

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