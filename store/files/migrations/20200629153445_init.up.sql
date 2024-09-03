CREATE TABLE users (
	id uuid NOT NULL,
	firstname text NOT NULL,
	lastname text NOT NULL,
	created_at timestamp default current_timestamp,
	CONSTRAINT "pk_user_id" PRIMARY KEY (id)
);

insert into users values('3410daa9-5443-4285-b960-0964ca8b973b','john', 'doe');
