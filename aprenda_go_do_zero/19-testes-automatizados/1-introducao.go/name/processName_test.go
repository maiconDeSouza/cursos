package name

import "testing"

type TestFullName struct {
	name        string
	processName string
}

var names = [...]TestFullName{
	{"Dante parrudo kiko", "Dante Parrudo Kiko"},
	{"ana maria silva", "Ana Maria Silva"},
	{"BRUNO ALVES", "Bruno Alves"},
	{"carlos da silva", "Carlos Da Silva"},
	{"dEiSe sOuZa", "Deise Souza"},
	{"EVALDO", "Evaldo"},
	{"felipe joão dávila", "Felipe João Dávila"},
	{"gabriela santos-dumont", "Gabriela Santos-Dumont"},
	{"HÉLIO", "Hélio"},
	{"iago", "Iago"},
}

func TestProcessName(t *testing.T) {
	for _, n := range names {
		r := ProcessName(n.name)

		if n.processName != r {
			t.Errorf("Teste não passou! foi recebido %s e era esperado %s\n", r, n.processName)
		}
	}
}
