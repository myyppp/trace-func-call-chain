package generator_test

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/myyppp/functrace/pkg/generator"
)

var update = flag.Bool("update", true, "update .golden files")

func TestRewrite(t *testing.T) {
	tests := []struct {
		goldFile string
		srcFile  string
	}{
		{
			goldFile: "no_import.golden",
			srcFile:  "no_import.go",
		},
		{
			goldFile: "with_import_no_trace.golden",
			srcFile:  "with_import_no_trace.go",
		},
		{
			goldFile: "with_import_with_trace.golden",
			srcFile:  "with_import_with_trace.go",
		},
	}

	for _, tt := range tests {
		golden := filepath.Join("testdata", tt.goldFile)
		got, err := generator.Rewrite(filepath.Join("testdata", tt.srcFile))
		if err != nil {
			t.Fatalf("rewrite failed: %v\n", err)
		}
		if *update {
			os.WriteFile(golden, got, 0644)
		}
		want, err := os.ReadFile(golden)
		if err != nil {
			t.Fatalf("open file %s failed: %v", tt.goldFile, err)
		}
		if !bytes.Equal(got, want) {
			t.Errorf("want %s, got %s", string(want), string(got))
		}
	}
}