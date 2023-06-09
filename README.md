Git clone this repo

Git clone the Nokia fork for the restconf project (https://github.com/nokia/restconf)

edit go.mod from this repo set the restconf folder correctly on the line that starts with "replace"
- for example: if you've cloned restconf under the folder "restconf" then go.mod is already ok.

Run the example: `./go.sh run .`

Then from another shell, do `curl -H'Content-Type:application/yang-data+json' -H'Accept:application/yang-data+json' http://localhost:8080/restconf/operations/animals:test -d '{"Animals:input":{"cats":"meow","dogs":"woof"}}' -vvv`
You'll see an output like: `{"Animals:output":{"cats":"meow","dogs":"woof"}}`. It means it's working.

What won't work is if you use xml: `curl -H'Content-Type:application/yang-data+xml' -H'Accept:application/yang-data+xml' http://localhost:8080/restconf/operations/animals:test -d '<input><cats>meow</cats><dogs>woof</dogs></input>' -vvv`
That's what needs to be implemented.

Hint:
I think that the function readInput in browser_handler.go (in the restconf lib) needs to be changed to check the Content-Type from r.Headers.
if the content Type is yang-data+xml, then you must parse xml instead of json.
But then there is a call to nodeutil.ReadJSONValues(payload). You'll probably need to implement another one of those method to handle xml.
I think that's the function that converts xml to Yang.

Background:
Restconf is a web server library that serves models written in Yang.
According to RFC 8040, it should support xml and json.
Parsing xml in the library is not as straightforward as you might think. Restconf enforces you to adhere to the yang model. 
It currently does this with json so it must work the same way woth xml.

The task is only to parse xml. Not to output xml. So it's ok to still receive a json response.
