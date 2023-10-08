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
-- Name: driver_status; Type: TYPE; Schema: gog_demo; Owner: postgres
--

CREATE TYPE gog_demo.driver_status AS ENUM (
    'busy',
    'waiting',
    'afw'
);


ALTER TYPE gog_demo.driver_status OWNER TO postgres;

--
-- Name: taxi_request_status; Type: TYPE; Schema: gog_demo; Owner: postgres
--

CREATE TYPE gog_demo.taxi_request_status AS ENUM (
    'findingDriver',
    'waitingForDriver',
    'inProgress',
    'completed',
    'canceled'
);


ALTER TYPE gog_demo.taxi_request_status OWNER TO postgres;

--
-- Name: CAST (character varying AS gog_demo.taxi_request_status); Type: CAST; Schema: -; Owner: -
--

CREATE CAST (character varying AS gog_demo.taxi_request_status) WITH INOUT AS IMPLICIT;


--
-- Name: calculate_distance(double precision, double precision, double precision, double precision); Type: FUNCTION; Schema: gog_demo; Owner: postgres
--

CREATE FUNCTION gog_demo.calculate_distance(lat1 double precision, lon1 double precision, lat2 double precision, lon2 double precision) RETURNS double precision
    LANGUAGE plpgsql
    AS $$
    DECLARE
        dist double precision = 0;
        radlat1 double precision;
        radlat2 double precision;
        theta double precision;
        radtheta double precision;
    BEGIN
        IF lat1 = lat2 AND lon1 = lon2
            THEN RETURN dist;
        ELSE
            radlat1 = pi() * lat1 / 180;
            radlat2 = pi() * lat2 / 180;
            theta = lon1 - lon2;
            radtheta = pi() * theta / 180;
            dist = sin(radlat1) * sin(radlat2) + cos(radlat1) * cos(radlat2) * cos(radtheta);

            IF dist > 1 THEN dist = 1; END IF;

            dist = acos(dist);
            dist = dist * 180 / pi();
            dist = dist * 60 * 1.1515;
            dist = dist * 1.609344;

            RETURN dist;
        END IF;
    END;
$$;


ALTER FUNCTION gog_demo.calculate_distance(lat1 double precision, lon1 double precision, lat2 double precision, lon2 double precision) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: customer; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.customer (
    customer_id integer NOT NULL,
    phone character varying(16) NOT NULL,
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
    active boolean NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL
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
-- Name: taxi_request; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.taxi_request (
    taxi_request_id integer NOT NULL,
    customer_id integer,
    driver_id integer,
    departure_id integer,
    destination_id integer,
    price numeric(19,4) NOT NULL,
    status gog_demo.taxi_request_status NOT NULL
);


ALTER TABLE gog_demo.taxi_request OWNER TO postgres;

--
-- Name: taxi_request_taxi_request_id_seq; Type: SEQUENCE; Schema: gog_demo; Owner: postgres
--

CREATE SEQUENCE gog_demo.taxi_request_taxi_request_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gog_demo.taxi_request_taxi_request_id_seq OWNER TO postgres;

--
-- Name: taxi_request_taxi_request_id_seq; Type: SEQUENCE OWNED BY; Schema: gog_demo; Owner: postgres
--

ALTER SEQUENCE gog_demo.taxi_request_taxi_request_id_seq OWNED BY gog_demo.taxi_request.taxi_request_id;


--
-- Name: vessel; Type: TABLE; Schema: gog_demo; Owner: postgres
--

CREATE TABLE gog_demo.vessel (
    vessel_id integer NOT NULL,
    model character varying(100) NOT NULL,
    seats integer NOT NULL,
    is_approved boolean DEFAULT false NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL
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
-- Name: taxi_request taxi_request_id; Type: DEFAULT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request ALTER COLUMN taxi_request_id SET DEFAULT nextval('gog_demo.taxi_request_taxi_request_id_seq'::regclass);


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

COPY gog_demo.dock (dock_id, name, active, latitude, longitude) FROM stdin;
\.


--
-- Data for Name: driver; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.driver (driver_id, first_name, last_name, vessel_id, status, balance, cert_first_aid, cert_driving) FROM stdin;
\.


--
-- Data for Name: taxi_request; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.taxi_request (taxi_request_id, customer_id, driver_id, departure_id, destination_id, price, status) FROM stdin;
\.


--
-- Data for Name: vessel; Type: TABLE DATA; Schema: gog_demo; Owner: postgres
--

COPY gog_demo.vessel (vessel_id, model, seats, is_approved, latitude, longitude) FROM stdin;
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
-- Name: taxi_request_taxi_request_id_seq; Type: SEQUENCE SET; Schema: gog_demo; Owner: postgres
--

SELECT pg_catalog.setval('gog_demo.taxi_request_taxi_request_id_seq', 1, false);


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
-- Name: taxi_request taxi_request_pkey; Type: CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request
    ADD CONSTRAINT taxi_request_pkey PRIMARY KEY (taxi_request_id);


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
-- Name: taxi_request taxi_request_customer_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request
    ADD CONSTRAINT taxi_request_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES gog_demo.customer(customer_id);


--
-- Name: taxi_request taxi_request_departure_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request
    ADD CONSTRAINT taxi_request_departure_id_fkey FOREIGN KEY (departure_id) REFERENCES gog_demo.dock(dock_id);


--
-- Name: taxi_request taxi_request_destination_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request
    ADD CONSTRAINT taxi_request_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES gog_demo.dock(dock_id);


--
-- Name: taxi_request taxi_request_driver_id_fkey; Type: FK CONSTRAINT; Schema: gog_demo; Owner: postgres
--

ALTER TABLE ONLY gog_demo.taxi_request
    ADD CONSTRAINT taxi_request_driver_id_fkey FOREIGN KEY (driver_id) REFERENCES gog_demo.driver(driver_id);


--
-- PostgreSQL database dump complete
--

