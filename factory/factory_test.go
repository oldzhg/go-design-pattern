package factory

import (
	"reflect"
	"testing"
)

func Test_getGun(t *testing.T) {
	type args struct {
		gunType string
	}
	tests := []struct {
		name    string
		args    args
		want    IGun
		wantErr bool
	}{
		{
			name:    "ak47",
			args:    args{gunType: "ak47"},
			want:    newAk47(),
			wantErr: false,
		},
		{
			name:    "musket",
			args:    args{gunType: "musket"},
			want:    newMusket(),
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{gunType: "error"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGun(tt.args.gunType)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGun() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGun_getName(t *testing.T) {
	type fields struct {
		name  string
		power int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "ak47",
			fields: fields{name: "ak47"},
			want:   "ak47",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gun{
				name:  tt.fields.name,
				power: tt.fields.power,
			}
			if got := g.getName(); got != tt.want {
				t.Errorf("getName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGun_getPower(t *testing.T) {
	type fields struct {
		name  string
		power int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"ak47",
			fields{power: 4},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gun{
				name:  tt.fields.name,
				power: tt.fields.power,
			}
			if got := g.getPower(); got != tt.want {
				t.Errorf("getPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGun_setName(t *testing.T) {
	type fields struct {
		name  string
		power int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"ak47",
			fields{"old", 4},
			args{"ak47"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gun{
				name:  tt.fields.name,
				power: tt.fields.power,
			}
			g.setName(tt.args.name)
			if g.name != tt.args.name {
				t.Errorf("name set fail")
			}
		})
	}
}

func TestGun_setPower(t *testing.T) {
	type fields struct {
		name  string
		power int
	}
	type args struct {
		power int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"ak47",
			fields{"ak47", 2},
			args{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gun{
				name:  tt.fields.name,
				power: tt.fields.power,
			}
			g.setPower(tt.args.power)
			if g.power != tt.args.power {
				t.Errorf("power set fail")
			}
		})
	}
}
