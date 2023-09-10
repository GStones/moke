package service

import (
	"context"
	pb "moke/proto/gen/buddy/api"
)

func (s *Service) AddBuddy(ctx context.Context, request *pb.AddBuddyRequest) (*pb.AddBuddyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveBuddy(ctx context.Context, request *pb.RemoveBuddyRequest) (*pb.Nothing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetBuddies(ctx context.Context, nothing *pb.Nothing) (*pb.Buddies, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ReplyAddBuddy(ctx context.Context, request *pb.ReplyAddBuddyRequest) (*pb.ReplyAddBuddyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) WatchBuddies(nothing *pb.Nothing, server pb.BuddyService_WatchBuddiesServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Remark(ctx context.Context, request *pb.RemarkRequest) (*pb.Nothing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetBlockedProfiles(ctx context.Context, nothing *pb.Nothing) (*pb.ProfileIds, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AddBlockedProfiles(ctx context.Context, ids *pb.ProfileIds) (*pb.Nothing, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RemoveBlockedProfiles(ctx context.Context, ids *pb.ProfileIds) (*pb.Nothing, error) {
	//TODO implement me
	panic("implement me")
}
