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

	for i := 0; i < 8; i++ {
		num := r.Int() % 10
		cpf += strconv.Itoa(num)
	}

	var num int
	if region == -1 {
		num = r.Int() % 10
	} else {
		num = region
	}
	cpf += strconv.Itoa(num)

  // primeiro digito de verificação
  cpf += generateDigit(cpf[:9])
  // segundo digito de verificação
  cpf += generateDigit(cpf[1:10])

	if punctuated {
		return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:])
	}

	return cpf
}

func generateDigit(digits string) string {
  var sum int
  for i := 0; i < 9; i++ {
    num, _ := strconv.Atoi(string(digits[i]))
    //fmt.Printf("%d * %d\n", 10-i, num)

    sum += (10-i) * num
  }

  rest := sum % 11
  if rest < 2 {
    return "0"
  }

  return strconv.Itoa(11-rest)
}
