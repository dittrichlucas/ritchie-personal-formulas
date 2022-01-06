package formula

import (
	"bytes"
	"testing"

	"github.com/gookit/color"
)

func TestFormula_Run(t *testing.T) {
	type fields struct {
		Number1 string
		Number2 string
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
	}{
		{
			name: "Run with success",
			fields: fields{
				Number1: "1",
				Number2: "2",
			},
			wantWriter: func() string {
				return color.FgGreen.Render("Your sum is: 3\n")
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Formula{
				Number1: tt.fields.Number1,
				Number2: tt.fields.Number2,
			}
			writer := &bytes.Buffer{}
			h.Run(writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Run() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
