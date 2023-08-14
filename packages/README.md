# Go packages
- A package is the tiniest form of a reusable code
- Even the simplest Go program is its own package
- There are even a more high level form of reusable code, that is the Go Module.
- Go has its own packages (https://pkg.go.dev/std)
- Some Go functions are builtin (println for example)
- Go packages can have packages ("math/rand for example")

# Third party packages
- Before using a third part package, you need to initiate your module (it must be done in the root)
- Then, you can use ```go get github.com/google/uuid```
- Every downloaded package is in your $GOPATH (usually the $HOME/go)
- Run ```go env GOMODCACHE``` to see it
- Notice that the module’s version is included in its directory name. This allows you to develop against and test multiple versions of the same package, either within one program or across different programs you are writing.