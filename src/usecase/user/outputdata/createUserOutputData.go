package outputdata

type CreateUserOutputData struct {
	CreatedId      string
	IsSuccess      bool
	Err            error
	HttpStatusCode int
}
