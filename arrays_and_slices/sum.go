package main

func Sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func SumAll(numsToSum ...[]int) []int {
    var sums []int

    for _, nums := range numsToSum {
        sums = append(sums, Sum(nums))
    }
    return sums
}