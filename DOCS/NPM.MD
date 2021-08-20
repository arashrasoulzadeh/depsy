#### Npm

Using this type, you can `install` npm packages:

note packges will be installed in current directory.
#### Install Packages

```yaml
- name: "installing packages"
  type: "npm"
  command: "install"
  path: ""
  args:
    - key: "package"
      value: "sarafi"
```

You may also install as many packages as you want in same step:

```yaml
- name: "installing packages"
  type: "npm"
  command: "install"
  path: ""
  args:
    - key: "package"
      value: "sarafi"
    - key: "package"
        value: "underline"
    - key: "package"
        value: "nuxt"
    ...
```
 
