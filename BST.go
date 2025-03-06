package main

import (
	"fmt"
)

type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

type BinarySearchTree struct {
    Root *TreeNode
}

// Insert вставляет новое значение в дерево
func (bst *BinarySearchTree) Insert(value int) {
    newNode := &TreeNode{Value: value}
    if bst.Root == nil {
        bst.Root = newNode
    } else {
        bst.Root.insert(newNode)
    }
}

// Вспомогательная функция вставки
func (node *TreeNode) insert(newNode *TreeNode) {
    if newNode.Value < node.Value {
        if node.Left == nil {
            node.Left = newNode
        } else {
            node.Left.insert(newNode)
        }
    } else {
        if node.Right == nil {
            node.Right = newNode
        } else {
            node.Right.insert(newNode)
        }
    }
}

// Search ищет значение в дереве и возвращает true, если оно найдено
func (bst *BinarySearchTree) Search(value int) bool {
    return bst.Root != nil && bst.Root.search(value)
}

// Вспомогательная функция поиска
func (node *TreeNode) search(value int) bool {
    if node == nil {
        return false
    }
    if value < node.Value {
        return node.Left.search(value)
    } else if value > node.Value {
        return node.Right.search(value)
    } else {
        return true
    }
}

// Remove удаляет значение из дерева
func (bst *BinarySearchTree) Remove(value int) {
	bst.Root = bst.Root.remove(value)
}

// Вспомогательная функция удаления
func (node *TreeNode) remove(value int) *TreeNode {
	if node == nil {
        return nil
    }

    if value < node.Value {
        node.Left = node.Left.remove(value)
    } else if value > node.Value {
        node.Right = node.Right.remove(value)
    } else {
        if node.Left == nil {
            return node.Right
        } else if node.Right == nil {
            return node.Left
        }

        minNode := node.Right.findMinNode()
        node.Value = minNode.Value
        node.Right = node.Right.remove(minNode.Value)
    }

    return node
}

// Найти минимальное значение в дереве (используется при удалении)
func (node *TreeNode) findMinNode() *TreeNode {
	if node.Left == nil {
        return node
    }
    return node.Left.findMinNode()
}

// Прямой обход (Preorder) - Корень, Лево, Право
func (bst *BinarySearchTree) PreorderTraversal() {
    preorder(bst.Root)
    fmt.Println()
}

func preorder(node *TreeNode) {
    if node != nil {
        fmt.Printf("%d ", node.Value)
        preorder(node.Left)
        preorder(node.Right)
    }
}

// Обратный обход (Postorder) - Лево, Право, Корень
func (bst *BinarySearchTree) PostorderTraversal() {
    postorder(bst.Root)
    fmt.Println()
}

func postorder(node *TreeNode) {
    if node != nil {
        postorder(node.Left)
        postorder(node.Right)
        fmt.Printf("%d ", node.Value)
    }
}

// Центрированный обход (Inorder) - Лево, Корень, Право
func (bst *BinarySearchTree) InorderTraversal() {
    inorder(bst.Root)
    fmt.Println()
}

func inorder(node *TreeNode) {
    if node != nil {
        inorder(node.Left)
        fmt.Printf("%d ", node.Value)
        inorder(node.Right)
    }
}

func main() {
    bst := &BinarySearchTree{}
    bst.Insert(50)
    bst.Insert(30)
    bst.Insert(20)
    bst.Insert(40)
    bst.Insert(70)
    bst.Insert(60)
    bst.Insert(80)
    bst.Insert(80)
    bst.Insert(80)
    bst.Insert(80)
	bst.PostorderTraversal()
    fmt.Println("Search for 40:", bst.Search(40))
    fmt.Println("Search for 100:", bst.Search(100))
}
