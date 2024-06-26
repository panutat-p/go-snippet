# Resty

https://github.com/go-resty/resty

```sh
go get github.com/go-resty/resty/v2
```

```go
type Users []struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Address  struct {
        Street  string `json:"street"`
        Suite   string `json:"suite"`
        City    string `json:"city"`
        Zipcode string `json:"zipcode"`
        Geo     struct {
            Lat string `json:"lat"`
            Lng string `json:"lng"`
        } `json:"geo"`
    } `json:"address"`
    Phone   string `json:"phone"`
    Website string `json:"website"`
    Company struct {
        Name        string `json:"name"`
        CatchPhrase string `json:"catchPhrase"`
        Bs          string `json:"bs"`
    } `json:"company"`
}

var users Users
client := resty.New()
res, err := client.
    R().
    SetResult(&users).
    Get("https://jsonplaceholder.typicode.com/users")
if err != nil {
    panic(err)
}
if res.IsError() {
    fmt.Println("❌", res.Status(), res.Error())
    t.FailNow()
}
fmt.Printf("%+v\n", users)
```
