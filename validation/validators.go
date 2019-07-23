// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"time"
	"unicode/utf8"
)

// MessageTmpls store commond validate template
var MessageTmpls = map[string]string{
	"Required":     "Can not be empty",
	"Min":          "Minimum is %d",
	"Max":          "Maximum is %d",
	"Range":        "Range is %d to %d",
	"MinSize":      "Minimum size is %d",
	"MaxSize":      "Maximum size is %d",
	"Length":       "Required length is %d",
	"Alpha":        "Must be valid alpha characters",
	"Numeric":      "Must be valid numeric characters",
	"AlphaNumeric": "Must be valid alpha or numeric characters",
	"Match":        "Must match %s",
	"NoMatch":      "Must not match %s",
	"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
	"Email":        "Must be a valid email address",
	"IP":           "Must be a valid ip address",
	"Base64":       "Must be valid base64 characters",
	"Mobile":       "Must be valid mobile number",
	"Tel":          "Must be valid telephone number",
	"Phone":        "Must be valid telephone or mobile phone number",
	"ZipCode":      "Must be valid zipcode",
}

// Validator interface
type Validator interface {
	IsSatisfied(interface{}) bool
	DefaultMessage() string
}

// Required struct
type Required struct {
}

// IsSatisfied judge whether obj has value
func (r Required) IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}

	switch v := obj.(type) {
	case bool:
		return v
	case string:
		return len(v) > 0
	case int:
		return v != 0
	case uint:
		return v != 0
	case int8:
		return v != 0
	case uint8:
		return v != 0
	case int16:
		return v != 0
	case uint16:
		return v != 0
	case int32:
		return v != 0
	case uint32:
		return v != 0
	case int64:
		return v != 0
	case uint64:
		return v != 0
	case time.Time:
		return !v.IsZero()
	default:
		val := reflect.ValueOf(obj)
		if val.Kind() == reflect.Slice {
			return val.Len() > 0
		}
	}

	return false
}

// DefaultMessage return the default error message
func (r Required) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Required"])
}

// Min check struct
type Min struct {
	Min int
}

// IsSatisfied judge whether obj is valid
func (m Min) IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}

	switch v := obj.(type) {
	case int:
		return v >= int(m.Min)
	case uint:
		return v >= uint(m.Min)
	case int8:
		return v >= int8(m.Min)
	case uint8:
		return v >= uint8(m.Min)
	case int16:
		return v >= int16(m.Min)
	case uint16:
		return v >= uint16(m.Min)
	case int32:
		return v >= int32(m.Min)
	case uint32:
		return v >= uint32(m.Min)
	case int64:
		return v >= int64(m.Min)
	case uint64:
		return v >= uint64(m.Min)
	}

	return false
}

// DefaultMessage return the default min error message
func (m Min) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["Min"], m.Min)
}

// Max validate struct
type Max struct {
	Max int
}

// IsSatisfied judge whether obj is valid
func (m Max) IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}

	switch v := obj.(type) {
	case int:
		return v <= int(m.Max)
	case uint:
		return v <= uint(m.Max)
	case int8:
		return v <= int8(m.Max)
	case uint8:
		return v <= uint8(m.Max)
	case int16:
		return v <= int16(m.Max)
	case uint16:
		return v <= uint16(m.Max)
	case int32:
		return v <= int32(m.Max)
	case uint32:
		return v <= uint32(m.Max)
	case int64:
		return v <= int64(m.Max)
	case uint64:
		return v <= uint64(m.Max)
	}

	return false
}

// DefaultMessage return the default max error message
func (m Max) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["Max"], m.Max)
}

// Range Requires an integer to be within Min, Max inclusive.
type Range struct {
	Min
	Max
}

// IsSatisfied judge whether obj is valid
func (r Range) IsSatisfied(obj interface{}) bool {
	return r.Min.IsSatisfied(obj) && r.Max.IsSatisfied(obj)
}

// DefaultMessage return the default Range error message
func (r Range) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["Range"], r.Min.Min, r.Max.Max)
}

// MinSize Requires an array or string to be at least a given length.
type MinSize struct {
	Min int
}

// IsSatisfied judge whether obj is valid
func (m MinSize) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) >= m.Min
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() >= m.Min
	}
	return false
}

// DefaultMessage return the default MinSize error message
func (m MinSize) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["MinSize"], m.Min)
}

// MaxSize Requires an array or string to be at most a given length.
type MaxSize struct {
	Max int
}

// IsSatisfied judge whether obj is valid
func (m MaxSize) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) <= m.Max
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() <= m.Max
	}
	return false
}

