package gin_config

var __fxDeps []interface{}

func provide(dep interface{}) {
	__fxDeps = append(__fxDeps, dep)
}

func FxDeps() []interface{} {
	return __fxDeps
}
