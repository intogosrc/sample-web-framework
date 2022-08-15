all:
	go build -o ./build/FrameworkTest .

run:
	./build/FrameworkTest ./build/conf.json

clean:
	rm -rf ./build/FrameworkTest
