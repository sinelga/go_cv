GOPATH=$GOPATH:/home/juno/git/go_cv go test -v

GOPATH=$GOPATH:/home/juno/git/go_cv go test -v ./...

bin/sitemap_maker -locale en_US -themes programming -linksdir  /home/juno/git/cv/version_desk_react_00/links -mapsdir /home/juno/git/go_cv/maps -contentsdir /home/juno/git/cv/version_desk_react_00/dist/www

bin/loader -locale en_US -themes programming -file programming.csv









