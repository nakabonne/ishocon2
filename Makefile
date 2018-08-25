MAKEFLAGS = --no-builtin-rules --no-builtin-variables --always-make
.DEFAULT_GOAL := help

SHELL  = /usr/bin/env bash

alp:
	sudo alp -r --sum -f $(file)

restart:
	sh scripts/restart.sh

rotate:
	sh scripts/rotate_alplog.sh

set-slow-log:
	sudo mysql -uishocon -pishocon -e "set global slow_query_log = 1"
	sudo mysql -uishocon -pishocon -e "set global long_query_time = 0.1"
	sudo mysql -uishocon -pishocon -e "set global log_queries_not_using_indexes = 1"

mysqldumpslow:
	sudo mysqldumpslow -s t /var/lib/mysql/ip-10-0-1-227-slow.log > ~/tmp/slow.log

restart-mysql:
	sudo mysql restart
