# igsr5's Portfolio Protocol Buffers ð


[![Auto release](https://github.com/igsr5/portfolio-proto/actions/workflows/create-release.yml/badge.svg)](https://github.com/igsr5/portfolio-proto/actions/workflows/create-release.yml)
[![Release gem](https://github.com/igsr5/portfolio-proto/actions/workflows/gem-release.yml/badge.svg)](https://github.com/igsr5/portfolio-proto/actions/workflows/gem-release.yml)
[![Release npm](https://github.com/igsr5/portfolio-proto/actions/workflows/npm-release.yml/badge.svg)](https://github.com/igsr5/portfolio-proto/actions/workflows/npm-release.yml)
[![Lint Code Base](https://github.com/igsr5/portfolio-proto/actions/workflows/linter.yml/badge.svg)](https://github.com/igsr5/portfolio-proto/actions/workflows/linter.yml)
![](https://camo.githubusercontent.com/029166d85f92969845201e59c3fcd8c8345556036155ff18140f6a9e796173a3/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d677265656e)


@igsr5 ã®ãã¼ããã©ãªãªãµã¤ãã«é¢ããProtobufã¹ã­ã¼ããprotoã®å®ç¾©ã¨åè¨èªå®è£ã®ã©ã¤ãã©ãªéä¿¡ãè¡ãã


Swagger UI - https://igsr5.github.io/portfolio-proto/

Frontend Repo - https://github.com/igsr5/portfolio-frontend  
Backend Repo - https://github.com/igsr5/portfolio-backend  


## Installation

### Ruby
https://rubygems.org/gems/igsr5-portfolio-proto
```sh
$ gem install igsr5-portfolio-proto
```

### Nodejs
https://www.npmjs.com/package/igsr5-portfolio-proto
```sh
$ npm i igsr5-portfolio-proto
```

### Golang
```sh
$ go get github.com/igsr5/portfolio-proto/go
```

## Design

```
.
âââ proto # protoå®ç¾©ãéç½®ãããæ°ããã¹ã­ã¼ããè¿½å ãããã¨ãã¯ãã®ãã£ã¬ã¯ããªãæ´æ°ãã¦PRãåºã
âââ ruby # protoå®ç¾©ã®Rubyå®è£ãèªåçæã³ã¼ãã§æ§æããã
âââ golang # protoå®ç¾©ã®Golangå®è£ãèªåçæã³ã¼ãã§æ§æããã
âââ nodejs # protoå®ç¾©ã®Nodejså®è£ãèªåçæã³ã¼ãã§æ§æããã
âââ swagger # protoå®ç¾©ããçæããã swagger å®ç¾©
```


## Development

æ°ããã¹ã­ã¼ããå®ç¾©ããããªã£ãéã¯ `/proto` ãã£ã¬ã¯ããªéä¸ã«protoå®ç¾©ãå®è£ããã
åè) https://github.com/igsr5/portfolio-proto/pull/6

1. `/proto`éä¸ã«ã¹ã­ã¼ããå®ç¾©
2. PR ãåºã
3. protoå®ç¾©ã«åé¡ããªããã°åè¨èªå®è£ãçæãããèªåã§PRãä½æãã(PRã«URLãã³ã¡ã³ãããã)
4. ã¬ãã¥ã¼OKãªãPRããã¼ã¸ãã
5. åè¨èªå®è£ãmasterã«ãã¼ã¸ãããã¨èªåGitHub Releaseãçæãã gem, npmãªã©ã«ãnew versionããªãªã¼ã¹ããã
6. ãªãªã¼ã¹ ð

## Contributing
https://github.com/igsr5/protosum
> GitHub Action x Protocol Buffers is an attempt to develop a miracle schema experience!
