package lzma

import (
	"bytes"
	"testing"
)

var (
	tFilterNo   = FilterNo
	tFilterYes  = FilterYes
	tFilterAuto = FilterAuto
	tLevel0     = Level0
	tLevel9     = Level9
	dictSize    = uint32(1 << 17)
)

var testCases = []struct {
	uncompressed []byte
	level        *CompressionLevel
	dictSize     *uint32
	filterMode   *FilterMode
	compressed   []byte
}{
	{[]byte("Input uncompressed string...."), nil, nil, &tFilterNo, []byte{0, 93, 0, 0, 128, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 149, 88, 155, 88, 219, 213, 115, 77, 160, 23, 15, 128, 43, 222, 80, 0}},
	{[]byte("Input uncompressed string...."), nil, nil, &tFilterYes, []byte{1, 93, 0, 0, 128, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 149, 88, 155, 88, 219, 213, 115, 77, 160, 23, 15, 128, 43, 222, 80, 0}},
	{[]byte("Input uncompressed string...."), &tLevel0, nil, &tFilterYes, []byte{1, 93, 0, 0, 128, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 148, 206, 206, 218, 241, 146, 204, 83, 215, 50, 31, 23, 148, 133, 228, 0}},
	{[]byte("Input uncompressed string...."), &tLevel9, nil, &tFilterYes, []byte{1, 93, 0, 0, 128, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 149, 88, 155, 88, 219, 213, 115, 77, 160, 23, 15, 128, 43, 222, 80, 0}},
	{[]byte("Input uncompressed string...."), nil, &dictSize, nil, []byte{0, 93, 0, 0, 2, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 149, 88, 155, 88, 219, 213, 115, 77, 160, 23, 15, 128, 43, 222, 80, 0}},
	{[]byte("Input uncompressed string...."), &tLevel0, nil, &tFilterAuto, []byte{0, 93, 0, 0, 128, 0, 29, 0, 0, 0, 0, 0, 0, 0, 0, 36, 155, 138, 7, 129, 236, 92, 204, 166, 68, 249, 86, 40, 3, 61, 148, 206, 206, 218, 241, 146, 204, 83, 215, 50, 31, 23, 148, 133, 228, 0}},
}

func TestLzma86Encode(t *testing.T) {
	output := make([]byte, 1000)

	for i := range testCases {
		tc := testCases[i]
		params := NewEncodeParameters()
		if tc.level != nil {
			params.CompressionLevel(*tc.level)
		}

		if tc.dictSize != nil {
			params.DictionarySize(*tc.dictSize)
		}

		if tc.filterMode != nil {
			params.FilterMode(*tc.filterMode)
		}

		n, err := Lzma86Encode(tc.uncompressed, output, params)
		if err != nil {
			t.Fatalf("Failed case %d: Lzma86Encode - %s\n", i, err.Error())
		}

		if n != len(tc.compressed) {
			t.Fatalf("Failed case %d: compressed size incorrect\n", i)
		}

		if !bytes.Equal(output[:n], tc.compressed) {
			t.Fatalf("Failed case %d: compressed data not equals\n", i)
		}
	}
}

func TestLzma86GetUnpackSize(t *testing.T) {
	for i := range testCases {
		tc := testCases[i]

		n, err := Lzma86GetUnpackSize(tc.compressed)
		if err != nil {
			t.Fatalf("Failed case %d: Lzma86GetUnpackSize - %s\n", i, err.Error())
		}

		if n != len(tc.uncompressed) {
			t.Fatalf("Failed case %d: uncompressed size incorrect\n", i)
		}
	}
}

func TestLzma86Decode(t *testing.T) {
	for i := range testCases {
		tc := testCases[i]

		n, err := Lzma86GetUnpackSize(tc.compressed)
		if err != nil {
			t.Fatalf("Failed case %d: GetUnpackSize - %s\n", i, err.Error())
		}

		output := make([]byte, n)
		n, err = Lzma86Decode(tc.compressed, output)
		if err != nil {
			t.Fatalf("Failed case %d: Lzma86Decode - %s\n", i, err.Error())
		}

		if n != len(tc.uncompressed) {
			t.Fatalf("Failed case %d: uncompressed size incorrect\n", i)
		}

		if !bytes.Equal(output, tc.uncompressed) {
			t.Fatalf("Failed case %d: uncompressed data not equals\n", i)
		}
	}
}
