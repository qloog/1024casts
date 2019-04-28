package repository

import (
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/constvar"
)

type TopicRepo struct {
	db *model.Database
}

func NewTopicRepo() *TopicRepo {
	return &TopicRepo{
		db: model.DB,
	}
}

func (repo *TopicRepo) CreateTopic(Topic model.TopicModel) (id uint64, err error) {
	err = repo.db.Self.Create(&Topic).Error
	if err != nil {
		return 0, err
	}

	return Topic.Id, nil
}

func (repo *TopicRepo) GetTopicById(id int) (*model.TopicModel, error) {
	Topic := model.TopicModel{}
	result := repo.db.Self.Where("id = ?", id).First(&Topic)

	return &Topic, result.Error
}

func (repo *TopicRepo) GetTopicList(TopicMap map[string]interface{}, offset, limit int) ([]*model.TopicModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	Topics := make([]*model.TopicModel, 0)
	var count uint64

	if err := repo.db.Self.Model(&model.TopicModel{}).Where(TopicMap).Count(&count).Error; err != nil {
		return Topics, count, err
	}

	if err := repo.db.Self.Where(TopicMap).Offset(offset).Limit(limit).Order("id desc").Find(&Topics).Error; err != nil {
		return Topics, count, err
	}

	return Topics, count, nil
}

func (repo *TopicRepo) UpdateTopic(userMap map[string]interface{}, id int) error {

	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(Topic).Updates(userMap).Error
}

func (repo *TopicRepo) DeleteTopic(id int) error {
	Topic, err := repo.GetTopicById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&Topic).Error
}

func (repo *TopicRepo) Store(Topic *model.TopicModel) (id uint64, err error) {
	//users := model.TopicModel{}

	return 0, nil
}
