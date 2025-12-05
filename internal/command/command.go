package command

type Command interface {
	Execute(args any) error
}
