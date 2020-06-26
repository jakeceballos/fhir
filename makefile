.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: generate_r4
generate_r4: 
	bin/schema-generate -d R4 -p R4 schemas/fhir-r4.schema.json
