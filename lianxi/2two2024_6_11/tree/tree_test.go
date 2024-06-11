package tree

import "testing"

var root *tree

func TestMain(m *testing.M) {
	InitTree()
	m.Run()
}

func InitTree() {
	t1 := &tree{
		data:  1,
		left:  nil,
		right: nil,
	}

	t2 := &tree{
		data:  2,
		left:  nil,
		right: nil,
	}

	t3 := &tree{
		data:  3,
		left:  nil,
		right: nil,
	}

	t4 := &tree{
		data:  4,
		left:  nil,
		right: nil,
	}
	t5 := &tree{
		data:  5,
		left:  nil,
		right: nil,
	}

	t6 := &tree{
		data:  6,
		left:  nil,
		right: nil,
	}

	t2.left = t1
	t2.right = t3
	t6.left = t5

	t4.left = t2
	t4.right = t6
	root = t4
}

func TestTree_PreRangeRec(t *testing.T) {
	root.PreRangeRec()
}

func TestTree_MidRangeRec(t *testing.T) {
	root.MidRangeRec()
}

func TestTree_AfterRangeRec(t *testing.T) {
	root.AfterRangeRec()
}

func TestTree_PreRange(t *testing.T) {
	PreRange(root)
}

func TestAfterRange(t *testing.T) {
	AfterRange(root)
}

func TestDemo(t *testing.T) {
	demo(root)
}

func TestDemoA(t *testing.T) {
	demoA(root)
}
