package server

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	pb "github.com/solcates/postmind/apis"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server will serve all reminders
type Server struct {
	port      int
	host      string
	reminders map[int64]*pb.Reminder
}

//NewServer returns a fresh Server instance for the address of host:port
func NewServer(host string, port int) *Server {

	// If no port provided, let's try another
	if port == 0 {
		port = 8080
	}
	s := &Server{
		host: host,
		port: port,
	}
	return s
}

//GetReminders returns a stream of Reminders that you need to be aware of.
func (s *Server) GetReminders(req *pb.GetRemindersRequest, stream pb.Reminders_GetRemindersServer) (err error) {
	for _, r := range s.reminders {
		if err = stream.Send(r); err != nil {
			return
		}
	}
	return
}

//ClearReminder clears a reminder
func (s *Server) ClearReminder(ctx context.Context, rem *pb.Reminder) (res *pb.Reminder, err error) {
	r, ok := s.reminders[rem.Id]
	if ok {
		r.Cleared = true
		res = r
	} else {
		res = nil
		err = errors.New("Unknown Reminder")
	}
	return
}

//NewReminder creates and saves an existing reminder...
func (s *Server) NewReminder(context.Context, *pb.Reminder) (*pb.Reminder, error) {
	panic("implement me")

}

//Run starts the gRPC Server
func (s *Server) Run() (err error) {

	// Load the reminders from disk
	go s.loadReminders()

	// Serve the gRPC server
	return s.serve()
}

// serve serves the gRPC Server
func (s *Server) serve() (err error) {
	var lis net.Listener
	address := fmt.Sprintf("%v:%v", s.host, s.port)
	lis, err = net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Bootstrap the Server and run it
	ss := grpc.NewServer()
	pb.RegisterRemindersServer(ss, s)
	return ss.Serve(lis)
}

func (s *Server) loadReminders() {
	for k, v := range s.reminders {
		logrus.Info(k)
		logrus.Info(v)
	}
}
