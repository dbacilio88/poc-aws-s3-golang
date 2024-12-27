package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

/**
*
* properties
* <p>
* properties file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 18/12/2024
*
 */

var YAML Properties

type Properties struct {
	Server    Server    `mapstructure:"server" yaml:"server"`
	Scheduler Scheduler `mapstructure:"scheduler" yaml:"scheduler"`
	Rabbitmq  Rabbitmq  `mapstructure:"rabbitmq" yaml:"rabbitmq" validate:"required"`
}

// Server use mapstructure in document github.com/go-viper/mapstructure/v2
type Server struct {
	Host        string `mapstructure:"host" yaml:"host" validate:"required"`
	Port        int    `mapstructure:"port" yaml:"port" validate:"required"`
	Name        string `mapstructure:"name" yaml:"name" validate:"required"`
	Timeout     int    `mapstructure:"timeout" yaml:"timeout" validate:"required"`
	Logging     string `mapstructure:"logging" yaml:"logging" validate:"required"`
	Environment string `mapstructure:"environment" yaml:"environment" validate:"required"`
	Logs        string `mapstructure:"logs" yaml:"logs" validate:"required"`
}

type Scheduler struct {
	Enable bool `mapstructure:"enable" yaml:"enable" validate:"required"`
}

type Rabbitmq struct {
	Exchange   Exchange   `mapstructure:"exchange" yaml:"exchange" validate:"required"`
	Host       string     `mapstructure:"host" yaml:"host"`
	Password   string     `mapstructure:"password" yaml:"password"`
	Port       int        `mapstructure:"port" yaml:"port"`
	Protocol   string     `mapstructure:"protocol" yaml:"protocol"`
	Queue      Queue      `mapstructure:"queue" yaml:"queue"`
	RoutingKey RoutingKey `mapstructure:"routing_key" yaml:"routing_key"`
	TlsEnabled bool       `mapstructure:"tls_enabled" yaml:"tls_enabled"`
	User       string     `mapstructure:"user" yaml:"user"`
	Vhost      string     `mapstructure:"vhost" yaml:"vhost"`
}
type Exchange struct {
	Durable bool   `mapstructure:"durable" yaml:"durable"`
	Name    string `mapstructure:"name" yaml:"name"`
	Type    string `mapstructure:"type" yaml:"type"`
}
type Queue struct {
	Name       string `mapstructure:"name" yaml:"name"`
	Durable    bool   `mapstructure:"durable" yaml:"durable"`
	MessageTtl int    `mapstructure:"message_ttl" yaml:"message_ttl"`
	Type       string `mapstructure:"type" yaml:"type"`
}
type RoutingKey struct {
	Request  string `mapstructure:"request" yaml:"request"`
	Response string `mapstructure:"response" yaml:"response"`
}

type ParameterBroker struct {
	Port       int
	Uri        string
	Exchange   string
	Queue      string
	Vhost      string
	RoutingKey string
}

type IParameterBroker interface {
	GetUri() string
	GetVhost() string
	GetQueueName() string
	GetRoutingKey() string
	GetExchange() string
}

func (r *Rabbitmq) GetUri() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/",
		YAML.Rabbitmq.Protocol,
		YAML.Rabbitmq.User,
		YAML.Rabbitmq.Password,
		YAML.Rabbitmq.Host,
		YAML.Rabbitmq.Port)
}
func (r *Rabbitmq) GetVhost() string { return r.Vhost }
func (r *Rabbitmq) GetQueueName() string {
	return r.Queue.Name
}
func (r *Rabbitmq) GetExchange() string {
	return r.Exchange.Name
}
func (r *Rabbitmq) GetRoutingKey() string {
	return r.RoutingKey.Request
}

func LoadProperties() error {

	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		return errors.New("la variable de entorno CONFIG_PATH no está definida")
	}

	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	// Si se desea pasar el archivo por variable de entorno:
	//viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var file viper.ConfigFileNotFoundError
		if errors.As(err, &file) {
			log.Fatalf("Error reading config file, %s", file)
			return err
		}
		return err
	}

	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Error writing config file, %s", err)
		return err
	}

	if err := viper.Unmarshal(&YAML); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
		return err
	}

	log.Println("Successfully loaded config")
	return nil
}
