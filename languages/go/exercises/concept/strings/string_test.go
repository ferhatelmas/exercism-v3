package strings_concept

import "testing"

func TestMessage(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Extract message from error message",
			args: args{
				line: "[ERROR]: Stack overflow",
			},
			want: "Stack overflow",
		},
		{
			name: "Extract message from warning message",
			args: args{
				line: "[WARNING]: Disk almost full",
			},
			want: "Disk almost full",
		},
		{
			name: "Extract message from info message",
			args: args{
				line: "[INFO]: File moved",
			},
			want: "File moved",
		},
		{
			name: "Extract message without extra whitespace",
			args: args{
				line: "[WARNING]:   \tTimezone not set  \r\n",
			},
			want: "Timezone not set",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Message(tt.args.line); got != tt.want {
				t.Errorf("Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Extract log level from error",
			args: args{
				line: "[ERROR]: Disk full",
			},
			want: "error",
		},
		{
			name: "Extract log level from warning",
			args: args{
				line: "[WARNING]: Unsafe password",
			},
			want: "warning",
		},
		{
			name: "Extract log level from info",
			args: args{
				line: "[INFO]: Timezone changed",
			},
			want: "info",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogLevel(tt.args.line); got != tt.want {
				t.Errorf("LogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReformat(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Reformat error message",
			args: args{
				line: "[ERROR]: Segmentation fault",
			},
			want: "Segmentation fault (error)",
		},
		{
			name: "Reformat warning message",
			args: args{
				line: "[WARNING]: Decreased performance",
			},
			want: "Decreased performance (warning)",
		},
		{
			name: "Reformat info message",
			args: args{
				line: "[INFO]: Disk defragmented",
			},
			want: "Disk defragmented (info)",
		},
		{
			name: "Reformat message with extra whitespace",
			args: args{
				line: "[ERROR]: \t Corrupt disk\t \t \r\n",
			},
			want: "Corrupt disk (error)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reformat(tt.args.line); got != tt.want {
				t.Errorf("Reformat() = %v, want %v", got, tt.want)
			}
		})
	}
}
