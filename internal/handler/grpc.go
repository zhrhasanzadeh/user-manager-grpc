package handler

import (
	"context"
	"grpcprj/internal/model"
	"grpcprj/internal/pb"
	"time"
)

type GrpcHandler struct {
	pb.UserServiceServer
	service model.Service
}

func NewGRPCHandler(service model.Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (g GrpcHandler) Users(ctx context.Context, temp *pb.Temp) (*pb.GetUsers, error) {
	users := g.service.GetUsers()
	return &pb.GetUsers{Users: users}, nil
}

func (g GrpcHandler) SignUp(ctx context.Context, data *pb.User) (*pb.Temp, error) {
	convertedDate, err := time.Parse("2006-01-02", data.BirthDate)
	err = g.service.SignUp(model.User{
		Username:  data.UserName,
		Email:     data.Email,
		MobileNo:  data.MobileNo,
		Password:  data.Password,
		Birthdate: convertedDate,
	})
	return &pb.Temp{}, err
}

func (g GrpcHandler) Login(ctx context.Context, data *pb.LoginData) (*pb.LoginStatus, error) {
	login, err := g.service.Login(data.UserName, data.Password)
	if login {
		return &pb.LoginStatus{Message: "user logged in successfully"}, nil
	}
	return &pb.LoginStatus{Message: "user not found"}, err
}

func (g GrpcHandler) Delete(ctx context.Context, data *pb.DeleteData) (*pb.DeleteStatus, error) {
	del, err := g.service.Delete(data.Username)
	if del {
		return &pb.DeleteStatus{Message: "user deleted successfully"}, nil
	}
	return &pb.DeleteStatus{Message: "user not found"}, err
}
