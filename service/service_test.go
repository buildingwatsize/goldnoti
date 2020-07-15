package service_test

import (
	"goldnoti/service"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// *****************
// ** SetupConfig **
// *****************

func Test_SetupConfig_ENV_dev_GetENV_Is_dev(t *testing.T) {
	input := "dev"
	expected := "dev"

	// start actual zone
	os.Setenv("ENV", input)
	service.SetupConfig("../config")
	// end actual zone

	assert.Equal(t, expected, viper.GetString("ENV"), "Expecting result should be TRUE")
}

func Test_SetupConfig_ENV_dev_Wrong_ConfigFile_Should_be_Panics(t *testing.T) {
	expectedPanics := func() {
		os.Setenv("ENV", "develop") // wrong env "develop"
		service.SetupConfig("../config")
	}

	assert.Panics(t, expectedPanics, "Expecting result should be PANICS")
}
