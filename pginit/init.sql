CREATE TABLE public.books
    (
      id Serial,
      name character varying(150),
      author character varying(100),
      year numeric(4,0),
      CONSTRAINT books_pkey PRIMARY KEY (id)
    );