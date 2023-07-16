package generator

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCpf(punctuated bool, region int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0)

	for i := 0; i < 8; i++ {
		num := r.Int() % 10
		nums[i] = num
	}

	var num int
	if region == -1 {
		num = r.Int() % 10
	} else {
		num = region
	}
	nums[9] = num
	//cpf += strconv.Itoa(num)

	// primeiro digito de verificação
	generateDigit(nums[:9])
	// segundo digito de verificação
	generateDigit(nums[1:10])

  cpf := convertToString(nums)

	if punctuated {
		return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])
	}

	return cpf
}

func convertToString(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
    char := strconv.Itoa(nums[i])
    res += char
	}

  return res

}

func generateDigit(digits []int) string {
	var sum int
	for i := 0; i < 9; i++ {
		num := digits[i]
		sum += (10 - i) * num
	}

	rest := sum % 11
	if rest < 2 {
		return "0"
	}

	return strconv.Itoa(11 - rest)
}
