# quickmenu

An extremely simple menu-creator for CLI-apps.

# changes

## v0.0.4

Add 4 *trigger-points**:

- `BeforeMenuFunc`
- `AfterMenuFunc`
- `BeforeExecFunc`
- `AfterExecFunc`

# examples

- [simple example](examples/simple/main.go) (quits after one choice has been made and uses `AfterMenuFunc`)
- [simple loop example](examples/simpleloop/main.go) (keeps looping until the user quits explicitely)
- [triggerpoints example](examples/triggerpoints/main.go) (keeps looping and uses all 4 trigger-points to style the menu and add some behaviour)
