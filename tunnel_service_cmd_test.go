package main_test

import (
	. "github.com/benlaplanche/cf-tunnel-service"
	fake_rpc_handlers "github.com/cloudfoundry/cli/testhelpers/rpc_server/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const validPluginPath = "./service-tunnel"

var _ = Describe("TunnelServiceCmd", func() {
	var (
		rpcHandlers *fake_rpc_handlers.FakeHandlers
		ts          *test_rpc_server.TestServer
		err         error
	)
})
