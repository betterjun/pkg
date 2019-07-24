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
	"testing"
	"time"
)

func TestRequired(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate(nil, "nil").Required().Passed() {
		t.Error("nil object should be false")
	}
	if !vg.Validate(true, "true").Required().Passed() {
		t.Error("bool true should be true")
	}
	if vg.Validate(false, "false").Required().Passed() {
		t.Error("bool false should be false")
	}
	if vg.Validate("", "emptyString").Required().Passed() {
		t.Error("emptyString should be false")
	}
	if !vg.Validate("string", "string").Required().Passed() {
		t.Error("string should be true")
	}
	if vg.Validate(0, "zero").Required().Passed() {
		t.Error("number 0 should be false")
	}
	if !vg.Validate(1, "one").Required().Passed() {
		t.Error("number 1 should be true")
	}
	if !vg.Validate(time.Now(), "time").Required().Passed() {
		t.Error("time should be true")
	}
	if vg.Validate([]string{}, "emptyStringSlice").Required().Passed() {
		t.Error("emptyStringSlice should be false")
	}
	if vg.Validate([]int{}, "emptyIntSlice").Required().Passed() {
		t.Error("emptyIntSlice should be false")
	}
	if vg.Validate([]interface{}{}, "emptyInterfaceSlice").Required().Passed() {
		t.Error("emptyInterfaceSlice should be false")
	}
	if !vg.Validate([]string{"s1", "s2"}, "stringSlice").Required().Passed() {
		t.Error("stringSlice should be false")
	}
	if !vg.Validate([]int{1, 2}, "intSlice").Required().Passed() {
		t.Error("intSlice should be false")
	}
	if !vg.Validate([]interface{}{"str", 1}, "interfaceSlice").Required().Passed() {
		t.Error("interfaceSlice should be false")
	}

}

func TestMin(t *testing.T) {
	vg := ValidationGroup{}
	var i int = 1
	var i8 int8 = 8
	var ui uint = 30

	if !vg.Validate(i, "int").Min(1).Passed() {
		t.Error("failed")
	}
	if !vg.Validate(i8, "int8").Min(3).Passed() {
		t.Error("failed")
	}
	if vg.Validate(ui, "uint").Min(300).Passed() {
		t.Error("failed")
	}
}

func TestMax(t *testing.T) {
	vg := ValidationGroup{}
	var i int = 1
	var i8 int8 = 8
	var ui uint = 30

	if !vg.Validate(i, "int").Max(1).Passed() {
		t.Error("failed")
	}
	if vg.Validate(i8, "int8").Max(3).Passed() {
		t.Error("failed")
	}
	if !vg.Validate(ui, "uint").Max(300).Passed() {
		t.Error("failed")
	}
}

func TestRange(t *testing.T) {
	vg := ValidationGroup{}
	var i int = 1
	var i8 int8 = 8
	var ui uint = 30

	if !vg.Validate(i, "int").Range(0, 3).Passed() {
		t.Error("failed")
	}
	if vg.Validate(i8, "int8").Range(1, 3).Passed() {
		t.Error("failed")
	}
	if vg.Validate(ui, "uint").Range(1, 3).Passed() {
		t.Error("failed")
	}
}

func TestMinLength(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("", "emptyString").MinLength(1).Passed() {
		t.Error("failed")
	}
	if !vg.Validate("str", "string").MinLength(1).Passed() {
		t.Error("failed")
	}
	if vg.Validate([]string{}, "stringSlice").MinLength(1).Passed() {
		t.Error("failed")
	}
	if vg.Validate([]interface{}{"ok"}, "interfaceSlice").MinLength(2).Passed() {
		t.Error("failed")
	}
}

func TestMaxLength(t *testing.T) {
	vg := ValidationGroup{}
	if !vg.Validate("", "emptyString").MaxLength(1).Passed() {
		t.Error("failed")
	}
	if vg.Validate("str", "string").MaxLength(2).Passed() {
		t.Error("failed")
	}
	if !vg.Validate([]string{}, "stringSlice").MaxLength(1).Passed() {
		t.Error("failed")
	}
	if vg.Validate([]interface{}{"ok1", "ok2", "ok3"}, "interfaceSlice").MaxLength(2).Passed() {
		t.Error("failed")
	}
}

func TestLength(t *testing.T) {
	vg := ValidationGroup{}
	if !vg.Validate("", "emptyString").Length(0).Passed() {
		t.Error("failed")
	}
	if vg.Validate("str", "string").Length(2).Passed() {
		t.Error("failed")
	}
	if vg.Validate([]string{}, "stringSlice").Length(1).Passed() {
		t.Error("failed")
	}
	if !vg.Validate([]interface{}{"ok1", "ok2", "ok3"}, "interfaceSlice").Length(3).Passed() {
		t.Error("failed")
	}
}

func TestAlpha(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("a,1-@ $", "alpha").Alpha().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("abCD", "alpha").Alpha().Passed() {
		t.Error("failed")
	}
}

func TestNumeric(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("a,1-@ $", "numeric").Numeric().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("1234", "numeric").Numeric().Passed() {
		t.Error("failed")
	}
}

func TestAlphaNumeric(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("a,1-@ $", "alphaNumeric").AlphaNumeric().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("1234aB", "alphaNumeric").AlphaNumeric().Passed() {
		t.Error("failed")
	}
}

