# web response

## What does it do

Webresponse is a simple web server listening for your requests and answering back in JSON with some information 
about the host and the incoming request. Hereafter is an example of request answer :

```json
{
  "count": 3,
  "header": {
    "Accept": [
      "*/*"
    ],
    "Accept-Encoding": [
      "gzip, deflate"
    ],
    "Cache-Control": [
      "no-cache"
    ],
    "Postman-Token": [
      "8ec38e79-a970-4fdc-bf46-df45afa5a47c"
    ],
    "User-Agent": [
      "PostmanRuntime/2.4.1"
    ],
    "X-Forwarded-For": [
      "192.168.99.1"
    ],
    "X-Forwarded-Host": [
      "api1:8020"
    ],
    "X-Forwarded-Proto": [
      "http"
    ],
    "X-Forwarded-Server": [
      "7f07ebd1e153"
    ]
  },
  "host": "22ac75d0b467",
  "ips": [
    "127.0.0.1/8",
    "::1/128",
    "172.17.0.4/16",
    "fe80::42:acff:fe11:4/64"
  ],
  "url": {
    "Scheme": "",
    "Opaque": "",
    "User": null,
    "Host": "",
    "Path": "/v3/test/1234123412341234",
    "RawPath": "",
    "RawQuery": "start=2&end=3",
    "Fragment": ""
  }
}
```

## Docker integration with compose

Docker image can be found on [Docker Hub](https://hub.docker.com/r/sebastienfr/webresponse)
It integrates easily with Docker-compose as following :

```yaml
# simulated service 1
api1:
  container_name: api1
  image: sebastienfr/webresponse:latest
  restart: always
  ports:
    - "8040:8020"
  command: /go/bin/webresponse -port 8020
```

## Build

Standard local build

```shell
   make all
   make dockerBuild
```

### Git Commit Guidelines

#### Commit Message Format
Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special
format that includes a **type**, a **scope** and a **subject**:

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

Any line of the commit message cannot be longer 100 characters! This allows the message to be easier
to read on github as well as in various git tools.

#### Type
Must be one of the following:

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **refactor**: A code change that neither fixes a bug or adds a feature
* **perf**: A code change that improves performance
* **test**: Adding missing tests
* **chore**: Changes to the build process or auxiliary tools and libraries such as documentation
  generation

#### Scope
The scope could be anything specifying place of the commit change. For example `logger`, `db`, etc...

#### Subject
The subject contains succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize first letter
* no dot (.) at the end

#### Body
Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes"
The body should include the motivation for the change and contrast this with previous behavior.

#### Footer
The footer should contain any information about **Breaking Changes** and is also the place to
reference issues that this commit **Closes**.
