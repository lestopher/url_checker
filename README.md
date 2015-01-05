# Install
`go get github.com/lestopher/url_checker`

# Usage
Binary should be installed as url_checker, use with the following syntax:  
```
$ url_checker -file=/path/to/urls
```

# Caveats
* Urls file should delimit urls with newlines
* If the protocol is missing, program will use http as the default
* Developed on go 1.3 but 1.x should be OK
