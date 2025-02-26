package levyfs

type PageFile struct {
	RootNode *Node
}

type Node struct {
	Key     string
	Value   []byte
	Children map[string]*Node
}

func NewPageFile() *PageFile {
	return &PageFile{
		RootNode: NewRootNode(),
	}
}

func NewRootNode() *Node {
	return &Node{
		Key:     "/",
		Value:   nil,
		Children: make(map[string]*Node),
	}
}

func NewNode(key string, value string) *Node {

	valueBytes := []byte{}

	if value != "" {
		valueBytes = []byte(value)
	}

	return &Node{
		Key:     key,
		Value:   valueBytes,
		Children: make(map[string]*Node),
	}
}

func (n *Node) AddChildNode(child *Node) {
	n.Children[child.Key] = child
}

func (p *PageFile) FindNodeWithKey(key string) *Node {
	return findNodeRecursive(p.RootNode, key)
}

func findNodeRecursive(node *Node, key string) *Node {
	if node.Key == key {
		return node
	}
	for _, child := range node.Children {
		if found := findNodeRecursive(child, key); found != nil {
			return found
		}
	}
	return nil
}

func (p *PageFile) GetStringValueForNodeKey(key string) *string {
	node := p.FindNodeWithKey(key)
	if node != nil {
		var value = string(node.Value)
		return &value
	}
	return nil
}
