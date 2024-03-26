package user

import "github.com/pallat/webdev/app"

func (handler Handler) User(c app.Context) {
	id := c.Param("id")

	user, err := handler.storage.User(id)
	if err != nil {
		c.StoreError(err)
		return
	}
	c.OK(user.User)
}
