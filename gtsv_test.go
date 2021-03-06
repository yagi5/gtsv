package gtsv

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int
		hasError bool
	}{
		{
			name: "int",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
		},
		{
			name: "int but contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int{[]int{1, 2, 3}, []int{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int

			for gr.Next() {
				rowCnt++
				var line []int
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint
		hasError bool
	}{
		{
			name: "uint",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint{[]uint{1, 2, 3}, []uint{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint

			for gr.Next() {
				rowCnt++
				var line []uint
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int8
		hasError bool
	}{
		{
			name: "int8",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t128\n",
			row:      2,
			col:      3,
			result:   [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int8{[]int8{1, 2, 3}, []int8{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int8

			for gr.Next() {
				rowCnt++
				var line []int8
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int8())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint8
		hasError bool
	}{
		{
			name: "uint8",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t256\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint8{[]uint8{1, 2, 3}, []uint8{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint8

			for gr.Next() {
				rowCnt++
				var line []uint8
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint8())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int16
		hasError bool
	}{
		{
			name: "int16",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t32768\n",
			row:      2,
			col:      3,
			result:   [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\t65536\n",
			row:      2,
			col:      3,
			result:   [][]int16{[]int16{1, 2, 3}, []int16{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int16

			for gr.Next() {
				rowCnt++
				var line []int16
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int16())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint16
		hasError bool
	}{
		{
			name: "uint16",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t65536\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint16{[]uint16{1, 2, 3}, []uint16{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint16

			for gr.Next() {
				rowCnt++
				var line []uint16
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint16())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int32
		hasError bool
	}{
		{
			name: "int32",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t2147483648\n",
			row:      2,
			col:      3,
			result:   [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int32{[]int32{1, 2, 3}, []int32{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int32

			for gr.Next() {
				rowCnt++
				var line []int32
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int32())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint32
		hasError bool
	}{
		{
			name: "uint32",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t4294967296\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint32{[]uint32{1, 2, 3}, []uint32{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint32

			for gr.Next() {
				rowCnt++
				var line []uint32
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint32())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int64
		hasError bool
	}{
		{
			name: "int64",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 6}},
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t9223372036854775808\n",
			row:      2,
			col:      3,
			result:   [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]int64{[]int64{1, 2, 3}, []int64{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int64

			for gr.Next() {
				rowCnt++
				var line []int64
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int64())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]uint64
		hasError bool
	}{
		{
			name: "uint64",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:    2,
			col:    3,
			result: [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 6}},
		},
		{
			name: "contains negative number",
			tsv: "1\t2\t3\n" +
				"4\t5\t-1\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains too large number",
			tsv: "1\t2\t3\n" +
				"4\t5\t18446744073709551616\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
		{
			name: "contains string",
			tsv: "1\t2\t3\n" +
				"4\t5\ta\n",
			row:      2,
			col:      3,
			result:   [][]uint64{[]uint64{1, 2, 3}, []uint64{4, 5, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]uint64

			for gr.Next() {
				rowCnt++
				var line []uint64
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Uint64())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]float32
		hasError bool
	}{
		{
			name: "float32",
			tsv: "1.0\t2.1\t3.2\n" +
				"1.234567\t1.2345678\t1.23456789\n",
			row:    2,
			col:    3,
			result: [][]float32{[]float32{1.0, 2.1, 3.2}, []float32{1.234567, 1.2345678, 1.2345679}},
		},
		{
			name: "contains string",
			tsv: "1.0\t2.1\t3.2\n" +
				"1.234567\t1.2345678\ta\n",
			row:      2,
			col:      3,
			result:   [][]float32{[]float32{1.0, 2.1, 3.2}, []float32{1.234567, 1.2345678, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]float32

			for gr.Next() {
				rowCnt++
				var line []float32
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Float32())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]float64
		hasError bool
	}{
		{
			name: "float64",
			tsv: "1.0\t2.1\t3.2\n" +
				"1.234567\t1.2345678\t1.23456789012345\n",
			row:    2,
			col:    3,
			result: [][]float64{[]float64{1.0, 2.1, 3.2}, []float64{1.234567, 1.2345678, 1.23456789012345}},
		},
		{
			name: "contains string",
			tsv: "1.0\t2.1\t3.2\n" +
				"1.234567\t1.2345678\ta\n",
			row:      2,
			col:      3,
			result:   [][]float64{[]float64{1.0, 2.1, 3.2}, []float64{1.234567, 1.2345678, 0}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]float64

			for gr.Next() {
				rowCnt++
				var line []float64
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Float64())
				}
				ret = append(ret, line)
			}
			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][][]byte
		hasError bool
	}{
		{
			name: "bytes",
			tsv: "aaa\tbbb\tccc\n" +
				"ddd\teee\tfff\n",
			row: 2,
			col: 3,
			result: [][][]byte{[][]byte{[]byte{97, 97, 97}, []byte{98, 98, 98}, []byte{99, 99, 99}},
				[][]byte{[]byte{100, 100, 100}, []byte{101, 101, 101}, []byte{102, 102, 102}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][][]byte

			for gr.Next() {
				rowCnt++
				var line [][]byte
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Bytes())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}

}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]string
		hasError bool
	}{
		{
			name: "string",
			tsv: "aaa\tbbb\tccc\n" +
				"ddd\teee\tfff\n",
			row:    2,
			col:    3,
			result: [][]string{[]string{"aaa", "bbb", "ccc"}, []string{"ddd", "eee", "fff"}},
		},
		{
			name: "contains unescaped value",
			tsv: "a \\b a\tb\\tb\n" +
				"d\\nd\te\\e\n",
			row: 2,
			col: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]string

			for gr.Next() {
				rowCnt++
				var line []string
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.String())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if tt.result != nil && !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]bool
		hasError bool
	}{
		{
			name: "bool",
			tsv: "1\tt\tT\tTRUE\ttrue\tTrue\n" +
				"0\tf\tF\tFALSE\tfalse\tFalse\n", // https://golang.org/pkg/strconv/#ParseBool
			row:    2,
			col:    6,
			result: [][]bool{[]bool{true, true, true, true, true, true}, []bool{false, false, false, false, false, false}},
		},
		{
			name:     "invalid as bool",
			tsv:      "a\n",
			row:      1,
			col:      1,
			result:   [][]bool{[]bool{false}},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]bool

			for gr.Next() {
				rowCnt++
				var line []bool
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Bool())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			if tt.row != rowCnt {
				t.Fatalf("row check failed expected: %d, actual: %d", tt.row, rowCnt)
			}

			if !reflect.DeepEqual(tt.result, ret) {
				t.Fatalf("returned value check failed expected: %v, actual: %v", tt.result, ret)
			}
		})
	}
}

func TestReadError(t *testing.T) {
	tests := []struct {
		name   string
		tsv    string
		row    int
		col    int
		result [][]int
		error  Error
	}{
		{
			name: "error",
			tsv: "1\t2\t3\n" +
				"4\ta\t6\n",
			row:    2,
			col:    3,
			result: [][]int{[]int{1, 2, 3}, []int{4, 0, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int

			for gr.Next() {
				rowCnt++
				var line []int
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if err == nil || err == io.EOF {
				t.Fatalf("invalid error %s", err)
			}

			er, ok := err.(Error)
			if !ok {
				t.Fatalf("invalid error %s", er)
			}
			if er.Row() != 2 || er.Col() != 2 {
				t.Fatalf("invalid error tracer row: %d, col: %d", er.Row(), er.Col())
			}
		})
	}
}

func TestInvalidReadError(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		row      int
		col      int
		result   [][]int
		hasError bool
		errRow   int
		errCol   int
	}{
		{
			name: "error",
			tsv: "1\t2\t3\n" +
				"4\t5\t6\n",
			row:      2,
			col:      2, // invalid
			result:   [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			hasError: true,
			errRow:   1,
			errCol:   3,
		},
		{
			name: "error",
			tsv: "1\t2\t3" + // \n is missing
				"4\t5\t6\n",
			row:      2,
			col:      3,
			result:   [][]int{[]int{1, 2, 3}, []int{4, 5, 6}},
			hasError: true,
			errRow:   1,
			errCol:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gr := New(bytes.NewBufferString(tt.tsv))
			var rowCnt int
			var ret [][]int

			for gr.Next() {
				rowCnt++
				var line []int
				for i := 0; i < tt.col; i++ {
					line = append(line, gr.Int())
				}
				ret = append(ret, line)
			}

			err := gr.Error()
			if (err != nil) != tt.hasError && err != io.EOF {
				t.Fatalf("error is not io.EOF but %s", err)
			}

			er, ok := err.(Error)
			if !ok {
				t.Fatalf("invalid error %s", er)
			}
			if er.Row() != tt.errRow || er.Col() != tt.errCol {
				t.Fatalf("invalid error tracer row: %d, col: %d", er.Row(), er.Col())
			}
		})
	}
}

func TestError(t *testing.T) {
	e := gtsverror{row: 1, col: 2}

	if e.Row() != 1 {
		t.Errorf("invalid error row")
	}

	if e.Col() != 2 {
		t.Errorf("invalid error col")
	}

	errmsg := "Parse failed at row #1, col #2"
	if e.Error() != errmsg {
		t.Errorf("invalid error message")
	}
}

func TestReader(t *testing.T) {
	gr := New(bytes.NewBufferString(""))
	gr.Int()
	err := gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Int8()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Int16()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Int32()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Int64()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Uint()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Uint8()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Uint16()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Uint32()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}

	gr = New(bytes.NewBufferString(""))
	gr.Uint64()
	err = gr.Error()
	if err == io.EOF {
		t.Fatalf("error was io.EOF but %s", err)
	}
}
