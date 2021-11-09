test:
	@mkdir -p test-reports
	@go test ./... -mod=vendor -json | tee test-reports/reports.jsons
	@go test ./... -mod=vendor -coverprofile=test-reports/cover.out 