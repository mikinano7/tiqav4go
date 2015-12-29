# About
tiqav4go is a [tiqav](http://dev.tiqav.com/) API binding library for the Go language.

# Usage
```go
client := NewTiqavClient()

// Search API
if res, err := client.Search("ちくわ"); err != nil {
	fmt.Print(err)
} else {
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].Image().Glitch)
	}
}

// Image API
if res, err := client.GetImages("1eb"); err != nil {
	fmt.Print(err)
} else {
	fmt.Println(res.Image().Original)
}
```

# Progress
- [x] Image Entity
- [x] Search API
    - [x] GET search
    - [x] GET search/newest
    - [x] GET search/random
- [ ] Image API
    - [x] GET images/:id
    - [ ] POST images
- [ ] Tag API
    - [ ] GET images/:id/tags
    - [ ] GET tags
- [ ] I'm Feeling Lucky
