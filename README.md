# Caster

[![CircleCI](https://circleci.com/gh/tomocy/caster.svg?style=svg)](https://circleci.com/gh/tomocy/caster)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Caster enables us to extend html/template so easily.   
If you want to use same html files such as "header.html", "sidebar.html", Caster is the tool for it.   

## Useage
Directory is for example like the below.   
Note: master template (in this case "master.html") should be defined as "master" to start with there.   
├── view   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── layout   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── master.html   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── sidebar.html   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── index.html   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── new.html   
├── main.go   

You can cast view 
main.go
```go
package main

import (
	"log"
	"net/http"

	"github.com/tomocy/caster"
)

func main() {
    // Ready to cast htmls
    // - create a new Caster instance with common parts
    caster, err := caster.New(
        "layout/master.html",
        "layout/sidebar.html",
    )
    if err != nil {
        panic(err)
    }

    // - name specific htmls and extend them
    viewMap := map[string][]string{
        "index":   {"index.html"},
        "new":     {"new.html"},
    }
    if err := caster.ExtendAll(viewMap); err != nil {
        panic(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Cast view using the key in the viewMap
        if err := caster.Cast(w, "index", nil); err != nil {
		log.Printf("failed to cast view: %s\n", err)
	}
    })
    http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
        // Cast view using the key in the viewMap
        if err := caster.Cast(w, "new", nil); err != nil {
		log.Printf("failed to cast view: %s\n", err)
	}
    })

	http.ListenAndServe(":8080", nil)
}
```

In this example,
- when you visit "/",  
you will see view composed of "master.html", "sidebar.html", and "***index.html***".
- when you visit "/new",  
you will see view composed of "master.html", "sidebar.html", and "***new.html***".

## Installation
```
go get github.com/tomocy/caster
```

## License
Licensed under [MIT License](/LICENSE)

## Author
[tomocy](https://github.com/tomocy)



