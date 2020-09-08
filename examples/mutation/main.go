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
	  mutation UserClusterReg($clusterInput: ClusterInput!) {
	      userClusterReg (clusterInput:$clusterInput)
	  }
	`)

	clusterInput := make(map[string]interface{})
	clusterInput["project_id"] = "66e40ea7-3a7b-4352-a2c7-a0de2816f4e0"
	clusterInput["cluster_name"] = "example-cluster"
	clusterInput["platform_name"] = "minikube"
	clusterInput["cluster_type"] = "dev"

	reqBody, err := gql.Var("clusterInput", clusterInput).GenerateRequestBody()
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
