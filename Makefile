

.PHONY: gen-api
gen-api:
	goctl api go --api ./desc/website.api --dir ./


# no cache, only database
.PHONY: gen-model
gen-model:
	goctl model mysql ddl -src="./doc/website-v1.0.sql" -dir="./model"


.PHONY: gen-dockerfile
gen-dockerfile:
	goctl docker -go ./waitlist.go



.PHONY: website-api
website-api:
	go run ./waitlist.go
