package trie

import "testing"

func TestTrieInsert(t *testing.T) {
	words := []string{"a", "b", "c", "ab", "abc"}
	trie := New()
	for _, word := range words {
		trie.Insert(word)
	}

	for _, char := range []byte{'a', 'b', 'c'} {
		if !trie.root.Conns[char].IsWord {
			t.Errorf("%b was not found in trie", char)
		}
	}
	for _, char := range []byte{'d', 'e', 'A'} {
		if trie.root.Conns[char] != nil {
			t.Errorf("%b was not supposed to be in trie", char)
		}
	}
}

func TestTrieSearch(t *testing.T) {
	words := []string{"a", "b", "c", "ab", "abc"}
	trie := New()
	for _, word := range words {
		trie.Insert(word)
	}

	if !trie.SearchWord("") {
		t.Errorf("empty string did not return a postive search in trie")
	}

	for _, word := range words {
		if !trie.SearchWord(word) {
			t.Errorf("word: %s could not be found in trie", word)
		}
	}

	invalidWords := []string{"ABC", "AB", "abC", "A"}
	for _, word := range invalidWords {
		if trie.SearchWord(word) {
			t.Errorf("word: %s was not supposed to be found in trie", word)
		}
	}
}

func TestTriePrefix(t *testing.T) {
	words := []string{"abc", "aab", "aaa", "aaaa"}
	trie := New()
	for _, word := range words {
		trie.Insert(word)
	}

	validPrefix := []string{"aa", "ab", "aaa"}
	for _, prefix := range validPrefix {
		if !trie.IsPrefix(prefix) {
			t.Errorf("prefix: %s did not return true", prefix)
		}
	}

	invalidPrefix := []string{"A", "Aaa", "bc"}
	for _, prefix := range invalidPrefix {
		if trie.IsPrefix(prefix) {
			t.Errorf("prefix: %s did return true", prefix)
		}
	}
}

func TestTrieGetSuggestions(t *testing.T) {
	words := []string{"aab", "aaa", "aab", "aabcd", "aac", "aabd"}
	trie := New()
	for _, word := range words {
		trie.Insert(word)
	}

	autoCompletedWords := trie.GetSuggestions("aab")
	expected := map[string]bool{
		"aab":   true,
		"aabcd": true,
		"aabd":  true,
	}

	for _, word := range autoCompletedWords {
		if !expected[word] {
			t.Errorf("word %s was not supposed to be found in suggestions", word)
		}
	}
}
