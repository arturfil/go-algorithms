package integers

import "math"

func reverse(x int) int {
    reversed := 0

    for x != 0 {
        digit := x % 10
        reversed = (reversed) * 10 + digit
        x = x / 10 // because it's an int so decimal is truncated
    }

    outOfRange := reversed > math.MaxInt32 || math.MinInt32 > reversed
    if outOfRange { return 0 }

    return reversed
}
