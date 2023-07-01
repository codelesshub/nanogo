package redis

import (
	"github.com/codelesshub/nanogo/config/env"
	logger "github.com/codelesshub/nanogo/config/log"

	"sync"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type RedisQueue struct {
	Enqueuer   *work.Enqueuer
	WorkerPool *work.WorkerPool
	redisAddr  string
	namespace  string
	redisPass  string
	mu         sync.Mutex // Mutex para sincronização de threads
	connected  bool       // Marcador de conexão
}

// Variável global para a instância de RedisQueue
var RedisQ *RedisQueue

func StartRedisQueue() {
	redisAddr := env.GetEnv("REDIS_ADDR")
	redisNamespace := env.GetEnv("REDIS_NAMESPACE")
	redisPass := env.GetEnv("REDIS_PASSWORD", "")

	RedisQ = &RedisQueue{
		redisAddr: redisAddr,
		namespace: redisNamespace,
		redisPass: redisPass,
	}

	RedisQ.connect()
}

func (rq *RedisQueue) connect() {
	if rq.connected {
		return
	}

	rq.mu.Lock()
	defer rq.mu.Unlock()

	if rq.connected {
		return
	}

	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				rq.redisAddr,
				redis.DialPassword(rq.redisPass),
			)
		},
	}

	rq.Enqueuer = work.NewEnqueuer(rq.namespace, redisPool)
	rq.WorkerPool = work.NewWorkerPool(struct{}{}, 10, rq.namespace, redisPool)

	rq.connected = true
}

func Enqueue(queueName string, params map[string]interface{}) bool {

	job, err := RedisQ.Enqueuer.Enqueue(queueName, work.Q(params))

	if err != nil {
		logger.Fatal("Erro ao enfileirar tarefa:", err)
		return false
	}

	logger.Debug("Tarefa enfileirada com sucesso. ID do job:", job.ID)

	return true
}
