NAME=sakura_vps_client_go

all: download_spec extract_cc_spec generate_models modify_gitignore put_readme

put_readme:
	./put_readme.sh;\
	rm -rv spec

modify_gitignore:
	echo /.idea >> .gitignore
	

generate_models: spec/spec.json
	openapi-generator generate \
		-i spec/spec.json \
		-g go \
		--package-name sakura_vps_client_go \
		--git-repo-id sakura_vps_client_go \
		--git-user-id g1eng \
		-o .

download_spec:
	[ -d spec ] || mkdir spec \
	&& curl -fSL -o spec/openapi.json https://manual.sakura.ad.jp/vps/api/api-doc/api-json.json

extract_cc_spec:
	./extract_spec_cc.sh

test:
	go test -test.v ./...

clean:
	rm -rv *.go README.md go.* api docs spec test
