# Fix for the protoc

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
installs only the Go Plugin for ProtoBuf but not the actual `protoc` compiler.

On Mac
```
brew install protobuf
```
correctly installed it.

But as pointed out during the classes something was off with my Env Vars.

So by adding the following lines to my `$HOME/.zshrc`
```bash
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"
```
and then sourcing it by running
```bash
source ~/.zshrc
```
I was able to fix my problems.

After that
```bash
protoc --go_out=. kv.proto
```
generated the current `kvproto` Go package.

Next
```bash
go get google.golang.org/protobuf
```
was needed to install the ProtoBuf's Go deps.