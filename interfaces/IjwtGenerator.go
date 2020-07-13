package interfaces

type IJwtGenerator interface {
	GenerateToken(content map[string]interface{}) (string, error)
}
