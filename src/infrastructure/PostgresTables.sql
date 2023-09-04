CREATE TABLE IF NOT EXISTS public."Car"
(
    vin character varying(255) COLLATE pg_catalog."default" NOT NULL,
    brand character varying(255) COLLATE pg_catalog."default" NOT NULL,
    model character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Car_pkey" PRIMARY KEY (vin)
)
