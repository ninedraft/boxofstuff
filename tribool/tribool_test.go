package tribool

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStringer(test *testing.T) {
	var testCases = []struct {
		In       triboolImpl
		Expected string
		Panic    bool
	}{
		{True, "true", false},
		{False, "false", false},
		{Inderminate, "inderminate", false},
		{triboolImpl(123), "", true},
		{0, "false", false},
	}

	for i, testCase := range testCases {
		func() {
			defer func() {
				if err := recover(); err != nil && !testCase.Panic {
					test.Fatal(err)
				} else if err == nil && testCase.Panic {
					test.Errorf("expected panic on case %d %#v", i, testCase)
				}
			}()
			var got = testCase.In.String()
			if got != testCase.Expected {
				test.Errorf("test case %d: expected %q, got %q", i, testCase.Expected, got)
			}
		}()
	}
}

func Test_Tribool_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		tribool triboolImpl
		want    []byte
		wantErr bool
	}{
		{
			tribool: True,
			want:    []byte(`{"tribool_value":"true"}`),
		},
		{
			tribool: False,
			want:    []byte(`{"tribool_value":"false"}`),
		},
		{
			tribool: Inderminate,
			want:    []byte(`{"tribool_value":"inderminate"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.tribool)
			if (err != nil) != tt.wantErr {
				t.Errorf("_Tribool.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_Tribool.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tribool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		args     []byte
		expected triboolImpl
		wantErr  bool
	}{

		{
			args:     []byte(`{"tribool_value": "true"}`),
			expected: True,
		},
		{
			args:     []byte(`{"tribool_value": "false"} `),
			expected: False,
		},
		{
			args:     []byte(`{"tribool_value": "inderminate"} `),
			expected: Inderminate,
		},
		{
			args:    []byte(`asd as e`),
			wantErr: true,
		},
		{
			args:    []byte(`{"tribool_value": "doot"} `),
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr triboolImpl
			if err := json.Unmarshal(tt.args, &tr); (err != nil) != tt.wantErr {
				t.Errorf("test №%d _Tribool.UnmarshalJSON() error = %q", i, err)
			} else if tr != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, tr)
			}
		})
	}
}

func Test_triboolImpl_IsTrue(t *testing.T) {
	tests := []struct {
		name    string
		tribool triboolImpl
		want    bool
	}{
		{
			tribool: True,
			want:    true,
		},
		{
			tribool: False,
			want:    false,
		},
		{
			tribool: Inderminate,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tribool.IsTrue(); got != tt.want {
				t.Errorf("triboolImpl.IsTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_triboolImpl_IsFalse(t *testing.T) {
	tests := []struct {
		name    string
		tribool triboolImpl
		want    bool
	}{
		{
			tribool: True,
			want:    false,
		},
		{
			tribool: False,
			want:    true,
		},
		{
			tribool: Inderminate,
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tribool.IsFalse(); got != tt.want {
				t.Errorf("triboolImpl.IsFalse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_triboolImpl_IsInderminate(t *testing.T) {
	tests := []struct {
		name    string
		tribool triboolImpl
		want    bool
	}{
		{
			tribool: True,
			want:    false,
		},
		{
			tribool: False,
			want:    false,
		},
		{
			tribool: Inderminate,
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tribool.IsInderminate(); got != tt.want {
				t.Errorf("triboolImpl.IsInderminate() = %v, want %v", got, tt.want)
			}
		})
	}
}
