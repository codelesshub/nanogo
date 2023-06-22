package redis

import (
	logger "github.com/codelesshub/nanogo/config/log"

	"os"
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
	redisAddr := getRedisAddress()
	redisNamespace := getRedisNamespace()
	redisPass := getRedisPassword()

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

	job, err := RedisQ.Enqueuer.Enqueue(queueName, params)

	if err != nil {
		logger.Fatal("Erro ao enfileirar tarefa:", err)
		return false
	}

	logger.Debug("Tarefa enfileirada com sucesso. ID do job:", job.ID)

	return true
}

func getRedisAddress() string {
	redisAddr := os.Getenv("REDIS_ADDR")

	if redisAddr == "" {
		logger.Fatal("O endereço do Redis não foi definido no arquivo .env")
	}

	return redisAddr
}

func getRedisNamespace() string {
	redisNamespace := os.Getenv("REDIS_NAMESPACE")
	if redisNamespace == "" {
		logger.Fatal("O namespace do Redis não foi definido no arquivo .env")
	}

	return redisNamespace
}

func getRedisPassword() string {
	redisPass := os.Getenv("REDIS_PASSWORD")
	if redisPass == "" {
		logger.Fatal("A senha do Redis não foi definida no arquivo .env")
	}

	return redisPass
}
