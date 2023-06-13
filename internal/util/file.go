package util

import (
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func LoadFile(fpath string) ([]byte, error) {
	// check if file exist
	// if not exist, return error
	_, err := IsFileExist(fpath)
	if err != nil {
		assert(err.Error(), err)
		return nil, err
	}

	f, err := os.Open(fpath)
	if err != nil {
		assert("failed to open file", err)
		return nil, err
	}
	defer f.Close()

	// @return []byte, err
	// io.ReadAll will read all data from reader
	// and return []byte and error
	return io.ReadAll(f)
}

func IsFileExist(fpath string) (bool, error) {
	f, err := os.Stat(fpath)
	if err != nil {
		assert("file not found", err)
		return false, err
	}

	if f.IsDir() {
		assert("file is directory", err)
		return false, ErrFileIsDir
	}

	return true, nil
}

func GetFilePath(fpath string, fname string) string {
	return fpath + "/" + fname
}

func assert(msg string, e error) {
	if e != nil {
		log.Fatal().Err(e).Msg(msg)
	}

	log.Info().Msg(msg)
}
