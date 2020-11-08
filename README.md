# Goby
A super lightweight Go project base framework.

# Framework project structure
```
$GOPATH/src/goby$ tree -d
.
├── api
├── assets
├── build
├── cmd
├── configs
├── docs
├── examples
├── githooks
├── pkg
│   ├── conf
│   ├── crypto
│   ├── db
│   ├── dict
│   ├── httpx
│   ├── image
│   ├── network
│   ├── ratelimit
│   ├── redis
│   ├── template
│   └── util
├── scripts
├── services
│   └── helloworld
│       ├── api
│       ├── db
│       └── model
├── test
├── third_party
├── tools
└── web
    └── app
        ├── pages
        └── static
            ├── css
            ├── images
            │   └── icons
            └── js
```

# How to use
## Setup
## The default.yaml configuration file
`default.yaml` lives at `/goby/conf/default.yaml`. It is the base configuration file that specifies the basic services  and other essential configs used in the framework.

```
cd $GOPATH/src/goby/pkg/configs/
sudo mkdir -p /goby/conf
sudo copy default.yaml /goby/conf/
```

## Build the binaries
```
cd $GOPATH/src/goby/services/helloworld/
go build -o helloworld
```

## Run your services
```
cd $GOPATH/src/goby/services/helloworld/
./helloworld
```
Copy the compiled binaries to where you whould like to run.

## Examples
Checkout the `helloworld` service and the `web/app`.

# License

