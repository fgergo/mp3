package mp3

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkDecode(t *testing.B) {
	r := MakeSilence()
	d := NewDecoder(r)
	for i := 0; i < t.N; i++ {
		f := Frame{}
		d.Decode(&f)
	}
}

func ExampleDecoder_Decode() {
	r, err := os.Open("file.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}

	d := NewDecoder(r)
	var f Frame
	for {

		if err := d.Decode(&f); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(&f)
	}
}
