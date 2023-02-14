package backendapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const postEndpointUrl string = "/auth/post"

func TestListPosts(t *testing.T) {
	t.Run("without login should failed as unauthorized", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		e.GET(postEndpointUrl).
			Expect().
			Status(401)

	})

	t.Run("with login should succeed", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		e.GET(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			Expect().
			Status(http.StatusOK).JSON().Array().NotEmpty()
	})

	t.Run("with author_id and login should succeed", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		expectedPost := map[string]interface{}{
			"id":        1,
			"author_id": 1,
			"Author": map[string]interface{}{
				"created_at": "2020-12-31T23:59:59Z",
				"email":      "test@integration.int",
				"id":         1,
				"updated_at": "2020-12-31T23:59:59Z",
				"username":   "testint1",
				"password":   "",
			},
			"body":       "This is a wonderful post that",
			"created_at": "2020-12-31T23:59:59Z",
			"title":      "A nice post",
			"updated_at": "2020-12-31T23:59:59Z",
		}

		e.GET(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithQuery("author_id", "1").
			Expect().
			Status(http.StatusOK).JSON().Array().Elements(expectedPost)
	})

	t.Run("With pagination", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		var respLen float64

		//page=1 and page_size=1
		respLen = e.GET(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithQuery("author_id", "2").
			WithQuery("page", "1").
			WithQuery("page_size", "1").
			Expect().
			Status(http.StatusOK).JSON().Array().Length().Raw()

		assert.Equal(t, 1.0, respLen, "The count of returned posts should be 1")

		//page=1 and page_size=3
		respLen = e.GET(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithQuery("author_id", "2").
			WithQuery("page", "1").
			WithQuery("page_size", "3").
			Expect().
			Status(http.StatusOK).JSON().Array().Length().Raw()

		assert.Equal(t, 3.0, respLen, "The count of returned posts should be 3")

		//page=10 and page_size=1
		e.GET(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithQuery("author_id", "2").
			WithQuery("page", "10").
			WithQuery("page_size", "1").
			Expect().
			Status(http.StatusOK).JSON().Array().Empty()
	})

}

func TestCreatePost(t *testing.T) {
	t.Run("with login should succeed", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		newPostParam := map[string]interface{}{
			"title":     "A new start",
			"body":      "From the sand of this paradaise",
			"author_id": 1,
		}

		e.POST(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithJSON(newPostParam).
			Expect().
			Status(http.StatusCreated)
	})

}

func TestDeletePost(t *testing.T) {
	t.Run("with login should succeed", func(t *testing.T) {
		e := HelperBuildHttpexpect(t)

		jwt := HelperLogin(e)

		newPostParam := map[string]interface{}{
			"title":     "A new start",
			"body":      "From the sand of this paradaise",
			"author_id": 1,
		}

		result := e.POST(postEndpointUrl).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			WithJSON(newPostParam).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		newPostId := result.Value("id").Number().Raw()

		url := fmt.Sprintf("%s/%.0f", postEndpointUrl, newPostId)

		e.DELETE(url).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", jwt)).
			Expect().
			Status(http.StatusFound)

	})
}
