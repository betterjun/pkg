package validation

import "regexp"

type VI interface {
	// 全部
	Required()

	// 数字
	Max(int)
	Min(int)
	Range(min, max int)

	// 字符串和数组
	Length(int)
	MinLength(int)
	MaxLength(int)

	// 字符串
	Alpha()
	Numeric()
	AlphaNumeric()
	AlphaDash()
	Email()
	IP()
	Base64()
	Mobile()
	Tel()
	Phone()
	ZipCode()

	Match(regex *regexp.Regexp)
	NoMatch(regex *regexp.Regexp)
}

// A Validation context manages data validation and error messages.
type Validation struct {
	Target interface{}
	Name   string
	Errors []string
}

func (v *Validation) setError(errMsg string) {
	v.Errors = append(v.Errors, errMsg)
}

func (v *Validation) Passed() (ok bool) {
	return len(v.Errors) == 0
}

func (v *Validation) GetErrors() []string {
	return v.Errors
}

// Required Test that the argument is non-nil and non-empty (if string or list)
func (v *Validation) Required() *Validation {
	return v.apply(Required{}, v.Target)
}

func (v *Validation) Min(val int) *Validation {
	return v.apply(Min{val}, v.Target)
}

func (v *Validation) Max(val int) *Validation {
	return v.apply(Max{val}, v.Target)
}

func (v *Validation) Range(min, max int) *Validation {
	return v.apply(Range{Min{min}, Max{max}}, v.Target)
}

// MinLength Test that the obj is longer than min size if type is string or slice
func (v *Validation) MinLength(min int) *Validation {
	return v.apply(MinLength{min}, v.Target)
}

// MaxLength Test that the obj is shorter than max size if type is string or slice
func (v *Validation) MaxLength(max int) *Validation {
	return v.apply(MaxLength{max}, v.Target)
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
