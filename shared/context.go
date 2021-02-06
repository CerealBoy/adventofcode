package shared

const (
	InvalidDay = "An invalid day was referenced for the year."
)

// Context procides a general and shared context to be used between all commands.
type Context struct {
	// Debug will define if additional output should be generated.
	Debug bool

	// Input is the actual file to be used as the input for the day.
	Input string
}
