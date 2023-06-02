up:
	docker compose up --build -d

down:
	docker compose down

pub:
	cd src && go run publisher.go

sub:
	cd src && go run consumer.go

check:
	cd src && go run checker.go

create-superuser:
	bash create_superuser.sh