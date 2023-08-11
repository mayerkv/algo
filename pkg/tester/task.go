package tester

type Task interface {
	Run(args []string) (string, error)
}
