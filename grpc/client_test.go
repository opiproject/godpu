package grpc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	grpcOpi "github.com/opiproject/godpu/grpc"
	"google.golang.org/grpc"
)

var _ = Describe("gRPC", func() {
	var addr string
	var dialler grpcOpi.Dialler

	When("we want to create a new client", func() {
		var c grpcOpi.Connector
		var err error

		Context("using a non-empty address", func() {
			BeforeEach(func() {
				addr = "localhost:1234"
				c, err = grpcOpi.New(addr)
			})

			It("should return a client", func() {
				Expect(c).NotTo(BeNil())
			})
			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})

			Context("and specifying a valid dialler", func() {
				BeforeEach(func() {
					dialler = diallerNoError
					c, err = grpcOpi.NewWithDialler(addr, dialler)
				})

				It("should return a client", func() {
					Expect(c).NotTo(BeNil())
				})
				It("should not return an error", func() {
					Expect(err).To(BeNil())
				})
			})
			Context("and specifying an invalid dialler", func() {
				BeforeEach(func() {
					dialler = nil
					c, err = grpcOpi.NewWithDialler(addr, dialler)
				})

				It("should not return a client", func() {
					Expect(c).To(BeNil())
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
		})

		Context("using an empty address", func() {
			BeforeEach(func() {
				addr = ""
				c, err = grpcOpi.New(addr)
			})

			It("should not return a client", func() {
				Expect(c).To(BeNil())
			})
			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})

			Context("and specifying a valid dialler", func() {
				BeforeEach(func() {
					dialler = diallerNoError
					c, err = grpcOpi.NewWithDialler(addr, dialler)
				})

				It("should not return a client", func() {
					Expect(c).To(BeNil())
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})

			Context("and specifying an invalid dialler", func() {
				BeforeEach(func() {
					dialler = nil
					c, err = grpcOpi.NewWithDialler(addr, dialler)
				})

				It("should not return a client", func() {
					Expect(c).To(BeNil())
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
		})
	})

	When("we want to create a new gRPC connection", func() {
		var conn grpc.ClientConnInterface
		var closer grpcOpi.Closer
		var err error

		BeforeEach(func() {
			addr = "localhost:1234"
		})

		Context("and a connection can be created", func() {
			BeforeEach(func() {
				dialler = diallerNoError
				c, _ := grpcOpi.NewWithDialler(addr, dialler)
				conn, closer, err = c.NewConn()
			})

			It("should return a valid connection", func() {
				Expect(conn).ToNot(BeNil())
			})

			It("should return a valid closer", func() {
				Expect(closer).ToNot(BeNil())
			})

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("and a connection cannot be created", func() {
			BeforeEach(func() {
				dialler = diallerWithError
				c, _ := grpcOpi.NewWithDialler(addr, dialler)
				conn, closer, err = c.NewConn()
			})

			It("should not return connection", func() {
				Expect(conn).To(BeNil())
			})

			It("should not return a closer", func() {
				Expect(closer).To(BeNil())
			})

			It("should return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
