# Convert Binary Search Tree to Sorted Doubly Linked List


Time complexity: O(N). Each node is processed exactly once.
Space complexity: O(N). We have to keep a recursion stack of the size of the tree height 

We use Depth First Search In Order Traversal.
  * inOrder(node.Left) node.Val inOrder(node.Right)

Here is the algorithm :

* Initiate the first and the last nodes as nulls.
* Call the standard inorder recursion helper(root) :
    * If node is not null :
      * Call the recursion for the left subtree helper(node.left).
      * If the last node is not null, link the last and the current node nodes.
      * Else initiate the first node.
      * Mark the current node as the last one : last = node.
      * Call the recursion for the right subtree helper(node.right).
      * Link the first and the last nodes to close DLL ring and then return the first node.





