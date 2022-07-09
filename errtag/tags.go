package errtag

type Tag struct {
	Code string

	Err error
}

func New(code string, err error) error {
	return &Tag{
		Code: code,
		Err:  err,
	}
}

func (c *Tag) Error() string {
	return c.Err.Error()
}

const (
	NotFound   string = "404"
	Unexpected string = "500"
)
