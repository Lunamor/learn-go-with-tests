package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {

    t.Run("collection of any size", func(t *testing.T) {
        nums := []int{1, 2, 3}    
        
        got := Sum(nums)
        want := 6
       if got != want {
            t.Errorf("want %d given %v, but got %q", want, nums, got)
        }
    })
}

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1, 2}, []int{0, 9})
    want := []int{3, 9}

    if !reflect.DeepEqual(want, got) {
        t.Errorf("want %v but got %v", want, got)
    }
}