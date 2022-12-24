# Golang API Example
A golang containerized API example
## Table of Contents

- [Documentation](#documentation)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Contributing](#contributing)

## Documentation (TODO)

## Prerequisites 

- Install Go
    - via Chocolately for Windows
    - via Homebrew for Mac
- Install Docker
    - via Chocolately for Windows
    - via Homebrew for Mac
- Create AWS CLI user via AWS IAM on AWS console 
- (Optional) Install [Docker VSCode Extension](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker). Helps with Syntax highlighting and autocompletion. 


## Getting Started

- Clone repo
- Run
    ```bash
        docker build -t task-api .
        docker run -t -i -p 8090:8090 task-api
    ```

## Contributing

Some of the plans I have for this repo are documented here in [GitHub Projects](https://github.com/users/codeherk/projects/1).

We appreciate feedback and contribution to this repo! Feel free to create a PR with your changes. I plan to add additional things like a PR template, etc. to make the contribution process more formal.  

