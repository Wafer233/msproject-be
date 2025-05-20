package service

import pb "github.com/Wafer233/msproject-be/user-service/proto/login"

type LoginService interface {
}
type DefaultLoginService struct {
	pb.UnimplementedLoginServiceServer // 保证向后兼容
}
