# Table Writer

## Introduction

Table Writer is a simple utility program to print tabular data in pretty format.

## Installation

	go get github.com/prvn/table_writer

## Examples

### Initialization

```go
// Header for table
columnHeaders := []string{"Id", "Count", "Timestamp"}

// Id will not be bigger than 4 digits, for better readability, I give it 6 character width
// Count will be up to 6 digits, 9 sounds reasonable readability wise
// Timestamp is epoch, so never beyond 10 digits, 12 looks reasonable space
columnMaxWidths := []int{6, 9, 12}

tableWriter, err := NewTableWriterWithWriter(os.Stdout, headers, widths)
```

### Printing Table (All at once)

```go
row1 := []string{"1", "11", "1461910429"}
row2 := []string{"2", "12", "1461910429"}
row3 := []string{"1000", "99999", "1461910429"}
tableWriter.PrintTable([][]string{row1, row2, row3}, table_writer.AlignCenter)
```

will print table like

```
+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
|  1   |   11    | 1461910429 |
|  2   |   12    | 1461910429 |
| 1000 | 999999  | 1461910429 |
+------+---------+------------+
```

### Printing Table (Stream data)

Initialize table writer as mentioned in [Initialization](#Initialization)

```go
// Print header
tableWriter.PrintHeader()
row1 := []string{"1", "11", "1461910429"}
row2 := []string{"2", "12", "1461910429"}
row3 := []string{"1000", "99999", "1461910429"}
rows := [][]string{row1, row2, row3}

// Add rows one by one
for _, r := range rows {
	tableWriter.PrintRow(r, table_writer.AlignCenter)
}

// OR Add all rows at once
tableWriter.PrintRows(rows, table_writer.AlignCenter)

// Print footer
tableWriter.PrintFooter()
```
Will print table like
```
+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
|  1   |   11    | 1461910429 |
|  2   |   12    | 1461910429 |
| 1000 | 999999  | 1461910429 |
+------+---------+------------+
```

### Print Table &amp; Summarize

Follow steps from above and add following code

```go
summaryRow1 := "Id's: 3"
summaryRow2 := "Count: 1000022"
tableWriter.PrintRowAsOneColumn(summaryRow1, table_writer.AlignRight)
tableWriter.PrintRowAsOneColumn(summaryRow2, table_writer.AlignRight)
tableWriter.PrintFooter()
```

Will print table like
```
+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
|  1   |   11    | 1461910429 |
|  2   |   12    | 1461910429 |
| 1000 | 999999  | 1461910429 |
+------+---------+------------+
|Id's: 3                      |
|Count: 1000022               |
+------+---------+------------+
```
