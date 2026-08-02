package pagamento

import "testing"

func TestPagamentos(t *testing.T) {
	c, b := NewPagamento()

	t.Run("Cartão", func(t *testing.T) {
		v := Taxa(c, 300)

		if v != 285.00 {
			t.Fatalf("O valor veio %.2f, mas erá esperado %.2f", v, 285.00)
		}
	})

	t.Run("Boleto", func(t *testing.T) {
		v := Taxa(b, 300)

		if v != 297.00 {
			t.Fatalf("O valor veio %.2f, mas erá esperado %.2f", v, 297.00)
		}
	})
}
