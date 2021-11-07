package supervisor

/*
import (
	"fmt"
	"github.com/immortal/immortal"
	"log"
	"os"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
}

// Supervisor for the process
type Supervisor struct {
	sync.RWMutex
	cfg            *immortal.Config
	count          int
	lock, lockOnce uint32
	quit, run      chan struct{}
	sTime          time.Time
	wg             sync.WaitGroup

	process immortal.Process
	pid     int
	wait    time.Duration
}

// Supervise keep daemon process up and running
func Supervise(cfg *Config) error {
	config := &immortal.Config{}

	supervisor := &Supervisor{
		process: immortal.NewProcess(config),
	}

	// start a new process
	err := supervisor.RunProcess()
	if err != nil {
		return err
	}

	return supervisor.Start()
}

func (s *Supervisor) RunProcess() error {
	s.Lock()
	defer s.Unlock()

	var err error

	// return if process is running
	if atomic.SwapUint32(&s.lock, uint32(1)) != 0 {
		return fmt.Errorf("cannot start, process still running or waiting to be started")
	}

	// increment count by 1
	s.count++

	time.Sleep(time.Duration(s.cfg.Wait) * time.Second)
	if _, err = s.process.Start(); err != nil {
		atomic.StoreUint32(&s.lock, s.lockOnce)
		return err
	}

	return nil
}

// Start loop forever
func (s *Supervisor) Start() error {
	for {
		select {
		case <-s.quit:
			s.wg.Wait()
			return fmt.Errorf("supervisor stopped, count: %d", s.count)
		case <-s.run:
			s.ReStart()
		case err := <-s.process.errch:
			// get exit code
			// TODO check EXIT from kqueue since we don't know the exit code there
			exitcode := 0
			if exitError, ok := err.(*exec.ExitError); ok {
				exitcode = exitError.ExitCode()
			}
			log.Printf("PID: %d exit code: %d", s.pid, exitcode)
			// Check for post_exit command
			if len(s.cfg.PostExit) > 0 {
				var shell = "sh"
				if sh := os.Getenv("SHELL"); sh != "" {
					shell = sh
				}
				if err := exec.Command(shell, "-c", fmt.Sprintf("%s %d", s.cfg.PostExit, exitcode)).Run(); err != nil {
					log.Printf("post exit command failed: %s", err)
				}
			}
			// stop or exit based on the retries
			if s.Terminate(err) {
				if s.cfg.cli || os.Getenv("IMMORTAL_EXIT") != "" {
					close(s.quit)
				} else {
					// stop don't exit
					atomic.StoreUint32(&s.lock, 1)
				}
			} else {
				// follow the new pid instead of trying to call run again unless the new pid dies
				if s.cfg.Pid.Follow != "" {
				} else {
					s.ReStart()
				}
			}
		}
	}
}

// ReStart create a new process
func (s *Supervisor) ReStart() {
	var err error
	time.Sleep(s.wait)
	if s.lock == 0 {
		np := immortal.NewProcess(s.cfg)
		if s.process, err = s.Run(); err != nil {
			close(np.quit)
			log.Print(err)
			// loop again but wait 1 seccond before trying
			s.wait = time.Second
			s.run <- struct{}{}
		}
	}
}

// Terminate handle process termination
func (s *Supervisor) Terminate(err error) bool {
	s.Lock()
	defer s.Unlock()

	// set end time
	s.process.eTime = time.Now()
	// unlock, or lock once
	atomic.StoreUint32(&s.lock, s.lockOnce)
	// WatchPid returns EXIT
	if err != nil && err.Error() == "EXIT" {
		log.Printf("PID: %d (%s) exited", s.pid, s.process.cmd.Path)
	} else {
		log.Printf("PID %d (%s) terminated, %s [%v user  %v sys  %s up]\n",
			s.process.Pid(),
			s.process.cmd.Path,
			s.process.cmd.ProcessState,
			s.process.cmd.ProcessState.UserTime(),
			s.process.cmd.ProcessState.SystemTime(),
			time.Since(s.process.sTime),
		)
		// calculate time for next reboot (avoids high CPU usage)
		uptime := s.process.eTime.Sub(s.process.sTime)
		s.wait = 0 * time.Second
		if uptime < time.Second {
			s.wait = time.Second - uptime
		}
	}
	// behavior based on the retries
	if s.cfg.Retries >= 0 {
		//  0 run only once (don't retry)
		if s.cfg.Retries == 0 {
			return true
		}
		// +1 run N times
		if s.count > s.cfg.Retries {
			return true
		}
	}
	// -1 run forever
	return false
}
*/