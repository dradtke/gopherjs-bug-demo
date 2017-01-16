package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/albrow/vdom"
	"honnef.co/go/js/dom"
)

var (
	pageTemplate = template.Must(template.New("").Delims("[[", "]]").Parse(`
<table>
  <tbody>
    <tr>
      <td>Name:</td>
      <td>
        [[if .Editing]]<input type="text" value="[[.User.Name]]">[[else]][[.User.Name]][[end]]
      </td>
    </tr>
  </tbody>
</table>
<button id="toggleEdit" type="button">Edit/Cancel</button>
`))
)

type User struct {
	Name string
}

type State struct {
	User    *User
	Editing bool

	Root dom.Element
	tree *vdom.Tree
}

func (s *State) Render() {
	if s.tree == nil {
		var buf bytes.Buffer
		if err := pageTemplate.Execute(&buf, s); err != nil {
			panic(err)
		}

		if tree, err := vdom.Parse(buf.Bytes()); err != nil {
			panic(err)
		} else {
			s.tree = tree
		}

		s.Root.SetInnerHTML(buf.String())
		return
	}

	var buf bytes.Buffer
	if err := pageTemplate.Execute(&buf, s); err != nil {
		panic(err)
	}

	newTree, err := vdom.Parse(buf.Bytes())
	if err != nil {
		panic(err)
	}

	patches, err := vdom.Diff(s.tree, newTree)
	if err != nil {
		panic(err)
	}

	if err := patches.Patch(s.Root); err != nil {
		panic(err)
	}

	s.tree = newTree
}

func WalkNode(node dom.Node) {
	var walk func(dom.Node, string, int)
	walk = func(node dom.Node, indent string, level int) {
		fmt.Printf("%s%d. %v\n", indent, level, node)
		for _, child := range node.ChildNodes() {
			walk(child, indent+"  ", level+1)
		}
	}
	walk(node, "", 0)
}

func main() {
	state := &State{
		User: &User{Name: "Damien"},
		Root: dom.GetWindow().Document().GetElementByID("page"),
	}
	state.Render()

	dom.GetWindow().Document().GetElementByID("toggleEdit").AddEventListener("click", true, func(dom.Event) {
		state.Editing = !state.Editing
		state.Render()
	})
}
