package trie

import "github.com/zainkai/go-collections/queue"

type TrieNode struct {
	Key    byte
	Conns  map[byte]*TrieNode
	IsWord bool
	Word   string
}

// Trie implementation
type Trie struct {
	root *TrieNode
}

// New create a Trie
func New() *Trie {
	root := &TrieNode{
		Key:    byte(0),
		Conns:  make(map[byte]*TrieNode),
		IsWord: true,
	}

	return &Trie{root}
}

func toBytes(s string) []byte {
	return []byte(s)
}

func (t *Trie) Insert(word string) {
	wordArray := toBytes(word)
	t.InsertBytes(wordArray)
}

// InsertBytes adds new word to trie
// O(k)
func (t *Trie) InsertBytes(word []byte) {
	curNode := t.root

	for _, char := range word {
		nextNode, ok := curNode.Conns[char]
		if !ok {
			nextNode = &TrieNode{
				Key:    char,
				Conns:  make(map[byte]*TrieNode),
				IsWord: false,
				Word:   string(word),
			}
			curNode.Conns[char] = nextNode
		}
		curNode = nextNode
	}

	curNode.IsWord = true
}

// SearchWord search if full word was inserted into trie
// O(k)
func (t *Trie) SearchWord(word string) bool {
	wordArray := toBytes(word)
	return t.SearchBytes(wordArray)
}

// SearchBytes search if full word was inserted into trie
// O(k)
func (t *Trie) SearchBytes(word []byte) bool {
	n := t.searchTrie(word)
	return n != nil && n.IsWord
}

// IsPrefix search if prefix path exists in trie
// O(k)
func (t *Trie) IsPrefix(word string) bool {
	wordArray := toBytes(word)
	return t.IsPrefixBytes(wordArray)
}

// IsPrefixBytes search if prefix path exists in trie
// O(k)
func (t *Trie) IsPrefixBytes(word []byte) bool {
	return t.searchTrie(word) != nil
}

func (t *Trie) searchTrie(word []byte) *TrieNode {
	curNode := t.root
	for _, char := range word {
		nextNode, ok := curNode.Conns[char]
		if !ok {
			return nil
		}
		curNode = nextNode
	}
	return curNode
}

// GetSuggestionsBytes from a word look for all valid connecting words in trie
// O(K + N)
func (t *Trie) GetSuggestionsBytes(word []byte) []string {
	foundNode := t.searchTrie(word)
	connectedWords := []string{}
	if foundNode == nil {
		return connectedWords
	}

	queue := queue.New()
	queue.Enqueue(foundNode)

	for queue.Length > 0 {
		curNode := queue.Dequeue().(*TrieNode)
		if curNode.IsWord {
			connectedWords = append(connectedWords, curNode.Word)
		}
		for _, conn := range curNode.Conns {
			queue.Enqueue(conn)
		}
	}

	return connectedWords
}

// GetSuggestions from a word look for all valid connecting words in trie
// O(K + N)
func (t *Trie) GetSuggestions(word string) []string {
	wordArray := toBytes(word)
	return t.GetSuggestionsBytes(wordArray)
}
