package stackoverflow

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

// 接口
type StackOverflowRepository interface {
	DoQuery(queryStr string) (*string, error)
}

type StackOverflowNeo4jRepository struct {
	Driver neo4j.Driver
}

// 对类 StackOverflowNeo4jRepository 实现了接口 StackOverflowRepository
func (s *StackOverflowNeo4jRepository) DoQuery(queryStr string) (queryResult *string, err error) {
	res := "hello"
	return &res, nil
}
