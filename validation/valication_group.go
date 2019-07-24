package validation

// A Validation context manages data validation and error messages.
type ValidationGroup struct {
	Fields []*Validation
}

func (vg *ValidationGroup) Validate(obj interface{}, name string) (v *Validation) {
	v = &Validation{Target: obj, Name: name}
	vg.Fields = append(vg.Fields, v)
	return v
}

func (vg *ValidationGroup) Passed() (ok bool) {
	for _, v := range vg.Fields {
		ret := v.Passed()
		if !ret {
			return false
		}
	}

	return true
}

func (vg *ValidationGroup) GetErrors() (errMsgs []string) {
	for _, v := range vg.Fields {
		ret := v.Passed()
		if !ret {
			errs := v.GetErrors()
			errMsgs = append(errMsgs, errs...)
		}
	}

	return errMsgs
}

func (vg *ValidationGroup) Clear() {
	vg.Fields = nil
}
