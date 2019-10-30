package bst

// New Initializes BST
func New(less, isEqual CompareKeys) *BST {
	return &BST{
		Head:       nil,
		LessThan:   less,
		isKeyEqual: isEqual,
	}
}

func (tree *BST) Insert(key, data interface{}) error {
	newNode := &NodeBST{key, data, nil, nil}
	if tree.Head == nil {
		tree.Head = newNode
		return nil
	}

	temp := tree.Head
	for {
		if tree.isKeyEqual(newNode.Key, temp.Key) {
			return ErrDuplicateNode
		}

		lessThanTemp := tree.LessThan(newNode.Key, temp.Key)
		if temp.Left == nil && lessThanTemp {
			temp.Left = newNode
			return nil
		} else if temp.Left != nil && lessThanTemp {
			temp = temp.Left
		} else if temp.Right == nil && !lessThanTemp {
			temp.Right = newNode
			return nil
		} else {
			temp = temp.Right
		}
	}
}

func (tree *BST) SearchNode(key interface{}) (*NodeBST, error) {
	temp := tree.Head
	for temp != nil {
		if tree.isKeyEqual(key, temp.Key) {
			return temp, nil
		}

		lettThanTemp := tree.LessThan(key, temp.Key)
		if lettThanTemp {
			temp = temp.Left
		} else {
			temp = temp.Right
		}
	}
	return nil, ErrNotFoundNode
}

func (tree *BST) Search(key interface{}) (interface{}, error) {
	node, err := tree.SearchNode(key)
	if err != nil {
		return nil, err
	}
	return node.Data, nil
}

func (root *NodeBST) SubTreeMin() *NodeBST {
	temp := root

	for temp != nil {
		if temp.Left == nil {
			break
		}
		temp = temp.Left
	}
	return temp
}

func (tree *BST) FindMin() (key interface{}, data interface{}) {
	min := tree.Head.SubTreeMin()
	if min == nil {
		return nil, nil
	}

	return min.Key, min.Data
}

func (root *NodeBST) SubTreeMax() *NodeBST {
	temp := root

	for temp != nil {
		if temp.Right == nil {
			break
		}
		temp = temp.Right
	}
	return temp
}

func (tree *BST) FindMax() (key interface{}, data interface{}) {
	max := tree.Head.SubTreeMax()
	if max == nil {
		return nil, nil
	}

	return max.Key, max.Data
}

// func (tree *BST) Delete(key interface{}) (*NodeBST, error) {
// 	temp, err := tree.Search(key)
// 	if err != nil || temp == nil {
// 		return nil, err
// 	}

// 	if temp.Left == nil && temp.Right == nil {
// 		temp = nil
// 	} else if temp.Left == nil {
// 		temp = temp.Right
// 	} else if temp.Right == nil {
// 		temp = temp.Left
// 	} else {
// 		replacement := Min(temp.Right)
// 		temp.Key, temp.Data = replacement.Key, replacement.Data

// 		temp.Right, err = tree.Delete(temp.Right.Key)
// 		if err != nil {
// 			return temp, err
// 		}
// 	}

// 	return temp, nil
// }