// func TestMatch(t *testing.T) {
// 	vg := ValidationGroup{}
// 	if vg.Validate("a,1-@ $", "alphaNumeric").Match().Passed() {
// 		t.Error("failed")
// 	}
// 	if !vg.Validate("1234aB", "alphaNumeric").Match().Passed() {
// 		t.Error("failed")
// 	}

// 	if valid.Match("suchuangji@gmail", regexp.MustCompile("^\\w+@\\w+\\.\\w+$"), "match").Ok {
// 		t.Error("\"suchuangji@gmail\" match \"^\\w+@\\w+\\.\\w+$\"  should be false")
// 	}
// 	if !valid.Match("suchuangji@gmail.com", regexp.MustCompile("^\\w+@\\w+\\.\\w+$"), "match").Ok {
// 		t.Error("\"suchuangji@gmail\" match \"^\\w+@\\w+\\.\\w+$\"  should be true")
// 	}
// }

// func TestNoMatch(t *testing.T) {
// 	valid := Validation{}

// 	if valid.NoMatch("123@gmail", regexp.MustCompile("[^\\w\\d]"), "nomatch").Ok {
// 		t.Error("\"123@gmail\" not match \"[^\\w\\d]\"  should be false")
// 	}
// 	if !valid.NoMatch("123gmail", regexp.MustCompile("[^\\w\\d]"), "match").Ok {
// 		t.Error("\"123@gmail\" not match \"[^\\w\\d@]\"  should be true")
// 	}
// }

func TestAlphaDash(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("a,1-@ $", "alphaDash").AlphaDash().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("1234aB-_", "alphaDash").AlphaDash().Passed() {
		t.Error("failed")
	}
}

func TestEmail(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("not@a email", "email").Email().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("suchuangji@gmail.com", "email").Email().Passed() {
		t.Error("failed")
	}
}

func TestIP(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("11.255.255.256", "IP").IP().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("01.11.11.11", "IP").IP().Passed() {
		t.Error("failed")
	}
}

func TestBase64(t *testing.T) {
	vg := ValidationGroup{}
	if vg.Validate("suchuangji@gmail.com", "base64").Base64().Passed() {
		t.Error("failed")
	}
	if !vg.Validate("c3VjaHVhbmdqaUBnbWFpbC5jb20=", "base64").Base64().Passed() {
		t.Error("failed")
	}
}

// func TestMobile(t *testing.T) {
// 	vg := ValidationGroup{}
// 	if vg.Validate("19800008888", "mobile").Mobile().Passed() {
// 		t.Error("failed")
// 	}
// 	if !vg.Validate("c3VjaHVhbmdqaUBnbWFpbC5jb20=", "base64").Mobile().Passed() {
// 		t.Error("failed")
// 	}

// 	if valid.Mobile("19800008888", "mobile").Ok {
// 		t.Error("\"19800008888\" is a valid mobile phone number should be false")
// 	}
// 	if !valid.Mobile("18800008888", "mobile").Ok {
// 		t.Error("\"18800008888\" is a valid mobile phone number should be true")
// 	}
// 	if !valid.Mobile("18000008888", "mobile").Ok {
// 		t.Error("\"18000008888\" is a valid mobile phone number should be true")
// 	}
// 	if !valid.Mobile("8618300008888", "mobile").Ok {
// 		t.Error("\"8618300008888\" is a valid mobile phone number should be true")
// 	}
// 	if !valid.Mobile("+8614700008888", "mobile").Ok {
// 		t.Error("\"+8614700008888\" is a valid mobile phone number should be true")
// 	}
// }

// func TestTel(t *testing.T) {
// 	valid := Validation{}

// 	if valid.Tel("222-00008888", "telephone").Ok {
// 		t.Error("\"222-00008888\" is a valid telephone number should be false")
// 	}
// 	if !valid.Tel("022-70008888", "telephone").Ok {
// 		t.Error("\"022-70008888\" is a valid telephone number should be true")
// 	}
// 	if !valid.Tel("02270008888", "telephone").Ok {
// 		t.Error("\"02270008888\" is a valid telephone number should be true")
// 	}
// 	if !valid.Tel("70008888", "telephone").Ok {
// 		t.Error("\"70008888\" is a valid telephone number should be true")
// 	}
// }

// func TestPhone(t *testing.T) {
// 	valid := Validation{}

// 	if valid.Phone("222-00008888", "phone").Ok {
// 		t.Error("\"222-00008888\" is a valid phone number should be false")
// 	}
// 	if !valid.Mobile("+8614700008888", "phone").Ok {
// 		t.Error("\"+8614700008888\" is a valid phone number should be true")
// 	}
// 	if !valid.Tel("02270008888", "phone").Ok {
// 		t.Error("\"02270008888\" is a valid phone number should be true")
// 	}
// }

// func TestZipCode(t *testing.T) {
// 	valid := Validation{}

// 	if valid.ZipCode("", "zipcode").Ok {
// 		t.Error("\"00008888\" is a valid zipcode should be false")
// 	}
// 	if !valid.ZipCode("536000", "zipcode").Ok {
// 		t.Error("\"536000\" is a valid zipcode should be true")
// 	}
// }
