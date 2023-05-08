.PHONY: run
run: ## run main.
	go build -o plugin_type1/plugin_type1.so -buildmode=plugin plugin_type1/plugin_name1.go && \
	go build -o plugin_type2/plugin_type2.so -buildmode=plugin plugin_type2/plugin_name2.go && \
	go run main.go