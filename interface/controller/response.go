package controller

// ResponseAndLogError returns response and log error.
func ResponseAndLogError(ctx Context, err error) {
	he := handleError(err)
	ctx.JSON(he.Status, he)
}
