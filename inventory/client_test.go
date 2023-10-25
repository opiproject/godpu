// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

package inventory_test

import (
	"context"
	"errors"

	//nolint:revive
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive
	. "github.com/onsi/gomega"
	grpcOpi "github.com/opiproject/godpu/grpc"
	"github.com/opiproject/godpu/inventory"
	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/inventory/v1/gen/go"
	"google.golang.org/grpc"
)

var _ = Describe("Inventory", func() {

	var addr string
	var g inventory.PbInvClientGetter
	var m grpcOpi.Connector

	When("we want to create a new client", func() {
		var c inventory.InvClient
		var err error

		Context("using a non-empty address", func() {
			BeforeEach(func() {
				addr = "localhost:1234"
				c, err = inventory.New(addr)
			})

			It("should return a client", func() {
				Expect(c).NotTo(BeNil())
			})
			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("using an empty address", func() {
			BeforeEach(func() {
				addr = ""
				c, err = inventory.New(addr)
			})

			It("should not return a client", func() {
				Expect(c).To(BeNil())
			})
			It("should return an error", func() {
				Expect(err).NotTo(BeNil())
			})
		})

		Context("using a connector and getter", func() {

			Context("when the connector is valid", func() {
				BeforeEach(func() {
					m = mocks.NewConnector(GinkgoT())
				})

				Context("and the getter is valid", func() {
					BeforeEach(func() {
						g = func(c grpc.ClientConnInterface) pb.InventoryServiceClient {
							return &mocks.InventorySvcClient{}
						}
						c, err = inventory.NewWithArgs(m, g)
					})

					It("should return a client", func() {
						Expect(c).ToNot(BeNil())
					})
					It("should return no error", func() {
						Expect(err).To(BeNil())
					})
				})

				Context("and the getter is not valid", func() {
					BeforeEach(func() {
						g = nil
						c, err = inventory.NewWithArgs(m, g)
					})

					It("should not return a client", func() {
						Expect(c).To(BeNil())
					})
					It("should return an error", func() {
						Expect(err).ToNot(BeNil())
					})
				})
			})

			Context("when the connector is not valid", func() {
				BeforeEach(func() {
					m = nil
				})

				Context("and the getter is valid", func() {
					BeforeEach(func() {
						g = func(c grpc.ClientConnInterface) pb.InventoryServiceClient {
							return &mocks.InventorySvcClient{}
						}
						c, err = inventory.NewWithArgs(m, g)
					})

					It("should not return a client", func() {
						Expect(c).To(BeNil())
					})
					It("should return an error", func() {
						Expect(err).ToNot(BeNil())
					})
				})

				Context("and the getter is not valid", func() {
					BeforeEach(func() {
						g = nil
						c, err = inventory.NewWithArgs(m, g)
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
	})

	When("we want to get tne inventory", func() {
		var c inventory.InvClient

		Context("and we can create a connection", func() {
			BeforeEach(func() {
				mockConn := mocks.NewConnector(GinkgoT())
				mockConn.EXPECT().NewConn().Return(&grpc.ClientConn{}, testCloser, nil).Once()
				m = mockConn
			})

			Context("and the inventory can be retrieved", func() {
				BeforeEach(func() {
					g = getterWithResponse
					c, _ = inventory.NewWithArgs(m, g)
				})

				It("should return a valid inventory", func() {
					r, _ := c.Get(context.Background())
					Expect(r).To(Equal(getTestInvResponse()))
				})
				It("should not return an error", func() {
					_, err := c.Get(context.Background())
					Expect(err).To(BeNil())
				})
			})

			Context("and the inventory can not be retrieved", func() {
				BeforeEach(func() {
					g = getterWithError
					c, _ = inventory.NewWithArgs(m, g)
				})

				It("should not return an inventory", func() {
					r, _ := c.Get(context.Background())
					Expect(r).To(BeNil())
				})
				It("should return an error", func() {
					_, err := c.Get(context.Background())
					Expect(err).NotTo(BeNil())
				})
			})
		})

		Context("and we can not create a connection", func() {
			BeforeEach(func() {
				mockConn := mocks.NewConnector(GinkgoT())
				mockConn.EXPECT().NewConn().
					Return(nil, nil, errors.New("cannot create new connection")).Once()
				m = mockConn
				g = getterWithResponse
				c, _ = inventory.NewWithArgs(m, g)
			})

			It("should not return an inventory", func() {
				r, _ := c.Get(context.Background())
				Expect(r).To(BeNil())
			})
			It("should return an error", func() {
				_, err := c.Get(context.Background())
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
