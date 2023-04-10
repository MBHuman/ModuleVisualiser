package visualiser

import "fmt"

const (
	CAP_SIZE = 3
	ROOT_URL = "root"
	INDIRECT = 1
	DIRECT   = 2
)

type RequirementNodeType uint8

// RequirementNode represents a node in a requirement tree.
type RequirementNode struct {
	reqType RequirementNodeType
	url     string
	childs  map[string]*RequirementNode
}

// NewRequirementNode creates a new RequirementNode instance with the given URL.
func NewRequirementNode(url string) *RequirementNode {
	return &RequirementNode{
		url:     url,
		childs:  make(map[string]*RequirementNode, CAP_SIZE),
		reqType: DIRECT,
	}
}

// addChild creates a new child node for this RequirementNode instance with the given URL.
func (n *RequirementNode) addChild(newNode *RequirementNode) {
	n.childs[newNode.url] = newNode
}

// Requirement represents a requirement tree.
type Requirement struct {
	root            *RequirementNode
	elementsParents map[string]map[string]*RequirementNode
	elementsSet     map[string]*RequirementNode
}

// NewRequirement creates a new Requirement instance.
func NewRequirement() *Requirement {
	return &Requirement{
		root:            NewRequirementNode(ROOT_URL),
		elementsParents: make(map[string]map[string]*RequirementNode, CAP_SIZE),
		elementsSet:     make(map[string]*RequirementNode, CAP_SIZE),
	}
}

// Add link between newNode and parentNode in key value graph "Requirements",
// parentNode can be nil, or other newNode can't be nil and can't contains "root" url.
func (req *Requirement) AddRequirementNodes(
	newNode,
	parentNode *RequirementNode,
) error {
	if newNode == nil {
		return fmt.Errorf("NEW NODE CAN'T BE nil")
	}
	if newNode.url == "root" {
		return fmt.Errorf("NEW NODE CAN'T BE root")
	}

	nodeN := newNode
	nodeP := parentNode
	if _, ok := req.elementsSet[nodeN.url]; !ok {
		req.elementsSet[nodeN.url] = nodeN
	} else {
		nodeN = req.elementsSet[nodeN.url]
	}

	if nodeP == nil {
		nodeP = req.root
	}

	if _, ok := req.elementsSet[nodeP.url]; !ok {
		req.elementsSet[nodeP.url] = nodeP
	} else {
		nodeP = req.elementsSet[nodeP.url]
	}

	nodeP.addChild(nodeN)

	if _, ok := req.elementsParents[nodeN.url]; !ok {
		req.elementsParents[nodeN.url] = make(map[string]*RequirementNode, CAP_SIZE)
	}
	req.elementsParents[nodeN.url][nodeP.url] = nodeP
	if _, ok := req.elementsParents[nodeN.url][req.root.url]; len(req.elementsParents[nodeN.url]) > 1 && ok {
		delete(req.elementsParents[nodeN.url], req.root.url)
	}
	if _, ok := req.elementsParents[nodeN.url][req.root.url]; !ok {
		nodeN.reqType = INDIRECT
	}

	return nil
}

// Add requirement to Requirements key value graph
func (req *Requirement) AddRequirement(
	newNodeUrl, parentUrl string,
) error {
	newNode := NewRequirementNode(newNodeUrl)
	parentNode := NewRequirementNode(parentUrl)
	if err := req.AddRequirementNodes(newNode, parentNode); err != nil {
		return err
	}
	return nil
}

// Add requirement to Requirement key value graph, but parentUrl automaticaly root
func (req *Requirement) AddSingleRequirement(
	newNodeUrl string,
) error {
	newNode := NewRequirementNode(newNodeUrl)
	if err := req.AddRequirementNodes(newNode, nil); err != nil {
		return err
	}
	return nil
}
