package stackoverflow

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// 接口
type StackOverflowRepository interface {
	DoQueryAll() (QueryAllResp, error)
	DoQuery(queryStr string) (*string, error)
}

type StackOverflowNeo4jRepository struct {
	Driver neo4j.Driver
}

// 对类 StackOverflowNeo4jRepository 实现了接口 StackOverflowRepository
type QueryAllResp struct {
	Nodes         []QueryAllNode         `json:"nodes"`
	Relationships []QueryAllRelationship `json:"relationships"`
}

type QueryAllNode struct {
	Id     int      `json:"id"`
	Labels []string `json:"labels"`
}

type QueryAllRelationship struct {
	Id     int      `json:"id"`
	Labels []string `json:"labels"`
}

func (s *StackOverflowNeo4jRepository) DoQueryAll() (queryResult QueryAllResp, err error) {

	// new a session
	session := s.Driver.NewSession(neo4j.SessionConfig{})
	defer func() {
		err = session.Close()
	}()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return s.queryAll(tx)
	})
	if result == nil {
		return QueryAllResp{}, err
	}
	// res := "DoQueryAll"
	return QueryAllResp{}, nil
}

func (s *StackOverflowNeo4jRepository) queryAll(tx neo4j.Transaction) (QueryAllResp, error) {
	queryNodesStatement := "MATCH (n) return n LIMIT 10"

	resultNodes, err := tx.Run(
		queryNodesStatement,
		map[string]interface{}{},
	)
	if err != nil {
		return QueryAllResp{}, err
	}

	record := resultNodes.Record()
	println(record.Keys[len(record.Keys)-1])
	// if value, ok := record.Get("id"); ok {
	// 	node := value.(neo4j.Node)
	// 	println(node.Labels)
	// }

	println("hello")

	// queryRelationshipsStatement := "Match (n1)-[r]->(n2) return r LIMIT 10"
	// resultRelationships, err := tx.Run(
	// 	queryRelationshipsStatement,
	// 	map[string]interface{}{},
	// )
	// if err != nil {
	// 	return QueryAllResp{}, err
	// }

	return QueryAllResp{}, nil
}

func (s *StackOverflowNeo4jRepository) DoQuery(queryStr string) (queryResult *string, err error) {
	res := "DoQuery"
	return &res, nil
}
