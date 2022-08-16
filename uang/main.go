package main

import "fmt"

func main() {
	rs := countOfMoney(145000)
	fmt.Println("Result", rs)
}

func countOfMoney(money int64) (result []string) {
	var a, b, c, d, e, f, g, h, i, j int

	if money > 100 {
		for money >= 100000 {
			money -= 100000
			a++
		}
		for money >= 50000 {
			money -= 50000
			b++
		}

		for money >= 20000 {
			money -= 20000
			c++
		}

		for money >= 10000 {
			money -= 10000
			d++
		}

		if money >= 5000 {
			money -= 5000
			e++
		}

		for money > 2000 {
			money -= 2000
			f++
		}

		for money >= 1000 {
			money -= 1000
			g++
		}

		for money >= 500 {
			money -= 500
			h++
		}

		for money >= 200 {
			money -= 200
			i++
		}

		if money >= 100 {
			money -= 100
			j++
		} else {
			if money > 0 {
				j++
			}
		}

	}

	if a > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 100000, a)
		result = append(result, txt)
	}

	if b > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 50000, b)
		result = append(result, txt)
	}

	if c > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 20000, c)
		result = append(result, txt)
	}

	if d > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 10000, d)
		result = append(result, txt)
	}

	if e > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 5000, e)
		result = append(result, txt)
	}

	if f > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 2000, f)
		result = append(result, txt)
	}

	if g > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 1000, g)
		result = append(result, txt)
	}

	if h > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 500, h)
		result = append(result, txt)
	}

	if i > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 200, i)
		result = append(result, txt)
	}

	if j > 0 {
		txt := fmt.Sprintf("'Rp. %d': %d", 100, j)
		result = append(result, txt)
	}

	return
}
