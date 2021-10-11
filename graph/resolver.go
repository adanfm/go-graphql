package graph

import "github.com/adanfm/go-graphql/graph/model"

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
