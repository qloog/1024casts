package repository

import (
	"1024casts/backend/model"
	"1024casts/backend/pkg/constvar"
)

type CourseRepo struct {
	db *model.Database
}

func NewCourseRepo() *CourseRepo {
	return &CourseRepo{
		db: model.DB,
	}
}

func (repo *CourseRepo) CreateCourse(course model.CourseModel) (id uint64, err error) {
	err = repo.db.Self.Create(&course).Error
	if err != nil {
		return 0, err
	}

	return course.Id, nil
}

func (repo *CourseRepo) GetCourseById(id int) (*model.CourseModel, error) {
	course := model.CourseModel{}
	result := repo.db.Self.Where("id = ?", id).First(&course)

	return &course, result.Error
}

func (repo *CourseRepo) GetCourseList(courseMap map[string]interface{}, offset, limit int) ([]*model.CourseModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	courses := make([]*model.CourseModel, 0)
	var count uint64

	if err := repo.db.Self.Model(&model.CourseModel{}).Where(courseMap).Count(&count).Error; err != nil {
		return courses, count, err
	}

	if err := repo.db.Self.Where(courseMap).Offset(offset).Limit(limit).Order("id desc").Find(&courses).Error; err != nil {
		return courses, count, err
	}

	return courses, count, nil
}

func (repo *CourseRepo) UpdateCourse(userMap map[string]interface{}, id int) error {

	course, err := repo.GetCourseById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Model(course).Updates(userMap).Error
}

func (repo *CourseRepo) DeleteCourse(id int) error {
	course, err := repo.GetCourseById(id)
	if err != nil {
		return err
	}

	return repo.db.Self.Delete(&course).Error
}

func (repo *CourseRepo) Store(course *model.CourseModel) (id uint64, err error) {
	//users := model.CourseModel{}

	return 0, nil
}
