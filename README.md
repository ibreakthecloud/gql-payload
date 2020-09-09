# GQL Payload 
gql-payload returns the request body for the provided query or mutation in Go. It is the strip down version of go based gql-client client only till query builder.

### Example: 

**Query**
```go
gql := gql.NewField(`
	query GetUserByID($id: String!) {
	    getUser (id:$id) {
	        name
	        email
	        phone
	        company
	    }
	}
`)

reqBody, _ := gql.Var("id", "b1f6366f-5263-4f9e-864b-dde9af14ea840").GenerateRequestBody()
r, _ := http.NewRequest("POST", GQL_SERVER_URL, reqBody)
```

