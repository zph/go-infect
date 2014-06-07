lint:
	./bin/go-lint

hooks:
	cp -f hooks/* .git/hooks/

build: clean
	go build -o dist/infect main.go; \
		GOARCH=amd64 GOOS=linux go build -o dist/infect.linux_amd64 main.go; \
		GOARCH=amd64 GOOS=freebsd go build -o dist/infect.freebsd_amd64 main.go; \
		go build -gcflags '-N' -o dist/infect.debug main.go;

clean:
	rm -f dist/*

install: clean build
	cp dist/infect ~/bin/infect

readme: clean build
	ruby -rerb -e "puts ERB.new(File.read('src/README.md.erb')).result" > README.md \
		&& cat README.md

