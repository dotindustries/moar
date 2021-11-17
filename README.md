[![Test](https://github.com/dotindustries/moar/actions/workflows/test.yml/badge.svg)](https://github.com/dotindustries/moar/actions/workflows/test.yml)

## moar

moar (pronounce "more") is a modular augmentation registry for VueJS and ReactJS apps.
The registry is a central hub for managing module (remote component) versions.

Grants the ability to use remote components and switch between versions without redeploying the frontend application.

## features
- Multiple versions of a module
- Use version selector when accessing a module [Semver constraints](docs/semver.md)

## try it

1. run docker-compose up in the ./docker directory to start the registry
2. create your first module
3. upload a version of the module
4. making sure everything is in place

## Install command line tool

### macos
```
brew install dotindustries/tap/moarctl
```

## credit
Based on [Distributed vue applications](https://markus.oberlehner.net/blog/distributed-vue-applications-loading-components-via-http/)
by Markus Oberlehner.

## contribute

If you find an issue or want to contribute please file an [issue](https://github.com/dotindustries/moar/issues)
or [create a pull request](https://github.com/dotindustries/moar/pulls).