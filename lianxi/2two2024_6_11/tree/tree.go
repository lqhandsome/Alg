package tree

import log "github.com/sirupsen/logrus"

// tree 二叉查找树
type tree struct {
	data  int
	left  *tree
	right *tree
}

func NewTree() *tree {
	return new(tree)
}

//不论先中后怎么遍历，每个节点被遍历到的顺序和次数都是一样的，
//先自己再左节点再右节点再自己，只是选择打印的时机不同

// PreRange 先序遍历
func PreRange(root *tree) {
	if root == nil {
		return
	}
	// 用切片模拟栈
	var stack []*tree
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			log.Info(root.data)
			root = root.left
		} else {
			if len(stack) > 0 {
				root = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				root = root.right
			}
		}
	}
}

// MidRange 中序遍历
func MidRange(root *tree) {
	if root == nil {
		return
	}
	// 用切片模拟栈
	var stack []*tree
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			log.Info(root.data)
			stack = stack[:len(stack)-1]
			root = root.right
		}
	}
}

// AfterRange 后续遍历
// 如果知道左右子树遍历完了，或者没有左右节点
func AfterRange(root *tree) {
	if root == nil {
		return
	}
	// 用切片模拟栈
	var stack []*tree
	var pre *tree
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.right == pre || root.right == nil {
				log.Info(root.data)
				pre = root
				root = nil
			} else {
				stack = append(stack, root)
				root = root.right
			}
		}
	}
}

// PreRangeRec 先序遍历递归方式
func (t *tree) PreRangeRec() {
	if t == nil {
		return
	}

	log.Info(t.data)
	t.left.PreRangeRec()
	t.right.PreRangeRec()
}

// MidRangeRec 中序遍历递归方式
func (t *tree) MidRangeRec() {
	if t == nil {
		return
	}

	t.left.MidRangeRec()
	log.Info(t.data)
	t.right.MidRangeRec()
}

// AfterRangeRec 后续遍历递归方式
func (t *tree) AfterRangeRec() {
	if t == nil {
		return
	}

	t.left.AfterRangeRec()
	t.right.AfterRangeRec()
	log.Info(t.data)

}

func demo(root *tree) {
	var stack []*tree
	var pre *tree
	for root != nil || len(stack) != 0 {
		// 将所有左节点压入
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		// 取出一个节点遍历右子树
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 如果右子树刚刚访问过或者右子树为空
			if root.right == pre || root.right == nil {
				log.Info(root.data)
				pre = root
				root = nil
			} else {
				// 重新压入等待右子树遍历完成时弹出
				stack = append(stack, root)
				root = root.right
			}
		}
	}
}

func demoA(root *tree) {
	var stack []*tree
	for root != nil || len(stack) != 0 {
		// 将所有左节点压入
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		// 取出一个节点遍历右子树
		if len(stack) != 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			log.Info(root.data)
			root = root.right
		}
	}
}
