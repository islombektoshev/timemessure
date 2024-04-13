package app_props

import (
	"github.com/magiconair/properties"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"strings"
)

const FileName = "app"
const FileExtension = ".properties"

var RelativePath = ""

type AppProperties map[string]string

func init() {
	provide(NewAppProperties)
}

func NewAppProperties() AppProperties {
	fileProperties := properties.NewProperties()
	loadAndMerge(RelativePath+FileName+FileExtension, fileProperties)
	profile := findProfile(fileProperties)
	log.Info().Str("profile", profile).Msg("Active profile selected")

	if len(profile) > 0 {
		loadAndMerge(RelativePath+FileName+"-"+profile+FileExtension, fileProperties)
	}

	var appProps AppProperties = fileProperties.Map()
	appProps.mergeEnvVariables()
	return appProps
}

func (props *AppProperties) mergeEnvVariables() {
	for _, env := range os.Environ() {
		if len(env) > 1 {
			keyValue := strings.Split(env, "=")
			(*props)[keyValue[0]] = keyValue[1]
		}
	}
}

func (props *AppProperties) MustParseInt(key string) int {
	val := (*props)[key]
	i, e := strconv.ParseInt(val, 10, 32)
	if e != nil {
		log.Err(e).Str("key", key).Str(val, val).Msg("Cannot parse prop")
		panic(e)
	}
	return int(i)
}

func (props *AppProperties) MustParseIntOrDefault(key string, def int) int {
	val := (*props)[key]
	if len(val) == 0 {
		return def
	}
	i, e := strconv.ParseInt(val, 10, 32)
	if e != nil {
		log.Err(e).Str("key", key).Str(val, val).Msg("Cannot parse prop")
		panic(e)
	}
	return int(i)
}

func (props *AppProperties) MustParseInt64(key string) int64 {
	val := (*props)[key]
	i, e := strconv.ParseInt(val, 10, 64)
	if e != nil {
		log.Err(e).Str("key", key).Str(val, val).Msg("Cannot parse prop")
		panic(e)
	}
	return i
}

func findProfile(appProps *properties.Properties) string {
	profile := appProps.GetString("profile", "")
	if envProfile := os.Getenv("profile"); len(envProfile) > 0 {
		profile = envProfile
	}
	if len(profile) == 0 {
		profile = os.Getenv("SPRING_PROFILES_ACTIVE")
	}
	return profile
}

func loadAndMerge(filename string, props *properties.Properties) {
	appProps, err := properties.LoadFile(filename, properties.UTF8)
	if err == nil {
		log.Info().Str("file", filename).Msg("Properties Loaded")
		props.Merge(appProps)
	} else {
		log.Error().Str("file", filename).Msg("Cannot load properties")
	}
}
