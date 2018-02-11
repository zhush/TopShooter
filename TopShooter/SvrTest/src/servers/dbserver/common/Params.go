package Params

type KeyValue struct {
	Key   string
	Value interface{}
}

type AddParam struct {
	TableName string
	Values    []KeyValue
}

type AddResult struct {
	Result          int
	AutoIncrementId int64
	ErrorMsg        string
}

type DelParam struct {
	TableName  string
	Conditions []KeyValue
}

type DelResult struct {
	Result   int
	ErrorMsg string
}

type UpdateParam struct {
	TableName  string
	Values     []KeyValue
	Conditions []KeyValue
}

type UpdateResult struct {
	Result   int
	ErrorMsg string
}

type ReadParam struct {
	TableName  string
	Keys       []string
	Conditions []KeyValue
}

type RowResult struct {
	Values map[string]string
}

type ReadResult struct {
	Result    int
	Rows      int
	RowValues []RowResult
}
