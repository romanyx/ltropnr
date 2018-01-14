package ltropnr

import (
	"html/template"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Opener
	}{
		{
			name: "default",
			args: args{
				options: []Option{},
			},
			want: &Opener{
				cmdName: cmdName(),
				layout:  template.Must(template.New("layout").Parse(defaultTemplate)),
			},
		},
		{
			name: "light mode",
			args: args{
				options: []Option{
					LightMode(),
				},
			},
			want: &Opener{
				cmdName: cmdName(),
				layout:  template.Must(template.New("layout").Parse(lightTemplate)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
