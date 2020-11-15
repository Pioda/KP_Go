package lexer

import "testing"

func TestNext(t *testing.T) {
	lexer := New("A & B")
	if lexer.NextToken() != "A" {
		t.Error("expected A")
	}
	c := lexer.NextToken()
	if  c != "&" {
		t.Error("expected & but was ",c)
	}
	b := lexer.NextToken()
	if b != "B" {
		t.Error("expected B but was ", b)
	}
}
