CREATE TABLE public.books (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	authorid int4 NOT NULL,
	"year" int4 NOT NULL,
	isbn varchar(50) NOT NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.books OWNER TO pguser;
GRANT ALL ON TABLE public.books TO pguser;


-- public.books внешние включи

ALTER TABLE public.books ADD CONSTRAINT books_authorid_fkey FOREIGN KEY (authorid) REFERENCES public.authors(id) ON DELETE CASCADE;

-- Insert

INSERT INTO books (name, authorid, year, isbn)
VALUES ('Книга 01', 1, 2000, '123'),
('Книга 02', 1, 2001, '1234'),
('Книга 03', 2, 2000, '12345'),
('Книга 04', 2, 2001, '123456');
