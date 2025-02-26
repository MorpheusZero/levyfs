package levyfs

import (
	"fmt"
	"testing"
)

func TestNodeCreation(t *testing.T) {
	rootNode := NewRootNode()

	if rootNode.Key != "/" {
		t.Errorf("Expected Key to be '/', got %s", rootNode.Key)
	}

	if rootNode.Value != nil {
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

	if string(root.Children["child"].Value) != "childValue" {
		t.Errorf("Expected child Value to be 'childValue', got %s", root.Children["child"].Value)
	}
}

func TestAddingMultipleChildNodes(t *testing.T) {
	root := NewRootNode()

	children := make([]struct {
		key   string
		value string
	}, 1000)

	for i := 0; i < 1000; i++ {
		children[i] = struct {
			key   string
			value string
		}{
			key:   fmt.Sprintf("child%d", i+1),
			value: fmt.Sprintf("childValue%d", i+1),
		}
	}

	for _, child := range children {
		root.AddChildNode(NewNode(child.key, child.value))
	}

	if len(root.Children) != len(children) {
		t.Errorf("Expected %d children, got %d", len(children), len(root.Children))
	}

	for _, child := range children {
		if string(root.Children[child.key].Value) != child.value {
			t.Errorf("Expected child Value to be '%s', got %s", child.value, root.Children[child.key].Value)
		}
	}
}

func TestPageFileCreation(t *testing.T) {
	pageFile := NewPageFile()

	if pageFile.RootNode.Key != "/" {
		t.Errorf("Expected Key to be '/', got %s", pageFile.RootNode.Key)
	}

	if pageFile.RootNode.Value != nil {
		t.Errorf("Expected Key to be 'EMPTY', got %s", pageFile.RootNode.Key)
	}

	if len(pageFile.RootNode.Children) != 0 {
		t.Errorf("Expected Children to be empty, got %d", len(pageFile.RootNode.Children))
	}
}

func TestFindNodeWithKey(t *testing.T) {
	pageFile := NewPageFile()

	pageFile.RootNode.AddChildNode(NewNode("child1", "childValue1"))
	pageFile.RootNode.AddChildNode(NewNode("child2", "childValue2"))
	pageFile.RootNode.AddChildNode(NewNode("child3", "childValue3"))

	foundNode := pageFile.FindNodeWithKey("child2")

	if foundNode == nil {
		t.Errorf("Expected to find a node, got nil")
	}

	if string(foundNode.Value) != "childValue2" {
		t.Errorf("Expected Value to be 'childValue2', got %s", foundNode.Value)
	}
}

func TestGetValueForNodeKey(t *testing.T) {
	pageFile := NewPageFile()

	pageFile.RootNode.AddChildNode(NewNode("child1", "childValue1"))
	pageFile.RootNode.AddChildNode(NewNode("child2", "childValue2"))
	pageFile.RootNode.AddChildNode(NewNode("child3", "childValue3"))

	value := pageFile.GetStringValueForNodeKey("child2")

	if value == nil {
		t.Errorf("Expected to find a value, got nil")
	}

	if *value != "childValue2" {
		t.Errorf("Expected Value to be 'childValue2', got %s", *value)
	}
}
