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

### Endpoints testing

After deployment, if you would like to test the API, see

```
https://<application-project-id>.appspot.com/_ah/api/explorer/
```

in the case of local development (in MacOs), please reboot Chrome by

```
$ open ~/Applications/Google\ Chrome.app --args --user-data-dir=test --unsafely-treat-insecure-origin-as-secure=http://localhost:8080
```

and then, access to

```
http://localhost:8080/_ah/api/explorer/
```

## Structure
- ```direnv``` will export the new ```GOROOT``` and ```GOPATH``` according to ```.envrc```.
- ```modules``` includes the modules of GAE
- ```gopath``` imports external libraries (```gopath``` is separated from our own codes. See [5] for this reason)

## References

- [1] [GAE/Go Tutorial](https://cloud.google.com/appengine/training/go-plus-appengine/hello-appengine)
- [2] [CONFIGURING YOUR GOPATH WITH GO AND GOOGLE APP ENGINE](http://www.compoundtheory.com/configuring-your-gopath-with-go-and-google-app-engine/)
- [3] [eureka tech blog](https://developers.eure.jp/tech/go-appengine-sql-waf/) (in Japanese)
- [4] [Qiita article by sinmetal](http://qiita.com/sinmetal/items/71cfba4ae27cc2366572) (in Japanese)
- [5] [Blog post by knightso](http://knightso.hateblo.jp/entry/2014/11/26/103637) (in Japanese)
- [6] [github.com/GoogleCloudPlatform/go-endpoints](https://github.com/GoogleCloudPlatform/go-endpoints)
- [7] [github.com/GoogleCloudPlatformTraining/cpd200-hello-endpoints-go](https://github.com/GoogleCloudPlatformTraining/cpd200-hello-endpoints-go)
