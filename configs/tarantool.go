package configs

import "github.com/caarlos0/env/v6"

// Tarantool содержит настройки для работы с Tarantool.
type Tarantool struct {
	// Addr содержит адрес для соединения с Tarantool.
	Address string `env:"RELATIONS_TARANTOOL_ADDRESS,required"`

	// User это имя пользователя для подключения.
	User string `env:"RELATIONS_TARANTOOL_USER,required"`

	// Pass это пароль для подключения.
	Pass string `env:"RELATIONS_TARANTOOL_PASS,required"`
}

// LoadTarantool - загружает настройки для работы с tarantool.
// возвращает ошибку, при неудачной попытке  загрузки.
func LoadTarantool() (*Tarantool, error) {
	tarantool := &Tarantool{}
	if err := env.Parse(tarantool); err != nil {
		return nil, err
	}

	return tarantool, nil
}
