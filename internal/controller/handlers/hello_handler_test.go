package handlers_test

import (
	"encoding/json"
	model "github.com/camerontrewbase/ginkgo/internal/domain/hello"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"log"
	"net/http"
)

var _ = Describe("HelloHandler", Ordered, func() {
	It("can return hello from the end point", func() {
		resp, err := http.Get("http://localhost:8080/hello")
		if err != nil {
			log.Fatalln(err)
		}

		var response model.Response
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		gomega.Expect(resp.StatusCode).To(gomega.Equal(200))
		gomega.Expect(response.Message).To(gomega.Equal("Hello World"))
	})
})
