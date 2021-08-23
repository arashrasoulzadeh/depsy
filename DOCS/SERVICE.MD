#### Service

Using this type, you can `restart` or `stop` or `start` or `reload` any linux service :

#### Sample

```yaml
- name: "working on services"
  type: "service"
  command: "restart"
  path: ""
  args:
    - key: "restart"
      value: "mysql"
    - key: "reload"
      value: "nginx"
    ...
  pass_on_error: false
```
