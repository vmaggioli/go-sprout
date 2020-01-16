# Sprout

## Summary
Sprout is a tool to allow developers to organize multiple repositories into a directory structure while also allowing execution of commands across all cloned repositories.

## Installation Instructions

**MacOS:**
1. Curl and unzip the package

```bash
curl -L https://github.com/Fair2Dare/go-sprout/releases/download/VERSION_HERE/kamino_macos_x64.tar.gz | tar xz -C /usr/local/bin
```

2. Dowload kamino_macos_x64.tar.gz from [releases](https://github.com/Fair2Dare/go-sprout/releases), extract the executable, and add it to your $PATH

**Windows:**

Download the latest release [releases](https://github.com/Fair2Dare/go-sprout/releases), extract the executable, and add it to your $PATH

**Linux:**
1. Curl and unzip the package

```bash
curl -L https://github.com/Fair2Dare/go-sprout/releases/download/VERSION_HERE/kamino_linux_x64.tar.gz | tar xz -C /usr/local/bin
```

2. Dowload kamino_linux_x64.tar.gz from [releases](https://github.com/Fair2Dare/go-sprout/releases), extract the executable, and add it to your $PATH

## Usage
### sprout_config.yml
Create a `.sprout_config.yml` inside your $HOME folder. This will contain your configuration for the folder structure, as well as what repos to select from.

```yaml
projects:
  - name: Project1
    projects:
      - name: Project2
        repos:
          - https://github.com/Fair2Dare/repo3.git
    repos:
      - https://github.com/Fair2Dare/repo1.git
      - git@github.com:Fair2Dare/repo2.git
  - name: Project3
    repos:
      - https://github.com/Fair2Dare/repo4.git
```

### Syntax
The format for inputting commands is `sprout {flags} {command} {args...}`

#### Flags
* -X or --verbose - Enable all logs including debug logs
* -h or --help - Print the help prompt

#### Commands
* create - generate the folder structure and clone all selected repos
* spread - all input after the word _spread_ will be treated as a command and ran in each repository
  ** NOTE: Currently can only be run in the root directory of the project (where you ran `sprout create`)