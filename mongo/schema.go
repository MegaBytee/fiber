package mongo

import (
	"fmt"

	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Schema struct {
	Name        string
	IndexModels []mongo.IndexModel
	Mongo       *Mongo
	FilterID    primitive.D
}

func NewSchema(name string) *Schema {
	return &Schema{
		Name: name,
	}
}

func (s *Schema) SetMongo(m *Mongo) *Schema {
	//fmt.Println("SetMongo:m:>", m)
	s.Mongo = m
	return s
}

func (s *Schema) SetIndex(idx mongo.IndexModel) *Schema {
	s.IndexModels = append(s.IndexModels, idx)
	return s
}
func (s *Schema) SetFilterID(filter primitive.D) *Schema {
	s.FilterID = filter
	return s
}

func (s *Schema) CreateIndex() {
	for _, y := range s.IndexModels {
		err := s.Mongo.CreateIndex(s.Name, y)
		if err.Code != 0 {
			fmt.Println("Plugin-SetupSchema:err>", err.Msg)
		}
	}
}

func (s *Schema) CheckSaved() bool {
	return s.Mongo.CheckSaved(s.Name, s.FilterID)
}

func (s *Schema) Save(data any) map[string]any {
	r := s.Mongo.Save(s.Name, data)
	return map[string]any{
		"code": r.Code,
		"msg":  r.Msg,
	}
}

func (s *Schema) Update(data primitive.D) map[string]any {
	r := s.Mongo.Update(s.Name, s.FilterID, data)
	return map[string]any{
		"code": r.Code,
		"msg":  r.Msg,
	}
}

func (s *Schema) Delete() map[string]any {
	r := s.Mongo.Delete(s.Name, s.FilterID)
	return map[string]any{
		"code": r.Code,
		"msg":  r.Msg,
	}
}

func (s *Schema) Incr(key, value string) map[string]any {
	r := s.Mongo.Incr(s.Name, key, value, s.FilterID)
	return map[string]any{
		"code": r.Code,
		"msg":  r.Msg,
	}
}

func (s *Schema) Paginate(filter primitive.D, limit, page int64, x any) mongopagination.PaginationData {
	pagination, _ := s.Mongo.Paginate(s.Name, filter, limit, page, x)
	return pagination
}
