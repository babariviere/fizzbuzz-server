# Fizzbuzz Server

This is a simple golang service to generate fizzbuzz sequences. It consists of writing a sequence of numbers and replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of both 3 and 5 by "fizzbuzz".

The API will generate sequences like this:

```
1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...
```

The API also provides custom parameters to modify the output:
- int1: set the first multiple (3 by default)
- int2: set the second multiple (5 by default)
- limit: the size of the sequence
- str1: the string to put when the number match is a multiple of int1
- str2: the string to put when the number match is a multiple of int2

## Running the server

The server will be available at http://localhost:3000 when run. You can also find the documentation for all API endpoints at http://localhost:3000/swagger/index.html

### With go

```sh
go run cmd/fizzbuzz-server/main.go
```

### With nix flakes

```sh
nix run github:babariviere/fizzbuzz
```

### With docker

```sh
docker build -t fizzbuzz .
docker run -p 3000:3000 fizzbuzz
```
