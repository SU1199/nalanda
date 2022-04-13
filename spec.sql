create table student(
	id serial not null,
	enum bigint primary key not null,
	expDate text,
	homeLib text,
	category text,
	fName text,
	lName text,
	dob text,
	gender text,
	contactNote  text,
	bookCategory text,
	rollNo text
)

create table address(
    id serial not null,
    enum bigint references student(enum),
    street text,
    address text,
    addressTwo text,
    city text,
    state text,
    country text,
    zip text
)
  
create table contact(
    id serial not null,
    enum bigint references student(enum),
    primaryPhone text,
    secondaryPhone text,
    otherPhone text,
    primaryEmail text,
    secondaryEmail text,
    fax text
)

create table raw(
  id serial not null,
  enum bigint references student(enum),
  profile text,
  checkout text
 )


create table permutations(
    id serial not null,
    enum bigint not null,
    year int not null,
    ts timestamp,
    hit int,
    complete int
);