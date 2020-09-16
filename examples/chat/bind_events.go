package chat

import (
	"github.com/inklabs/rangedb"
)

// TODO: Generate code below

func BindEvents(binder rangedb.EventBinder) {
	binder.Bind(
		&UserWasOnBoarded{},
		&UserWasWarned{},
		&RoomWasOnBoarded{},
		&RoomWasJoined{},
		&MessageWasSentToRoom{},
		&PrivateMessageWasSentToRoom{},
	)
}