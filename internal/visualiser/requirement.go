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
		url:    url,
		childs: make(map[string]*RequirementNode, CAP_SIZE),
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
func (r *Requirement) AddRequirementNodes(newNode, parentNode *RequirementNode) error {
	if newNode == nil {
		return fmt.Errorf("NEW NODE CAN'T BE nil")
	}
	if newNode.url == "root" {
		return fmt.Errorf("NEW NODE CAN'T BE root")
	}

	nodeN := newNode
	nodeP := parentNode
	if _, ok := r.elementsSet[nodeN.url]; !ok {
		r.elementsSet[nodeN.url] = nodeN
	} else {
		nodeN = r.elementsSet[nodeN.url]
	}

	if nodeP == nil {
		nodeP = r.root
	}

	if _, ok := r.elementsSet[nodeP.url]; !ok {
		r.elementsSet[nodeP.url] = nodeP
	} else {
		nodeP = r.elementsSet[nodeP.url]
	}

	nodeP.addChild(nodeN)

	if _, ok := r.elementsParents[nodeN.url]; !ok {
		r.elementsParents[nodeN.url] = make(map[string]*RequirementNode, CAP_SIZE)
	}
	r.elementsParents[nodeN.url][nodeP.url] = nodeP
	if _, ok := r.elementsParents[nodeN.url][r.root.url]; len(r.elementsParents[nodeN.url]) > 1 && ok {
		delete(r.elementsParents[nodeN.url], r.root.url)
	}

	return nil
}

// Add requirement to Requirements key value graph
func (r *Requirement) AddRequirement(newNodeUrl, parentUrl string) error {
	newNode := NewRequirementNode(newNodeUrl)
	parentNode := NewRequirementNode(parentUrl)
	err := r.AddRequirementNodes(newNode, parentNode)
	if err != nil {
		return err
	}
	return nil
}

// Add requirement to Requirement key value graph, but parentUrl automaticaly root
func (r *Requirement) AddSingleRequirement(newNodeUrl string) error {
	newNode := NewRequirementNode(newNodeUrl)
	err := r.AddRequirementNodes(newNode, nil)
	if err != nil {
		return err
	}
	return err
}
