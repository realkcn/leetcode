package leetcode

type DFANode struct {
	matchChar int
	matchNext *DFANode
}

var (
	startNode = DFANode{matchChar: -1, matchNext: nil}
	endNode   = DFANode{matchChar: -1, matchNext: nil}

	anyMatchChar      = -255
	anyMultiMatchChar = -256
	endChar           = uint8('$')
)

func compile(pattern string) {
	pattern += string(endChar)
	pPrev := &startNode
	for i := 0; i < len(pattern); i++ {
		var ch int
		ch = int(pattern[i])
		switch pattern[i] {
		case '*':
			if pPrev.matchChar == anyMatchChar {
				pPrev.matchChar = anyMultiMatchChar
			} else {
				pPrev.matchChar = -pPrev.matchChar
			}
		case '.':
			ch = anyMatchChar
			fallthrough
		default:
			if !(pPrev.matchChar < 0 && ch != anyMatchChar && pPrev.matchChar == ch) {
				curNode := new(DFANode)
				pPrev.matchNext = curNode
				curNode.matchChar = ch
				pPrev = curNode
			}
		}
	}
	pPrev.matchNext = &endNode
	//优化一下，合并.*.* a*a*这类情况
	//for p := &startNode; p != &endNode; p = p.matchNext {
	//	for p.matchChar < 0 && p.matchChar != anyMatchChar && p.matchChar == p.matchNext.matchChar {
	//		p.matchNext = p.matchNext.matchNext
	//	}
	//}
}

func addNodeToSet(node *DFANode, nextNodeSet *nodeSet) bool {
	for true {
		if node == &endNode {
			return true
		}
		found := false
		for _, item := range *nextNodeSet {
			if item == node {
				found = true
				break
			}
		}
		if !found {
			*nextNodeSet = append(*nextNodeSet, node)
			if node.matchChar < 0 && node.matchChar != anyMatchChar {
				node = node.matchNext
			} else {
				break
			}
		} else {
			break
		}
	}
	return false
}

func transfer(node *DFANode, ch uint8, nextNodeSet *nodeSet) bool {
	if (node.matchChar == anyMatchChar && ch != endChar) || (node.matchChar > 0 && ch == uint8(node.matchChar)) {
		if addNodeToSet(node.matchNext, nextNodeSet) {
			return true
		}
	} else {
		if (node.matchChar == anyMultiMatchChar && ch != endChar) || (node.matchChar < 0 && ch == uint8(-node.matchChar)) {
			if addNodeToSet(node.matchNext, nextNodeSet) {
				return true
			}
			addNodeToSet(node, nextNodeSet)
		}
	}
	return false
}

type nodeSet []*DFANode

func isMatch(s string, p string) bool {
	compile(p)

	//currentNodeSet和nextNodeSet循环使用两个集合，优化一下
	var dataNodeSet [2]*nodeSet
	dataNodeSet[0] = new(nodeSet)
	dataNodeSet[1] = new(nodeSet)
	currentNodeSet := dataNodeSet[0]
	nextNodeSet := dataNodeSet[1]
	addNodeToSet(startNode.matchNext, currentNodeSet)
	s = s + string(endChar)
	for _, ch := range s {
		*nextNodeSet = (*nextNodeSet)[:0]
		for _, node := range *currentNodeSet {
			if transfer(node, uint8(ch), nextNodeSet) {
				return true
			}
		}
		//currentNodeSet和nextNodeSet循环互换
		var tmp *nodeSet
		tmp = currentNodeSet
		currentNodeSet = nextNodeSet
		nextNodeSet = tmp
		if len(*currentNodeSet) == 0 {
			return false
		}
	}
	return false
}
