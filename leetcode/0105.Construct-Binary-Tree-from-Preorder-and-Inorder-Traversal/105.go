package leetcode


func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val:   preorder[0],
	}

	leftLength := 0
	for i, v := range inorder {
		if v == preorder[0] {
			leftLength = i+1
			break
		}
	}

	leftPre := preorder[1:leftLength]
	rightPre := preorder[leftLength:]

	leftIn := inorder[0:leftLength-1]
	rightIn := inorder[leftLength:]

	node.Left = buildTree2(leftPre, leftIn)
	node.Right = buildTree2(rightPre, rightIn)
	return node
}