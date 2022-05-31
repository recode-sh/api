# API

This repository contains the source code of the Recode API.

The Recode API is hosted in a private server in the sole goal of implementing the GitHub OAuth flow in a secured manner (ie: without exposing our GitHub secret in the <a href="https://github.com/recode-sh/cli">CLI</a>).

The base URL for the API is: <a href="https://api.recode.sh">https://api.recode.sh</a>

## Table of contents
- [Requirements](#requirements)
- [Usage](#usage)
- [API](#api)
  - [Routes](#routes)
  - [GitHub OAuth flow](#github-oauth-flow)
- [The future](#the-future)
- [License](#license)

## Requirements

The Recode API only requires `go >= 1.17` to run.

## Usage

To use the Recode API, the following steps need to be taken:

1. Clone the GitHub repository using: `git clone https://github.com/recode-sh/api.git`.

2. Create a file named `.env` at the root of the repository. (This file will contain the environment variables required by the API).

3. Use the `.env.dist` file to add the required environment variables to your `.env` file. 

4. Launch the API server using the command `go run main.go`. The server will listen on `http://127.0.0.1:8080` by default.


## API

The Recode API is implemented using the [Gin Web Framework](https://github.com/gin-gonic/gin).

### Routes

```bash
GET /github/oauth/callback
```

This is the sole route defined. This is where you will be redirected after authorizing the Recode application on GitHub.

The GitHub OAuth code will be exchanged for an access token and transmited to the Recode CLI using an HTTP redirect.

### GitHub OAuth flow

The GitHub OAuth flow is a multi-step process between the **Recode CLI**, **GitHub** and the **Recode API**:

1. First, the **Recode CLI** will create an HTTP server that will listen for connections on `127.0.0.1:${RANDOM_PORT}`.

2. Then, the **Recode CLI** will redirect you to **GitHub** to authorize the Recode application.

3. Once authorized, **GitHub** will redirect you to the **Recode API** that will exchange to passed OAuth code for an access token.

4. Once done, the **Recode API** will pass the received access token to the **Recode CLI** by redirecting you to `http://127.0.0.1:${RANDOM_PORT}?access_token=${GITHUB_ACCESS_TOKEN}`.

Given that GitHub doesn't support dynamic OAuth redirect uris, the random port used by the CLI is passed to the API via the OAuth `state` parameter. 

## The future

This project is **100% community-driven**, meaning that except for bug fixes <ins>**no more features will be added**</ins>. 

The only features that will be added are the ones that will be [posted as an issue](https://github.com/recode-sh/cli/issues/new) and that will receive a significant amount of upvotes **(>= 10 currently)**.

## License

Recode is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
