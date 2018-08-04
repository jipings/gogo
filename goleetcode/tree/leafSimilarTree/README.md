## [LeetCode 872 Leaf-Similar Trees](https://leetcode.com/problems/leaf-similar-trees/description/)


### 题目大意
根据叶子节点从左到右的序列判断两颗二叉树是否相似


### 解题
- 思路1：先序遍历二叉树依次记录叶子节点，比较两棵树的叶子节点序列。（LeetCode 4ms）
- 思路2：利用go语言特性优化思路1，并发遍历两棵树, 在遍历的同时就比较叶子节点序列。(LeetCode 0ms)
    - receiver通过ch接收sender传过来的值
    - done指示叶子节点是否一致
```Go
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
```
- 思路3：利用栈记录深度优先遍历的路径。（C++实现LeetCode 0ms）
``` C++
     bool leafSimilar(TreeNode* root1, TreeNode* root2) {
        stack<TreeNode*> s1 , s2;
        s1.push(root1); s2.push(root2);
        while (!s1.empty() && !s2.empty())
            if (dfs(s1) != dfs(s2)) return false;
        return s1.empty() && s2.empty();
    }

    int dfs(stack<TreeNode*>& s) {
        while (true) {
            TreeNode* node = s.top(); s.pop();
            if (node->right) s.push(node->right);
            if (node->left) s.push(node->left);
            if (!node->left && !node->right) return node->val;
        }
    }
```



### 参考
1. [LeetCode Discuss O(logN)Space](https://leetcode.com/problems/leaf-similar-trees/discuss/152329/C++JavaPython-O(logN)-Space)
2. [872. Leaf-Similar Trees](https://csxuejin.gitbooks.io/leetcode/content/algorithms/872.html)
