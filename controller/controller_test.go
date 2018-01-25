package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

type ProviderMock struct {}

func (t *ProviderMock) Get() ([]byte, error) {
 return []byte("{ foo: 1 }"), nil
}


var _ = Describe("Controller Tests", func() {

	var p DataProvider
	var c Controller

	BeforeEach(func() {
		p = &ProviderMock{}
		c = NewController(p)
	})

	It("Should get root", func() {

		// given
		req, _ := http.NewRequest("GET", "", nil)
		rr := httptest.NewRecorder()
		controller := http.HandlerFunc(c.RootRoute)

		// when
		controller.ServeHTTP(rr, req)

		// then
		Expect(rr.Body.String()).To(Equal("Welcome home!"))
		Expect(rr.Code).To(Equal(200))
	})


	It("Should get JSON data", func() {

		// given
		req, _ := http.NewRequest("GET", "", nil)
		rr := httptest.NewRecorder()
		controller := http.HandlerFunc(c.JsonRoute)

		// when
		controller.ServeHTTP(rr, req)

		// then
		Expect(rr.Body.String()).To(Equal("{ foo: 1 }"))
		Expect(rr.Code).To(Equal(200))
	})

})
