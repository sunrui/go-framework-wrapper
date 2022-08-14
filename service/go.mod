module service

go 1.19

require gorm.io/gorm v1.23.4 // indirect

require framework v0.0.0-00010101000000-000000000000

replace framework => ../framework

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/garyburd/redigo v1.6.3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/matoous/go-nanoid/v2 v2.0.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
)
