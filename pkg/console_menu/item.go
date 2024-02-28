package consolemenu

type MenuItem interface {
	ID() int
	Name() string
	ShouldExit() bool
	Action() *any
}

type FuncItem struct {
	id         int
	name       string
	shouldExit bool
	args       []*any
	fn         func(...any) *any
}

func NewFuncItem(
	id int,
	name string,
	shouldExit bool,
	fn func(...any) *any,
	args ...*any,
) MenuItem {
	return &FuncItem{
		id:         id,
		name:       name,
		shouldExit: shouldExit,
		fn:         fn,
		args:       args,
	}
}

func (f *FuncItem) ID() int {
	return f.id
}
func (f *FuncItem) Name() string {
	return f.name
}
func (f *FuncItem) ShouldExit() bool {
	return f.shouldExit
}

func (f *FuncItem) Action() *any {
	if f.fn == nil {
		return nil
	}
	return f.fn(f.args)
}

func NewExitItem() MenuItem {
	return NewFuncItem(0, "SAIR", true, nil)
}

func NewMenuItem(id int, name string) MenuItem {
	return NewFuncItem(id, name, false, nil)
}
