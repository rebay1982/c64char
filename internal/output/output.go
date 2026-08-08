package output

type FormatterType int

const (
	AcmeFormatter FormatterType = iota
)

type Formatter interface {
	Output(buf []byte) string
}

func NewFormatter(t FormatterType) Formatter {
	switch t {
	case AcmeFormatter:
		return acmeFormatter{}
	default:
		return nil
	}
}
