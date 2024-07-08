CREATE TABLE public.authors (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	sirname varchar(50) NOT NULL,
	biography varchar(50) NOT NULL,
	birthday date NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (id)
);

-- Permissions

ALTER TABLE public.authors OWNER TO pguser;
GRANT ALL ON TABLE public.authors TO pguser;

-- Insert

INSERT INTO authors (name, sirname, biography, birthday)
VALUES ('Автор01', 'А', 'что то пишет', '2024-07-01'),
('Автор02', 'А', 'и эта что то пишет', '2024-07-01');

