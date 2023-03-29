package common_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/opiproject/godpu/common"
	"google.golang.org/grpc"
)

var _ = Describe("Client", func() {
	var addr string
	var dialler common.Dialler

	Describe("NewClient", func() {
		var c common.Client
		var err error
		Context("When creating a new client", func() {

			Context("using a non-empty address", func() {
				BeforeEach(func() {
					addr = "localhost:1234"
					c, err = common.NewClient(addr)
				})

				It("should return a non nil client implementation", func() {
					Expect(c).NotTo(BeNil())
				})
				It("should not return an error", func() {
					Expect(err).To(BeNil())
				})

				Context("and specifying the dialler", func() {
					BeforeEach(func() {
						dialler = diallerNoError
						c, err = common.NewClientWithDialler(addr, dialler)
					})

					It("should return a non nil client implementation", func() {
						Expect(c).NotTo(BeNil())
					})
					It("should not return an error", func() {
						Expect(err).To(BeNil())
					})
				})
			})

			Context("using an empty address", func() {
				BeforeEach(func() {
					addr = ""
					c, err = common.NewClient(addr)
				})

				It("should return a nil client implementation", func() {
					Expect(c).To(BeNil())
				})
				It("should return an error", func() {
					Expect(err).NotTo(BeNil())
				})

				Context("and specifying the dialler", func() {
					BeforeEach(func() {
						dialler = diallerNoError
						c, err = common.NewClientWithDialler(addr, dialler)
					})

					It("should return a nil client implementation", func() {
						Expect(c).To(BeNil())
					})
					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})

		})
	})

	Describe("NewGrpcConn", func() {
		var conn grpc.ClientConnInterface
		var closer common.Closer
		var err error

		BeforeEach(func() {
			addr = "localhost:1234"
		})

		Context("When creating a new gRPC connection", func() {
			Context("and a connection can be created", func() {
				BeforeEach(func() {
					dialler = diallerNoError
					c, _ := common.NewClientWithDialler(addr, dialler)
					conn, closer, err = c.NewGrpcConn()
				})

				It("should return a valid connection", func() {
					Expect(conn).ToNot(BeNil())
				})

				It("should return a valid closer", func() {
					Expect(closer).ToNot(BeNil())
				})

				It("should return no error", func() {
					Expect(err).To(BeNil())
				})
			})

			Context("and a connection cannot be created", func() {
				BeforeEach(func() {
					dialler = diallerWithError
					c, _ := common.NewClientWithDialler(addr, dialler)
					conn, closer, err = c.NewGrpcConn()
				})

				It("should return a nil connection", func() {
					Expect(conn).To(BeNil())
				})

				It("should return a nil closer", func() {
					Expect(closer).To(BeNil())
				})

				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})
})
