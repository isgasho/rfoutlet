// Package config provides the config file schema and utilities to load the
// config into concrete outlet and outlet group types.
package config

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/imdario/mergo"
	"github.com/martinohmann/rfoutlet/internal/outlet"
	"github.com/martinohmann/rfoutlet/internal/schedule"
	"github.com/martinohmann/rfoutlet/pkg/gpio"
)

const (
	// DefaultListenAddress defines the default address to listen on.
	DefaultListenAddress = ":3333"

	// DefaultTransmitPin defines the default gpio pin for transmitting rf codes.
	DefaultTransmitPin uint = 17

	// DefaultReceivePin defines the default gpio pin for receiving rf codes.
	DefaultReceivePin uint = 27

	// DefaultProtocol defines the default rf protocol.
	DefaultProtocol int = 1

	// DefaultPulseLength defines the default pulse length.
	DefaultPulseLength uint = 189
)

// DefaultConfig contains the default values which are chosen if a file is
// omitted in the config file.
var DefaultConfig = Config{
	ListenAddress: DefaultListenAddress,
	GPIO: GPIOConfig{
		ReceivePin:         DefaultReceivePin,
		TransmitPin:        DefaultTransmitPin,
		DefaultPulseLength: DefaultPulseLength,
		DefaultProtocol:    DefaultProtocol,
		TransmissionCount:  gpio.DefaultTransmissionCount,
	},
}

// Config is the structure of the config file.
type Config struct {
	ListenAddress    string              `json:"listenAddress"`
	StateFile        string              `json:"stateFile"`
	DetectStateDrift bool                `json:"detectStateDrift"`
	GPIO             GPIOConfig          `json:"gpio"`
	OutletGroups     []OutletGroupConfig `json:"outletGroups"`
}

// GPIOConfig is the structure of the gpio config section.
type GPIOConfig struct {
	ReceivePin         uint `json:"receivePin"`
	TransmitPin        uint `json:"transmitPin"`
	DefaultPulseLength uint `json:"defaultPulseLength"`
	DefaultProtocol    int  `json:"defaultProtocol"`
	TransmissionCount  int  `json:"transmissionCount"`
}

// OutletGroupConfig is the structure of the config for a single outlet group.
type OutletGroupConfig struct {
	ID          string         `json:"id"`
	DisplayName string         `json:"displayName"`
	Outlets     []OutletConfig `json:"outlets"`
}

// OutletConfig is the structure of the config for a single outlet.
type OutletConfig struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	CodeOn      uint64 `json:"codeOn"`
	CodeOff     uint64 `json:"codeOff"`
	Protocol    int    `json:"protocol"`
	PulseLength uint   `json:"pulseLength"`
}

// BuildOutletGroups builds outlet groups from c.
func (c Config) BuildOutletGroups() []*outlet.Group {
	groups := make([]*outlet.Group, len(c.OutletGroups))

	for i, gc := range c.OutletGroups {
		outlets := make([]*outlet.Outlet, len(gc.Outlets))

		for j, oc := range gc.Outlets {
			o := &outlet.Outlet{
				ID:          oc.ID,
				DisplayName: oc.DisplayName,
				CodeOn:      oc.CodeOn,
				CodeOff:     oc.CodeOff,
				Protocol:    oc.Protocol,
				PulseLength: oc.PulseLength,
				Schedule:    schedule.New(),
				State:       outlet.StateOff,
			}

			if o.DisplayName == "" {
				o.DisplayName = o.ID
			}

			if o.PulseLength == 0 {
				o.PulseLength = c.GPIO.DefaultPulseLength
			}

			if o.Protocol == 0 {
				o.Protocol = c.GPIO.DefaultProtocol
			}

			outlets[j] = o
		}

		g := &outlet.Group{
			ID:          gc.ID,
			DisplayName: gc.DisplayName,
			Outlets:     outlets,
		}

		if g.DisplayName == "" {
			g.DisplayName = g.ID
		}

		groups[i] = g
	}

	return groups
}

// LoadWithDefaults loads config from file and merges in the default config for
// unset fields.
func LoadWithDefaults(file string) (*Config, error) {
	config, err := Load(file)
	if err != nil {
		return nil, err
	}

	err = mergo.Merge(config, DefaultConfig)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Load loads the config from a file.
func Load(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return LoadWithReader(f)
}

// LoadWithReader loads the config using reader.
func LoadWithReader(r io.Reader) (*Config, error) {
	c, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal(c, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
