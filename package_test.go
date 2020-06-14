package main

import "testing"

// TestMessage ensure the message function works. by implication
// this also tests isPositive
func TestMessage(t *testing.T) {
	tables := []struct {
		answer string
		reply  string
	}{
		{"yes", "Glad to hear it mate!"},
		{"Y", "Glad to hear it mate!"},
		{"AYE", "Glad to hear it mate!"},
		{"yES", "Glad to hear it mate!"},
		{"ayE", "Glad to hear it mate!"},
		{"Indeed", "Glad to hear it mate!"},
		{"No", "I'm sorry about that!"},
		{"saUSage", "I'm sorry about that!"},
		{"inded", "I'm sorry about that!"},
		{"Yeah no", "I'm sorry about that!"},
	}
	for _, table := range tables {
		funcReply := message(table.answer)
		if funcReply != table.reply {
			t.Errorf("Message of (%s) was incorrect, got: %s, want: %s.", table.answer, funcReply, table.reply)
		}
	}
}

// TestBytesToString make sure we get what we came for
func TestBytesToString(t *testing.T) {
	tables := []struct {
		bytestring []byte
		reply      string
	}{
		{[]byte{'H', 'E', 'L', 'P'}, "HELP"},
		{[]byte{'s', 'a', 'u', 's', 'a', 'g', 'e'}, "sausage"},
		{[]byte{'1', '2'}, "12"},
	}
	for _, table := range tables {
		funcReply := bytesToString(table.bytestring)
		if funcReply != table.reply {
			t.Errorf("Message of (%s) was incorrect, got: %s, want: %s.", table.bytestring, funcReply, table.reply)
		}
	}
}
