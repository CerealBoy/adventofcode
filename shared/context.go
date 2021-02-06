package shared

// Context procides a general and shared context to be used between all commands.
type Context struct {
	// Debug will define if additional output should be generated.
	Debug bool

	// Input is the actual file to be used as the input for the day.
	Input string
}
