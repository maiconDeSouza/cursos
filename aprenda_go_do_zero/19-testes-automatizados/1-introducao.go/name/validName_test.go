package name

import "testing"

type TestValidNameStruct struct {
	name      string
	validName bool
}

var listValidNames = [...]TestValidNameStruct{
	{"Dante", true},
	{"Cat", false},
	{"Ana Maria", true},
	{"Jo", false},
	{"   ", false},
	{"Hélio", true},
	{"", false},
	{"Lucas", true},
	{"  a  ", false},
	{"A", false},
	{"Ana", false},
	{"Felipe", true},
}

func TestValidName(t *testing.T) {
	for _, n := range listValidNames {
		r := ValidName(n.name)

		if n.validName != r {
			t.Errorf("Teste não passou! foi recebido %v e era esperado %v\n", r, n.validName)
		}
	}
}
