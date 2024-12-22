# Path

Package path offers classes representing filesystem paths with semantics appropriate for different operating systems.

```go
p := Wrap("/parent/stem.suffix")
print(p.Full)   // /parent/stem.suffix
print(p.Name)   //         stem.suffix
print(p.Parent) // /parent
print(p.Stem)   //         stem
print(p.Suffix) //             .suffix
```
