# hipchat_message_parsing REST service 

This was to be my first exposure to _GO_. I've used this project to learn about GO language constructs, _GO_ style of concurrent processing, dockerizing _GO_ applications, and will be looking next at deploying the dockerized app to [Google Compute Engine](http://cloud.google.com/compute). 


## Getting Started

go get github.com/gorilla/mux
go get github.com/palpatov/hipchar_message_parsing
go install github.com/palpatov/hipchat_message_parsing

you can now run the hipchat_message_parsing binary from your $GOBIN. The index page on localhost:8080 will give you usage examples 


## Service specification

see _api_spec.txt_ for service specs. The service follows the spec exactly to the t. This means that emoticons are considered to be all (<alphanumeric string>) patterns where the length of the string is 15 charaqcters or less as specified. The formatting of the output JSON is exactly that of the service spec file. You don't like that, get your own cooking show. 
Having said that...the output formatting was factored out into the _hcjson/Formatter.go_ class, so if you care to change it change it there. Word of caution, that would most likely break the unit tests as they check the output format to the character.

## API design

This is a RESTfull service. POSTs are used to service the requests. The input format is JSON. See below curl statements for an example.

    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Olympics are starting soon; http://www.nbcolympics.com\"}" http://localhost:8080/hipchat_message_parsing
    
    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Good morning! (megusta) (coffee)\"}" http://localhost:8080/hipchat_message_parsing
    
    curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016\"}" http://localhost:8080/hipchat_message_parsing
    
## Implementation details

I've split the processing into two goroutines, due to the remote page lookup that has to happen, see _MessageParser.go_. I've used mutex to synchronise the routines. The word on the street is that is more efficient than channels. See [here](http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/) for discussion. 

## Dockerizing the app

There is a docker file in the root. It exposes the app on 8080, so use --publish 80:8080 to change the port to whatever you want it to be (ex 80) in your deployment.
