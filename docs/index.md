# moar

moar (pronounce "more") is a modular augmentation registry for VueJS and ReactJS.

# But why?

Based on [Distributed vue applications](https://markus.oberlehner.net/blog/distributed-vue-applications-loading-components-via-http/)
by Markus Oberlehner this registry is a central hub for managing UMD modules for VueJS and ReactJS applications.

# Try it

1. Clone this repository
2. Start a dev instance of moar with docker compose:
    ```bash
    docker-compose -f ./docker/docker-compose.yml up
    ```
3. Clone the Vue example repository
    ```bash
    git clone github.com/nadilas/moar-vue-example
    ```
   1. Install moar-cli
      ```bash
      brew install moar
      ```
   2. Create your first module: 
       ```bash
       moar m c HelloWorld -a author@domain.com -l vue
       ```
   3. Upload a version of the module
       ```bash
       moar v upload -m HelloWorld -v 0.0.1+1394b72e1ef0fdc7b047 server/components/HelloWorld/HelloWorld.1394b72e1ef0fdc7b047.umd.min.js server/components/HelloWorld/HelloWorld.1394b72e1ef0fdc7b047.css
       ```
4. Making sure everything is in place:
   <details>
    <summary>Click to expand: via curl</summary>
   
    ```bash
      curl --request POST \
      --url http://localhost:8000/moarpb.ModuleRegistry/GetModule \
      --header 'Content-Type: application/json' \
      --data '{"moduleName": "HelloWorld"}' | jq
    ```

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
   </details>

    Or by running the hello world Vue application:
    ```bash
    npm run serve
    ```

