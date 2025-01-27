alias b := build

host := `uname -a`

build:
	go build -o bgotop main.go

clean:
	rm -f bgotop
