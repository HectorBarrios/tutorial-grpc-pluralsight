package main

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/HectorBarrios/grpcdemo-pluralsight-server/pb"
	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var employees = []pb.Employee{
	pb.Employee{
		Id:                  1,
		BadgeNumber:         2080,
		FirstName:           "Hector",
		LastName:            "Barrios",
		VacationAccrualRate: 2,
		VacationAccrued:     30,
	},
	pb.Employee{
		Id:                  2,
		BadgeNumber:         5835,
		FirstName:           "Amity",
		LastName:            "Fuller",
		VacationAccrualRate: 2.3,
		VacationAccrued:     32.5,
	},
	pb.Employee{
		Id:                  3,
		BadgeNumber:         7342,
		FirstName:           "Shit",
		LastName:            "Happens",
		VacationAccrualRate: 23,
		VacationAccrued:     23,
	},
	pb.Employee{
		Id:                  4,
		BadgeNumber:         3827,
		FirstName:           "Hppy",
		LastName:            "Ending",
		VacationAccrualRate: 22.3,
		VacationAccrued:     23.23,
	},
}

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &employeeService{})
	log.Println("Starting server on port " + port)
	s.Serve(lis)
}

type employeeService struct{}

func (s *employeeService) GetByBadgeNumber(ctx context.Context,
	req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		fmt.Printf("metadata received: %v\n", md)
	}

	for _, e := range employees {
		if req.BadgeNumber == e.BadgeNumber {
			return &pb.EmployeeResponse{Employee: &e}, nil
		}
	}

	return nil, errors.New("Employee not found")
}

func (s *employeeService) GetAll(req *pb.GetAllRequest,
	stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}

	return nil
}

func (s *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	return nil
}

func (s *employeeService) Save(ctx context.Context,
	req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	return nil
}
