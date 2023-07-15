--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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
-- Name: gog_demo; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA gog_demo;


ALTER SCHEMA gog_demo OWNER TO postgres;

--
-- Name: postgis; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS postgis WITH SCHEMA public;


--
-- Name: EXTENSION postgis; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION postgis IS 'PostGIS geometry and geography spatial types and functions';


--
-- Name: driver_status; Type: TYPE; Schema: gog_demo; Owner: postgres
--

CREATE TYPE gog_demo.driver_status AS ENUM (
    'busy',
    'waiting',
    'afw'
);


ALTER TYPE gog_demo.driver_status OWNER TO postgres;

--
-- Name: request_status; Type: TYPE; Schema: gog_demo; Owner: postgres
--

CREATE TYPE gog_demo.request_status AS ENUM (
    'findingDriver',
    'waitingForDriver',
    'inProgress',
    'completed',
    'canceled'
);


ALTER TYPE gog_demo.request_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: customer; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.customer (
    customer_id integer NOT NULL,
    phone integer NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL
);


ALTER TABLE gog_demo.customer OWNER TO postgres;

--
-- Name: customer_customer_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.customer_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.customer_customer_id_seq OWNER TO postgres;

--
-- Name: customer_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.customer_customer_id_seq OWNED BY gog_demo.customer.customer_id;


--
-- Name: dock; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.dock (
    dock_id integer NOT NULL,
    name character varying(100) NOT NULL,
    location public.geometry(Point,4326) NOT NULL,
    lat double precision NOT NULL,
    lon double precision NOT NULL,
    working boolean NOT NULL
);


ALTER TABLE gog_demo.dock OWNER TO postgres;

--
-- Name: dock_dock_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.dock_dock_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.dock_dock_id_seq OWNER TO postgres;

--
-- Name: dock_dock_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.dock_dock_id_seq OWNED BY gog_demo.dock.dock_id;


--
-- Name: driver; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.driver (
    driver_id integer NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    vessel_id integer,
    status gog_demo.driver_status NOT NULL,
    balance numeric(19,4) DEFAULT 0 NOT NULL,
    cert_first_aid integer,
    cert_driving integer
);


ALTER TABLE gog_demo.driver OWNER TO postgres;

--
-- Name: driver_driver_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.driver_driver_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.driver_driver_id_seq OWNER TO postgres;

--
-- Name: driver_driver_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.driver_driver_id_seq OWNED BY gog_demo.driver.driver_id;


--
-- Name: request; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.request (
    request_id integer NOT NULL,
    customer_id integer,
    driver_id integer,
    departure_id integer,
    destination_id integer,
    price numeric(19,4) NOT NULL,
    status gog_demo.request_status NOT NULL
);


ALTER TABLE gog_demo.request OWNER TO postgres;

--
-- Name: request_request_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.request_request_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.request_request_id_seq OWNER TO postgres;

--
-- Name: request_request_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.request_request_id_seq OWNED BY gog_demo.request.request_id;


--
-- Name: vessel; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.vessel (
    vessel_id integer NOT NULL,
    model character varying(100) NOT NULL,
    seats integer NOT NULL,
    is_approved boolean DEFAULT false NOT NULL,
    location public.geometry(Point,4326) NOT NULL,
    lat double precision NOT NULL,
    lon double precision NOT NULL
);


ALTER TABLE gog_demo.vessel OWNER TO postgres;

--
-- Name: vessel_vessel_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.vessel_vessel_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.vessel_vessel_id_seq OWNER TO postgres;

--
-- Name: vessel_vessel_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.vessel_vessel_id_seq OWNED BY gog_demo.vessel.vessel_id;


--
-- Name: customer customer_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.customer ALTER COLUMN customer_id SET DEFAULT nextval('gog_demo.customer_customer_id_seq'::regclass);


--
-- Name: dock dock_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.dock ALTER COLUMN dock_id SET DEFAULT nextval('gog_demo.dock_dock_id_seq'::regclass);


--
-- Name: driver driver_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.driver ALTER COLUMN driver_id SET DEFAULT nextval('gog_demo.driver_driver_id_seq'::regclass);


--
-- Name: request request_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request ALTER COLUMN request_id SET DEFAULT nextval('gog_demo.request_request_id_seq'::regclass);


--
-- Name: vessel vessel_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.vessel ALTER COLUMN vessel_id SET DEFAULT nextval('gog_demo.vessel_vessel_id_seq'::regclass);


--
-- Data for Name: customer; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.customer (customer_id, phone, first_name, last_name) FROM stdin;
\.


--
-- Data for Name: dock; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.dock (dock_id, name, location, lat, lon, working) FROM stdin;
\.


--
-- Data for Name: driver; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.driver (driver_id, first_name, last_name, vessel_id, status, balance, cert_first_aid, cert_driving) FROM stdin;
\.


--
-- Data for Name: request; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.request (request_id, customer_id, driver_id, departure_id, destination_id, price, status) FROM stdin;
\.


--
-- Data for Name: vessel; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.vessel (vessel_id, model, seats, is_approved, location, lat, lon) FROM stdin;
\.


--
-- Data for Name: spatial_ref_sys; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.spatial_ref_sys (srid, auth_name, auth_srid, srtext, proj4text) FROM stdin;
\.


--
-- Name: customer_customer_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.customer_customer_id_seq', 1, false);


--
-- Name: dock_dock_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.dock_dock_id_seq', 1, false);


--
-- Name: driver_driver_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.driver_driver_id_seq', 1, false);


--
-- Name: request_request_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.request_request_id_seq', 1, false);


--
-- Name: vessel_vessel_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.vessel_vessel_id_seq', 1, false);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);


--
-- Name: driver driver_pkey; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.driver
    ADD CONSTRAINT driver_pkey PRIMARY KEY (driver_id);


--
-- Name: request request_pkey; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request
    ADD CONSTRAINT request_pkey PRIMARY KEY (request_id);


--
-- Name: dock vessel_pkey; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.dock
    ADD CONSTRAINT vessel_pkey PRIMARY KEY (dock_id);


--
-- Name: vessel vessel_pkey1; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.vessel
    ADD CONSTRAINT vessel_pkey1 PRIMARY KEY (vessel_id);


--
-- Name: driver driver_vessel_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.driver
    ADD CONSTRAINT driver_vessel_id_fkey FOREIGN KEY (vessel_id) REFERENCES gog_demo.vessel(vessel_id);


--
-- Name: request request_customer_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request
    ADD CONSTRAINT request_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES gog_demo.customer(customer_id);


--
-- Name: request request_departure_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request
    ADD CONSTRAINT request_departure_id_fkey FOREIGN KEY (departure_id) REFERENCES gog_demo.dock(dock_id);


--
-- Name: request request_destination_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request
    ADD CONSTRAINT request_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES gog_demo.dock(dock_id);


--
-- Name: request request_driver_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.request
    ADD CONSTRAINT request_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES gog_demo.driver(driver_id);


--
-- PostgreSQL database dump complete
--

-- cast for spring boot
CREATE CAST (CHARACTER VARYING as gog_demo.request_status) WITH INOUT AS IMPLICIT;