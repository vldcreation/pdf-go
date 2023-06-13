package constants

import "path/filepath"

var (
	// Path to the config file
	ConfigPath              = filepath.Join("config.env")
	ROOT_FROM_CONFIG        = "../../"
	TEST_DATA               = filepath.Join("internal", "test_data")
	INPUT_QRCODE_TEST_DATA  = filepath.Join(TEST_DATA, "qrcode", "input") + "/"
	OUTPUT_QRCODE_TEST_DATA = filepath.Join(TEST_DATA, "qrcode", "output") + "/"
	INPUT_TEST_DATA         = filepath.Join(TEST_DATA, "input") + "/"
	OUTPUT_TEST_DATA        = filepath.Join(TEST_DATA, "output") + "/"
)
