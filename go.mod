module github.com/hochfrequenz/go-local-days

go 1.16

require (
	github.com/corbym/gocrest v1.0.5
	github.com/stretchr/testify v1.8.1
)

replace (
	github.com/hochfrequenz/go-local-days/germany => ./germany
	github.com/hochfrequenz/go-local-days/local_days => ./local_days
)
