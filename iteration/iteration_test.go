package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
    got := Repeat("a", 5)
    want:= "aaaaa"
    
    if got != want {
        t.Errorf("want %q but got %q", want, got)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i:= 0; i< b.N; i++ {
        Repeat("a", 5)
    }
}

func ExampleRepeat() {
    repeated := Repeat("z", 7)
    fmt.Println(repeated)
    // Output: zzzzzzz
}

