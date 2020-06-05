package sms

type multi struct {
	to   string
	vars map[string]string
}

func CreateMulti() *multi {
	return &multi{"", make(map[string]string)}
}

func (this *multi) SetTo(to string) {
	this.to = to
}

func (this *multi) AddVar(key string, val string) {
	this.vars[key] = val
}

func (this *multi) Get() map[string]interface{} {
	item := make(map[string]interface{})
	item["to"] = this.to
	item["vars"] = this.vars
	return item
}
