package stackoverflow

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// 接口
type StackOverflowRepository interface {
	DoQueryAll() (*string, error)
	DoQuery(queryStr string) (*string, error)
}

type StackOverflowNeo4jRepository struct {
	Driver neo4j.Driver
}

// 对类 StackOverflowNeo4jRepository 实现了接口 StackOverflowRepository
func (s *StackOverflowNeo4jRepository) DoQueryAll() (queryResult *string, err error) {
	res := "DoQueryAll"
	return &res, nil
}

func (s *StackOverflowNeo4jRepository) DoQuery(queryStr string) (queryResult *string, err error) {
	res := "DoQuery"
	return &res, nil
}
