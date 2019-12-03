test:
	docker build -t testing .
	docker run -e GITHUB_REPOSITORY=syncromatics/proto-schema-registry \
		-e GITHUB_REF=refs/heads/master \
		testing