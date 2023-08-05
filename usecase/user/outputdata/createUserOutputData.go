package outputdata

type CreateUserOutputDataInterface interface {
	SetError(error)
	SetIsSuccess(bool)

	IsSuccess(bool)
	Error() error
}

type CreateUserOutputData struct {
	isSuccess bool
	err       error
}

func (c *CreateUserOutputData) SetError(err error) {
	c.err = err
}

func (c *CreateUserOutputData) SetIsSuccess(isSuccess bool) {
	c.isSuccess = isSuccess
}

func (c *CreateUserOutputData) IsSuccess() bool {
	return c.isSuccess
}

func (c *CreateUserOutputData) Error() error {
	return c.err
}
