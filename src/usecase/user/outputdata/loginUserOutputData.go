package outputdata

type LoginUserOutputData struct {
	JwtToken       string
	IsSuccess      bool
	Err            error
	HttpStatusCode int
}
