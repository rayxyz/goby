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
MIT License

Copyright (c) 2020 Ray

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
