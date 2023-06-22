package tasks

import (
	"fmt"

	redis "github.com/codelesshub/nanogo/config/redis"
	"github.com/gocraft/work"
)

type CurlTask struct {
	// Seu código pode ter outras propriedades aqui
}

func NewCurlTask() *CurlTask {
	return &CurlTask{}
}

// Implementa a interface RedisQueueConsumer
func (ct *CurlTask) Run(job *work.Job) error {
	fmt.Println("Executando a tarefa curl...")
	// Seu código para executar a tarefa curl aqui
	return nil
}

var _ redis.RedisQueueConsumerInterface = (*CurlTask)(nil)
