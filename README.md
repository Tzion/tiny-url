This is a tiny-url server implementation written in Go!

To start the server execute from the terminal
```bash
$ 'go run server.go'
```
The default configuration is to run locally and listen to port 8080 (localhost:8080).


- Upon POST request the server responds with the possix (HASHCODE) of the tiny url matched the website that was given as part of the POST request (in its body as text/plain).

- Upon GET requests to '{HOST_ADDRESS}/HASHCODE' the servers riderects to the website represented by that hash-code.



Example:

http://localhost:8080/ (POST) {Body:'https://github.com/Tzion/tiny-url'} -> Response {'S1BoCNyOo1M='}
http://localhost:8080/S1BoCNyOo1M= (GET) -> redirect to 'https://github.com/Tzion/tiny-url' 



#### Tests
This project contains 3 packages, each has its own unit test.
To run the tests navigate to the package and run 
```bash
$ go test
````
