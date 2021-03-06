package ec2

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
)

func catchInterrupt(lg *zap.Logger, stopc chan struct{}, once *sync.Once, sigc chan os.Signal, run func() error) (err error) {
	errc := make(chan error)
	go func() {
		errc <- run()
	}()
	select {
	case <-stopc:
		lg.Info("interrupting")
		serr := <-errc
		lg.Info("interrupted", zap.Error(serr))
		err = fmt.Errorf("interrupted (run function returned %v)", serr)
	case sig := <-sigc:
		once.Do(func() { close(stopc) })
		err = fmt.Errorf("received os signal %v, closed stopc (interrupted %v)", sig, <-errc)
	case err = <-errc:
	}
	return err
}
