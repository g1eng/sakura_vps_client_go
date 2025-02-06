NAME=sakura_vps_client_go

all: download_spec diff_spec extract_cc_spec generate_models modify_gitignore put_readme

check_diff: download_spec diff_spec 

put_readme:
	./scripts/put_readme.sh; 
	rm -rv spec/openapi.json

modify_gitignore:
	echo /.idea >> .gitignore
	echo /.*.swp >> .gitignore
	

generate_models: spec/spec.json
	openapi-generator generate \
		-i spec/spec.json \
		-g go \
		--package-name sakura_vps_client_go \
		--git-repo-id sakura_vps_client_go \
		--git-user-id g1eng \
		-o . ;\
	go mod tidy

diff_spec:
	[ -f spec/openapi.json ] \
		&& diff spec/openapi.json spec/openapi-next.json && exit 1 \
		|| mv -v spec/openapi-next.json spec/openapi.json

download_spec:
	[ -d spec ] || mkdir spec \
	&& curl -fSL -o spec/openapi-next.json https://manual.sakura.ad.jp/vps/api/api-doc/api-json.json

extract_cc_spec: spec/openapi.json
	./scripts/extract_spec_cc.sh spec/openapi.json spec/spec-tmp.json
	mv -v spec/spec-tmp.json spec/spec.json

test:
	go test -test.v ./...

clean:
	rm -rv *.go README.md go.* api docs spec test
