sample web framework

This is a sample web framework, it implements basic usage of web application  
The external packages inuse:  

1. gorm: use to operate db
2. gin: use to start a web server

api docement and table creation SQLs can found in docs dir  

build & run

for macOS, you can use make to build and run   

for windows:

build: go build -o ./build/FrameworkTest.exe .  
run: ./build/FrameworkTest.exe ./build/conf.json  

tip: remove ./build/FrameworkTest.exe befour build the application

test page  
http://127.0.0.1:3000/auth/test?auth_token=1234567890