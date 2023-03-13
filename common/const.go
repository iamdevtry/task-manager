package common

const CurrentUser = "user"

type Requester interface {
	GetUserId() int64
}
