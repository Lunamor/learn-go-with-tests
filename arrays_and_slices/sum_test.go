package main

import "testing"

func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t Testing.T) {
        nums := [5]int{1, 2, 3, 4, 5}
        got := Sum(nums)
        want := 15

        if got != want {
            t.Errorf("want %d given %v, but got %q", want, nums, got)
        }
    })

    t.Run("collection of any size", func(t *testing.T) {
        nums := []int{1, 2, 3}    
        
        got := Sum(nums)
        want := 6

       if got != want {
            t.Errorf("want %d given %v, but got %q", want, nums, got)
        }
    })
}