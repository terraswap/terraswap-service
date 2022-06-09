package repeater

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/delight-labs/terraswap-service/internal/pkg/logging"
)

type Repeater interface {
	Repeat()
	Recover()
}

type Runner interface {
	Run()
}

type repeaterImpl struct {
	logging.Logger
	Runner
	name            string
	exponential     int
	update_interval time.Duration
}

var _ Repeater = &repeaterImpl{}

func Enroll(logger logging.Logger, runable Runner, name string, exponential, update_interval int) Repeater {
	r := &repeaterImpl{
		logger, runable, name, exponential, time.Duration(update_interval),
	}

	go func() {
		r.Repeat()
	}()

	return r
}

func (r *repeaterImpl) Recover() {
	recovered := recover()
	if recovered == nil {
		return
	}

	err, ok := recovered.(error)
	if !ok {
		err = fmt.Errorf("could not convert recovered error into error: %s", spew.Sdump(recovered))
	}

	stack := string(debug.Stack())
	r.WithField("err", logging.NewErrorField(err)).WithField("stack", stack).Errorf("panic caught during repeat(%s)", r.name)
	r.exponential = r.exponential * 2
	if r.exponential == 0 || r.exponential > 256 {
		r.exponential = 2
	}
	t := time.NewTimer(time.Second * time.Duration(r.exponential))
	<-t.C

}

func (r *repeaterImpl) Repeat() {
	for {
		func() {
			defer r.Recover()

			r.Run()

			t := time.NewTimer(time.Second * r.update_interval)
			<-t.C
		}()
	}
}
