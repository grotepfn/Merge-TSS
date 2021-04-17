# Merge function

## Usage

Run the program with the default values defined in the input slice
```bash
go run merge.go
```
Run with custom boundaries with providing arrays that contain two integer numbers
```bash
go run merge.go [value1, value2] [value1, value2]
```

Test the program with
```bash
go test
```
## Complexity
This code runs in O(n log n). This is due to the sort function. The actual merging step only need O(n) time. An out-of-place-sort-function could therefore speed this up.

## Robustness 
This code accepts all valid inputs and exits otherwise.

## Size
The needed program space increases linearly with input length.
