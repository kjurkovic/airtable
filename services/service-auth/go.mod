module github.com/kjurkovic/airtable/service/auth

go 1.18

require (
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/rs/cors v1.8.2
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/postgres v1.3.8
	gorm.io/gorm v1.23.8
	xorm.io/xorm v1.3.1
)

require (
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/kjurkovic/airtable/clients/service-audit v0.0.0-00010101000000-000000000000 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/kjurkovic/airtable/clients/service-audit => ../../clients/service-audit
