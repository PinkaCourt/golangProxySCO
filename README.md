# golangProxySCO

sco test proxy

/my-proxy-server
│
├── cmd
│ └── my-proxy-server
│ └── main.go # Точка входа в приложение
│
├── internal
│ ├── config
│ │ └── config.go # Конфигурация приложения
│ │
│ ├── handler
│ │ ├── http_handler.go # Обработка HTTP запросов
│ │ └── ws_handler.go # Обработка WebSocket соединений
│ │
│ ├── proxy
│ │ └── proxy.go # Логика проксирования
│ │
│ └── model
│ └── message.go # Модели данных (например, для JSON)
│
├── pkg
│ ├── logger
│ │ └── logger.go # Логирование
│ │
│ └── utils
│ └── utils.go # Утилиты и вспомогательные функции
│
├── test
│ └── integration_test.go # Интеграционные тесты
│
├── go.mod # Файл зависимостей Go
└── README.md # Документация проекта
