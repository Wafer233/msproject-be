package handler

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	domainSvc "github.com/Wafer233/msproject-be/project-service/internal/domain/service"
	pb "github.com/Wafer233/msproject-be/project-service/proto/project"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"time"
)

type ProjectGRPCHandler struct {
	pb.UnimplementedProjectServiceServer
	indexSvc   service.IndexService
	projectSvc service.ProjectService
}

func (handler *ProjectGRPCHandler) Index(ctx context.Context, req *pb.IndexRequest) (*pb.IndexResponse, error) {
	ctx = context.Background()

	domainMenus, err := handler.indexSvc.Index(ctx)
	if err != nil {
		zap.L().Warn("调用项目菜单失败")
		return nil, errors.New("调用Index服务失败")
	}

	domainChild := model.CovertChild(domainMenus)

	var pbMenu []*pb.Menu
	er := copier.Copy(&pbMenu, domainChild)
	if er != nil {
		zap.L().Warn("复制菜单失败")
		return nil, errors.New("复制菜单失败")
	}

	pbResp := &pb.IndexResponse{
		Menus: pbMenu,
	}

	zap.L().Info("调用菜单成功")
	return pbResp, nil
}

func (handler *ProjectGRPCHandler) SelfProject(ctx context.Context, pbReq *pb.SelfProjectRequest) (*pb.SelfProjectResponse, error) {
	ctx = context.Background()

	var domainReq model.SelfProjectRequest

	err := copier.Copy(&domainReq, pbReq)
	if err != nil {
		zap.L().Warn("复制project请求失败")
		return nil, errors.New("复制参数失败")
	}

	// PaM -> projectAndMenu
	domainProjAndMems, total, er := handler.projectSvc.SelfProject(ctx, &domainReq)
	if er != nil {
		zap.L().Warn("调用自己项目服务失败")
		return nil, errors.New("调用项目服务失败")
	}
	var pbProjects []*pb.Project
	err = copier.Copy(&pbProjects, domainProjAndMems)
	if err != nil {
		zap.L().Warn("复制project响应失败")
		return nil, errors.New("复制响应失败")
	}

	for _, pbProject := range pbProjects {
		pbProject.Code, _ = domainSvc.EncryptInt64(pbProject.ProjectCode, model.AESKey)
		maps := model.ToMap(domainProjAndMems)[pbProject.Id]
		pbProject.AccessControlType = maps.GetAccessControlType()
		pbProject.OrganizationCode, _ = domainSvc.EncryptInt64(maps.OrganizationCode, model.AESKey)
		pbProject.JoinTime = time.UnixMilli(maps.JoinTime).Format("2006-01-02 15:04:05")
		pbProject.OwnerName = pbReq.MemberName
		pbProject.Order = int32(maps.Sort)
		pbProject.CreateTime = time.UnixMilli(maps.CreateTime).Format("2006-01-02 15:04:05")
	}

	zap.L().Info("调用自己项目成功")
	return &pb.SelfProjectResponse{
		Projects: pbProjects,
		Total:    total,
	}, nil
}

func NewProjectGRPCHandler(
	indexSvc service.IndexService,
	projectSvc service.ProjectService,
) *ProjectGRPCHandler {
	return &ProjectGRPCHandler{
		indexSvc:   indexSvc,
		projectSvc: projectSvc,
	}
}
