package levyfs

import (
	"testing"
)

func TestNodeCreation(t *testing.T) {
	rootNode := NewRootNode()

	if rootNode.Key != "/" {
		t.Errorf("Expected Key to be '/', got %s", rootNode.Key)
	}

	if rootNode.Value != "" {
		t.Errorf("Expected Key to be 'EMPTY', got %s", rootNode.Key)
	}

	if len(rootNode.Children) != 0 {
		t.Errorf("Expected Children to be empty, got %d", len(rootNode.Children))
	}
}

func TestAddChildNode(t *testing.T) {
	root := NewRootNode()

	child := NewNode("child", "childValue")

	root.AddChildNode(child)

	if len(root.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(root.Children))
	}

	if root.Children["child"].Value != "childValue" {
		t.Errorf("Expected child Value to be 'childValue', got %s", root.Children["child"].Value)
	}
}
