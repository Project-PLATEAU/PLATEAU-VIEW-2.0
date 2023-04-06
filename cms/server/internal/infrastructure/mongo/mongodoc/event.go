package mongodoc

import (
	"time"

	"github.com/reearth/reearth-cms/server/pkg/event"
	"github.com/reearth/reearth-cms/server/pkg/id"
	"github.com/reearth/reearth-cms/server/pkg/operator"
	"github.com/reearth/reearthx/mongox"
)

type EventDocument struct {
	ID          string
	Timestamp   time.Time
	User        *string
	Integration *string
	Machine     bool
	Type        string
	Object      Document
}

func NewEvent(e *event.Event[any]) (*EventDocument, string, error) {
	eId := e.ID().String()
	objDoc, _, err := NewDocument(e.Object())
	if err != nil {
		return nil, "", err
	}
	return &EventDocument{
		ID:          eId,
		Timestamp:   e.Timestamp(),
		User:        e.Operator().User().StringRef(),
		Integration: e.Operator().Integration().StringRef(),
		Machine:     e.Operator().Machine(),
		Type:        string(e.Type()),
		Object:      objDoc,
	}, eId, nil
}

func (d *EventDocument) Model() (*event.Event[any], error) {
	eID, err := event.IDFrom(d.ID)
	if err != nil {
		return nil, err
	}

	m, err := ModelFrom(d.Object)
	if err != nil {
		return nil, err
	}

	var o operator.Operator
	if d.User != nil {
		if uid := id.UserIDFromRef(d.User); uid != nil {
			o = operator.OperatorFromUser(*uid)
		}
	} else if d.Integration != nil {
		if iid := id.IntegrationIDFromRef(d.Integration); iid != nil {
			o = operator.OperatorFromIntegration(*iid)
		}
	} else if d.Machine {
		o = operator.OperatorFromMachine()
	}

	e, err := event.New[any]().
		ID(eID).
		Type(event.Type(d.Type)).
		Timestamp(d.Timestamp).
		Operator(o).
		Object(m).
		Build()
	if err != nil {
		return nil, err
	}

	return e, nil
}

type EventConsumer = mongox.SliceFuncConsumer[*EventDocument, *event.Event[any]]

func NewEventConsumer() *EventConsumer {
	return NewComsumer[*EventDocument, *event.Event[any]]()
}
