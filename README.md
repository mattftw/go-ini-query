## Features:
* Get key `./ini-query --file=testing.ini get --section baz --property foo`
* Set key `./ini-query --file=testing.ini set --section baz --property foo --value bar`
* Delete key `./ini-query --file=testing.ini delete --section baz --property foo`

## TO DO:
* Unit tests
* github ci
* package up for distros?

## Why does this exist?
1. I wanted something that would allow me to create / edit values in `.ini` files via commands
2. I wanted no dependencies