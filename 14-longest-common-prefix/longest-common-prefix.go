const NUM_ALPHABET = 26

type TrieNode struct {
	Childs    [NUM_ALPHABET]*TrieNode
	EndOfWord bool
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(s string) {
	t.insert(t.Root, s)
}

func (t *Trie) insert(currNode *TrieNode, s string) bool {
	if s == "" {
		return true
	}
	idx := s[0] - 'a'
	if currNode.Childs[s[0]-'a'] == nil {
		currNode.Childs[idx] = &TrieNode{}
	}
	currNode.EndOfWord = t.insert(currNode.Childs[idx], s[1:])
	return false
}

func (t *Trie) HasCommonPrefix(s string) bool {
	return t.hasCommonPrefix(t.Root, s)
}

func (t *Trie) hasCommonPrefix(currNode *TrieNode, prefix string) bool {
	if prefix == "" {
		return true
	}
	idx := prefix[0] - 'a'
	if currNode.Childs[idx] == nil {
		return false
	}
	for i := 0; i < NUM_ALPHABET; i++ {
		if byte(i) == idx {
			continue
		}
		if currNode.Childs[i] != nil {
			return false
		}
	}
	return t.hasCommonPrefix(currNode.Childs[idx], prefix[1:])
}

func longestCommonPrefix(strs []string) string {
	shortestStr, shortestLen := strs[0], len(strs[0])
	trie := &Trie{Root: &TrieNode{}}
	for _, str := range strs {
		if len(str) < shortestLen {
			shortestLen = len(str)
			shortestStr = str
		}
		trie.Insert(str)
	}

	var res string
	for i := 1; i <= shortestLen; i++ {
		if trie.HasCommonPrefix(shortestStr[:i]) {
			res = shortestStr[:i]
		} else {
			break
		}
	}
	return res
}