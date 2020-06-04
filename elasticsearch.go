package goelasticsearch

import "github.com/olivere/elastic/v7"

// NewESClient 创建新的 elasticsearch 连接客户端
// discovery.type=single-node 時設置 elastic.SetSniff(false)
// 其他选项参考：https://pkg.go.dev/github.com/olivere/elastic?tab=doc#example-NewClient-ManyOptions
func NewESClient(urls []string, sniff bool, options ...elastic.ClientOptionFunc) (*elastic.Client, error) {
	options = append(options, elastic.SetURL(urls...), elastic.SetSniff(sniff))
	return elastic.NewClient(options...)
}
