package scheduler

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/model"
)

type Service struct {
	ID             string // service id
	ProxyServiceID string
	Node           model.Address // the node that service exists
	Shadow         model.Address // node that may replace the origin node, useful in migration
	logger         *model.Logger
}

func init() {
	PseudoRegister()
}

// NewService new a storage service, keep it in map
func NewService(opts model.ServiceOpts) *Service {
	return &Service{
		ID:             opts.ID,
		ProxyServiceID: opts.ProxyServiceID,
		Node:           opts.NodeAddr,
		Shadow:         model.Address{},
		logger:         model.NewLogger(opts.ProxyServiceID),
	}
}

// PseudoRegister register services
func PseudoRegister() {
	opts1 := model.ServiceOpts{
		ID:             "service1.1",
		ProxyServiceID: "service1",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service1", opts1)

	opts2 := model.ServiceOpts{
		ID:             "service2.1",
		ProxyServiceID: "service2",
		NodeAddr: model.Address{
			IP:   "127.0.0.1",
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service2", opts2)
}

// Register register a service to DefaultScheduler
func DefaultRegister(ProxyService string, opts model.ServiceOpts) {
	Default().AddService(ProxyService, NewService(opts))
}

func (s *Service) AddShadow(addr model.Address) {
	s.Shadow.IP = addr.IP
	s.Shadow.Port = addr.Port
}

/* DataLog log data to logger of the service.
if data size exceeds the logger capacity, send log to dst node */
func (s *Service) DataLog(ser *Service, data string) error {
	s.logger.Lock()
	s.logger.Count++
	s.logger.LogQueue = append(s.logger.LogQueue, data)

	if s.logger.Count == s.logger.Capacity {
		// todo send to dst by goroutine
		cli := client.Client{
			Dest: ser.Shadow,
		}
		logWithID := model.LogWithServiceID{
			Log:            s.logger.Log,
			ProxyServiceID: ser.ProxyServiceID,
		}
		cli.SendLog(logWithID)
		s.logger.TotalSend++
		s.logger.ClearQueue()
		s.logger.Count = 0
	}
	s.logger.Unlock()
	return nil
}

// LockAndGetSentConsumed return sent and consumed
func (s *Service) LockAndGetSentConsumed() (int, int) {
	s.logger.Lock()
	sent := s.logger.TotalSend
	consumed := s.logger.TotalConsumed
	return sent, consumed
}

func (s *Service) UnlockLogger() {
	s.logger.Unlock()
}

// SendLastLog send the last log to dst
func (s *Service) SendLastLog(ProxyServiceID string, addr model.Address) error {
	logrus.Info("send the last log")
	cli := client.Client{
		Dest: addr,
	}

	s.logger.Lock()
	defer s.logger.Unlock()

	s.logger.SetLastFlag()
	logWithID := model.LogWithServiceID{
		Log:            s.logger.Log,
		ProxyServiceID: ProxyServiceID,
	}
	err := cli.SendLog(logWithID)
	if err != nil {
		return err
	}

	logrus.Infof("SetLastLog finished, ProxyServiceID: %v", ProxyServiceID)
	return nil
}

// ConsumedAdder
func (s *Service) ConsumedAdder() {
	s.logger.Lock()
	s.logger.TotalConsumed++
	s.logger.Unlock()
}
