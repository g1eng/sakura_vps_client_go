NAME=sakura_vps_client_go

all: download_spec extract_cc_spec diff_spec generate_models modify_gitignore put_readme

check_diff: download_spec extract_cc_spec diff_spec 

put_readme:
	./scripts/put_readme.sh; 
	rm -rv spec/openapi-next.json

modify_gitignore:
	echo /.idea >> .gitignore
	echo /.*.swp >> .gitignore
	

generate_models: spec/spec.json
	openapi-generator generate \
		-i spec/spec.json \
		-g go \
		--package-name sakura_vps \
		--git-repo-id sakura_vps_client_go \
		--git-user-id g1eng \
		-o . ;\
	go mod tidy

diff_spec:
	[ -f spec/spec.json ] \
		&& diff spec/spec.json spec/spec-tmp.json && exit 1 \
		|| mv -v spec/spec-tmp.json spec/spec.json

download_spec:
	[ -d spec ] || mkdir spec \
	&& curl -fSL -o spec/openapi-next.json https://manual.sakura.ad.jp/vps/api/api-doc/api-json.json

extract_cc_spec: spec/openapi-next.json 
	./scripts/extract_spec_cc.sh spec/openapi-next.json spec/spec-tmp.json

test:
	go test -test.v ./...

clean:
	rm -rv *.go README.md go.* api docs spec test
