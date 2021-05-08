# learning-go
Learnning to use Golang for tools

A cloud Guru course - System Tooling with Go


## Commends

1. Running the code directly
    ```shell
    go run main.co
    ```
1. Building a binary
    ```shell
    go build
    ```

## Comments

Allo to generate Go documentation - [documentation](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
```go
/*
A multi line comment
*/

// A single line comment

func main() {
    fmt.println("somenthing") // This is trailing comment
}
```


## Third-Party Packages

go get - download (and compile if needed) 3rd-party libraries
go install - build and install your code as binary to the /go/bin folder

```shell
go get github.com/spf13/cobra@v1.1.3
```


### Known libraries

[cobra](https://github.com/spf13/cobra) - Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.