/*
	CMPS4191 - Quiz #1
	Rene Sanchez - 2018118383
*/
ALTER TABLE entries
ADD COLUMN created_at timestamp(0) with time zone NOT NULL DEFAULT NOW();