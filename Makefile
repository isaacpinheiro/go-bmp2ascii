.PHONY: build clean install uninstall

build:
	go build

clean:
	rm go-bmp2ascii

install:
	go build
	sudo mv go-bmp2ascii /usr/local/bin
	echo "\n\ngo-bmp2ascii was successfully installed."

uninstall:
	sudo rm /usr/local/bin/go-bmp2ascii
	echo "\n\ngo-bmp2ascii was uninstalled."

