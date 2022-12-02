package bootstrap

import "testing"

func TestValidate(t *testing.T) {
	t.Parallel()

	type arg struct {
		name    string
		data    Config
		wantErr bool
	}

	tests := []arg{
		{
			name: "positive case",
			data: Config{
				HTTPPort: "1",
				DB: DBConfig{
					Scheme:   "1",
					Host:     "1",
					Port:     "1",
					Name:     "1",
					Username: "1",
					Password: "1",
				},
			},
			wantErr: false,
		},
		{
			name:    "all fields empty",
			data:    Config{},
			wantErr: true,
		},
		{
			name: "part fields empty",
			data: Config{
				HTTPPort: "1",
				DB: DBConfig{
					Scheme: "1",
					Port:   "1",
				},
			},
			wantErr: true,
		},
	}

	for _, obj := range tests {
		obj := obj

		t.Run(obj.name, func(t *testing.T) {
			t.Parallel()

			err := obj.data.Validate()
			if err != nil && !obj.wantErr {
				t.Errorf(
					"expected: nil\ngot: %s", err,
				)
			}

			if err == nil && obj.wantErr {
				t.Errorf(
					"expected: %s\ngot: %s", ErrValidate, err,
				)
			}
		})
	}
}
