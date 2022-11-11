package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCpf(punctuated bool, region int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cpf := ""
	sum := 0
	sum2 := 0

	for i := 0; i < 8; i++ {
		num := r.Int() % 10
		cpf += strconv.Itoa(num)

		sum += (10 - i) * num
	}

	var num int
	if region == -1 {
		num = r.Int() % 10
	} else {
		num = region
	}

	cpf += strconv.Itoa(num)
	sum += 2 * num

	cpf += strconv.Itoa(generateDigit(sum))

	// 2 loops pro primeiro nao ficar ilegivel
	for i := 1; i < 10; i++ {
		mult := 10 - (i - 1)
		num, _ := strconv.Atoi(string(cpf[i]))
		sum2 += mult * num
	}
	cpf += strconv.Itoa(generateDigit(sum2))

	if punctuated {
		return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])
	}
	return cpf
}

func generateDigit(sum int) int {
	res := sum % 11
	if res < 2 {
		return 0
	} else {
		return 11 - res
	}
}
