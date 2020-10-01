package formula

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gookit/color"
)

func TestFormula_Run(t *testing.T) {
	type fields struct {
		FirstEntry  string
		FirstExit   string
		SecondEntry string
		SecondExit  string
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
	}{
		{
			name: "Run with success",
			fields: fields{
				FirstEntry:  "",
				FirstExit:   "",
				SecondEntry: "",
				SecondExit:  "",
			},
			wantWriter: func() string {
				return fmt.Sprintf("Hello world!\n") +
					color.FgGreen.Render("You receive Hello in text.\n") +
					color.FgRed.Render("You receive World in list.\n") +
					color.FgYellow.Render("You receive true in boolean.\n")
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Formula{
				FirstEntry:  tt.fields.FirstEntry,
				FirstExit:   tt.fields.FirstExit,
				SecondEntry: tt.fields.SecondEntry,
				SecondExit:  tt.fields.SecondExit,
			}

			writer := &bytes.Buffer{}
			f.Run(writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Run() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
