
CREATE TABLE hengwei_users (
	id serial   PRIMARY KEY,
	name 		character varying(100) NOT NULL UNIQUE,
	password 	character varying(500) ,
	description text,
	attributes 	jsonb,
	source      character varying(50),
	locked_at   timestamp,
	created_at  timestamp,
	updated_at  timestamp
);

CREATE TABLE  hengwei_user_groups (
		id 		    serial  PRIMARY KEY UNIQUE,
		name   		character varying(100) NOT NULL,
		description text,
		parent_id 	integer REFERENCES hengwei_user_groups ON DELETE CASCADE,
		created_at 	timestamp,
		updated_at 	timestamp
);

CREATE TABLE  hengwei_users_and_user_groups (
		id 		    serial  PRIMARY KEY,
		user_id 	integer	REFERENCES hengwei_users ON DELETE CASCADE,
		group_id 	integer REFERENCES hengwei_user_groups ON DELETE CASCADE,
		UNIQUE(user_id,group_id)
);

CREATE TABLE  hengwei_roles (
		id 			serial PRIMARY KEY,
		name 		character varying(100) NOT NULL UNIQUE,
		description text,
		created_at 	timestamp,
		updated_at 	timestamp
);

CREATE TABLE  hengwei_users_and_roles (
		id 			serial  PRIMARY KEY,
		user_id  	integer REFERENCES hengwei_users ON DELETE CASCADE,
		role_id 	integer REFERENCES hengwei_roles ON DELETE CASCADE,
		UNIQUE(user_id,role_id)
);

CREATE TABLE  hengwei_permission_groups (
		id 		    serial  PRIMARY KEY UNIQUE,
		name   		character varying(100) NOT NULL ,
		parent_id 	integer REFERENCES hengwei_permission_groups ON DELETE CASCADE,
		description text,
		is_default  bool, 
		created_at 	timestamp,
		updated_at 	timestamp,
		UNIQUE(parent_id,name)
);

CREATE TABLE  hengwei_permission_groups_and_roles (
		id 					  serial  PRIMARY KEY,
		role_id 		      integer REFERENCES hengwei_roles ON DELETE CASCADE,
		group_id 		      integer REFERENCES hengwei_permission_groups ON DELETE CASCADE,
		create_operation  	  bool,
		delete_operation 	  bool,
		update_operation	  bool,
		query_operation 	  bool,
		UNIQUE(role_id,group_id)
);

CREATE TABLE  hengwei_permissions_and_groups (
		id               	serial  PRIMARY KEY,
		group_id 	   		integer REFERENCES hengwei_permission_groups ON DELETE CASCADE,
		permission_object 	varchar(50),
		type 	            integer,
		UNIQUE(group_id,permission_object)
);

CREATE TABLE  hengwei_online_users (
		user_id             serial  PRIMARY KEY,
		address 	   		text ,
		created_at 	timestamp,
		updated_at 	timestamp
);


