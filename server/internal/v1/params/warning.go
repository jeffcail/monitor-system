package params

// IgnoreServeCheckRecordParams
type IgnoreServeCheckRecordParams struct {
	ServeId int64 `json:"serve_id" validate:"required"`
}

// IgnoreMachineParams
type IgnoreMachineParams struct {
	Id int64 `json:"id" validate:"required"`
}
