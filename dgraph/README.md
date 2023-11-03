# Schema

- GraphQL
- RDF
- DQL

が使えますが、GraphQLで作ってみます。

## Notes

schemaを反映

```shell
> curl -X POST localhost:8080/admin/schema -H "Content-Type: application/graphql" --data-binary '@schema.graphql
{"data":{"code":"Success","message":"Done"}}
```

ユーザを作成

```shell
> curl -X POST 'localhost:8080/graphql' -H 'Content-Type: application/graphql' -d 'mutation MyMutation {
  addUser(input: [
    { name: "makoto", description: "makoto description", tweets: [], follow: []}
  ]) {
    numUids
    user {
      name
    }
  }
}'
{"data":{"addUser":{"numUids":1,"user":[{"name":"makoto"}]}},"extensions":{"touched_uids":7,"tracing":{"version":1,"startTime":"2023-11-03T06:43:23.405885334Z","endTime":"2023-11-03T06:43:23.41094743Z","duration":5062096,"execution":{"resolvers":[{"path":["addUser"],"parentType":"Mutation","fieldName":"addUser","returnType":"AddUserPayload","startOffset":129447,"duration":4930002,"dgraph":[{"label":"preMutationQuery","startOffset":264073,"duration":1290961},{"label":"mutation","startOffset":1614574,"duration":1423462},{"label":"query","startOffset":3943627,"duration":1111060}]}]}}}}⏎
```

ユーザを取得(http://localhost:8000/のconsole -> Queryで実行)

```dql
{
  users(func: has(<dgraph.type>)) @filter(eq(<dgraph.type>, "User")) {
    expand(_all_) {
      expand(_all_)
    }
  }
}
```
