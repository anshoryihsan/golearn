package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

////bachmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("ihsan")
	}
}
func BenchmarkHelloWorldAnshory1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("anshory")
	}
}

////end bachmark

////tabel test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "ihsan",
			request:  "ihsan",
			expected: "hello ihsan",
		}, {
			name:     "anshory",
			request:  "anshory",
			expected: "hello anshory",
		}, {
			name:     "wasis",
			request:  "wasis",
			expected: "hello wasis",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

////end tabel test
////tes sub tes
///uuntuk runing bisa go test -v -run=TestSubtes/ihsan *-run=namafunctiontest/namasubtest
func TestSubTest(t *testing.T) {
	t.Run("ihsan", func(t *testing.T) {
		result := HelloWorld("ihsan")
		require.Equal(t, "hello ihsan", result, "result must be 'hello ihsan'")
	})
	t.Run("anshory", func(t *testing.T) {
		result := HelloWorld("anshory")
		require.Equal(t, "hello anshory", result, "result must be 'hello anshory'")
	})
}

////end tes sub tes

//// test before /a after
//// berlaku hanya di satu package saja
func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")
	m.Run()
	//after
	fmt.Println("After unit test")
}

//// end test before /a after

//test skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { //goos (go oprating system) //darwin -> macOs
		t.Skip("unit test tidak bisa jalan di Mac")
	}
}

//end test skip

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("ihsan")
	if result != "hello ihsan" {
		//error
		// t.Fail() //t.Fail() akan melanjutkan ke test selanjutnya meskipun menemukan fail
		t.Error("result must be 'hello ihsan'") //memanggil t.fatal()
	}
	fmt.Println("test hello world")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("ihsan")
	assert.Equal(t, "hello ihsan", result, "result must be 'hello ihsan'")
	fmt.Println("testHellowWorld with Assert Done")
}

func TestHelloWorld1(t *testing.T) {
	result := HelloWorld("anshory")
	if result != "hello anshory" {
		//error
		// t.FailNow() //t,FailNow() mengeluarkan pesan error tetapi proses selanjutnya tidak akan di eksekusi fmt.println
		t.Fatal("result must be 'hello anshory'") //memanggil t.failnow()
	}
	fmt.Println("test hello world 1")
}

func TestHelloWorldRequire(t *testing.T) {
	//require memanggil t.failnow()
	result := HelloWorld("ihsan")
	require.Equal(t, "hello ihsan", result, "result must be 'hello ihsan'")
	fmt.Println("testHellowWorld with Require Done")
}
