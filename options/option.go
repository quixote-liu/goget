package options

type Option interface {
	Parse(c string) error
}

type OptionKeyVal struct {
	key string
	val string
}
