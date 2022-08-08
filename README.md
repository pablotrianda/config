# Config 🪚
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)
[![Open Source Love png1](https://badges.frapsoft.com/os/v1/open-source.png?v=103)](https://github.com/ellerbrock/open-source-badges/)

This package allow to read configurations an a easy and fast way from a `yaml` file

This mini tool have the follow actions:
* [x] Read a config file
* [ ] Sync with a remote repository (git)

## Installation
To install Gen package, you need to install Go and set your Go workspace first.

1.The first need Go installed(version 1.14+ is required), then you can use the below Go command to install Gen.

```
go get -u github.com/pablotrianda/config
```

2.Import it in your code:
```golang
import "github.com/pablotrianda/config"
```
## Usage example
Having a `yaml` file with this data:
```yaml
some: thing
here: goes
```
```golang
import "github.com/pablotrianda/config"

func main(){
  // Load the configuration file
  cfg := config.Load("to/yaml/file/conf.yaml")
  // Get values
  cfg.Get("some")
}
```
## Contact
You can reach me on [Twitter @pablotrianda](https://www.twitter.com/pablotrianda)
