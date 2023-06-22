package console

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/schedule"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	events := []schedule.Event{
		//facades.Schedule().Command("send:emails name").Daily(),
	}

	//event := facades.Schedule().Call(func() {
	//	//facades.Orm.Query().Where("1 = 1").Delete(&models.User{})
	//	facades.Log().Info("执行任务调度测试...") // 每分钟执行
	//}).EveryMinute().Cron(`*/1 * * * *`).DelayIfStillRunning()
	//events = append(events, event)

	return events
}

func (kernel *Kernel) Commands() []console.Command {
	commands := []console.Command{
		//&commands.Migrate{},
		//&commands.SendEmails{},
	}

	//commands = append(commands, &commands.SendEmails{})

	return commands
}
