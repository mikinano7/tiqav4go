# About
tiqav4go is a [tiqav](http://dev.tiqav.com/) API binding library for the Go language.

# Usage
```go
client := tiqav4go.NewTiqavClient()

if results, err := client.Search("query"); err != nil {
	for i := 0; i < len(results); i++ {
		fmt.Println(results[i].Url())
	}
}
```

# Progress
- [ ] Search API
    - [x] search by specified word
    - [ ] error handling
    - [ ] specify the number of response
- [ ] Image API
- [ ] Tag API

