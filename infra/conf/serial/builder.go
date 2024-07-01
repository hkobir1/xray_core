package serial

import (
	"context"
	"io"

	"github.com/hkobir1/xray_core/common/errors"
	creflect "github.com/hkobir1/xray_core/common/reflect"
	"github.com/hkobir1/xray_core/core"
	"github.com/hkobir1/xray_core/infra/conf"
	"github.com/hkobir1/xray_core/main/confloader"
)

func MergeConfigFromFiles(files []string, formats []string) (string, error) {
	c, err := mergeConfigs(files, formats)
	if err != nil {
		return "", err
	}

	if j, ok := creflect.MarshalToJson(c); ok {
		return j, nil
	}
	return "", errors.New("marshal to json failed.").AtError()
}

func mergeConfigs(files []string, formats []string) (*conf.Config, error) {
	cf := &conf.Config{}
	for i, file := range files {
		errors.LogInfo(context.Background(), "Reading config: ", file)
		r, err := confloader.LoadConfig(file)
		if err != nil {
			return nil, errors.New("failed to read config: ", file).Base(err)
		}
		c, err := ReaderDecoderByFormat[formats[i]](r)
		if err != nil {
			return nil, errors.New("failed to decode config: ", file).Base(err)
		}
		if i == 0 {
			*cf = *c
			continue
		}
		cf.Override(c, file)
	}
	return cf, nil
}

func BuildConfig(files []string, formats []string) (*core.Config, error) {
	config, err := mergeConfigs(files, formats)
	if err != nil {
		return nil, err
	}
	return config.Build()
}

type readerDecoder func(io.Reader) (*conf.Config, error)

var ReaderDecoderByFormat = make(map[string]readerDecoder)

func init() {
	ReaderDecoderByFormat["json"] = DecodeJSONConfig
	ReaderDecoderByFormat["yaml"] = DecodeYAMLConfig
	ReaderDecoderByFormat["toml"] = DecodeTOMLConfig

	core.ConfigBuilderForFiles = BuildConfig
	core.ConfigMergedFormFiles = MergeConfigFromFiles
}
