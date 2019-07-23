package validation

import "regexp"

// A Validation context manages data validation and error messages.
type Validation struct {
	Target interface{}
	Name   string
	Errors []string
}

func (v *Validation) setError(errMsg string) {
	v.Errors = append(v.Errors, errMsg)
}

func (v *Validation) Passed() (ok bool, errMsgs []string) {
	return len(v.Errors) == 0, v.Errors
}

// Required Test that the argument is non-nil and non-empty (if string or list)
func (v *Validation) Required() *Validation {
	return v.apply(Required{}, v.Target)
}

// Min Test that the obj is greater than min if obj's type is int
func (v *Validation) Min(min int) *Validation {
	return v.apply(Min{min}, v.Target)
}

// Max Test that the obj is less than max if obj's type is int
func (v *Validation) Max(max int) *Validation {
	return v.apply(Max{max}, v.Target)
}

// Range Test that the obj is between mni and max if obj's type is int
func (v *Validation) Range(min, max int) *Validation {
	return v.apply(Range{Min{Min: min}, Max{Max: max}}, v.Target)
}

// MinSize Test that the obj is longer than min size if type is string or slice
func (v *Validation) MinSize(min int) *Validation {
	return v.apply(MinSize{min}, v.Target)
}

// MaxSize Test that the obj is shorter than max size if type is string or slice
func (v *Validation) MaxSize(max int) *Validation {
	return v.apply(MaxSize{max}, v.Target)
}

// Length Test that the obj is same length to n if type is string or slice
func (v *Validation) Length(n int) *Validation {
	return v.apply(Length{n}, v.Target)
}

// Alpha Test that the obj is [a-zA-Z] if type is string
func (v *Validation) Alpha() *Validation {
	return v.apply(Alpha{}, v.Target)
}

// Numeric Test that the obj is [0-9] if type is string
func (v *Validation) Numeric() *Validation {
	return v.apply(Numeric{}, v.Target)
}

// AlphaNumeric Test that the obj is [0-9a-zA-Z] if type is string
func (v *Validation) AlphaNumeric() *Validation {
	return v.apply(AlphaNumeric{}, v.Target)
}

// Match Test that the obj matches regexp if type is string
func (v *Validation) Match(regex *regexp.Regexp) *Validation {
	return v.apply(Match{regex}, v.Target)
}

// NoMatch Test that the obj doesn't match regexp if type is string
func (v *Validation) NoMatch(regex *regexp.Regexp) *Validation {
	return v.apply(NoMatch{Match{Regexp: regex}}, v.Target)
}

// AlphaDash Test that the obj is [0-9a-zA-Z_-] if type is string
func (v *Validation) AlphaDash() *Validation {
	return v.apply(AlphaDash{NoMatch{Match: Match{Regexp: alphaDashPattern}}}, v.Target)
}

// Email Test that the obj is email address if type is string
func (v *Validation) Email() *Validation {
	return v.apply(Email{Match{Regexp: emailPattern}}, v.Target)
}

// IP Test that the obj is IP address if type is string
func (v *Validation) IP() *Validation {
	return v.apply(IP{Match{Regexp: ipPattern}}, v.Target)
}

// Base64 Test that the obj is base64 encoded if type is string
func (v *Validation) Base64() *Validation {
	return v.apply(Base64{Match{Regexp: base64Pattern}}, v.Target)
}

// Mobile Test that the obj is chinese mobile number if type is string
func (v *Validation) Mobile() *Validation {
	return v.apply(Mobile{Match{Regexp: mobilePattern}}, v.Target)
}

// Tel Test that the obj is chinese telephone number if type is string
func (v *Validation) Tel() *Validation {
	return v.apply(Tel{Match{Regexp: telPattern}}, v.Target)
}

// Phone Test that the obj is chinese mobile or telephone number if type is string
func (v *Validation) Phone() *Validation {
	return v.apply(Phone{Mobile{Match: Match{Regexp: mobilePattern}},
		Tel{Match: Match{Regexp: telPattern}}}, v.Target)
}

// ZipCode Test that the obj is chinese zip code if type is string
func (v *Validation) ZipCode() *Validation {
	return v.apply(ZipCode{Match{Regexp: zipCodePattern}}, v.Target)
}

func (v *Validation) apply(chk Validator, obj interface{}) *Validation {
	if !chk.IsSatisfied(obj) {
		v.setError(v.Name + " " + chk.DefaultMessage())
	}

	return v
}

// A Validation context manages data validation and error messages.
type ValidationGroup struct {
	Fields []*Validation
}

func (vg *ValidationGroup) Validate(obj interface{}, name string) (v *Validation) {
	v = &Validation{Target: obj, Name: name}
	vg.Fields = append(vg.Fields, v)
	return v
}

func (vg *ValidationGroup) Passed() (ok bool, errMsgs []string) {
	for _, v := range vg.Fields {
		ret, errs := v.Passed()
		if !ret {
			errMsgs = append(errMsgs, errs...)
		}
	}

	return len(errMsgs) == 0, errMsgs
}
