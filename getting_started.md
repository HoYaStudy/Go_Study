# Installation

## on Windows

> $ choco install golang

# Configuration

## GOROOT

`GOROOT`는 `Go`의 설치 경로이다.

예전에는 따로 환경변수로 설정해줬어야 했지만, 현재는 go binary에 해당 정보가 포함되어 있으므로 명시적으로 지정할 필요는 없다. 만약, 사용자가 설치 디렉토리를 변경했다면 환경변수 `$GOROOT`를 설정해야 한다.

Windows에서는 기본적으로 `C:\go`에 설치가 된다.

## GOPATH

환경변수 `$GOPATH`는 `Go` 언어를 개발하는 작업 공간이며 반드시 설정해야 하는 부분이다. `GOROOT`와 다른 곳이면 어디든 괜찮다.

Windows에서는 기본적으로 `$HOME/go`로 설정된다.

`GOPATH` 내부는 다음과 같이 구성된다.

```
GOPATH
  \-- bin/
  \-- pkg/
  \-- src/
```

Folder는 기본적으로 다음의 구조를 지켜야 한다.

```
$GOPATH/src/{Repository FQDN}/{Repository Path}
```

## Go REPL (Read-Eval-Print Loop)

`Go`에는 기본제공되는 REPL이 없어서 다음과 같이 3rd party tool을 설치해서 사용한다.

> $ go get -u github.com/motemen/gore/cmd/gore

* 실행: `$ gore`
* 종료: `<CTRL-d>` or `:q`

## Code Formatter

`Go`에는 code formatter가 내장되어 있다. `Go`의 철학으로 단일 style만을 지원하여 code 통일성을 쉽게 달성할 수 있다.

`gofmt` 명령어를 사용하여 수행하고, `-w` option을 사용해서 결과를 출력하는 대신 file을 덮어쓸 수 있다.

> $ gofmt -w \<filename>

## Code Linter

기본적으로 내장된 `go vet`과 공식 tool인 `golint`가 있다.

`go vet`은 bug의 원인이 될 것 같은 code를 찾아주는 tool이고, `golint`는 `Go`답지 않은 coding style을 찾아주는 tool이다.

`golint`는 다음과 같이 설치를 해야 사용 가능하다.

> $ go get -u golang.org/x/lint/golint

작성한 code에 lint를 적용하려면 다음과 같은 명령어를 입력한다.

> $ go vet \<file> or \<folder>

> $ golint \<file> or \<folder>

## Task Runner

`Go`는 task runner로 `Makefile`을 사용한다.

# Usage

## env

### GOROOT를 확인한다.

> $ go env GOROOT
