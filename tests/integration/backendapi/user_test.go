package backendapi

import (
	"fmt"
	"net/http"
	"testing"
)

const userEndpointUrl string = "/auth/user"

func TestListUsers(t *testing.T) {

	t.Run("without login should failed as unauthorized", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		e.GET(userEndpointUrl).
			Expect().
			Status(401)

	})

	t.Run("with  login should   succeed", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		e.GET(userEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			Expect().
			Status(http.StatusOK).
			JSON().
			Array().
			NotEmpty()
	})
}

func TestCreateUser(t *testing.T) {

}

func TestDeleteUser(t *testing.T) {

}
