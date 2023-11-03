# Schema

- GraphQL
- RDF
- DQL

が使えますが、GraphQLで作ってみます。

## Notes

schema書き込み

```shell
> curl -X POST localhost:8080/admin/schema -H "Content-Type: application/graphql" --data-binary '@schema.graphql
{"data":{"code":"Success","message":"Done"}}
```

