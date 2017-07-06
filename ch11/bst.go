// 二叉搜索树

package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type BST struct {
	root *Node
}

// 二叉树插入操作
func (this *BST) insert(value int) {
	node := this.root
	newNode := Node{value: value}

	if node == nil {
		this.root = &newNode
		return
	}

	for {
		if node.value > value {
			if node.left == nil {
				node.left = &newNode
				break
			} else {
				node = node.left
			}
		} else if node.value < value {
			if node.right == nil {
				node.right = &newNode
				break
			} else {
				node = node.right
			}
		} else {
			break
		}
	}
}

// 查找节点,返回节点及其父节点
func (this *BST) find(value int) (*Node, *Node) {
	var parentNode *Node
	node := this.root

	if node == nil {
		return nil, nil
	}

	for {
		if value < node.value {
			if node.left == nil {
				return nil, nil
			}
			parentNode = node
			node = node.left
		} else if value > node.value {
			if node.right == nil {
				return nil, nil
			}
			parentNode = node
			node = node.right
		} else {
			return node, parentNode
		}
	}
}

// 二叉树删除操作
func (this *BST) remove(val int) bool {
	node, parentNode := this.find(val)
	leftFlag := node.left != nil
	rightFlag := node.right != nil
	if node != nil {
		if !leftFlag && !rightFlag {
			// 没有子节点
			if node.value > parentNode.value {
				parentNode.right = nil
			} else {
				parentNode.left = nil
			}
			return true
		} else if leftFlag && !rightFlag {
			// 左节点存在
			if node.value > parentNode.value {
				parentNode.right = node.left
			} else {
				parentNode.left = node.left
			}
			return true
		} else if !leftFlag && rightFlag {
			// 右节点存在
			if node.value > parentNode.value {
				parentNode.right = node.right
			} else {
				parentNode.left = node.right
			}
			return true
		} else {
			// 左右节点均存在
			newNode := node
			for {
				newNode = newNode.right
				if newNode.right == nil {
					break
				}
			}

			this.remove(newNode.value)
			if node.value > parentNode.value {
				parentNode.right = newNode
			} else {
				parentNode.left = newNode
			}

			newNode.left = node.left
			newNode.right = node.right
			return true
		}
	}

	return false
}

// 中序遍历
func (this *BST) show(node *Node) {
	if node == nil {
		return
	}
	this.show(node.left)
	fmt.Printf("%d\t", node.value)
	this.show(node.right)
}



func main() {
	tree := BST{}
	tree.insert(4)
	tree.insert(2)
	tree.insert(5)
	tree.insert(1)
	tree.insert(3)
	tree.insert(6)
	// tree.remove(6)
	// tree.remove(5)
	tree.remove(2)
	tree.show(tree.root)
}