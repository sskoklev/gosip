package csom

import (
	"bytes"
	"fmt"
	"text/template"
)

// Action CSOM XML action node builder interface
type Action interface {
	String() string
	SetID(id int)
	GetID() int
	SetObjectID(objectID int)
	GetObjectID() int
	CheckErr() error
}

type action struct {
	template string
	id       int
	objectID int
	err      error
}

// NewAction creates CSOM XML action node builder instance
func NewAction(template string) Action {
	a := &action{}
	a.template = template
	return a
}

// NewActionIdentityQuery creates CSOM XML action node builder instance
func NewActionIdentityQuery() Action {
	return NewAction(`<ObjectIdentityQuery Id="{{.ID}}" ObjectPathId="{{.ObjectID}}" />`)
}

// NewActionMethod creates CSOM XML action node builder instance
func NewActionMethod(methodName string, parameters []string) Action {
	params := ""
	for _, param := range parameters {
		params += param
	}
	return NewAction(fmt.Sprintf(`
		<Method Id="{{.ID}}" ObjectPathId="{{.ObjectID}}" Name="%s">
			<Parameters>%s</Parameters>
		</Method>
	`, methodName, trimMultiline(params)))
}

// String stringifies an action
func (a *action) String() string {
	a.err = nil

	t, _ := template.New("action").Parse(a.template)

	data := &struct {
		ID       int
		ObjectID int
	}{
		ID:       a.GetID(),
		ObjectID: a.GetObjectID(),
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		a.err = err
		return a.template
	}

	return trimMultiline(tpl.String())
}

// SetID sets action's ID
func (a *action) SetID(id int) { a.id = id }

// GetID gets action's ID
func (a *action) GetID() int { return a.id }

// SetObjectID sets action's object ID
func (a *action) SetObjectID(objectID int) { a.objectID = objectID }

// GetObjectID gets action's object ID
func (a *action) GetObjectID() int { return a.objectID }

// CheckErr checks if an action contains errors
func (a *action) CheckErr() error { return a.err }
