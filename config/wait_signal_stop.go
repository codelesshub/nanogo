package config

import (
	"os"
	"os/signal"
)

func WaitSignalStop() {
	// Aguarde um sinal de encerramento
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
