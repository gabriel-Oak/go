package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() // Indica que o teste pode ser executado em paralelo com outros testes.
	if runtime.GOOS == "linux" {
		t.Skip("Não funciona no Windows.")
	}

	t.Fail()
}
