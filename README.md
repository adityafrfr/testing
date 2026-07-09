# Go with Tests

TDD-based Go learning project covering fundamentals through test-driven development.

## Prerequisites

- Go 1.26.4+
- Git

## Getting Started

```bash
git clone <repo>
cd "go with tests"
go test ./...
```

## Topics Covered

### 1. Hello World — `helloworld/`

- Functions with parameters and return values
- Constants and switch statements
- Named return values
- TDD cycle: red → green → refactor

```go
Hello("Elodie", "spanish") // "Hola, Elodie"
```

```bash
go test ./helloworld/
```

### 2. Integers — `integers/`

- Basic types and function signatures
- Example functions for documentation
- `go doc` for viewing documentation

```go
Add(2, 2) // 4
```

```bash
go test ./integers/
```

### 3. Iteration — `iteration/`

- Go's single loop construct (`for`)
- `strings.Builder` for efficient string concatenation
- Benchmarking with `b.Loop()`

```go
Repeat("a", 5) // "aaaaa"
```

```bash
go test ./iteration/
go test -bench=. -benchmem ./iteration/
```

### 4. Arrays & Slices — `arrays-and-slices/`

- Slices vs arrays
- `range` for iteration
- `append` for dynamic growth
- Variadic parameters (`...[]int`)
- Tail slicing (`numbers[1:]`)
- `reflect.DeepEqual` and `slices.Equal` for slice comparison

```go
Sum([]int{1, 2, 3})         // 6
SumAll([]int{1, 2}, []int{3, 4}) // []int{3, 7}
SumAllTails([]int{1, 2, 3})    // []int{5}
```

```bash
go test ./arrays-and-slices/
```

### 5. Structs, Methods & Interfaces — `structs-methods-nd-interfaces/`

- Defining structs (`Rectangle`, `Circle`, `Triangle`)
- Value receiver methods
- Interfaces (`Shape`) for polymorphic behavior
- Table-driven tests with anonymous structs

```go
Rectangle{10, 10}.Area() // 100
Circle{10}.Area()        // 314.159...
Triangle{12, 6}.Area()   // 36
```

```bash
go test ./structs-methods-nd-interfaces/
```

### 6. Pointers & Errors — `pointers/`

- Pointer receivers to mutate state
- Custom types (`Bitcoin`) with `String()` method (the `Stringer` interface)
- Sentinel errors with `errors.Is`
- Test helpers for cleaner assertions

```go
wallet := Wallet{}
wallet.Deposit(Bitcoin(10))
wallet.Balance() // 10 BTC
wallet.Withdraw(Bitcoin(100)) // ErrInsufficientFunds
```

```bash
go test ./pointers/
```

### 7. Sorting — `sort_test.go` (root level)

- Implementing `sort.Interface` (Len, Swap, Less)
- Custom sort by age using `ByAge` type
- Package-level `Example()` test function

```go
people := []Person{{"Bob", 31}, {"John", 42}, {"Michael", 17}, {"Jenny", 26}}
sort.Sort(ByAge(people))
// [Michael: 17 Jenny: 26 Bob: 31 John: 42]
```

```bash
go test -v .
```

## Commands

| Command | Description |
|---------|-------------|
| `go test ./...` | Run all tests |
| `go test -v ./...` | Run all tests verbosely |
| `go test -bench=. -benchmem ./...` | Run all benchmarks |
| `go test -coverprofile=coverage.out ./...` | Generate coverage profile |
| `go tool cover -html=coverage.out` | View coverage in browser |
| `go vet ./...` | Static analysis |

## Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Package Reference

| Package | Path | Exports |
|---------|------|---------|
| `main` | `helloworld/` | `Hello(name, language string) string` |
| `integers` | `integers/` | `Add(x, y int) int` |
| `iteration` | `iteration/` | `Repeat(character string, count int) string` |
| `arrayAndSlices` | `arrays-and-slices/` | `Sum`, `SumAll`, `SumAllTails` |
| `smc` | `structs-methods-nd-interfaces/` | `Shape`, `Rectangle`, `Circle`, `Triangle`, `Perimeter`, `Area` |
| `pointers` | `pointers/` | `Bitcoin`, `Wallet`, `ErrInsufficientFunds`, `Stringer` |
| `sort_test` | `.` (root) | `Person`, `ByAge` |
