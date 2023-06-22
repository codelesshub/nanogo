package redis

import "github.com/gocraft/work"

type RedisQueueConsumerInterface interface {
	Run(job *work.Job) error
}
