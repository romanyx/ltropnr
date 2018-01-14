package ltropnr

import (
	"bytes"
	"testing"
)

const msg = `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

func Test_templateData_GetHeader(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{
			name: "subject",
			key:  "Subject",
			want: "Gophers at Gophercon",
		},
		{
			name: "date",
			key:  "Date",
			want: "Mon, 23 Jun 2015 11:40:36 -0400",
		},
		{
			name: "cc",
			key:  "Cc",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, _ := newTemplateData(bytes.NewBufferString(msg))
			if got := tmpl.GetHeader(tt.key); got != tt.want {
				t.Errorf("templateData.GetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
