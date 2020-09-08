package main

import (
	"fmt"
	gql "github.com/ibreakthecloud/gql-payload"
	"io/ioutil"
	"net/http"
)

// Change it with your own GraphQL server URL
const GQL_SERVER_URL = "http://172.17.0.2:31000/query"

func main() {
	gql := gql.NewField(`
	query ($project_id: String!) {
	    getCluster (project_id:$project_id) {
	        cluster_name
	        cluster_id
	    }
	}
`)

	reqBody, err := gql.Var("project_id", "66e40ea7-3a7b-4352-a2c7-a0de2816f4e0").GenerateRequestBody()
	if err != nil {
		//handle error
	}

	r, _ := http.NewRequest("POST", GQL_SERVER_URL, reqBody)

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		//handle error
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//handle error
	}
	fmt.Println(string(body))

}
