package levyfs

type PageFile struct {
	RootNode *Node
}

type Node struct {
	Key     string
	Value   string
	Children map[string]*Node
}

func NewRootNode() *Node {
	return &Node{
		Key:     "/",
		Value:   "",
		Children: make(map[string]*Node),
	}
}

func NewNode(key string, value string) *Node {
	return &Node{
		Key:     key,
		Value:   value,
		Children: make(map[string]*Node),
	}
}

func (n *Node) AddChildNode(child *Node) {
	n.Children[child.Key] = child
}
