package maps

type HashMap struct {
	Key      string
	Value    string
	hashCode int
	next     *HashMap
}

var table [16](*HashMap)

func initTable() {
	for i := range table {
		table[i] = &HashMap{"", "", i, nil}
	}
}

func getInstance() [16](*HashMap) {
	if table[0] == nil {
		initTable()
	}
	return table
}

func genHashCode(k string) int {
	if len(k) == 0 {
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i])) * 31
	}
	return hashCode
}

func (hashMap HashMap) indexTable(hashCode int) int {
	return hashCode % 16
}

func indexNode(hashCode int) int {
	return hashCode >> 4
}

func (hashMap HashMap) Put(k string, v string) string {
	var hashCode = genHashCode(k)
	var thisNode = HashMap{k, v, hashCode, nil}
	var tableIndex = hashMap.indexTable(hashCode)
	var nodeIndex = indexNode(hashCode)
	var headPtr [16](*HashMap) = getInstance()
	var headNode = headPtr[tableIndex]

	if (*headNode).Key == "" {
		*headNode = thisNode
		return ""
	}

	var lastNode *HashMap = headNode
	var nextNode *HashMap = (*headNode).next

	for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex) {
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue string = lastNode.Value
		lastNode.Value = thisNode.Value
		return oldValue
	}
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func (hashMap HashMap) Get(k string) string {
	var hashCode = genHashCode(k)
	var tableIndex = hashMap.indexTable(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var node *HashMap = headPtr[tableIndex]

	if (*node).Key == k {
		return (*node).Value
	}

	for (*node).next != nil {
		if k == (*node).Key {
			return (*node).Value
		}
		node = (*node).next
	}
	return ""
}
