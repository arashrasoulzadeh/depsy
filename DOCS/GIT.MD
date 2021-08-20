#### Git

Using this type, you can `pull` or `hard reset` cloned repository:

#### Pull

```yaml
- name: "pull the repository"
  type: "git"
  command: "pull"
  path: "/var/www/html/test"
```

#### Reset Hard

This will hard reset the repository.

```yaml
- name: "hard reset repo"
  type: "git"
  command: "hard_reset"
  path: "/var/www/html/test"
```

