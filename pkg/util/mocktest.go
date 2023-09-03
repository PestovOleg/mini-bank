package util

type MockLoggerConfig struct{}

func (m *MockLoggerConfig) GetAllConfig() []LogPathCfg {
	return []LogPathCfg{
		{
			Encoding: "console",
			Output:   "stdout",
			Level:    "debug",
		},
	}
}

func InitMockLogger() error {
	err := InitLogger(&MockLoggerConfig{})

	return err
}
