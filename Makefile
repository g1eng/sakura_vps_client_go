NAME=sakura_vps_client_go
OPENAPI_VERSION=7.13.0

all: download_spec spec/spec-tmp.json diff_spec generate_models modify_gitignore put_readme

renew: download_spec spec/spec-tmp.json generate_models modify_gitignore put_readme

check_diff: download_spec spec/spec-tmp.json diff_spec 

put_readme:
	./scripts/put_readme.sh; 
	rm -rv spec/openapi-next.json

modify_gitignore:
	echo /.idea >> .gitignore
	echo /.*.swp >> .gitignore
	echo /spec/spec-tmp.json >> .gitignore
	echo /spec/openapi-next.json >> .gitignore
	

generate_models: spec/spec.json
	openapi-generator version | grep $(OPENAPI_VERSION)  || exit 1 \;
	openapi-generator generate \
		-i spec/spec.json \
		-g go \
		--package-name sakuravps \
		--git-repo-id sakura_vps_client_go \
		--git-user-id g1eng \
		-o . ;\
	go mod tidy

diff_spec:
	[ -f spec/spec.json ] \
		&& diff spec/spec.json spec/spec-tmp.json || exit 1

download_spec:
	[ -d spec ] || mkdir spec \
	&& curl -fSL -o spec/openapi-next.json https://manual.sakura.ad.jp/vps/api/api-doc/api-json.json

spec/spec-tmp.json: spec/openapi-next.json
	./scripts/extract_spec_cc.sh spec/openapi-next.json spec/spec-tmp.json


test:
	go test -test.v ./...

clean:
	rm -rv *.go README.md go.* api docs spec test
