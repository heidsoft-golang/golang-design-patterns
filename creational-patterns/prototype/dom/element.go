package dom

import (
	"bytes"
	"fmt"
)

// Element represents an element in document object model
type Element struct {
	text     string
	parent   *Element
	children []*Element
}

// NewElement makes a new element
func NewElement(text string) *Element {
	return &Element{
		text:     text,
		parent:   nil,
		children: make([]*Element, 0),
	}
}

// String returns string representation of element
func (e *Element) String() string {
	buffer := bytes.NewBufferString(e.text)

	for _, c := range e.children {
		text := c.String()
		fmt.Fprintf(buffer, "\n %s", text)
	}

	return buffer.String()
}

// Add adds child to the root
func (e *Element) Add(child *Element) {
	copy := child.Clone()
	copy.parent = e
	e.children = append(e.children, copy)
}

// Clone makes a copy of particular element. Note that the element becomes a
// root of new orphan tree
func (e *Element) Clone() *Element {
	copy := &Element{
		text:     e.text,
		parent:   nil,
		children: make([]*Element, 0),
	}
	for _, child := range e.children {
		copy.Add(child)
	}
	return copy
}
