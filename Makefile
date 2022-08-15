all:
	go build -o ./build/FrameworkTest.exe .

run:
	./build/FrameworkTest.exe ./build/conf.json

clean:
	rm -rf ./build/FrameworkTest.exe
