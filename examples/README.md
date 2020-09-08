## GQL Payload Examples

This example assumes you have a GraphQL Server running and have the following schema and resolvers

### Mutation

```go

input ClusterInput {
  cluster_name: String!
  description: String
  platform_name: String!
  project_id: ID!
  cluster_type: String!
}

type Mutation{
  #It is used to create external cluster.
  userClusterReg(clusterInput: ClusterInput!): String!
}
```

### Query 

```go

type Cluster {
   cluster_id: ID!
   project_id: ID!
   cluster_name: String!
   description: String
   platform_name: String!
   access_key: String!
   is_registered: Boolean!
   is_cluster_confirmed: Boolean!
   is_active: Boolean!
   updated_at: String!
   created_at: String!
   cluster_type: String!
 }


type Query{  
  getCluster(project_id: String!, cluster_type: String): [Cluster!]!
}
```
