package ics

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/interchain-security/tests/integration"

	appConsumer "github.com/cosmos/interchain-security/app/consumer"
	icstestingutils "github.com/cosmos/interchain-security/testutil/ibc_testing"
	"github.com/stretchr/testify/suite"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	ibctesting "github.com/cosmos/interchain-security/legacy_ibc_testing/testing"

	gaiaApp "github.com/cosmos/gaia/v10/app"
)

func TestCCVTestSuite(t *testing.T) {
	// Pass in concrete app types that implement the interfaces defined in https://github.com/cosmos/interchain-security/testutil/e2e/interfaces.go
	// IMPORTANT: the concrete app types passed in as type parameters here must match the
	// concrete app types returned by the relevant app initers.
	ccvSuite := integration.NewCCVTestSuite[*gaiaApp.GaiaApp, *appConsumer.App](
		// Pass in ibctesting.AppIniters for gaia (provider) and consumer.
		GaiaAppIniter, icstestingutils.ConsumerAppIniter, []string{})

	// Run tests
	suite.Run(t, ccvSuite)
}

// GaiaAppIniter implements ibctesting.AppIniter for the gaia app
func GaiaAppIniter() (ibctesting.TestingApp, map[string]json.RawMessage) {
	encoding := gaiaApp.RegisterEncodingConfig()
	app := gaiaApp.NewGaiaApp(
		log.NewNopLogger(),
		tmdb.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		gaiaApp.DefaultNodeHome,
		encoding,
		gaiaApp.EmptyAppOptions{})

	testApp := ibctesting.TestingApp(app)
	return testApp, gaiaApp.NewDefaultGenesisState(encoding)
}
