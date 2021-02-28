package tire

type Trie struct {
	isWord bool
	next   []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.next[v-'a'] == nil {
			t.next[v-'a'] = new(Trie)
		}
		t = t.next[v-'a']
	}
	t.isWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.next[v-'a'] == nil {
			return false
		} else {
			t = t.next[v-'a']
		}
	}
	return t.isWord == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if t.next[v-'a'] == nil {
			return false
		} else {
			t = t.next[v-'a']
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
