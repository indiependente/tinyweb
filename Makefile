.PHONY: clean
clean: 
	rm -rf ./html
	mkdir ./html

.PHONY: datefmt
datefmt: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./export/wasm.go
	cp ./export/wasm.js ./html/
	cp ./export/index.html ./html/

.PHONY: export
export: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./export/wasm.go
	cp ./export/wasm.js ./html/
	cp ./export/index.html ./html/

.PHONY: main
main: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm -no-debug ./main/main.go
	cp ./main/index.html ./html/

.PHONY: run
run: 
	go run server.go

.PHONY: wasm_exec
wasm_exec:
	cp ./target/wasm_exec.js ./html/