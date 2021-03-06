# Wifi Password Service

### Overview

Randomly generates a password that you can email to your guests.

### Requirements

* Golang 1.9
* GNU Make

### Usage

Create a `~/.wps` folder.

Create a `~/.wps/config.json` file with the following shape:

```
{
  "users": [
    {
      "name": "guest0",
      "email": "guest0@yahoo.com"
    },
    {
      "name": "guest1",
      "email": "guest1@gmail.com"
    }
  ],
  "words": [
    "fringe",
    "essence",
    "star",
    "signum",
    "blankets",
    "aurelius",
    "schumpeter",
    "arrogance"
  ],
  "sender": {
    "name": "guestAdmin",
    "email": "guest_admin@gmail.com",
    "password": "dsadasddw98queqje9"
  },
  "smtp": {
    "server": "smtp.gmail.com",
    "port": "587"
  },
  "subject": "Weekly Password"
}
```

Create a similar file in the same folder named `config-5g.json` if you want to run one for a 5g network.

Run `make` to build the binary, then execute it: `bin/wps --help`.

*NOTE: For sending email using gmail, generate an App Password for the host machine. See [this](https://support.google.com/accounts/answer/185833) page for more info.
