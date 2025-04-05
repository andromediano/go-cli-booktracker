# CLI Book Tracker App

## go.mod, go.sum
`react`로 치자면 `package.json`과 `pnpm-lock.yaml`(또는 package-lock.json, yarn.lock)로 보면 되겠다.

### go.mod
go.mod 파일은 Go 모듈의 정의와 의존성 정보를 포함하는 핵심 파일입니다.

**주요 기능:**
- 모듈 정의: 모듈의 경로를 지정하며, 이는 프로젝트의 고유 식별자 역할을 합니다(예: 저장소 URL).
- Go 버전 지정: 해당 모듈에서 사용 가능한 최소 Go 버전을 명시합니다.
- 의존성 관리: 프로젝트에서 필요한 외부 모듈(패키지)과 그 버전을 나열합니다. require 키워드를 통해 의존성을 선언합니다.
- 모듈 대체 및 제외: 특정 버전의 모듈을 다른 버전으로 대체하거나, 특정 버전을 제외할 수 있습니다(replace, exclude 키워드 사용).

**생성 및 관리:**
go mod init <module-path> 명령으로 생성됩니다.
이후 go get, go mod tidy 등의 명령어를 통해 의존성 정보가 자동으로 업데이트됩니다.

**버전 업데이트**
```sh
go mod edit -go=1.24.2
```
버전을 변경한 후 의존성을 정리하기 위해 다음 명령을 내려준다.
```sh
go mod tidy
```

참고로 파일을 직접 수정해도 된다. 그 이후 `go mod tidy`실행해도 무방하다.
하지만 실수를 줄이기 위해서는 명령어에 익숙해지자.

### go.sum
go.sum 파일은 go.mod에 명시된 의존성에 대한 체크섬 정보를 저장하는 파일입니다.

**주요 기능:**
- 무결성 검증: 각 의존성(모듈)의 정확한 내용을 보장하기 위해 암호화된 체크섬을 기록합니다. 이를 통해 의존성이 변경되거나 손상되지 않았음을 확인할 수 있습니다.
- 모든 의존성 기록: 직접적인 의존성뿐만 아니라 간접적인 의존성(하위 모듈)까지 포함하여 체크섬 정보를 저장합니다.
버전 관리: 동일한 모듈의 여러 버전에 대한 체크섬도 기록하여 과거 버전과의 호환성을 유지합니다.

**생성 및 관리:**
- go mod tidy 또는 go get 명령을 실행하면 자동으로 생성 및 업데이트됩니다.
- 개발자가 직접 수정할 필요는 없으며, 항상 최신 상태로 유지됩니다
