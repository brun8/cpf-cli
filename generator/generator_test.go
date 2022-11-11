package generator

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateCpfLength(t *testing.T) {
	unpunctuated := GenerateCpf(false, -1)
	if len(unpunctuated) != 11 {
		t.Errorf("wrong cpf length. got=%d, want=%d", len(unpunctuated), 11)
	}

	punctuated := GenerateCpf(true, -1)
	if len(punctuated) != 14 {
		t.Errorf("wrong cpf length. got=%d, want=%d", len(punctuated), 14)
	}
}

func TestGenerateWithRegion(t *testing.T) {
	cpf := GenerateCpf(false, 1)

	digit := cpf[8]
	if digit != '1' {
		t.Errorf("invalid cpf region. got=%c, want=%c", digit, '1')
	}

	validateCpf(t, cpf)
}

func TestCpfIsValid(t *testing.T) {
	unpunctuated := GenerateCpf(false, -1)
	punctuated := GenerateCpf(true, -1)

	validateCpf(t, unpunctuated)
	validateCpf(t, punctuated)
}

func validateCpf(t *testing.T, cpf string) {
	cpf = normalizeCpf(cpf)

	var sum1, sum2 int
	for i := 0; i < 10; i++ {
		num, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			t.Error("invalid characters on cpf")
			return
		}

		// pula o ultimo digito pro calculo do primeiro digito verificador
		if i != 9 {
			sum1 += (10 - i) * num
		}

		// pula o primeiro digito pro calculo do segundo digito verificador
		if i != 0 {
			sum2 += (10 - (i - 1)) * num
		}
	}

	v1 := generateDigit(sum1)
	v2 := generateDigit(sum2)

	if strconv.Itoa(v1) != string(cpf[9]) || strconv.Itoa(v2) != string(cpf[10]) {
		t.Error("invalid cpf")
	}
}

func normalizeCpf(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	return strings.ReplaceAll(cpf, "-", "")
}
