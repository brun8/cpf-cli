package generator

import "testing"

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
}
