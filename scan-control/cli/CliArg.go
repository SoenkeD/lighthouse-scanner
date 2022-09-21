package lib

type CliArg struct {
	Id          string
	Default     string
	Description string
	SetOnly     bool
	SetValue    *bool
	Value       *string
}

func (arg *CliArg) InitVal() {
	initVal := ""
	arg.Value = &initVal

	setVal := false
	arg.SetValue = &setVal
}
