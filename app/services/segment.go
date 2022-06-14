package services

import (
	"challange/app/repository"
	"github.com/valyala/fastjson"
	"log"
	"time"
)

type SegmentService struct {
	userRepository *repository.UserRepository
}

func NewSegmentService(userRepository repository.UserRepository) SegmentService {
	return SegmentService{userRepository: &userRepository}
}

func (ss SegmentService) CreateUser(jsonStr []byte) error {
	var p fastjson.Parser
	v, err := p.Parse(string(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	id := string(v.GetStringBytes("ID"))
	segment := string(v.GetStringBytes("segment"))
	expire := time.Unix(v.GetInt64("expiredSegment"), 0)
	return ss.userRepository.Save(id, segment, expire)
}
