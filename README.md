# Overview
This codebase is a demo of the ability to use twemproxy to access memcache. It uses docker-compose to stand up a local memcache isntance and [twemproxy](https://github.com/markwunsch/twemproxy-docker) instance. 

The code in [main.go](https://github.com/markwunsch/twemproxy-memcache-demo/blob/master/main.go) uses a memcache client to access the memcache server through twemproxy. 

# Usage

    $ go get github.com/markwunsch/twemproxy-memcache-demo
    $ cd $GOPATH/src/github.com/markwunsch/twemproxy-memcache-demo
    $ docker-compose up
    $ go run main.go    
    $ docker-compose down


    