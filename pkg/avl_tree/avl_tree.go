package avltree

type Node[T any] struct {
	Key    int
	Value  T
	Height int
	Left   *Node[T]
	Right  *Node[T]
}

type AVLTree[T any] struct {
	Root *Node[T]
}

// Вспомогательные функции для работы с AVL деревом
func getHeight[T any](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance[T any](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return getHeight(n.Left) - getHeight(n.Right)
}

// Вращения для поддержания баланса
func rightRotate[T any](y *Node[T]) *Node[T] {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = getMax(getHeight(y.Left), getHeight(y.Right)) + 1
	x.Height = getMax(getHeight(x.Left), getHeight(x.Right)) + 1

	return x
}

func leftRotate[T any](x *Node[T]) *Node[T] {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = getMax(getHeight(x.Left), getHeight(x.Right)) + 1
	y.Height = getMax(getHeight(y.Left), getHeight(y.Right)) + 1

	return y
}

func (t *AVLTree[T]) insert(n **Node[T], key int, value T) {
	if *n == nil {
		*n = &Node[T]{Key: key, Value: value, Height: 1}
		return
	}

	if key < (*n).Key {
		t.insert(&(*n).Left, key, value)
	} else if key > (*n).Key {
		t.insert(&(*n).Right, key, value)
	} else {
		return
	}

	(*n).Height = 1 + getMax(getHeight((*n).Left), getHeight((*n).Right))

	balance := getBalance(*n)

	if balance > 1 && key < (*n).Left.Key {
		*n = rightRotate(*n)
		return
	}

	if balance < -1 && key > (*n).Right.Key {
		*n = leftRotate(*n)
		return
	}

	if balance > 1 && key > (*n).Left.Key {
		(*n).Left = leftRotate((*n).Left)
		*n = rightRotate(*n)
		return
	}

	if balance < -1 && key < (*n).Right.Key {
		(*n).Right = rightRotate((*n).Right)
		*n = leftRotate(*n)
		return
	}
}

// Метод для добавления элемента
func (t *AVLTree[T]) Insert(key int, value T) {
	t.insert(&t.Root, key, value)
}

// Нахождение минимального и максимального элементов
func findMin[T any](n *Node[T]) *T {
	currentNode := n
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return &currentNode.Value
}

// Функция обертка
func (t *AVLTree[T]) FindMin() *T {
	if t.Root != nil {
		return findMin(t.Root)
	}
	return nil // или возвращать ошибку, если дерево пустое
}

func findMax[T any](n *Node[T]) *T {
	currentNode := n
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return &currentNode.Value
}

func (t *AVLTree[T]) FindMax() *T {
	if t.Root != nil {
		return findMax(t.Root)
	}
	return nil
}

func (t *AVLTree[T]) ToSlice() []T {
	var elements []T
	inOrderTraverse(t.Root, &elements)
	return elements
}

// Обход дерева и добавление элементов в слайс
func inOrderTraverse[T any](node *Node[T], elements *[]T) {
	if node != nil {
		inOrderTraverse(node.Left, elements)
		*elements = append(*elements, node.Value)
		inOrderTraverse(node.Right, elements)
	}
}
