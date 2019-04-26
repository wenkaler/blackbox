create schema box;
create table box.projects(
	id serial primary key,
	name varchar(225) not null unique,
	token varchar(50) not null unique,
	basic_setting json not null
);
create table box.settings(
	id serial primary key,
	id_project int4 not null,
	indexer varchar(225) not null,
	"token" varchar(50) not null unique,
	"status" varchar(10),
	setting json,
	last_update_date timestamp,
	foreign key(id_project) references box.projects(id)
);

-- Basic function --
create or replace function box.generic_token_v1() 
returns text 
language plpgsql
as $function$
declare 
	_sequence  text[]= '{0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z}';
	result text := '';
	i int4 := 0;
begin
	for i in 1..50 loop
		result := result || _sequence[1+random()*(array_length(_sequence, 1)-1)];
	end loop;
	return result;
end
$function$;

-- Project function --

create or replace function box.create_project_v1("name" varchar, _data json)
returns json
language plpgsql
as $function$
declare
	id_project int;
	result json;
begin
	if "name" isnull or "name" = '' then
		raise exception 'name must not be empty.';
	end if;
	insert into box.projects("name", "token", basic_setting) values("name", box.generic_token_v1(), _data) returning id into id_project;
	select row_to_json(f) into result from (select * from box.projects where id = id_project ) f;
	return result;
end
$function$;

create or replace function box.update_project_v1(t varchar, "name" varchar, _data json)
returns json
language plpgsql
as $function$
declare
	result json;
begin
	if "name" isnull or "name" = '' then
		raise exception 'name must not be empty.';
	end if;
	if _data isnull then
		raise exception 'setting must not be null.';
	end if;
	update box.projects set "name" = "name", basic_setting = _data where "token" = t;
	select row_to_json(f) into result from (select * from box.projects where "token" = t ) f;
	return result;
end
$function$;
	

create or replace function box.get_project_v1(t varchar)
returns json
language plpgsql
as $function$
declare
	result json;
begin
	select row_to_json(f) into result from (select * from box.projects where "token" = t) f;
	if result isnull then
		raise exception 'The token (%) does not exist in the box.projects table.', t;
	end if;
	return result;
end
$function$;

create or replace function box.remove_project_v1(t varchar)
returns void
language plpgsql
as $function$
declare 
	_id int;
begin
	select id into _id from box.projects where "token" = t;
	delete from box.settings where id_project = _id;
	delete from box.projects where "token" = t;
end
$function$;

create or replace function box.list_project(t varchar)
returns json[]
language plpgsql
as $function$
declare
	projectID int;
	result json[];
begin
	select id into projectID from box.projects where token = t;
	result = array(select row_to_json(f) from (select * from box.settings where id_project = projectID) f);
	return result;
end
$function$;

-- Setting function --
create or replace function box.update_setting_v1(t varchar, d json)
returns json
language plpgsql
as $function$
declare 
	result json;
begin
	update box.settings set "status" = 'new', setting = d where "token" = t;
	select row_to_json(f) into result from (select * from box.settings where "token" = t ) f;
	if result isnull then 
		raise exception 'The token (%) does not exist in the box.settings table.', t;		
	end if;
	return result;
end
$function$;

create or replace function box.initial_setting_v1(t varchar, indexer varchar)
returns json
language plpgsql
as $function$
declare
	id_setting int;
	id_project int;
	result json;
	"basic" json;
begin
	if indexer isnull or indexer = '' then
		raise exception 'Indexer must not be empty.';
	end if;
	select id, basic_setting into id_project, "basic" from box.projects where token = t;
	if "basic" isnull then
		raise exception 'The token (%) does not exist in the projects table.', t;
	end if;
	insert into box.settings(id_project, indexer, "token", "status", setting, last_update_date) values(id_project, indexer, box.generic_token_v1(), 'new', "basic", current_timestamp) returning id into id_setting;
	select row_to_json(f) into result from (select * from box.settings where id = id_setting ) f;
	if result isnull then 
		raise exception 'Something happened wrong empty json result indexer(%) project_token(%)', indexer, t;
	end if;
	return result;
end
$function$;

create or replace function box.get_setting_v1(t varchar) 
returns json 
language plpgsql
as $function$
declare 
	result json;
begin
	select row_to_json(f) into result from (select * from box.settings where "token" = t) f ;
	if result isnull then
		raise exception 'The token (%) does not exist in the box.settings table.', t;
	end if;
	if result->>'status' = 'new' then
		update box.settings set "status" = 'pending' where "token" = t;
	end if;
	return result;
end
$function$;

create or replace function box.confirm_setting_v1(t varchar)
returns json
language plpgsql
as $function$
declare
	result json;
begin
	update box.settings set "status" = 'done', last_update_date = current_timestamp  where "token" = t;
	select row_to_json(f) into result from (select * from box.settings where "token" = t ) f;
	if result isnull then 
		raise exception 'Something happened wrong empty json result setting_token(%)', t;
	end if;
	return result;
end
$function$;

create or replace function box.clean_unused_settings_v1(t varchar, "time" interval)
returns void
language plpgsql
as $function$
declare 
	projectID int;
begin
	projectID = 0;
	select id into projectID from box.projects where "token" = t;
	if projectID = 0 then 
		raise exception 'The token (%) does not exist in the box.projects table.', t;		
	end if;
	delete from box.settings  where last_update_date <=  current_timestamp - "time" and project_id = projectID;
end
$function$;
