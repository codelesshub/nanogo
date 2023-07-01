// myConsumer.go
package main

import (
	"log"
)

type MyConsumer struct {
}

func (mc *MyConsumer) Consume(body map[string]interface{}, headers map[string]interface{}) {
	log.Printf("Headers: %v", headers)
	log.Printf("Body: %s", body)
}
