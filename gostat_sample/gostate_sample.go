package main

import (
	"fmt"

	"github.com/ematvey/gostat"
)

func showConfidenceLimit(imp int64, click int64, alpha float64) {
	ctr := float64(click) / float64(imp)
	lowerCL, upperCL := stat.Binom_p_ConfI(imp, ctr, alpha)

	fmt.Printf("IMP:%12d, Click:%10d, CTR:%.4f, ALPHA:%.2f, LowerConfidenceLimit:%.f, UpperConfidenceLimit:%f\n", imp, click, ctr, alpha, lowerCL, upperCL)
}

func main() {
	showConfidenceLimit(100, 1, 0.95)
	showConfidenceLimit(1000, 10, 0.95)
	showConfidenceLimit(10000, 100, 0.95)
	showConfidenceLimit(100000, 1000, 0.95)
	showConfidenceLimit(1000000, 10000, 0.95)
	showConfidenceLimit(10000000, 100000, 0.95)
	showConfidenceLimit(100000000, 1000000, 0.95)
}
