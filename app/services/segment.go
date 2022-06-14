package services

import (
	"challange/app/repository"
	"github.com/valyala/fastjson"
	"log"
	"time"
)

type SegmentService struct {
	segmentRepository *repository.SegmentRepository
}

func NewSegmentService(segmentRepository repository.SegmentRepository) SegmentService {
	return SegmentService{segmentRepository: &segmentRepository}
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
	return ss.segmentRepository.Save(id, segment, expire)
}
