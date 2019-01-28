# Enums
Go does not natively support enums. However, through the following design pattern, enums can be used.

```go
package foo

type enum int

const (
  // Foo1 ...
  Foo1 enum = iota
  // Foo2 ...
  Foo2 enum
)

// Enum is the exported enum that can be required by outside packages
type Enum interface {
  Type() enum
  String() string
}

// Type returns the enum
func (f enum) Type() enum {
  return f
}

// NOTE: the String() method is made with the 'stringer' command
```

```go
package bar

// RequiresFoo requires a Foo enum.
//
// In this way, only Foo1 or Foo2 can be passed to 'RequiresFoo'
func RequiresFoo(f foo.Enum) error {
  // do something with f...

  return nil
}
```

# Stringer
Each enum in this package includes a `String()` method. Rather then implement these methods by hand for each enum, the code is auto generated using `golang.org/x/tools/cmd/stringer`. Run `$ make check` to see if you have stringer installed and `$ make install/stringer` if the previous check errors.

After following the above steps, run `$ make` to autogenerate the stringer methods.
