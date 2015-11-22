# About
tiqav4go is a [tiqav](http://dev.tiqav.com/) API binding library for the Go language.

# Usage
```go
client := NewTiqavClient()
if res, er := client.Search("ちくわ"); er != nil {
	fmt.Print(er)
} else {
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].Image().Glitch)
	}
}
```

# Progress
- [x] Image Entity
- [x] Search API
    - [x] GET search
    - [x] GET search/newest
    - [ ] GET search/random
- [ ] Image API
    - [ ] GET images/:id
    - [ ] POST images
- [ ] Tag API
    - [ ] GET images/:id/tags
    - [ ] GET tags
- [ ] I'm Feeling Lucky
