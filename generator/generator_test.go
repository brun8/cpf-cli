package generator

import "testing"

func TestGenerateCpfLength(t *testing.T) {
  unpunctuated := GenerateCpf(false)
  if len(unpunctuated) != 11 {
    t.Errorf("wrong cpf length. got=%d, want=%d", len(unpunctuated), 11)
  }
  punctuated := GenerateCpf(true)
  if len(punctuated) != 14 {
    t.Errorf("wrong cpf length. got=%d, want=%d", len(punctuated), 14)
  }
}
