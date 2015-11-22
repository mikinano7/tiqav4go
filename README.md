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
- [ ] Image Entity
    - [x] Original
    - [ ] Thumbnail
    - [ ] Glitch
- [x] Search API
    - [x] GET search
    - [ ] GET search/newest
    - [ ] GET search/random
- [ ] Image API
    - [ ] GET images/:id
    - [ ] POST images
- [ ] Tag API
    - [ ] GET images/:id/tags
    - [ ] GET tags
- [ ] I'm Feeling Lucky
