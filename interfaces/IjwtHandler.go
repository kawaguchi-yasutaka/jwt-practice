package interfaces

type IJwtHandler interface {
	Verify(token string) error
}
