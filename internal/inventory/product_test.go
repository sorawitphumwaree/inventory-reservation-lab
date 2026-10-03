package inventory

import "testing"

func TestCreateProductInputValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateProductInput
		wantErr bool
	}{
		{
			name:    "positive stock",
			input:   CreateProductInput{Name: "Keyboard", InitialStock: 10},
			wantErr: false,
		},
		{
			name:    "zero stock",
			input:   CreateProductInput{Name: "Keyboard", InitialStock: 0},
			wantErr: false,
		},
		{
			name:    "empty name",
			input:   CreateProductInput{Name: "", InitialStock: 10},
			wantErr: true,
		},
		{
			name:    "whitespace name",
			input:   CreateProductInput{Name: "   ", InitialStock: 10},
			wantErr: true,
		},
		{
			name:    "negative stock",
			input:   CreateProductInput{Name: "Keyboard", InitialStock: -1},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.input.Validate()
			gotErr := err != nil

			if gotErr != test.wantErr {
				t.Fatalf("Validate() error = %v, want error = %t",
					err, test.wantErr)
			}
		})
	}
}
