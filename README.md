# GAE/Go template

## Requirements
- Google Appengine SDK: Downlaod from [here](https://cloud.google.com/appengine/downloads)
- direnv: ```$ brew install direnv``` for Mac

## How to use

### Run in local server

```
$ goapp serve default
```

### Deployment
Set your own ```<project_id>``` to app.yaml and then,

```
$ goapp deploy default
```

## References
This repository is highly inspired by these posts:

- [Tutorial](https://cloud.google.com/appengine/training/go-plus-appengine/hello-appengine)
- [eureka tech blog](https://developers.eure.jp/tech/go-appengine-sql-waf/) (in Japanese)
- [Qiita article by sinmetal](http://qiita.com/sinmetal/items/71cfba4ae27cc2366572) (in Japanese)
