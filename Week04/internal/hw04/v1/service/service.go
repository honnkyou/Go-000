/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package service

import (
	"context"

	pb "Week04/api/hw04/v1"

	"Week04/internal/hw04/v1/biz"
)

// server is used to implement helloworld.GreeterServer.
type service struct {
	u *biz.HWUseCase
	pb.UnimplementedHomeworkServer
}

func NewService(u *biz.HWUseCase) *service {
	return &service{u: u}
}

func (s *service) TakeHomework(ctx context.Context, r *pb.HWRequest) (*pb.HWReply, error) {
	//DTO->DO
	item := &biz.Homework{detail: r.Message}
	// call biz
	s.u.TakeHomework(item)

	//response
	return &pb.HWReply{Message: "Homework was taked"}, nil
}