// DefaultMessage return the default MaxSize error message
func (m MaxSize) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["MaxSize"], m.Max)
}

// Length Requires an array or string to be exactly a given length.
type Length struct {
	N int
}

// IsSatisfied judge whether obj is valid
func (l Length) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		return utf8.RuneCountInString(str) == l.N
	}
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		return v.Len() == l.N
	}
	return false
}

// DefaultMessage return the default Length error message
func (l Length) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["Length"], l.N)
}

// Alpha check the alpha
type Alpha struct {
}

// IsSatisfied judge whether obj is valid
func (a Alpha) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		for _, v := range str {
			if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
				return false
			}
		}
		return true
	}
	return false
}

// DefaultMessage return the default Length error message
func (a Alpha) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Alpha"])
}

// Numeric check number
type Numeric struct {
}

// IsSatisfied judge whether obj is valid
func (n Numeric) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		for _, v := range str {
			if '9' < v || v < '0' {
				return false
			}
		}
		return true
	}
	return false
}

// DefaultMessage return the default Length error message
func (n Numeric) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Numeric"])
}

// AlphaNumeric check alpha and number
type AlphaNumeric struct {
}

// IsSatisfied judge whether obj is valid
func (a AlphaNumeric) IsSatisfied(obj interface{}) bool {
	if str, ok := obj.(string); ok {
		for _, v := range str {
			if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
				return false
			}
		}
		return true
	}
	return false
}

// DefaultMessage return the default Length error message
func (a AlphaNumeric) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["AlphaNumeric"])
}

// Match Requires a string to match a given regex.
type Match struct {
	Regexp *regexp.Regexp
}

// IsSatisfied judge whether obj is valid
func (m Match) IsSatisfied(obj interface{}) bool {
	return m.Regexp.MatchString(fmt.Sprintf("%v", obj))
}

// DefaultMessage return the default Match error message
func (m Match) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["Match"], m.Regexp.String())
}

// NoMatch Requires a string to not match a given regex.
type NoMatch struct {
	Match
}

// IsSatisfied judge whether obj is valid
func (n NoMatch) IsSatisfied(obj interface{}) bool {
	return !n.Match.IsSatisfied(obj)
}

// DefaultMessage return the default NoMatch error message
func (n NoMatch) DefaultMessage() string {
	return fmt.Sprintf(MessageTmpls["NoMatch"], n.Regexp.String())
}

var alphaDashPattern = regexp.MustCompile("[^\\d\\w-_]")

// AlphaDash check not Alpha
type AlphaDash struct {
	NoMatch
}

// DefaultMessage return the default AlphaDash error message
func (a AlphaDash) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["AlphaDash"])
}

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

// Email check struct
type Email struct {
	Match
}

// DefaultMessage return the default Email error message
func (e Email) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Email"])
}

var ipPattern = regexp.MustCompile("^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$")

// IP check struct
type IP struct {
	Match
}

// DefaultMessage return the default IP error message
func (i IP) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["IP"])
}

var base64Pattern = regexp.MustCompile("^(?:[A-Za-z0-99+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")

// Base64 check struct
type Base64 struct {
	Match
}

// DefaultMessage return the default Base64 error message
func (b Base64) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Base64"])
}

// just for chinese mobile phone number
var mobilePattern = regexp.MustCompile("^((\\+86)|(86))?(1(([35][0-9])|[8][0-9]|[7][06789]|[4][579]))\\d{8}$")

// Mobile check struct
type Mobile struct {
	Match
}

// DefaultMessage return the default Mobile error message
func (m Mobile) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Mobile"])
}

// just for chinese telephone number
var telPattern = regexp.MustCompile("^(0\\d{2,3}(\\-)?)?\\d{7,8}$")

// Tel check telephone struct
type Tel struct {
	Match
}

// DefaultMessage return the default Tel error message
func (t Tel) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Tel"])
}

// Phone just for chinese telephone or mobile phone number
type Phone struct {
	Mobile
	Tel
}

// IsSatisfied judge whether obj is valid
func (p Phone) IsSatisfied(obj interface{}) bool {
	return p.Mobile.IsSatisfied(obj) || p.Tel.IsSatisfied(obj)
}

// DefaultMessage return the default Phone error message
func (p Phone) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["Phone"])
}

// just for chinese zipcode
var zipCodePattern = regexp.MustCompile("^[1-9]\\d{5}$")

// ZipCode check the zip struct
type ZipCode struct {
	Match
}

// DefaultMessage return the default Zip error message
func (z ZipCode) DefaultMessage() string {
	return fmt.Sprint(MessageTmpls["ZipCode"])
}
