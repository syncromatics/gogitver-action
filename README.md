# Gogitver Action

[![Action](https://github.com/syncromatics/gogitver-action/workflows/build/badge.svg)](https://github.com/syncromatics/gogitver-action/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/syncromatics/gogitver-action)](https://goreportcard.com/report/github.com/syncromatics/gogitver-action)

This action wraps [gogitver](https://github.com/syncromatics/gogitver). It will output the version of the repository
as determined by the git history. This makes version information available to other steps, like a docker build and push.

Example:

## Using gogitver for a github release and docker release:

```yml
name: build
on: [push]
jobs:
  release:
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/master'
    steps: 
      - uses: syncromatics/gogitver-action@v0.0.1
        id: gogitver

      - name: Create Release
        id: create_release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ steps.gogitver.outputs.version }}
          release_name: Release ${{ steps.gogitver.outputs.version }}
          draft: false
          prerelease: false

      - uses: actions/checkout@v1

      - name: Ship the Docker image
        run: make ship
        env:
          VERSION: ${{ steps.gogitver.outputs.version }}
          DOCKER_USERNAME: ${{ secrets.DOCKER_USERNAME }}
          DOCKER_PASSWORD: ${{ secrets.DOCKER_PASSWORD }}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
