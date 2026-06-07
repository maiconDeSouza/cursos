package main

import (
	"errors"
	"fmt"
)

// 1. O Contrato: Quem quiser notificar, precisa saber Enviar uma mensagem.
type Notificador interface {
	Enviar(mensagem string) error
}

// 2. A Função Central: Ela aceita QUALQUER struct, desde que ele assine o contrato Notificador.
func AvisarCliente(n Notificador, mensagem string) {
	err := n.Enviar(mensagem)
	if err != nil {
		fmt.Println("Falha na notificação:", err)
		return
	}
	fmt.Println("Sistema: Cliente avisado com sucesso!")
}

type SMS struct {
	NumeroTelefone string
}

func (s SMS) Enviar(msg string) error {
	err := errors.New("Não enviada")

	if true {
		return nil
	}
	return err
}

func main() {
	sms := SMS{
		NumeroTelefone: "11999990000",
	}

	AvisarCliente(sms, "Texto")

}
