package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	appListen          = "app.listen"
	appReadTimeout     = "app.read_timeout"
	appWriteTimeout    = "app.write_timeout"
	appShutdownTimeout = "app.shutdown_timeout"

	loggingLevel  = "logging.level"
	loggingOutput = "logging.output"

	dbURL                = "db.url"
	dbMaxOpenConnections = "db.max_open_connections"
	dbMaxIdleConnections = "db.max_idle_connections"
	dbConnMaxLifetime    = "db.conn_max_lifetime"
)

// Manager represents the app configuration manager
type Manager struct {
	CfgReader *viper.Viper
}

// Init initializes the application configuration
func Init(path string) (*Manager, error) {
	configManager := &Manager{
		CfgReader: viper.New(),
	}
	configManager.CfgReader.SetConfigFile(path)
	configManager.setDefaults()

	err := configManager.CfgReader.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = configManager.checkRequiredProps(configManager.requiredProps())
	if err != nil {
		return nil, err
	}
	return configManager, nil
}

// AppListen retrieves the app listen TCP address from configuration file
func (m *Manager) AppListen() string {
	return m.CfgReader.GetString(appListen)
}

// AppReadTimeout retrieves the application server read timeout from configuration file
func (m *Manager) AppReadTimeout() time.Duration {
	return m.CfgReader.GetDuration(appReadTimeout)
}

// AppWriteTimeout retrieves the application server write timeout from configuration file
func (m *Manager) AppWriteTimeout() time.Duration {
	return m.CfgReader.GetDuration(appWriteTimeout)
}

// AppShutdownTimeout retrieves the application server shutdown timeout from configuration file
func (m *Manager) AppShutdownTimeout() time.Duration {
	return m.CfgReader.GetDuration(appShutdownTimeout)
}

// LoggingLevel retrieves the application logging level from configuration file
func (m *Manager) LoggingLevel() string {
	return m.CfgReader.GetString(loggingLevel)
}

// LoggingOutput retrieves the application logging output types from configuration file
func (m *Manager) LoggingOutput() []string {
	return m.CfgReader.GetStringSlice(loggingOutput)
}

// DBUrl retrieves the mysql database url connection string
func (m *Manager) DBUrl() string {
	return m.CfgReader.GetString(dbURL)
}

// DBMaxOpenConnections retrieves the mysql database amount of max open connections
func (m *Manager) DBMaxOpenConnections() int {
	return m.CfgReader.GetInt(dbMaxOpenConnections)
}

// DBMaxIdleConnections retrieves the mysql database amount of max idle connections
func (m *Manager) DBMaxIdleConnections() int {
	return m.CfgReader.GetInt(dbMaxIdleConnections)
}

// DBConnMaxLifetime retrieves the mysql database connection max lifetime
func (m *Manager) DBConnMaxLifetime() time.Duration {
	return m.CfgReader.GetDuration(dbConnMaxLifetime)
}

// setDefaults sets application default configs
func (m *Manager) setDefaults() {
	m.CfgReader.SetDefault(appListen, "0.0.0.0:8080")
	m.CfgReader.SetDefault(appReadTimeout, 10*time.Second)
	m.CfgReader.SetDefault(appWriteTimeout, 10*time.Second)
	m.CfgReader.SetDefault(appShutdownTimeout, 15*time.Second)
	m.CfgReader.SetDefault(loggingLevel, zap.InfoLevel.String())
	m.CfgReader.SetDefault(loggingOutput, []string{"app.log"})
	m.CfgReader.SetDefault(dbMaxOpenConnections, 100)
	m.CfgReader.SetDefault(dbMaxIdleConnections, 10)
	m.CfgReader.SetDefault(dbConnMaxLifetime, 120*time.Second)
}

// requiredProps retrieves the list of required config props
func (m *Manager) requiredProps() map[string]func() string {
	return map[string]func() string{
		dbURL: m.DBUrl,
	}
}

// checkRequiredProps checks if all required props are present in config file
func (m *Manager) checkRequiredProps(requiredProps map[string]func() string) error {
	for key, prop := range requiredProps {
		if strings.Trim(prop(), "\n ") == "" {
			return fmt.Errorf("%s must be set and should not be empty", key)
		}
	}
	return nil
}
