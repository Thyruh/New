package main

func adder() func(int) int {
    var total_count int = 0
    return func(num int) int {
            total_count += num
            return total_count
    }
}

// WATCH THE COMMAS
