# config

This directory it's used from eggs install --custom to let an unattended
installation customized for you.

`sudo eggs install --unattended` is equivalent to `sudo install --custom un` so
it's easy to have more customizations.

For example, you can copy us.yaml to bliss.yaml and change user and password, if
you want, and call your custom configuration with:

`sudo eggs install --custom bliss`
