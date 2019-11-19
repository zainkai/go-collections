package trie

type node struct {
	key    byte
	conns  map[byte]*node
	isWord bool
}

// Trie implementation
type Trie struct {
	root *node
}

// New create a Trie
func New() *Trie {
	root := &node{
		key:    byte(0),
		conns:  make(map[byte]*node),
		isWord: true,
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
		nextNode, ok := curNode.conns[char]
		if !ok {
			nextNode = &node{
				key:    char,
				conns:  make(map[byte]*node),
				isWord: false,
			}
			curNode.conns[char] = nextNode
		}
		curNode = nextNode
	}

	curNode.isWord = true
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
	return n != nil && n.isWord
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

func (t *Trie) searchTrie(word []byte) *node {
	curNode := t.root
	for _, char := range word {
		nextNode, ok := curNode.conns[char]
		if !ok {
			return nil
		}
		curNode = nextNode
	}
	return curNode
}

func (t *Trie) GetSuggestions(word []byte) []string {
	curNode := t.searchTrie(word)
	if curNode == nil {
		return []string{}
	}

	return []string{}
}
