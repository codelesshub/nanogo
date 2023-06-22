package redis

type RedisConsumer struct {
	Queue *RedisQueue
}

func RedisQueueConsumer(consumer RedisQueueConsumerInterface, queueName string) *RedisConsumer {
	redisQueue := RedisQ

	// Registre a função de consumo
	redisQueue.WorkerPool.Job(queueName, consumer.Run)

	// Inicia o processamento de jobs
	redisQueue.WorkerPool.Start()

	return &RedisConsumer{
		Queue: redisQueue,
	}
}
