Reflection in Go is the ability of a program to **inspect and manipulate its own structure and behavior at runtime**. It is provided by the `reflect` package in the standard library.

###  Why Reflection?
Reflection is useful when:
- You don’t know the type at compile-time (e.g., serialization, deserialization, generic logic)
- You want to inspect struct tags, field names/types, or methods
- You need to call functions dynamically

---

###  Key Types in `reflect`
- `reflect.Type` – represents the type of a value
- `reflect.Value` – represents the actual value
- `reflect.Kind` – enumeration of basic kinds like `Int`, `Struct`, `Slice`, etc.

---

###  Common Reflection Operations in Go

| Operation                              | Example Code |
|----------------------------------------|--------------|
| Get the `Type` of a value              | `t := reflect.TypeOf(x)` |
| Get the `Value` of a value             | `v := reflect.ValueOf(x)` |
| Get the `Kind` of a value              | `v.Kind()` |
| Check if a value is valid              | `v.IsValid()` |
| Get number of fields in a struct       | `v.NumField()` |
| Get a struct field                     | `v.Field(i)` |
| Get a field by name                    | `v.FieldByName("FieldName")` |
| Get struct tags                        | `t.Field(i).Tag.Get("json")` |
| Set a field value (must be addressable)| `v.Field(i).Set(newVal)` |
| Create new value                       | `reflect.New(t)` |
| Call a function                        | `v.Call([]reflect.Value{...})` |
| Check if a field is settable           | `v.CanSet()` |
| Get method count                       | `v.NumMethod()` |
| Call method by index                   | `v.Method(i).Call(...)` |

---

###  Notes:
- Only exported fields/methods are accessible via reflection.
- To set values, the value must be **addressable** (e.g., use `reflect.ValueOf(&x).Elem()`).
- Reflection is powerful but slower and less type-safe—use only when necessary.

