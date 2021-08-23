module github.com/mocheer/nix

go 1.16

require (
	github.com/mocheer/pluto v1.0.0
	github.com/mocheer/xena v0.0.0-20210813111416-12a844dd8fc7
	github.com/urfave/cli/v2 v2.3.0
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.11
)

replace (
	github.com/mocheer/pluto => ../pluto
	github.com/mocheer/vesta => ../vesta
	github.com/mocheer/xena => ../xena
)
