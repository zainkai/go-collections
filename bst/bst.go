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
		} else if lessThanTemp := tree.LessThan(key, temp.Key); lessThanTemp {
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

func (root *NodeBST) SubTreeMin() (*NodeBST, *NodeBST) {
	var temp, parent *NodeBST = root, nil

	for temp != nil {
		if temp.Left == nil {
			break
		}
		parent, temp = temp, temp.Left
	}
	return temp, parent
}

func (tree *BST) FindMin() (key interface{}, data interface{}) {
	min, _ := tree.Head.SubTreeMin()
	if min == nil {
		return nil, nil
	}

	return min.Key, min.Data
}

func (root *NodeBST) SubTreeMax() (*NodeBST, *NodeBST) {
	var temp, parent *NodeBST = root, nil

	for temp != nil {
		if temp.Right == nil {
			break
		}
		parent, temp = temp, temp.Right
	}
	return temp, parent
}

func (tree *BST) FindMax() (key interface{}, data interface{}) {
	max, _ := tree.Head.SubTreeMax()
	if max == nil {
		return nil, nil
	}

	return max.Key, max.Data
}

func (tree *BST) Delete(key interface{}) error {
	if tree.Head == nil {
		return ErrNotFoundNode
	}

	return tree.Head.DeleteNode(tree, key)
}

func (n *NodeBST) replaceNode(parent, replacement *NodeBST) {
	if parent.Left == n {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}
}

func (n *NodeBST) DeleteNode(tree *BST, key interface{}) error {
	var temp, parent *NodeBST = n, nil

	for temp != nil {
		if tree.isKeyEqual(temp.Key, key) { // has both children
			if temp.Left != nil && temp.Right != nil {
				replacement, replacementParent := temp.Right.SubTreeMin() // replacement will be min most of right subtree

				temp.Key, temp.Data = replacement.Key, replacement.Data // transfer data from replacement node

				temp, parent = replacement, replacementParent
				key = replacement.Key // delete replacement
			} else if temp.Left == nil && temp.Right == nil { // has no children
				temp.replaceNode(parent, nil)
				return nil
			} else if temp.Left != nil { // only has left child
				temp.replaceNode(parent, temp.Left)
				return nil
			} else { // only has right child
				temp.replaceNode(parent, temp.Right)
				return nil
			}
		} else if lessThanTemp := tree.LessThan(key, temp.Key); lessThanTemp { // search for target node
			parent = temp
			temp = temp.Left
		} else {
			parent = temp
			temp = temp.Right
		}
	}

	return ErrNotFoundNode
}
