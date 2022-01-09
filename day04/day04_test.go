package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestPassport_IsValid(t *testing.T) {
	type fields struct {
		BirthYear      int
		IssueYear      int
		ExpirationYear int
		Height         int
		HeightUnit     string
		HairColor      string
		EyeColor       string
		PID            string
		CID            string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "all fields present - example 1",
			fields: fields{
				BirthYear:      1937,
				IssueYear:      2017,
				ExpirationYear: 2023,
				Height:         183,
				HeightUnit:     "cm",
				HairColor:      "#fffffd",
				EyeColor:       "gry",
				PID:            "860033327",
				CID:            "147",
			},
			want: true,
		},
		{
			name: "missing height - example 2",
			fields: fields{
				IssueYear:      2013,
				EyeColor:       "amb",
				CID:            "350",
				ExpirationYear: 2023,
				PID:            "028048884",
				HairColor:      "#cfa07d",
				BirthYear:      1929,
			},
			want: false,
		},
		{
			name: "missing CID - example 3",
			fields: fields{
				HairColor:      "#ae17e1",
				IssueYear:      2013,
				ExpirationYear: 2024,
				EyeColor:       "brn",
				PID:            "760753108",
				BirthYear:      1931,
				Height:         179,
				HeightUnit:     "cm",
			},
			want: true,
		},
		{
			name: "missing borh CID and BirthYear - example 4",
			fields: fields{
				HairColor:      "#cfa07d",
				ExpirationYear: 2025,
				PID:            "166559648",
				IssueYear:      2011,
				EyeColor:       "brn",
				Height:         59,
				HeightUnit:     "in",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			p := &Passport{
				BirthYear:      tt.fields.BirthYear,
				IssueYear:      tt.fields.IssueYear,
				ExpirationYear: tt.fields.ExpirationYear,
				Height:         tt.fields.Height,
				HeightUnit:     tt.fields.HeightUnit,
				HairColor:      tt.fields.HairColor,
				EyeColor:       tt.fields.EyeColor,
				PID:            tt.fields.PID,
				CID:            tt.fields.CID,
			}
			is.Equal(p.IsValid(), tt.want)
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    *Passport
		wantErr bool
	}{
		{
			name: "example 1",
			args: args{text: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm"},
			want: &Passport{
				EyeColor:       "gry",
				PID:            "860033327",
				ExpirationYear: 2020,
				HairColor:      "#fffffd",
				BirthYear:      1937,
				IssueYear:      2017,
				CID:            "147",
				Height:         183,
				HeightUnit:     "cm",
			},
			wantErr: false,
		},
		{
			name: "example 2",
			args: args{text: "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\nhcl:#cfa07d byr:1929"},
			want: &Passport{
				EyeColor:       "amb",
				PID:            "028048884",
				ExpirationYear: 2023,
				HairColor:      "#cfa07d",
				BirthYear:      1929,
				IssueYear:      2013,
				CID:            "350",
				Height:         0,  // Height is missing
				HeightUnit:     "", // Height is missing
			},
			wantErr: false,
		},
		{
			name: "example 3",
			args: args{text: "hcl:#ae17e1 iyr:2013\neyr:2024\necl:brn pid:760753108 byr:1931\nhgt:179cm"},
			want: &Passport{
				EyeColor:       "brn",
				PID:            "760753108",
				ExpirationYear: 2024,
				HairColor:      "#ae17e1",
				BirthYear:      1931,
				IssueYear:      2013,
				CID:            "", // CID is missing
				Height:         179,
				HeightUnit:     "cm",
			},
			wantErr: false,
		},
		{
			name: "example 4",
			args: args{text: "hcl:#cfa07d eyr:2025 pid:166559648\niyr:2011 ecl:brn hgt:59in"},
			want: &Passport{
				EyeColor:       "brn",
				PID:            "166559648",
				ExpirationYear: 2025,
				HairColor:      "#cfa07d",
				BirthYear:      0, // Birth year is missing
				IssueYear:      2011,
				CID:            "", // CID is missing
				Height:         59,
				HeightUnit:     "in",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			got, err := Parse(tt.args.text)
			is.True((err != nil) == tt.wantErr)
			is.Equal(got, tt.want)
		})
	}
}

func TestPassport_IsValid2(t *testing.T) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "invalid 1",
			text: "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			want: false,
		},
		{
			name: "invalid 2",
			text: "iyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946",
			want: false,
		},
		{
			name: "invalid 3",
			text: "hcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			want: false,
		},
		{
			name: "invalid 4",
			text: "hgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007",
			want: false,
		},
		{
			name: "valid 1",
			text: "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f",
			want: true,
		},
		{
			name: "valid 2",
			text: "eyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			want: true,
		},
		{
			name: "valid 3",
			text: "hcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022",
			want: true,
		},
		{
			name: "valid 4",
			text: "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			p, err := Parse(tt.text)

			t.Logf("%#v\n", p)

			is.NoErr(err)
			is.Equal(p.IsValid2(), tt.want)
		})
	}
}
