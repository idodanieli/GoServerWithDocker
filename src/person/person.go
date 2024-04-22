package person

type Person struct {
	Name    string   `mapstructure:"name"`
	Age     int      `mapstructure:"age"`
	Gender  string   `mapstructure:"gender"`
	Friends []string `mapstructure:"friends"`
	Family  *Family  `mapstructure:"family"`
}

type Family struct {
	Girlfriend *Person `mapstructure:"girlfriend"`
}
