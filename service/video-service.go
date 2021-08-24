package service

import (
	"github.com/shylabo/golang-gin-poc/entity"
	"github.com/shylabo/golang-gin-poc/repository"
)

type VideoService interface {
	Save(entity.Video) error
	Update(video entity.Video) error
	Delete(video entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	repository repository.VideoRepository
}

func New(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (service *videoService) Save(video entity.Video) error {
	service.repository.Save(video)
	return nil
}
func (service *videoService) Update(video entity.Video) error {
	service.repository.Update(video)
	return nil
}
func (service *videoService) Delete(video entity.Video) error {
	service.repository.Delete(video)
	return nil
}

func (service *videoService) FindAll() []entity.Video {
	return service.repository.FindAll()
}
