package visualiser

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestAddRequirementNodes(t *testing.T) {
	convey.Convey("Given a requirement tree", t, func() {
		req := NewRequirement()

		convey.Convey("When a new requirement node is added to the tree without a parent", func() {
			newNode := NewRequirementNode("newNode")
			err := req.AddRequirementNodes(newNode, nil)

			convey.Convey("Then the node should be added as a child of the root node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(req.root.childs[newNode.url], convey.ShouldEqual, newNode)
			})
		})

		convey.Convey("When a new requirement node is added to the tree with a parent", func() {
			newNode := NewRequirementNode("newNode")
			parentNode := NewRequirementNode("parentNode")
			err := req.AddRequirementNodes(newNode, parentNode)

			convey.Convey("Then the node should be added as a child of the parent node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(parentNode.childs[newNode.url], convey.ShouldEqual, newNode)
			})

			convey.Convey("And the parent node should be added as a parent of the new node", func() {
				convey.So(req.elementsParents[newNode.url][parentNode.url], convey.ShouldEqual, parentNode)
			})
		})

		convey.Convey("When a new requirement node with the same URL as an existing node is added", func() {
			existingNode := NewRequirementNode("existingNode")
			req.AddRequirementNodes(existingNode, nil)

			newNode := NewRequirementNode("existingNode")
			err := req.AddRequirementNodes(newNode, nil)

			convey.Convey("Then the new node should not be added as a child of the existing node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(req.elementsSet[existingNode.url].childs[newNode.url], convey.ShouldNotEqual, newNode)
			})

			convey.Convey("And the existing node should remain a child of the root node", func() {
				convey.So(req.elementsParents[existingNode.url][req.root.url], convey.ShouldEqual, req.root)
			})
		})

		convey.Convey("When a new node is added with the same URL as an existing node, and a parent node is provided", func() {
			existingNode := NewRequirementNode("existingNode")
			req.AddRequirementNodes(existingNode, nil)

			newNode := NewRequirementNode("existingNode")
			parentNode := NewRequirementNode("parentNode")
			req.AddRequirementNodes(parentNode, nil)
			err := req.AddRequirementNodes(newNode, parentNode)

			convey.Convey("Then the new node should be added as a child of the parent node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(req.elementsSet[parentNode.url].childs[newNode.url], convey.ShouldEqual, existingNode)
			})

			convey.Convey("And the parent of new node should not be root", func() {
				convey.So(req.elementsParents[newNode.url][req.root.url], convey.ShouldNotEqual, req.root)
			})
		})
	})
}

func TestAddRequirement(t *testing.T) {
	convey.Convey("Given a requirement tree", t, func() {
		req := NewRequirement()

		convey.Convey("When a new requirement node is added to the tree with a parent", func() {
			newElementUrl := "newNode"
			parrentUrl := "parrentNode"

			err := req.AddRequirement(newElementUrl, parrentUrl)

			convey.Convey("Then the node should be added as a child of the parent node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(req.elementsSet[parrentUrl].childs[newElementUrl], convey.ShouldEqual, req.elementsSet[newElementUrl])
			})

			convey.Convey("And the parent node should be added as a parent of the new node", func() {
				convey.So(req.elementsParents[newElementUrl][parrentUrl], convey.ShouldEqual, req.elementsSet[parrentUrl])
			})
		})
	})
}

func TestAddSingleRequirement(t *testing.T) {
	convey.Convey("Given a requirement tree", t, func() {
		req := NewRequirement()

		convey.Convey("When a new requirement node is added to the tree without a parent", func() {
			newNodeUrl := "newNode"
			err := req.AddSingleRequirement(newNodeUrl)

			convey.Convey("Then the node should be added as a child of the root node", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(req.root.childs[newNodeUrl], convey.ShouldEqual, req.elementsSet[newNodeUrl])
			})
		})
	})
}
