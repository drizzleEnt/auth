-- +goose Up
ALTER TABLE users ADD password VARCHAR(255) NOT NULL;
