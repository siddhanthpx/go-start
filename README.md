# Go Start

CLI tool for generating go boiler plate made with [cobra](https://github.com/spf13/cobra).

## Installation

```bash
go get github.com/siddhanthpx/go-start
```

Which will install the binary to your $GOPATH/bin

If you don't have your Go workspace variables configured properly it won't be accessible directly.

Here's how you could set them up :
```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```


## Usage

Once you have installed the tool you are required to export an environment variable to your Go src folder where your projects live 
### Example

```bash
$ export GOSTARTDIR="github.com/foobar"
```
###### Make sure to set this in your .zshrc or .bashrc


After which go-start would be able to initialize the folder. 

```bash
$ go-start workspace <workspace-folder-name>
```

For eg.

```bash
$ go-start workspace foo
$ Go workspace created successfully at: github.com/foobar/foo
$ cd foo
$ ll -a
total 20K
 drwx------  3  foobar foobar 4.0K Jul  5 22:34 .
 drwxr-xr-x 15  foobar foobar 4.0K Jul  5 22:34 ..
 drwxr-xr-x  7  foobar foobar 4.0K Jul  5 22:34 .git
-rw-r--r--   1  foobar foobar 43   Jul  5 22:34 go.mod
-rwx------   1  foobar foobar 99   Jul  5 22:34 main.go
```




## License
[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)
