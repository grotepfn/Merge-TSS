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
This code runs in O(n log n). This is due to the sort function. An out-of-place-sort-function would fasten it up to O(n) as 
the step after, the actual merging, only takes O(n). This sorting function would need O(n) space however.

## Robustness 
This code accepts all valid inputs and exits otherwise.

## Size
The needed program space increases linearly with input length.
