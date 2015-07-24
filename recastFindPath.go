package main

import "fmt"
import "math"

const DT_FAILURE = 1 << 31
const DT_SUCCESS = 1 << 30
const DT_NULL_LINK = 0xffffffff
const DT_INVALID_PARAM = 1 << 3
const DT_NODE_CLOSED = 0x01
const DT_NODE_OPEN = 0x02
const DT_BUFFER_TOO_SMALL = 1 << 4
const DT_PARTIAL_RESULT = 1 << 6

type dtNode struct {
	pos   []float64
	cost  float64
	total float64
	pidx  *dtNode
	flags uint64
	id    uint64
}

type dtTile struct {
	data uint64
	next *dtTile
}

type dtPoly struct {
	area      float64
	firstLink *dtTile
}

// Add a edge between a and b
func add(a *dtPoly, b uint64) {
	t := new(dtTile)
	t.data = b
	t.next = a.firstLink
	a.firstLink = t
}

func (instance *dtPoly) getArea() float64 {
	return instance.area
}

type openList struct {
	m_size int
	m_heap []dtNode
}

func (instance *openList) trickleDown(i int, node *dtNode) {
	child := (i * 2) + 1
	for child < instance.m_size {
		if child+1 < instance.m_size && instance.m_heap[child].total > instance.m_heap[child+1].total {
			child++
		}
		instance.m_heap[i] = instance.m_heap[child]
		i = child
		child = i*2 + 1
	}
	instance.bubbleUp(i, node)
}

func (instance *openList) pop() *dtNode {
	result := instance.m_heap[0]
	instance.m_size--
	instance.trickleDown(0, &instance.m_heap[instance.m_size])
	return &result
}

func (instance *openList) bubbleUp(i int, node *dtNode) {
	parent := (i - 1) >> 1
	for (i > 0) && instance.m_heap[parent].total > node.total {
		instance.m_heap[i] = instance.m_heap[parent]
		i = parent
		parent = (i - 1) >> 1
	}
	instance.m_heap[i] = *node
}

func (instance *openList) push(node *dtNode) {
	instance.m_size++
	startNode := new(dtNode)
	startNode.pidx = nil
	startNode.cost = 0
	startNode.total = 100
	startNode.flags = 1
	startNode.id = 1
	instance.m_heap = append(instance.m_heap, *startNode)
	instance.bubbleUp(instance.m_size-1, node)
}

func (instance *openList) empty() bool {
	if instance.m_size == 0 {
		return true
	} else {
		return false
	}
}

//neighbourRef, &neighbourTile, &neighbourPoly
func getTileAndPolyByRefUnsafe(ref uint64, tile *dtTile, poly *dtPoly) {

}

func dtVdist(a []float64, b []float64) float64 {
	sum := (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1]) + (a[2]-b[2])*(a[2]-b[2])
	return math.Sqrt(sum)
}

func getCost(a []float64, b []float64, cur *dtPoly) float64 {
	return 1.0 * dtVdist(a, b) * cur.getArea()
}

func findPath(startRef uint64, endRef uint64, startPos *[]float64, endPos *[]float64, path *[]uint64, pathCount *int, maxPath int) uint64 {
	*pathCount = 0

	if startRef == 0 || endRef == 0 {
		return DT_FAILURE | DT_INVALID_PARAM
	}

	if maxPath == 0 {
		return DT_FAILURE | DT_INVALID_PARAM
	}
	/*
		if !validRef(startRef) || !validRef(endRef) {
			return DT_FAILURE | DT_INVALID_PARAM
		}
	*/
	if startRef == endRef {
		*path = append(*path, startRef)
		*pathCount = 1
		return DT_SUCCESS
	}
	//var m_nodePool nodePool
	m_openList := &openList{}
	startNode := new(dtNode)
	startNode.pidx = nil
	startNode.cost = 0
	startNode.total = dtVdist(*startPos, *endPos) * 0.999
	startNode.flags = 1
	startNode.id = startRef
	lastBestNode := startNode
	lastBestNodeCost := startNode.total
	var status uint64
	status = DT_SUCCESS
	m_openList.push(startNode)
	for !m_openList.empty() {
		bestNode := m_openList.pop()
		if bestNode.id == endRef {
			lastBestNode = bestNode
			break
		}
		//bestRef := bestNode.id
		var bestPoly *dtPoly
		parent := bestNode.pidx
		//var bestTile *dtTile
		//getTileAndPolyByRefUnsafe(bestRef, bestTile, bestPoly)
		for i := bestPoly.firstLink; i != nil; i = i.next {
			neighbourRef := i.data
			if neighbourRef == 0 || neighbourRef == parent.id {
				continue
			}
			var neighbourTile *dtTile
			var neighbourPoly *dtPoly
			neighbourNode := new(dtNode)
			neighbourNode.id = neighbourRef
			neighbourNode.pidx = bestNode
			getTileAndPolyByRefUnsafe(neighbourRef, neighbourTile, neighbourPoly)
			var curCost, endCost, cost, heuristic float64
			//var heuristic int
			if neighbourRef == endRef {
				curCost = getCost(bestNode.pos, neighbourNode.pos, bestPoly)
				endCost = getCost(neighbourNode.pos, *endPos, neighbourPoly)
				cost = bestNode.cost + curCost + endCost
				heuristic = 0
			} else {
				curCost = getCost(bestNode.pos, neighbourNode.pos, bestPoly)
				cost = bestNode.cost + curCost
				heuristic = dtVdist(neighbourNode.pos, *endPos) * 0.99
			}
			total := cost + heuristic
			if (neighbourNode.flags&DT_NODE_OPEN) != 0 && total >= neighbourNode.total {
				continue
			}
			if (neighbourNode.flags&DT_NODE_CLOSED) != 0 && total >= neighbourNode.total {
				continue
			}
			if (neighbourNode.flags & DT_NODE_OPEN) != 0 {
				//m_openList.modify(neighbourNode);
			} else {
				neighbourNode.flags |= DT_NODE_OPEN
				m_openList.push(neighbourNode)
			}
			if heuristic < lastBestNodeCost {
				lastBestNodeCost = heuristic
				lastBestNode = neighbourNode
			}
		}
	}
	if lastBestNode.id != endRef {
		status |= DT_PARTIAL_RESULT
	}
	//prev := 0;
	node := lastBestNode
	n := 0
	for node != nil {
		*path = append(*path, node.id)
		n += 1
		if n >= maxPath {
			status |= DT_BUFFER_TOO_SMALL
			break
		}
		node = node.pidx
	}
	*pathCount = n
	return status
}
func push(m_openList *openList, num float64) {
	startNode := new(dtNode)
	startNode.pidx = nil
	startNode.cost = 0
	startNode.total = num
	startNode.flags = 1
	startNode.id = 1
	m_openList.push(startNode)
}

func pop(m_openList *openList) {
	s := m_openList.pop()
	fmt.Printf("%f\n", s.total)
}

func main() {
	/*
		// TEST 1
		// To test that does the priority queue module works well
		// all is OK!
		m_openList := &openList{}
		push(m_openList, 10)
		push(m_openList, 50)
		push(m_openList, 44)
		push(m_openList, 99)
		push(m_openList, 1)
		pop(m_openList)
		pop(m_openList)
		push(m_openList, 70)
		pop(m_openList)
		push(m_openList, 3)
		pop(m_openList)
	*/
}
