# Path

Package path offers classes representing filesystem paths with semantics appropriate for different operating systems.

## Usage

```go
package main

import (
	"log"

	"github.com/mohanson/godump/path"
)

func main() {
	p := path.Wrap("/parent/stem.suffix")
	log.Println(p.Full)   // /parent/stem.suffix
	log.Println(p.Name)   //         stem.suffix
	log.Println(p.Parent) // /parent
	log.Println(p.Stem)   //         stem
	log.Println(p.Suffix) //             .suffix
}
```
