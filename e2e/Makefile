.PHONY: test
test:
	@echo "Running e2e capability tests..."
	nomad job run capabilities_job.hcl
	nomad job run capabilities_add_job.hcl  
	nomad job run capabilities_drop_job.hcl
	@echo "E2E tests completed"
