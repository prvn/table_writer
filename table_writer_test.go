package table_writer

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	headers = []string{"Id", "Count", "Timestamp"}
	widths  = []int{6, 9, 12}
)

func TestInit(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
}

func TestPrintHeader(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	tw.PrintHeader()
	header := `+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
`
	assert.Equal(t, header, buf.String())
}

func TestPrintFooter(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	tw.PrintFooter()
	footer := `+------+---------+------------+
`
	assert.Equal(t, footer, buf.String())
}

func TestInvalidInit(t *testing.T) {
	wrongWidths := []int{6, 9, 8}
	tw, err := NewTableWriter(headers, wrongWidths)
	assert.NotNil(t, err)
	assert.Nil(t, tw)
}

func TestPrintRow(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	row := []string{"12", "10", "1461910429"}
	tw.PrintRow(row, AlignCenter)
	formatted := `|  12  |   10    | 1461910429 |
`
	assert.Equal(t, formatted, buf.String())
}

func TestPrintRows(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	row := []string{"12", "10", "1461910429"}
	assert.Equal(t, "|12|10|1461910429", tw.repeatWithValue("|", row))
}

func TestPrintTable(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	row1 := []string{"1", "11", "1461910429"}
	row2 := []string{"2", "12", "1461910429"}
	row3 := []string{"1000", "999999", "1461910429"}
	tw.PrintTable([][]string{row1, row2, row3}, AlignCenter)
	expected := `+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
|  1   |   11    | 1461910429 |
|  2   |   12    | 1461910429 |
| 1000 | 999999  | 1461910429 |
+------+---------+------------+
`
	assert.Equal(t, expected, buf.String())
}

func TestInvalidRow(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	row := []string{"1234567", "12", "1461910429"}
	err = tw.PrintRow(row, AlignRight)
	assert.NotNil(t, err)
}

func TestSummaryTable(t *testing.T) {
	var buf bytes.Buffer
	tw, err := NewTableWriterWithWriter(&buf, headers, widths)
	assert.NotNil(t, tw)
	assert.Nil(t, err)
	row1 := []string{"1", "11", "1461910429"}
	row2 := []string{"2", "12", "1461910429"}
	row3 := []string{"1000", "999999", "1461910429"}
	tw.PrintTable([][]string{row1, row2, row3}, AlignCenter)
	summaryRow1 := `Id's: 3`
	summaryRow2 := `Count: 1000022`
	tw.PrintRowAsOneColumn(summaryRow1, AlignRight)
	tw.PrintRowAsOneColumn(summaryRow2, AlignRight)
	tw.PrintFooter()
	expected := `+------+---------+------------+
|  Id  |  Count  | Timestamp  |
+------+---------+------------+
|  1   |   11    | 1461910429 |
|  2   |   12    | 1461910429 |
| 1000 | 999999  | 1461910429 |
+------+---------+------------+
|Id's: 3                      |
|Count: 1000022               |
+------+---------+------------+
`
	assert.Equal(t, expected, buf.String())
}