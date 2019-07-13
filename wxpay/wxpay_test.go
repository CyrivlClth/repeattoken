package wxpay

import (
	"testing"

	"github.com/CyrivlClth/repeattoken/token"
	"github.com/stretchr/testify/assert"
)

func TestDigest_Generate(t *testing.T) {
	data := token.Data{
		"mch_id": "asdf",
		"appid":  123,
		"a":      1.4,
		"sa":     nil,
		"vvv":    "",
		"111":    make([]string, 0),
	}
	d := new(digest)

	z, err := d.Generate("asdf", data)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "B55ED7A8104741DB428B633C3A47816C", z)
}

func TestDigest_Verify(t *testing.T) {
	datas := []token.Data{
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "B55ED7A8104741DB428B633C3A47816C",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "B55ED7A8104741DB428B633C3A47816D",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   nil,
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   123,
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
		},
	}
	expect := []bool{
		true,
		false,
		false,
		false,
		false,
		false,
	}
	for i := 0; i < len(datas); i++ {
		d := new(digest)
		assert.Equal(t, expect[i], d.Verify("asdf", datas[i]))
	}
}

func TestFastDigest_Verify(t *testing.T) {
	datas := []token.Data{
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "B55ED7A8104741DB428B633C3A47816C",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "B55ED7A8104741DB428B633C3A47816D",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "",
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   nil,
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   123,
		},
		{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
		},
	}
	expect := []bool{
		true,
		false,
		false,
		false,
		false,
		false,
	}
	for i := 0; i < len(datas); i++ {
		d := new(digest)
		assert.Equal(t, expect[i], d.Verify("asdf", datas[i]))
	}
}

func BenchmarkDigest_Verify(b *testing.B) {
	data := token.Data{
		"mch_id": "asdf",
		"appid":  123,
		"a":      1.4,
		"sa":     nil,
		"vvv":    "",
		"111":    make([]string, 0),
		"sign":   "B55ED7A8104741DB428B633C3A47816C",
	}
	d := new(digest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.Verify("asdf", data)
	}
}

func BenchmarkFastDigest_Verify(b *testing.B) {
	d := new(fastDigest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := token.Data{
			"mch_id": "asdf",
			"appid":  123,
			"a":      1.4,
			"sa":     nil,
			"vvv":    "",
			"111":    make([]string, 0),
			"sign":   "B55ED7A8104741DB428B633C3A47816C",
		}
		_ = d.Verify("asdf", data)
	}
}

func BenchmarkDigest_Generate(b *testing.B) {
	data := token.Data{
		"mch_id": "asdf",
		"appid":  123,
		"a":      1.4,
		"sa":     nil,
		"vvv":    "",
		"111":    make([]string, 0),
	}
	d := new(digest)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = d.Generate("asdf", data)
	}
}
