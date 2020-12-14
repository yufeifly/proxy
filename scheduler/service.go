package scheduler

import (
	"github.com/sirupsen/logrus"
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/api/types/logger"
	"github.com/yufeifly/proxy/api/types/svc"
	"github.com/yufeifly/proxy/client"
	"github.com/yufeifly/proxy/cluster"
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/ticket"
)

// Service ...
type Service struct {
	ID             string // service id
	ProxyServiceID string
	Node           types.Address  // the node that service exists
	MigTarget      types.Address  // node that may replace the origin node, useful in migration
	logger         *logger.Logger // log the data while migrating, useful in migration
	ticket         ticket.Ticket  // ticket interface
}

// NewService new a storage service, keep it in map
func NewService(opts svc.ServiceOpts) *Service {
	return &Service{
		ID:             opts.ID,
		ProxyServiceID: opts.ProxyServiceID,
		Node:           opts.NodeAddr,
		MigTarget:      types.Address{},
		logger:         logger.NewLogger(opts.ProxyServiceID),
		ticket:         ticket.NewTicket(),
	}
}

// PseudoRegister register services
func PseudoRegister() {
	proxyIP := cluster.DefaultCluster().GetProxy().IP
	opts1 := svc.ServiceOpts{
		ID:             "service1.1",
		ProxyServiceID: "service1",
		NodeAddr: types.Address{
			IP:   proxyIP, // fixme localhost
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service1", opts1)

	opts2 := svc.ServiceOpts{
		ID:             "service2.1",
		ProxyServiceID: "service2",
		NodeAddr: types.Address{
			IP:   proxyIP, // fixme localhost
			Port: config.DefaultMigratorListeningPort,
		},
	}
	DefaultRegister("service2", opts2)
}

// DefaultRegister register a service to defaultScheduler
func DefaultRegister(ProxyService string, opts svc.ServiceOpts) {
	Default().AddService(ProxyService, NewService(opts))
}

// AddMigTarget ...
func (s *Service) AddMigTarget(addr types.Address) {
	s.MigTarget.IP = addr.IP
	s.MigTarget.Port = addr.Port
}

// LogDataInJSON log data to logger of the service. if data size exceeds the logger capacity, send log to dst node
func (s *Service) LogDataInJSON(data string) error {
	s.logger.Lock()
	defer s.logger.Unlock()
	s.logger.Count++
	s.logger.LogQueue = append(s.logger.LogQueue, data)

	if s.logger.Count == s.logger.Capacity {
		// todo send to dst by goroutine
		cli := client.NewClient(s.MigTarget)
		logWithID := logger.LogWithServiceID{
			Log:            s.logger.Log,
			ProxyServiceID: s.ProxyServiceID,
		}
		err := cli.SendLog(logWithID)
		if err != nil {
			logrus.Errorf("scheduler.LogDataInJSON SendLog failed, err: %v", err)
			return err
		}
		s.logger.Sent++
		s.logger.ClearQueue()
		s.logger.Count = 0
	}
	return nil
}

// LockAndGetSentConsumed return sent and consumed
func (s *Service) LockAndGetSentConsumed() (int, int) {
	s.logger.Lock()
	return s.logger.Sent, s.logger.Consumed
}

// UnlockLogger ...
func (s *Service) UnlockLogger() {
	s.logger.Unlock()
}

// SendLastLog send the last log to dst
func (s *Service) SendLastLog() error {
	logrus.Info("send the last log")
	cli := client.NewClient(s.MigTarget)
	s.logger.Lock()
	defer s.logger.Unlock()

	s.logger.SetLastFlag()
	logWithID := logger.LogWithServiceID{
		Log:            s.logger.Log,
		ProxyServiceID: s.ProxyServiceID,
	}
	err := cli.SendLog(logWithID)
	if err != nil {
		return err
	}

	logrus.Infof("SetLastLog finished, ProxyServiceID: %v", s.ProxyServiceID)
	return nil
}

// ConsumedAdder ...
func (s *Service) ConsumedAdder() {
	s.logger.Lock()
	s.logger.Consumed++
	s.logger.Unlock()
}

// Ticket get ticket interface
func (s *Service) Ticket() ticket.Ticket {
	return s.ticket
}
