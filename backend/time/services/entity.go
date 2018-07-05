package services

import (
	"evolution/backend/time/models"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

type Entity struct {
	Basic
}

func NewEntity(engine *xorm.Engine, cache *redis.Client) *Entity {
	ret := Entity{}
	ret.Engine = engine
	ret.Cache = cache
	return &ret
}

func (s *Entity) One(id int) (interface{}, error) {
	model := models.Entity{}
	_, err := s.Engine.Where("id = ?", id).Get(&model)
	return model, err
}

func (s *Entity) Add(model models.Entity) (err error) {
	_, err = s.Engine.Insert(&model)
	return
}

func (s *Entity) Update(id int, model models.Entity) (err error) {
	_, err = s.Engine.Id(id).Update(&model)
	return
}

func (s *Entity) Delete(id int, model models.Entity) (err error) {
	_, err = s.Engine.Id(id).Get(&model)
	if err == nil {
		_, err = s.Engine.Id(id).Delete(&model)
	}
	return
}

func (s *Entity) ListWithCondition(entity *models.Entity) (entities []models.Entity, err error) {
	entitiesJoin := make([]models.EntityJoin, 0)
	sql := s.Engine.Unscoped().Table("entity").Join("INNER", "area", "area.id = entity.area_id")

	condition := entity.BuildCondition()
	sql = sql.Where(condition)
	err = sql.Find(&entitiesJoin)
	if err != nil {
		return
	}
	entities = make([]models.Entity, 0)
	for _, one := range entitiesJoin {
		one.Entity.Area = one.Area
		entities = append(entities, one.Entity)
	}
	return
}

func (s *Entity) List() (entites []models.Entity, err error) {
	entites = make([]models.Entity, 0)
	err = s.Engine.Find(&entites)
	return
}
