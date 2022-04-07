module service

go 1.18

require gorm.io/gorm v1.23.4
require "framework" v0.0.0
replace "framework" => "../framework"

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
)
