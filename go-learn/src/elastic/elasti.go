package elastic

import (
	"time"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"fmt"
	"reflect"
	"encoding/json"
)
//http://olivere.github.io/elastic/
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func RunElastic() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	PingNow(client, ctx)

	CheckEsVersion(client)

	//create index
	CreateIndex(client, ctx)

	//index tweet1
	tweet1 := Tweet{User: "chenjian", Message: "Fuck you", Retweets: 100}
	IndexWithStruct(client, tweet1, ctx)

	//index tweet2
	tweet2 := `{"user" : "olivere", "message" : "It's a Raggy Waltz"}`
	IndexWithString(client, tweet2, ctx)

	//get
	GetById(client, ctx)

	//flush
	FlushNow(client, ctx)

	//query
	QueryNow(client, ctx)

	//Update an index
	UpdateIndexNow(client, ctx)

	//Delete an index
	DeleteIndexNow(client, ctx)
}

func PingNow(client *elastic.Client, ctx context.Context) {
	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func CheckEsVersion(client *elastic.Client) {
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func CreateIndex(client *elastic.Client, ctx context.Context) {
	exists, err := client.IndexExists("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}

		if !createIndex.Acknowledged {
			fmt.Println("==== NOT Acknowledged =====")
		}
	}
}

func IndexWithStruct(client *elastic.Client, tweet1 Tweet, ctx context.Context) {
	put1, err := client.Index().Index("twitter").Type("tweet").Id("1").BodyJson(tweet1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func IndexWithString(client *elastic.Client, tweet2 string, ctx context.Context) {
	put2, err := client.Index().Index("twitter").Type("tweet").Id("2").BodyString(tweet2).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
}

func GetById(client *elastic.Client, ctx context.Context) {
	get1, err := client.Get().Index("twitter").Type("tweet").Id("1").Do(ctx)
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type: %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func FlushNow(client *elastic.Client, ctx context.Context) {
	_, err := client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
}

func QueryNow(client *elastic.Client, ctx context.Context) {
	termQuery := elastic.NewTermQuery("user", "chenjian")
	searchResult, err := client.Search().Index("twitter").Query(termQuery).Sort("user", true).From(0).Size(10).Pretty(true).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("QueryNow took %d milliseconds\n", searchResult.TookInMillis)
	var ttype Tweet
	//Method1: Each is a convenience function that iterates over hits in a search result.
	for _, item := range searchResult.Each(reflect.TypeOf(ttype)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s %d\n;", t.User, t.Message, t.Retweets)
		}
	}
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	//Method2
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		for _, hit := range searchResult.Hits.Hits {
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
}

func DeleteIndexNow(client *elastic.Client, ctx context.Context) {
	deleteIndex, err := client.DeleteIndex("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(deleteIndex.Acknowledged)
}

func UpdateIndexNow(client *elastic.Client, ctx context.Context) {
	update, err := client.Update().Index("twitter").Type("tweet").Id("1").
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)
}








