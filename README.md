# GAE/Go template

## Requirements
- [Google Cloud SDK](https://cloud.google.com/sdk/)
- [Google App Engine SDK](https://cloud.google.com/appengine/downloads)
- [direnv](https://github.com/direnv/direnv)

## How to use

### Run in local server

```
$ make serve
```

### Deploy

```
$ make deploy
```

## Structure
- ```direnv``` will export the new ```GOROOT``` and ```GOPATH``` according to ```.envrc```.
- ```lib``` has the codes which are commonly used in modules (e.g., ```import lib/xxx```)
- ```modules``` includes the modules of GAE
- ```gopath``` imports external libraries (```gopath``` is separated from our own codes. See [5] for this reason)

## References

- [1] [GAE/Go Tutorial](https://cloud.google.com/appengine/training/go-plus-appengine/hello-appengine)
- [2] [CONFIGURING YOUR GOPATH WITH GO AND GOOGLE APP ENGINE](http://www.compoundtheory.com/configuring-your-gopath-with-go-and-google-app-engine/)
- [3] [eureka tech blog](https://developers.eure.jp/tech/go-appengine-sql-waf/) (in Japanese)
- [4] [Qiita article by sinmetal](http://qiita.com/sinmetal/items/71cfba4ae27cc2366572) (in Japanese)
- [5] [Blog post by knightso](http://knightso.hateblo.jp/entry/2014/11/26/103637) (in Japanese)
