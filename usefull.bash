GOPATH=$GOPATH:/home/juno/git/go_cv go test -v

GOPATH=$GOPATH:/home/juno/git/go_cv go test -v ./...

bin/sitemap_maker -locale en_US -themes programming -linksdir  /home/juno/git/cv/version_desk_react_00/links -mapsdir /home/juno/git/go_cv/maps -contentsdir /home/juno/git/cv/version_desk_react_00/dist/www

# from wordstream.com
unzip \*.zip

for x in *.csv; do
	sed -i '1,5d' $x
	bin/loader -locale en_US -themes programming -dbloc '/home/juno/git/go_cv/en_US_programming.db' -file $x
	mv $x csv/

done	
for x in *.csv; do sed -i '1,5d' $x; bin/loader -locale en_US -themes programming -dbloc '/home/juno/git/go_cv/en_US_programming.db' -file $x; mv $x csv/; done

rm *.zip

scp  /home/juno/git/go_cv/en_US_programming.db 104.236.240.215:/home/juno/git/go_cv
#scp /home/juno/git/goFastCgi/goFastCgi/singo.db 104.236.240.215:/home/juno/git/goFastCgi/goFastCgi

bin/blogfeeder -locale en_US -themes programming -rootdir /home/juno/git/cv/version_desk_react_00 -topic golang  -title 'intro'
bin/blogfeeder -locale en_US -themes programming -rootdir /home/juno/git/cv/version_desk_react_00 -topic java  -title 'intro'
bin/blogfeeder -locale en_US -themes programming -rootdir /home/juno/git/cv/version_desk_react_00 -topic javascript  -title 'intro'












