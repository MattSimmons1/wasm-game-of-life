
# WASM Game Of Life

This project serves as a simple example of running Go functions in a web page with WebAssembly.

It shows off how fast WASM is by computing up to 2500 generations a second.

## How to run

First copy the _wasm_exec.js_ file into this project. It gets imported in index.html and is used to load wasm. 
The version of this file **must correspond to the version of go you are using**, so do this again if you update go.

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./js/wasm_exec.js
```

Build the WASM module:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

Then run a local web server:

```bash
go run server.go
```

Go to http://localhost:8080/

## How it works

When [index.html](index.html) loads it gets [main.wasm](main.wasm) from the server and runs it using 
[wasm_exec.js](js/wasm_exec.js), which registers some WebAssembly functions, allowing them to be called from JavaScript.

[main.wasm](main.wasm) is a WebAssembly module, which we compiled using go. [main.go](main.go) defines some functions 
and binds them to a JS object. It also tells the compiler which types it will be using and interacts with some JS via 
go's "syscall/js" module.

Our go function simulates the game of life, returns the result as a text string, and then updates the HTML in our web 
page.
