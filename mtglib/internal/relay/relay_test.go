package relay_test

import (
	"context"
	"testing"
	"time"

	"github.com/9seconds/mtg/v2/mtglib/internal/relay"
	"github.com/stretchr/testify/suite"
)

type RelayTestSuite struct {
	suite.Suite

	ctx       context.Context
	ctxCancel context.CancelFunc
	r         *relay.Relay
}

func (suite *RelayTestSuite) SetupTest() {
	suite.ctx, suite.ctxCancel = context.WithCancel(context.Background())
	suite.r = relay.AcquireRelay(suite.ctx, loggerMock{}, 4096, time.Second)
}

func (suite *RelayTestSuite) TearDownTest() {
	suite.ctxCancel()
	relay.ReleaseRelay(suite.r)
	suite.r = nil
}

func (suite *RelayTestSuite) TestCancelled() {
	suite.ctxCancel()

	eastConn := &rwcMock{}
	eastConn.Write([]byte{1, 2, 3, 4, 5}) // nolint: errcheck

	westConn := &rwcMock{}
	westConn.Write([]byte{100, 101, 102}) // nolint: errcheck

	suite.Nil(suite.r.Process(eastConn, westConn))
}

func (suite *RelayTestSuite) TestCopyFine() {
	eastConn := &rwcMock{}
	eastConn.Write([]byte{1, 2, 3, 4, 5}) // nolint: errcheck

	westConn := &rwcMock{}
	westConn.Write([]byte{100, 101, 102}) // nolint: errcheck

	// yes, this test is not good enough. but apparently, if it hangs,
	// we can debug most of possible issues.
	_ = suite.r.Process(eastConn, westConn)
}

func TestRelay(t *testing.T) {
	t.Parallel()
	suite.Run(t, &RelayTestSuite{})
}