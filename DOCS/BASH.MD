#### Bash
Using this type, you can execute bash commands:

```yaml
- name: "create a html file"
  type: "bash"
  command: "touch a.html"
  path: "/var/www/html"
```

You may also use arguments that rendered into variables in curly brackets:
```yaml
- name: "create a html file"
  type: "bash"
  command: "touch {name}"
  path: "/var/www/html"
  args:
    - key: "name"
      value: "a.html"
```