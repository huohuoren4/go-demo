@goto %1

@:hello
@REM hello.go
@REM 编译 go 代码
@go build -o .\target\hello.exe hello.go
@REM 执行 go 代码
@.\target\hello.exe
@exit

@:main
@REM main.go
@REM 编译 go 代码
@go build -o .\target\main.exe main.go
@REM 执行 go 代码
@.\target\main.exe
@exit
