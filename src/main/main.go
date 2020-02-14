package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {

	//elas
	client, e := elastic.NewClient(elastic.SetSniff(false))

	if e != nil {
		panic(e)
	}

	response, e := client.Index().Index("dating_profile").Type("zhenai").Id("fadffa").BodyJson(struct {
		name string
	}{name: "ceshi"}).Do(context.Background())

	if e != nil {
		panic(e)
	}

	fmt.Printf(response.Id);
	//log.Printf(client.ElasticsearchVersion("http://localhost:9200"))

}
