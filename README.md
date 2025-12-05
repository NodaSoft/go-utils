# go-utils

`go-utils` — is a utility library for the Go programming language. It provides convenient generic functions and tools
for working with slices, strings, time, and much more.

## Packages

- [generics](#generics)
- [maps](#maps)
- [math](#math)
- [models](#models)
- [other](#other)
- [slices](#slices)
- [strings](#strings)
- [time](#time)

## generics

A package that provides common types for use in other project packages. It generalizes functions and data types to work
with various value types.

### Main types:

- **[Numeric](#Numeric)** - combines all numeric types like `int`, `float32`, `uint` etc.

### Numeric

The `Numeric` interface combines all numeric types like `int`, `float32`, `uint`, and others.

**Usage example:**

```go
func Sum[T Numeric](n []T) T {
    var sum T
    for _, v := range n {
        sum += v
    }
    return sum
}
```

## maps

A package that provides functions for convenient map operations.

### Main functions:

- **[Has](#Has)**: Checks if a map contains a given key.
- **[Merge](#Merge)**: Merges two maps. Values from map "a" take precedence.
- **[DiffKeys](#DiffKeys)**: Returns map "a" without the elements from map "b".

### Has

Checks if a map contains a given key.

**Parameters:**

- `m` — a map of type `map[K]V`, where `K` are the keys, and `V` are the values.
- `key` — the key to check for.

**Return value:**

- `bool` — `true`, if the key is present in the map, otherwise `false`.

**Usage example:**

```go
m := map[string]int{"one": 1, "two": 2}
exists := maps.Has(m, "two")
// exists: true
```

### Merge

Merges two maps. Values from map "a" take precedence.

**Parameters:**

- `a` — first map of type map `map[K]V`.
- `b` — second map of type `map[K]V`.

**Return value:**

- `map[K]V` is a new map containing the combined elements.

**Usage example:**

```go
a := map[string]int{"one": 1, "two": 2}
b := map[string]int{"two": 3, "three": 4}
merged := maps.Merge(a, b)
// merged: {"one": 1, "two": 3, "three": 4}
```

### DiffKeys

Returns map "a" without elements from map "b".

**Parameters:**

- `a` is a map of type `map[K]V` from which elements will be removed.
- `b` is a map of type `map[K]V` containing the keys to be removed from the map "a".

**Return value:**

- `map[K]V` is a new map containing elements from "a" that are not in "b".

**Usage example:**

```go
a := map[string]int{"one": 1, "two": 2, "three": 3}
b := map[string]int{"two": 2, "four": 4}
diff := maps.DiffKeys(a, b)
// diff: {"one": 1, "three": 3}
```

## math

Package providing mathematical functions for working with numeric types.

### Main functions:

- **[Max](#Max)**: Returns the maximum value from the provided arguments.
- **[Min](#Min)**: Returns the minimum value from the provided arguments.
- **[Sum](#Sum)**: Returns the sum of all the provided values.

### Max

Returns the maximum value from the provided arguments.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` is the maximum value from the provided arguments.

**Usage example:**

```go
maxValue := math.Max(1, 2, 3, 4, 5)
// maxValue: 5
```

### Min

Returns the minimum value from the provided arguments.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` is the minimum value from the provided arguments.

**Usage example:**

```go
minValue := math.Min(1, 2, 3, 4, 5)
// minValue: 1
```

### Sum

Returns the sum of all the provided values.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` is the sum of all the provided arguments.

**Usage example:**

```go
total := math.Sum(1, 2, 3, 4, 5)
// total: 15
```

## models

Package providing functions for working with entities (models).

### Main functions:

- **[CollectIDs](#CollectIDs)**: Returns a slice of identifiers from a slice of entities.
- **[CollectIDsFromMap](#CollectIDsFromMap)**: Returns a slice of identifiers from a map of entities.
- **[UniqueValues](#UniqueValues)**: Method for collecting unique values from any field of a model into a slice of the desired result type.
- **[UniqueValuesFromMap](#UniqueValuesFromMap)**: Method for collecting unique values from any field of a model into a slice of the desired result type, using a map.

### CollectIDs

Returns a slice of identifiers from a slice of entities.

**Parameters:**

- `sl` is a slice of entities of type `T`, which have an id.

**Return value:**

- `[]uint` is a slice of unique identifiers.

**Usage example:**

```go
type User struct {
    ID uint 
	// other fields
}

func (u User) GetID() uint {
    return u.ID
}

users := []User{{ID: 1}, {ID: 2}, {ID: 1}}
ids := CollectIDs(users)
// ids: []uint{1, 2}
```

### CollectIDsFromMap

Returns a slice of identifiers from a map of entities that have an id.

**Parameters:**

- `m` is a map where the keys are of type `K`, and the values are of type `T`, which have an id.

**Return value:**

- `[]uint` is a slice of unique identifiers.

**Usage example:**

```go
type User struct {
    ID uint
    // other fields
}

func (u User) GetID() uint {
    return u.ID
}

userMap := map[string]User{
    "user1": {ID: 1},
    "user2": {ID: 2},
    "user3": {ID: 1},
}

ids := CollectIDsFromMap(userMap)
// ids: []uint{1, 2}
```

### UniqueValues

Method for collecting unique values from any field of a model into a slice of the desired result type.

**Parameters:**

- `slice` is a slice of entities of type `S`.
- `getter` is a function that takes an entity of type `S` and returns a value of type `R`, which will be added to the final slice.

**Return value:**

- `[]R` is a slice of unique values obtained from the provided slice.

**Usage example:**

```go
type Product struct {
    ID   uint
    Name string
}

func GetProductID(p Product) uint {
    return p.ID
}

products := []Product{
    {ID: 1, Name: "Product A"},
    {ID: 2, Name: "Product B"},
    {ID: 1, Name: "Product C"},
}

uniqueIDs := UniqueValues(products, GetProductID)
// uniqueIDs: []uint{1, 2}
```

### UniqueValuesFromMap

Method for collecting unique values from any field of a model into a slice of the desired result type, using a map.

**Parameters:**

- `m` is a map where the keys are of type `K`, and the values are of type `V`.
- `getter` is a function that takes a value of type `V` and returns a value of type `R`, which will be added to the final slice.

**Return value:**

- `[]R` is a slice of unique values obtained from the provided map.

**Usage example:**

```go
type User struct {
    ID   uint
    Name string
}

func GetUserName(u User) string {
    return u.Name
}

users := map[string]User{
    "user1": {ID: 1, Name: "Alice"},
    "user2": {ID: 2, Name: "Bob"},
    "user3": {ID: 3, Name: "Alice"},
}

uniqueNames := UniqueValuesFromMap(users, GetUserName)
// uniqueNames: []string{"Alice", "Bob"}
```

## other

Package that contains utility functions for working with various data types.

### Main functions:

- **[FirstNonEmpty](#FirstNonEmpty)**: Returns the first element with a non-zero value from the provided arguments.

### FirstNonEmpty

Returns the first element with a non-zero value from the provided arguments.

**Parameters:**

- `tt` is a variable number of arguments of type `T`, where `T` is any type that supports comparison.

**Return value:**

- `T` is the first non-zero element from the provided arguments. If all elements are zero, it returns the default value for type `T`.

**Usage example:**

```go
first := other.FirstNonEmpty("", "hello", "world")
// first: "hello"

firstNum := other.FirstNonEmpty(0, 1, 2)
// firstNum: 1

firstNil := other.FirstNonEmpty(nil, nil)
// firstNil: nil
```

### Short

Package that contains short functions for working with various data types.

### Main functions:

- **[If](#If)**: Returns the value of `then` if condition is true, otherwise returns the value of `otherwise`.
- **[IfFunc](#IfFunc)**: Returns the result of calling `then` if condition is true, otherwise returns the result of calling `otherwise` (lazy evaluation).
- **[IfFuncE](#IfFuncE)**: Returns the result of calling `then` if condition is true, otherwise returns the result of calling `otherwise` (lazy evaluation with error).

### If

Returns the value of `then` if condition is true, otherwise returns the value of `otherwise`.

**Parameters:**

- `condition` — a boolean value that determines which value to return.
- `then` — a value of type `T` that will be returned if `condition` is `true`.
- `otherwise` — a value of type `T` that will be returned if `condition` is `false`.

**Return value:**

- `T` — the value of `then` if `condition` is `true`, otherwise the value of `otherwise`.

**Usage example:**

```go
result := other.If(true, "yes", "no")
// result: "yes"

value := other.If(false, 10, 20)
// value: 20
```

### IfFunc

Returns the result of calling `then` if condition is true, otherwise returns the result of calling `otherwise` (lazy evaluation). This function allows for lazy evaluation, meaning the functions are only called when needed.

**Parameters:**

- `condition` — a boolean value that determines which function to call.
- `then` — a function that returns a value of type `T` and will be called if `condition` is `true`.
- `otherwise` — a function that returns a value of type `T` and will be called if `condition` is `false`.

**Return value:**

- `T` — the result of calling `then()` if `condition` is `true`, otherwise the result of calling `otherwise()`.

**Usage example:**

```go
result := other.IfFunc(true,
    func() string { return expensiveOperation() },
    func() string { return "default" },
)
// Only expensiveOperation() will be called

value := other.IfFunc(false,
    func() int { return 100 },
    func() int { return 200 },
)
// value: 200
```

### IfFuncE

Returns the result of calling `then` if condition is true, otherwise returns the result of calling `otherwise` (lazy evaluation with error). This function allows for lazy evaluation and error handling.

**Parameters:**

- `condition` — a boolean value that determines which function to call.
- `then` — a function that returns a value of type `T` and an error, and will be called if `condition` is `true`.
- `otherwise` — a function that returns a value of type `T` and an error, and will be called if `condition` is `false`.

**Return value:**

- `T` — the result of calling `then()` if `condition` is `true`, otherwise the result of calling `otherwise()`.
- `error` — an error returned by the called function, or `nil` if no error occurred.

**Usage example:**

```go
result, err := other.IfFuncE(true,
    func() (string, error) { return fetchData() },
    func() (string, error) { return "default", nil },
)
if err != nil {
    // error handling
}

value, err := other.IfFuncE(false,
    func() (int, error) { return 100, nil },
    func() (int, error) { return 0, errors.New("error") },
)
// value: 0, err: error
```

## slices

Package providing functions for working with slices.

### Main functions:

- **[ConvertSlice](#ConvertSlice)**: changes the type of elements in the slice.
- **[FilterNil](#FilterNil)**: returns a slice without empty values (e.g., 0, "", etc.), modifying the original slice.
- **[Unique](#Unique)**: returns a slice without duplicates, modifying the original slice.
- **[Union](#Union)**: combines two slices, excluding duplicates.
- **[Cross](#Cross)**: returns a slice with values present in both slices.
- **[IsEqual](#IsEqual)**: checks if the slices are identical, regardless of the order of elements.
- **[Has](#Has)**: checks if the slice contains a given value.
- **[TrimStrings](#TrimStrings)**: removes spaces from each element of a string slice.
- **[ToKeyMap](#ToKeyMap)**: returns a map with keys equal to the values of the slice.
- **[SliceDiff](#SliceDiff)**: returns a slice containing elements present in the first slice but absent in others.
- **[SliceIntersect](#SliceIntersect)**: returns a slice with unique values present in all provided slices.
- **[Max](#Max)**: returns the maximum value from the provided elements.
- **[Min](#Min)**: returns the minimum value from the provided elements.
- **[Sum](#Sum)**: returns the sum of all values.

### ConvertSlice

Function that changes the type of elements in a slice.

**Parameters:**

- `s` is a slice of elements of type `T` that needs to be converted.

**Return value:**

- `[]R` is a new slice of elements of type `R`, obtained by converting the elements from the slice `s`.

**Usage example:**

```go
newSlice := slices.ConvertSlice[int32, uint]([]int32{1, 2, 3})
// newSlice: []uint{1, 2, 3}
```

### FilterNil

Function that returns a slice without empty values (e.g., `0`, `""`, etc.). Note that it modifies the original slice.

**Parameters:**

- `sl` is a slice of elements of type `T` from which empty values will be removed.

**Return value:**

- `[]T` is a slice containing only non-empty values.

**Usage example:**

```go
values := []int{0, 1, 2, 0, 3}
filtered := slices.FilterNil(values)
// filtered: []int{1, 2, 3}
```

### Unique

Function that returns a slice without duplicates. Note that it modifies the original slice.

**Parameters:**

- `sl` is a slice of elements of type `T` from which duplicates will be removed.

**Return value:**

- `[]T` is a slice containing only unique values.

**Usage example:**

```go
values := []int{1, 2, 2, 3, 4, 4}
uniqueValues := slices.Unique(values)
// uniqueValues: []int{1, 2, 3, 4}
```

### Union

Function that combines two slices, excluding duplicates.

**Parameters:**

- `sl1` is the first slice of type `T`.
- `sl2` is the second slice of type `T`.

**Return value:**

- `[]T` is a new slice containing unique values from both input slices.

**Usage example:**

```go
slice1 := []int{1, 2, 3}
slice2 := []int{3, 4, 5}
result := slices.Union(slice1, slice2)
// result: []int{1, 2, 3, 4, 5}
```

### Cross

Function that returns a slice of values present in both input slices.

**Parameters:**

- `sl1` is the first slice of type `T`.
- `sl2` is the second slice of type `T`.

**Return value:**

- `[]T` is a new slice containing values that are present in both input slices.

**Usage example:**

```go
slice1 := []int{1, 2, 3}
slice2 := []int{2, 3, 4}
result := slices.Cross(slice1, slice2)
// result: []int{2, 3}
```

### IsEqual

Function that checks whether two slices are identical, regardless of the order of elements.

**Parameters:**

- `sl1` is the first slice of type `T`.
- `sl2` is the second slice of type `T`.

**Return value:**

- `bool` — `true` if the slices are identical (contain the same elements in any order), otherwise `false`.

**Usage example:**

```go
slice1 := []int{1, 2, 3}
slice2 := []int{3, 2, 1}
isEqual := slices.IsEqual(slice1, slice2)
// isEqual: true
```

### Has

Function that checks whether a slice contains a specified value.

**Parameters:**

- `sl` is a slice of type `T` in which to search.
- `n` is a value of type `T` that needs to be found in the slice.

**Return value:**

- `bool` — `true` if the value is present in the slice, otherwise `false`.

**Usage example:**

```go
slice := []string{"apple", "banana", "cherry"}
exists := slices.Has(slice, "banana")
// exists: true
```

### TrimStrings

Function that trims whitespace from the beginning and end of each string in a slice of strings.

**Parameters:**

- `ss` is a slice of strings to be processed.

**Return value:**

- `[]string` — a slice of strings where each string is trimmed of whitespace.

**Usage example:**

```go
strings := []string{"  apple  ", " banana ", "cherry  "}
trimmed := slices.TrimStrings(strings)
// trimmed: []string{"apple", "banana", "cherry"}
```

### ToKeyMap

Function that converts a slice of values into a map where the keys are the elements of the slice and the values are boolean indicating the presence of these keys.

**Parameters:**

- `sl` is a slice of values of type `T` to be converted into a map.

**Return value:**

- `map[T]bool` — a map where the keys are the elements from the slice and the values are `true`.

**Usage example:**

```go
values := []string{"apple", "banana", "cherry"}
keyMap := slices.ToKeyMap(values)
// keyMap: map[string]bool{"apple": true, "banana": true, "cherry": true}
```

### SliceDiff

Function that returns a slice containing elements that are present in the first slice but absent in the other provided slices.

**Parameters:**

- `slices` is a variable number of slices of type `T` from which the difference will be calculated.

**Return value:**

- `[]T` — a slice containing elements from the first slice that are not in the others.

**Usage example:**

```go
slice1 := []int{1, 2, 3, 4}
slice2 := []int{3, 4, 5}
slice3 := []int{4, 5, 6}

result := slices.SliceDiff(slice1, slice2, slice3)
// result: []int{1, 2}
```

### SliceIntersect

Function that returns a slice of unique values present in all provided slices.

**Parameters:**

- `slices` is a variable number of slices of type `T` from which the intersection will be calculated.

**Return value:**

- `[]T` — a slice containing unique elements that are present in all of the provided slices.

**Usage example:**

```go
slice1 := []int{1, 2, 3, 4}
slice2 := []int{3, 4, 5}
slice3 := []int{4, 5, 6}

result := slices.SliceIntersect(slice1, slice2, slice3)
// result: []int{4}
```

### Max

Function that returns the maximum value from the provided arguments.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` — the maximum value from the provided arguments.

**Usage example:**

```go
maxValue := slices.Max([]int{1, 2, 3, 4, 5})
// maxValue: 5
```

### Min

Function that returns the minimum value from the provided arguments.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` — the minimum value from the provided arguments.

**Usage example:**

```go
minValue := slices.Min([]int{1, 2, 3, 4, 5})
// minValue: 1
```

### Sum

Function that returns the sum of all provided values.

**Parameters:**

- `n` is a variable number of arguments of type `T`, where `T` is any numeric type supported by the `Numeric` interface.

**Return value:**

- `T` — the sum of all provided arguments.

**Usage example:**

```go
total := slices.Sum([]int{1, 2, 3, 4, 5})
// total: 15
```

# strings

Package providing functions for working with strings.

### Main functions:

- **[Truncate](#Truncate)**: Truncates a string to the specified number of runes.

#### Truncate

Truncates a string to the specified number of runes.

**Parameters:**

- `str` — the string to be truncated.
- `maxRunes` — the maximum number of runes to truncate the string to.

**Return value:**

- `string` — the truncated string if the length exceeds `maxRunes`, otherwise the original string.

**Usage example:**

```go
result := strings.Truncate("Hello, World!", 5)
// result: "Hello"
```

## time

Package providing functions for working with time values.

### Main functions:

- **[Midnight](#Midnight)**: Returns the time corresponding to midnight for the current date in the local time zone.
- **[MidnightByLocation](#MidnightByLocation)**: Returns midnight time for the specified location.
- **[MidnightByTimeZone](#MidnightByTimeZone)**: Returns midnight time for the specified time zone.

### Midnight

Returns the time corresponding to midnight for the current date in the local time zone.

**Return value:**

- `time.Time` — the time value corresponding to midnight.
- `error` — an error if there was a problem calculating the time (usually does not occur).

**Usage example:**

```go
midnight, err := time.Midnight()
if err != nil {
    // error handling
}
// midnight: 2024-09-20 00:00:00 +0000 UTC
```

### MidnightByLocation

Returns the midnight time for the specified location.

**Parameters:**

- `loc` — a pointer to a `time.Location` structure representing the time zone.

**Return value:**

- `time.Time` — the time value corresponding to midnight in the specified time zone.
- `error` — an error if there was a problem calculating the time (usually does not occur).

**Usage example:**

```go
loc, err := time.LoadLocation("Europe/Moscow")
if err != nil {
    // error handling
}

midnight, err := time.MidnightByLocation(loc)
if err != nil {
    // error handling
}
// midnight: 2024-09-20 00:00:00 +0300 MSK
```

### MidnightByTimeZone

Returns the midnight time for the specified time zone.

**Parameters:**

- `timeZone` — a string representing the name of the time zone (e.g., "Europe/Moscow").

**Return value:**

- `time.Time` — the time value corresponding to midnight in the specified time zone.
- `error` — an error if there was a problem loading the time zone (e.g., if the specified time zone does not exist).

**Usage example:**

```go
midnight, err := time.MidnightByTimeZone("Europe/Moscow")
if err != nil {
    // error handling
}
// midnight: 2024-09-20 00:00:00 +0300 MSK
```
