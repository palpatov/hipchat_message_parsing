# hipchat_message_parsing


Examples of testing routines:

curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Olympics are starting soon; http://www.nbcolympics.com\"}" http://localhost:8080/hipchat_message_parsing
curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Good morning! (megusta) (coffee)\"}" http://localhost:8080/hipchat_message_parsing
curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016\"}" http://localhost:8080/hipchat_message_parsing
