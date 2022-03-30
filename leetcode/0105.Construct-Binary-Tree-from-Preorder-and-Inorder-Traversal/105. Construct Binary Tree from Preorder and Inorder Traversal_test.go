package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question105 struct {
	para105
	ans105
}

// para 是参数
// one 代表第一个参数
type para105 struct {
	preorder []int
	inorder  []int
}

// ans 是答案
// one 代表第一个答案
type ans105 struct {
	one []int
}

func Test_Problem105(t *testing.T) {

	qs := []question105{

		{
			para105{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			ans105{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 105------------------------\n")

	for _, q := range qs {
		_, p := q.ans105, q.para105
		fmt.Printf("【input】:%v      ", p)
		tree2 := buildTree2(p.preorder, p.inorder)
		fmt.Print(tree2)
		//fmt.Printf("【output】:%v      \n", structures.Tree2ints())
	}
	fmt.Printf("\n\n\n")
}
