--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: reservations; Type: TABLE; Schema: public; Owner: arya
--

CREATE TABLE public.reservations (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    shuttle_id uuid NOT NULL,
    reserv_name character varying(255) NOT NULL,
    seat_number integer NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.reservations OWNER TO arya;

--
-- Name: shuttles; Type: TABLE; Schema: public; Owner: arya
--

CREATE TABLE public.shuttles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    shuttle_type character varying(100) NOT NULL,
    seats smallint NOT NULL,
    start_date timestamp without time zone NOT NULL,
    route_start character varying(100) NOT NULL,
    route_end character varying(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.shuttles OWNER TO arya;

--
-- Name: users; Type: TABLE; Schema: public; Owner: arya
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_password_check CHECK ((length((password)::text) >= 8)),
    CONSTRAINT users_username_check CHECK ((length((username)::text) >= 8))
);


ALTER TABLE public.users OWNER TO arya;

--
-- Data for Name: reservations; Type: TABLE DATA; Schema: public; Owner: arya
--

COPY public.reservations (id, shuttle_id, reserv_name, seat_number, user_id, created_at) FROM stdin;
6f7696bd-cece-4867-9e9f-b2c8b82cce45    155d0d1e-5d65-4ef6-83db-723b9051b7d7    arya rangga kusumas     15      70021222-3617-47d1-b68d-5316a1a96003    2024-03-25 04:10:59.227027
\.


--
-- Data for Name: shuttles; Type: TABLE DATA; Schema: public; Owner: arya
--

COPY public.shuttles (id, shuttle_type, seats, start_date, route_start, route_end, created_at) FROM stdin;
bed57f59-5022-40a0-8b1c-afdd6352e390    Space Shuttle   7       2024-04-01 08:00:00     Earth   Mars    2024-03-24 17:02:58.24925
155d0d1e-5d65-4ef6-83db-723b9051b7d7    SpaceX Starship 100     2024-05-15 12:00:00     Earth   Moon    2024-03-24 17:02:58.24925
69f3bc01-ad90-493f-927d-a981ad3c516b    Galactic Cruiser        50      2024-06-01 15:30:00     Mars    Jupiter 2024-03-24 17:02:58.24925
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: arya
--

COPY public.users (id, username, password, created_at) FROM stdin;
70021222-3617-47d1-b68d-5316a1a96003    arya12345       $2a$10$haA6UzBB7.bhxl1Zl6Pm0eB7xcKn7Jhmmv5iBrmsyviLVYzznkoki    2024-03-25 01:09:35.588563
\.


--
-- Name: reservations reservation_pkey; Type: CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservation_pkey PRIMARY KEY (id);


--
-- Name: shuttles shuttle_pkey; Type: CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.shuttles
    ADD CONSTRAINT shuttle_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: reservations reservation_shuttle_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservation_shuttle_id_fkey FOREIGN KEY (shuttle_id) REFERENCES public.shuttles(id);


--
-- Name: reservations reservation_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: arya
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservation_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--
