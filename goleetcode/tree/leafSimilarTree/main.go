package tree

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		preOrderSender(root1, ch)
	}()
	go func() {
		preOrderReceiver(root2, ch, done)
		done <- true
	}()
	return <-done

}

// receiver
func preOrderReceiver(root *TreeNode, ch <-chan int, done chan bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		v := <-ch
		if root.Val != v {
			done <- false
		}
	}
	preOrderReceiver(root.Left, ch, done)
	preOrderReceiver(root.Right, ch, done)

}

// sender
func preOrderSender(root *TreeNode, ch chan<- int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		ch <- root.Val
	}
	preOrderSender(root.Left, ch)
	preOrderSender(root.Right, ch)

}
