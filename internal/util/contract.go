package util

import (
	"errors"
	"net/http"
)

var (
	ErrFileIsDirMsg  = "file is a directory"
	ErrFileIsDirCode = http.StatusBadRequest
	ErrFileIsDir     = errors.New(ErrFileIsDirMsg)

	ErrFileNotFoundMsg  = "file not found"
	ErrFileNotFoundCode = http.StatusNotFound
	ErrFileNotFound     = errors.New(ErrFileNotFoundMsg)

	ErrFileInputNotValidMsg  = "file input not valid"
	ErrFileInputNotValidCode = http.StatusUnprocessableEntity
	ErrFileInputNotValid     = errors.New(ErrFileInputNotValidMsg)
)
