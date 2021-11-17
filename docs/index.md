## moar

moar (pronounce "more") is a modular augmentation registry for VueJS and ReactJS apps.
The registry is a central hub for managing module (remote component) versions.

Grants the ability to use remote components and switch between versions without redeploying the frontend application. The registry focuses on:
- Use of [Semantic Versions](http://semver.org/)
- Storing and serving multiple versions of a module
- Accessing modules by using version constraints (e.g. `1.2.x` is equivalent to `>= 1.2.0, < 1.3.0`) [Semver constraints](SEMVER.md)

## dependencies

- minio for module storage
- reverse-proxy for serving modules (e.g. nginx)

## try it

1. Clone this repository
2. Start an instance of moar with docker compose:
    ```bash
    cd docker & docker-compose up
    ```
3. Clone the Vue example repository
    ```bash
    git clone github.com/dotindustries/moar-vue-example
    ```
   1. Install moar-cli
      ```bash
      brew install dotindustries/tap/moarctl
      ```
   2. Create your first module: 
       ```bash
       moarctl m c HelloWorld -a author@domain.com -l vue
       ```
   3. Upload a version of the module
       ```bash
       moarctl v upload -m HelloWorld -v 0.0.1+1394b72e1ef0fdc7b047 server/components/HelloWorld/HelloWorld.1394b72e1ef0fdc7b047.umd.min.js server/components/HelloWorld/HelloWorld.1394b72e1ef0fdc7b047.css
       ```
4. Verify that everything is in place:
   
    Query:

    ```bash
      curl --request POST \
      --url http://localhost:8000/moarpb.ModuleRegistry/GetModule \
      --header 'Content-Type: application/json' \
      --data '{"moduleName": "HelloWorld"}' | jq
    ```
    Response:

    ```json
    {
      "module": {
        "name": "HelloWorld",
        "versions": [
          {
            "value": "0.0.1+1394b72e1ef0fdc7b047",
            "resources": {
              "scriptUri": "modules/HelloWorld/HelloWorld@0.0.1+1394b72e1ef0fdc7b047.js",
              "styleUri": "modules/HelloWorld/HelloWorld@0.0.1+1394b72e1ef0fdc7b047.css"
            }
          }
        ],
        "author": "author@domain.com",
        "language": "vue"
      }
    }
    ```

    Or by running the hello world Vue application:
    ```bash
    npm run serve
    ```

