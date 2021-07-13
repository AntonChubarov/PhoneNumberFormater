package infrastructure

type ValidRAMstorage struct {
	validNumbers []string
}

func (v *ValidRAMstorage) AddValidNumber(validNumber ...string) {
	for i := range validNumber {
		if v.contains(validNumber[i]) {
			continue
		} else {
			v.validNumbers = append(v.validNumbers, validNumber[i])
		}
	}
}

func (v *ValidRAMstorage) contains(element string) bool {
	for _, a := range v.validNumbers {
		if a == element {
			return true
		}
	}
	return false
}

func (v *ValidRAMstorage) GetAllNumbers() []string {
	return v.validNumbers
}

func NewValidRAMstorage() *ValidRAMstorage {
	v := &ValidRAMstorage{}
	v.validNumbers = []string{}
	return v
}