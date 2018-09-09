package server

import (
	"errors"
	"google.golang.org/grpc"
	"reflect"
	"testing"

	pb "github.com/solcates/postmind/apis"
	"golang.org/x/net/context"
)

type mockremindersGetremindersserver struct {
	grpc.ServerStream
	error bool
}

func (rgc *mockremindersGetremindersserver) Send(r *pb.Reminder) error {
	if rgc.error {
		return errors.New("Some Error")
	}
	return nil
}

func TestServer_GetReminders(t *testing.T) {

	s := &Server{
		reminders: map[int64]*pb.Reminder{
			1: {
				Id:        1,
				Text:      "Straighten Back",
				Action:    "Kick Self",
				Cleared:   false,
				Frequency: "@hourly",
			},
		},
	}
	type args struct {
		in0 *pb.GetRemindersRequest
		in1 pb.Reminders_GetRemindersServer
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			s:    s,
			args: args{
				in0: &pb.GetRemindersRequest{},
				in1: &mockremindersGetremindersserver{},
			},
			wantErr: false,
		}, {
			name: "ok",
			s:    s,
			args: args{
				in0: &pb.GetRemindersRequest{},
				in1: &mockremindersGetremindersserver{
					error: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.GetReminders(tt.args.in0, tt.args.in1); (err != nil) != tt.wantErr {
				t.Errorf("Server.GetReminders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_ClearReminder(t *testing.T) {
	s := &Server{
		reminders: map[int64]*pb.Reminder{
			1: {
				Id:        1,
				Text:      "Straighten Back",
				Action:    "Kick Self",
				Cleared:   false,
				Frequency: "@hourly",
			},
		},
	}
	type args struct {
		in0 context.Context
		in1 *pb.Reminder
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.Reminder
		wantErr bool
	}{
		{
			name: "Ok",
			s:    s,
			args: args{
				in0: context.Background(),
				in1: &pb.Reminder{
					Id:        1,
					Text:      "Straighten Back",
					Action:    "Kick Self",
					Cleared:   false,
					Frequency: "@hourly",
				},
			},
			want: &pb.Reminder{
				Id:        1,
				Text:      "Straighten Back",
				Action:    "Kick Self",
				Cleared:   true,
				Frequency: "@hourly",
			},
			wantErr: false,
		}, {
			name: "Error",
			s:    s,
			args: args{
				in0: context.Background(),
				in1: &pb.Reminder{
					Id:        2,
					Text:      "Straighten Back",
					Action:    "Kick Self",
					Cleared:   false,
					Frequency: "@hourly",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ClearReminder(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ClearReminder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ClearReminder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_NewReminder(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *pb.Reminder
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *pb.Reminder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.NewReminder(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.NewReminder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.NewReminder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Run(t *testing.T) {
	tests := []struct {
		name    string
		s       *Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Server.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewServer(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "Default",
			args: args{},
			want: &Server{
				port: 8080,
				host: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
