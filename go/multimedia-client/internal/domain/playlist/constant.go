package playlist

type OperationKey int32

const (
	OperationKeyAdd OperationKey = iota
	OperationKeyRemove
	OperationKeyGet
)
