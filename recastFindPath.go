package main

//import "fmt"
import "math"



type dtNode struct {
	pos   []float64
	cost  float64
	total float64
	pidx  uint64
	flags uint64
	id    uint64
}

type openList struct {
	m_size int
}

func (instance *openList) push(u *dtNode){
	
}
func (instance *openList) pop()*dtNode{
	
}

func (instance *openList) empty()bool{
	if instance.m_size == 0{
		return true;
	}else{
		return false;
	}
}



func dtVdist(a []float64, b []float64) float64 {
	sum := (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1]) + (a[2]-b[2])*(a[2]-b[2])
	return math.Sqrt(sum)
}

func findPath(startRef uint64, endRef uint64, startPos *[]float64, endPos *[]float64, path *[]uint64, pathCount *int, maxPath int) uint32 {
	*pathCount = 0

	if startRef == 0 || endRef == 0 {
		return 2147483656
	}

	if maxPath == 0 {
		return 2147483656
	}

	if !validRef(startRef) || !validRef(endRef) {
		return 2147483656
	}

	if startRef == endRef {
		*path = append(*path, startRef)
		*pathCount = 1
		return 1073741824
	}
	var m_nodePool nodePool
	m_openList := &openList{}

	startNode := new(dtNode)
	startNode.pidx = 0
	startNode.cost = 0
	startNode.total = dtVdist(*startPos, *endPos) * 0.999
	startNode.flags = 1
	startNode.id = startRef
	lastBestNodeCost := startNode
	status := DT_SUCCESS
	m_openList.push(startNode)
	for !m_openList.empty() {
		bestNode := m_openList.pop()
		if bestNode.id == endRef {
			lastBestNodeCost = bestNode
			break
		}

		bestRef := bestNode.id
		bestPoly := 0

		parentRef := getNodeAtIdx(bestNode.pidx)

		getTileAndPolyByRefUnsafe(bestRef, bestTile, bestPoly)

		for i := bestPoly.firstLink; i != DT_NULL_LINK; i = bestTile.links[i].next {
			neighbourRef := bestTile.links[i].ref

			if neighbourRef == 0 || neighbourRef == parentRef {
				continue
			}

			neighbourTile := 0
			neighbourPoly := 0

			getTileAndPolyByRefUnsafe(neighbourRef, &neighbourTile, &neighbourPoly)

			if neighbourRef == endRef {
				curCost = getCost(bestNode.pos, neighbourNode.pos, parentRef, parentTile, parentPoly, bestRef, bestTile, bestPoly, neighbourRef, neighbourTile, neighbourPoly)
				endCost = getCost(neighbourNode.pos, endPos, bestRef, bestTile, bestPoly, neighbourRef, neighbourTile, neighbourPoly,0, 0, 0)
				
				cost = bestNode.cost + curCost + endCost
				heuristic = 0;
			}else{
				curCost = getCost(bestNode.pos, neighbourNode.pos, parentRef, parentTile, parentPoly, bestRef, bestTile, bestPoly, neighbourRef, neighbourTile, neighbourPoly);
				cost = bestNode.cost + curCost
				heuristic = dtVdist(neighbourNode.pos, endPos) * 0.99
			}
			total := cost + heuristic;
			if (neighbourNode.flags & DT_NODE_OPEN) && total >= neighbourNode.total{
				continue;
			}
			if (neighbourNode.flags & DT_NODE_CLOSED) && total >= neighbourNode.total{
				continue;
			}
			if(neighbourNode.flags & DT_NODE_OPEN){
				m_openList.modify(neighbourNode);
			}else{
				neighbourNode.flags |= DT_NODE_OPEN;
				m_openList.push(neighbourNode);
			}
			if(heuristic < lastBestNodeCost)
			{
				lastBestNodeCost = heuristic;
				lastBestNode = neighbourNode;
			}
		}
	}
	if(lastBestNode.id != endRef){
		status |= DT_PARTIAL_RESULT;
	}
	prev := 0;
	node := lastBestNode
	n := 0
	for node!=0{
		*path = append(*path, node.id);
		n += 1;
		if(n >= maxPath){
			status |= DT_BUFFER_TOO_SMALL;
			break;
		}
		node = getNodeAtIdx(node.pidx);
	}
	*pathCount = n
	return status;
}
