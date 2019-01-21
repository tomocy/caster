# Caster

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![CircleCI](https://circleci.com/gh/tomocy/caster.svg?style=svg)](https://circleci.com/gh/tomocy/caster)

Caster enables us to extend html/template so easily.   
If you want to use same html files such as "header.html", "sidebar.html", Caster is the tool for it.   

## Demo
If we have "master.html", "sidebar.html" as common html files and "index.html" as a main html file,   

- master.html
```html
{{ define "master" }}
<html lang="en">
<head>
    <meta charset="UTF-8">
    {{ template "css" }}
    <title>Caster</title>
</head>
<body>
    {{ template "content" . }}
</body>
</html>
{{ end }}
```

- sidebar.html
```html
{{ define "sidebar" }}
<nav>
    <ul>
        <li>
            <a href="https://github.com">GitHub</a>
        </li>
        <li>
            <a href="https://twitter.com">Twitter</a>
        </li>
    </ul>
</nav>
{{ end }}
```

- index.html
```html
{{ define "css" }}
<link rel="stylesheet" href="/css/index.css">
{{ end }}

{{ define "content" }}
{{ include "sidebar" }}
<div>
    <h1>Caster</h1>
    <h3>
        Caster enables us to extend html/template so easily.
    </h3>
</div>
{{ end }}
```

So we should have the output

```html
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="/css/index.css">
    <title>Caster</title>
</head>
<body>
    <nav>
        <ul>
            <li>
                <a href="https://github.com">GitHub</a>
            </li>
            <li>
                <a href="https://twitter.com">Twitter</a>
            </li>
        </ul>
    </nav>
    <div>
        <h1>Caster</h1>
        <h3>
            Caster enables us to extend html/template so easily.
        </h3>
    </div>
</body>
</html>
```

## Requirement
Go1.11 ~   

## Installation
```
go get github.com/tomocy/caster
```

## Useage
Directory is for example like below.   
├── view   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── layout   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── master.html   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── sidebar.html   
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;├── index.html   
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

	http.ListenAndServe(":8080", nil)
}

```

## License
Licensed under [MIT License](/LICENSE)

## Author
[tomocy](https://github.com/tomocy)



