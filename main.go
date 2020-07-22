package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {

	if root == nil {
		return root
	}

	var first, last *Node

	//function declaration
	var dfs func(node *Node)

	//function implementation
	// dfs - in order
	// inorder(node.Left)
	// node.val
	// inorder(node.Right)
	dfs = func(node *Node) {

		if node != nil {
			dfs(node.Left)

			if last != nil {
				// draw last node and current node connections
				last.Right = node
				node.Left = last
			} else {
				// if there is no last node, you are dealing with the very first node!
				first = node
			}
			// all done, now the last node is the current node
			last = node

			dfs(node.Right)
		}
	}

	//function call
	dfs(root)

	// making it circular
	last.Right = first
	first.Left = last

	return first
}
