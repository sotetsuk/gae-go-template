show-goenv:
	@printenv | grep 'GO'

clean:
	@rm -rf gopath/*

fmt:
	@goapp fmt ./...

get:
	# @goapp get ...

serve:
	@read -p "Which module?: " module; \
	goapp serve $$module

deploy:
	@read -p "Which module?: " module; \
	echo "---"; \
	gcloud projects list; \
	echo "---"; \
	read -p "Which project?: " project; \
	goapp deploy -application $$project modules/$$module
