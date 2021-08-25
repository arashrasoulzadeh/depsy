![go](https://github.com/arashrasoulzadeh/deepzy/actions/workflows/go.yml/badge.svg)
![codeql-analytis](https://github.com/arashrasoulzadeh/deepzy/actions/workflows/codeql-analysis.yml/badge.svg)

Simple application to manage basic deployments

#### Usage
You need to create a deployment yaml file ([sample](deploy.yaml)). then run the binary file with `--config` flag to `path` or `url` of deployment file.

for example : 
```yaml
- name: folder
  become: true
  exec:
    - name: "test creating a new folder"
      type: "folder"
      command: "create"
      path: "/var/www/html/test"
      pass_on_error: true
```

#### Become Root
You need to specify `become : true` in each step that you want to be run only in root environment. default if false.

#### Pass On error
You may specify `pass_on_error: true` in each step that you want to be skipped if an error occured. default if false.



#### Supported types:

- [Bash](DOCS/BASH.MD)
- [Composer](DOCS/COMPOSER.MD)
- [Folder](DOCS/FOLDER.MD)
- [Git](DOCS/GIT.MD)
- [Maria](DOCS/MARIA.MD)
- [Mysql](DOCS/MYSQL.MD)
- [Nginx](DOCS/NGINX.MD)
- [NPM](DOCS/NPM.MD)
- [PM2](DOCS/PM2.MD)
- [Service](DOCS/SERVICE.MD)


#### Contribution & Donation
You are free to fork project and send me PR.
You may also [Buy me a Coffee](https://dono.gg/arashrasoulzadeh)