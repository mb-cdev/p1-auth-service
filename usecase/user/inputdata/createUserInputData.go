package inputdata

type CreateUserInputDataInterface interface {
	SetName(string)
	SetPassword(string)
	Name() string
	Password() string
}

type CreateUserInputData struct {
	name     string
	password string
}

func (c *CreateUserInputData) SetName(name string) {
	c.name = name
}

func (c *CreateUserInputData) SetPassword(password string) {
	c.password = password
}

func (c *CreateUserInputData) Name() string {
	return c.name
}

func (c *CreateUserInputData) Password() string {
	return c.password
}
