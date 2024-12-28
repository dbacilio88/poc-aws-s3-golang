package config

import (
	"fmt"
	"github.com/dbacilio88/poc-aws-s3-golang/pkg/utils"
	"os"
	"path/filepath"
	"testing"
)

/**
*
* properties_test
* <p>
* properties_test file
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
* @since 26/12/2024
*
 */

func TestRabbitmq_GetExchange(t *testing.T) {
	tests := []struct {
		name     string
		rabbitmq Rabbitmq
		want     string
	}{
		{
			name: "TestRabbitmq_GetExchange",
			rabbitmq: Rabbitmq{
				Exchange: Exchange{
					Durable: false,
					Name:    "exchange",
					Type:    "direct",
				},
			},
			want: "exchange",
		}, {
			name: "EmptyExchange",
			rabbitmq: Rabbitmq{
				Exchange: Exchange{
					Durable: false,
					Name:    "",
					Type:    "direct",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			YAML.Rabbitmq = tt.rabbitmq

			if got := tt.rabbitmq.GetExchange(); got != tt.want {
				t.Errorf("GetExchange() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestRabbitmq_GetRoutingKey(t *testing.T) {
	tests := []struct {
		name     string
		rabbitmq Rabbitmq
		want     string
	}{
		{
			name: "TestGetRoutingKey",
			rabbitmq: Rabbitmq{
				RoutingKey: RoutingKey{
					Request:  "rk.request",
					Response: "rk.response",
				},
			},
			want: "rk.request",
		},
		{
			name: "EmptyRoutingKey",
			rabbitmq: Rabbitmq{
				RoutingKey: RoutingKey{
					Request:  "",
					Response: "",
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			YAML.Rabbitmq = tt.rabbitmq
			if got := YAML.Rabbitmq.GetRoutingKey(); got != tt.want {
				t.Errorf("GetRoutingKey() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestRabbitmq_GetQueueName(t *testing.T) {
	tests := []struct {
		name     string
		rabbitmq Rabbitmq
		want     string
	}{
		{
			name: "TestGetQueueName",
			rabbitmq: Rabbitmq{
				Queue: Queue{
					Name:       "queue-test",
					Durable:    false,
					MessageTtl: 0,
					Type:       "topic",
				},
			},
			want: "queue-test",
		}, {
			name: "EmptyQueueName",
			rabbitmq: Rabbitmq{
				Queue: Queue{
					Name:       "",
					Durable:    false,
					MessageTtl: 0,
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			YAML.Rabbitmq = tt.rabbitmq
			got := tt.rabbitmq.GetQueueName()
			if got != tt.want {
				t.Errorf("GetQueueName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRabbitmq_GetVhost(t *testing.T) {
	tests := []struct {
		name     string
		rabbitmq Rabbitmq
		want     string
	}{
		{
			name: "TestGetVhost",
			rabbitmq: Rabbitmq{
				Vhost: "test",
			},
			want: "test",
		}, {
			name: "EmptyVhost",
			rabbitmq: Rabbitmq{
				Vhost: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			YAML.Rabbitmq = tt.rabbitmq
			got := tt.rabbitmq.GetVhost()
			if got != tt.want {
				t.Errorf("GetVhost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRabbitmq_GetUri(t *testing.T) {
	tests := []struct {
		name     string
		rabbitmq Rabbitmq
		want     string
	}{
		{
			name: "TestGetUri",
			rabbitmq: Rabbitmq{
				Host:     "localhost",
				Port:     5672,
				User:     "guest",
				Password: "guest",
				Protocol: "amqp",
			},
			want: "amqp://guest:guest@localhost:5672/",
		}, {
			name: "EmptyUri",
			rabbitmq: Rabbitmq{
				Host:     "",
				Port:     0,
				User:     "",
				Password: "",
				Protocol: "",
				Vhost:    "",
			},
			want: "://:@:0/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			YAML.Rabbitmq = tt.rabbitmq
			got := tt.rabbitmq.GetUri()
			if got != tt.want {
				t.Errorf("GetUri() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var rootPath = ""

func setup() {
	uh := utils.NewHelper("../")
	rootPath = uh.RootDir
}
func teardown() {
	// Limpia las variables de entorno después de las pruebas
	_ = os.Unsetenv("CONFIG_PATH")
}
func TestLoadProperties(t *testing.T) {
	setup()
	defer teardown()
	rootFilePath, _ := filepath.Abs(filepath.Join(rootPath, "test", "config", "valid"))
	fmt.Println("rootFilePath:", rootFilePath)
	rootInvalidFilePath, _ := filepath.Abs(filepath.Join(rootPath, "test", "config", "invalid"))
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		env     string
		args    args
		wantErr bool
	}{
		{
			name:    "TestValidPathLoadProperties",
			env:     "CONFIG_PATH",
			args:    args{path: rootFilePath},
			wantErr: false,
		},
		{
			name:    "TestInvalidPathLoadProperties",
			env:     "CONFIG_PATH",
			args:    args{path: rootInvalidFilePath},
			wantErr: true,
		}, {
			name:    "TestEmptyPathLoadProperties",
			env:     "CONFIG_PATH",
			args:    args{path: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.path != "" {
				_ = os.Setenv("CONFIG_PATH", tt.args.path)
				//setup(tt.env, tt.path)
				//err := LoadProperties()
				//defer teardown()
				//if tt.wantErr {
				//	require.Error(t, err)
				//} else {
				//	require.NoError(t, err)
				//}
			} else {
				_ = os.Unsetenv("CONFIG_PATH")
			}

			err := LoadProperties()
			fmt.Println("err", err)
			//if (err != nil) != tt.wantErr {
			//	require.Error(t, err)
			//}
		})
	}
}
