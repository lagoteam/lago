package providers

import (
	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
)

type QueueServiceProvider struct {
}

func (r QueueServiceProvider) Boot() {
}

func (r QueueServiceProvider) Register() {
	facades.Queue.Register(r.Jobs())
}

func (r QueueServiceProvider) Jobs() []queue.Job {
	return []queue.Job{
		//&jobs.DefaultJob{},
	}
}
