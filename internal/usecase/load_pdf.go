package usecase

import (
	"bytes"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/unidoc/unipdf/v3/model"
	"github.com/vldcreation/privy-pdf-go/internal/constants"
	"github.com/vldcreation/privy-pdf-go/internal/entity"
	"github.com/vldcreation/privy-pdf-go/internal/util"
)

func LoadPDF(inputPath string, opts *entity.PdfOpts) (*entity.PDF, error) {
	log.Info().Msgf("load pdf from %s", inputPath)
	log.Info().Msg("Identify pdf file")
	if opts != nil {
		log.Info().Msg("load pdf with opts")
		return loadWithOpts(inputPath, opts)
	}

	log.Info().Msg("load pdf without opts")
	return loadWithoutOpts(inputPath)
}

func loadWithoutOpts(inputPath string) (*entity.PDF, error) {
	br, err := getBytesReader(inputPath)

	pdf, err := model.NewPdfReader(br)
	if err != nil {
		return nil, err
	}

	return entity.NewPDF(pdf), nil
}

func loadWithOpts(inputPath string, opts *entity.PdfOpts) (*entity.PDF, error) {
	isEncrypted := opts.IsEncrypt()

	if !isEncrypted {
		return loadWithoutOpts(inputPath)
	}

	if opts.EncryptOpts() == nil {
		return nil, errors.New("encrypt opts is nil")
	}

	if opts.EncryptOpts().Password == "" {
		return nil, errors.New("encryption password is empty")
	}

	if opts.OutputPath() == nil {
		fn, ext, err := util.ParseFile(inputPath)
		if err != nil {
			return nil, err
		}

		fnWithExt := fn + "." + ext

		err = util.SetUnlockFilename(&fnWithExt)
		if err != nil {
			return nil, err
		}

		fpath := util.GetFilePath(constants.OUTPUT_TEST_DATA, fnWithExt)

		opts.SetOutputPath(&fpath)
	}

	log.Info().Msg("Start unlock encrypted pdf")
	err := UnlockPDF(inputPath, *opts.OutputPath(), []byte(opts.EncryptOpts().Password))
	if err != nil {
		log.Fatal().Msgf("error while unlock pdf %+v\n", err)
		return nil, err
	}
	log.Info().Msg("Finish unlock encrypted pdf")

	return loadWithoutOpts(*opts.OutputPath())
}

func getBytesReader(inputPath string) (*bytes.Reader, error) {
	bt, err := util.LoadFile(inputPath)
	if err != nil {
		log.Fatal().Msgf("error while load file %+v\n", err)
		return nil, err
	}

	return bytes.NewReader(bt), nil
}
