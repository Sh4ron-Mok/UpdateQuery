source = main.go
compiler = go
executable = cod.exe
debug_dir = build/debug
release_dir = build/release

$(executable): $(source)
	go build -o $(executable) -ldflags "-s -w"

clean:
	rm -rf $(executable)

run:$(executable)
	./$<

rerun: clean run