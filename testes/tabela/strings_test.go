package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Ola", "ola", -1},
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa de teste: %v", teste)
		retorno := strings.Index(teste.texto, teste.parte)
		if retorno != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, retorno)
		}
	}
}
