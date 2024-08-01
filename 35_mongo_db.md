# MongoDB

https://github.com/mongodb/mongo-go-driver

```sh
go get go.mongodb.org/mongo-driver/mongo
```

## BSON

https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson

```go
filter := bson.D{{"age", 8}}
```

https://raw.githubusercontent.com/mongodb/docs-golang/v1.10/source/includes/usage-examples/code-snippets/struct-tag.go

```go
filter := bson.D{{"author", "Sam Lee"}}

var result bson.M
_ = coll.FindOne(context.TODO(), filter).Decode(&result)
```
