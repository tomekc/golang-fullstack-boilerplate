-- +goose Up
CREATE TABLE IF NOT EXISTS clients (
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT NOT NULL,
    company  TEXT NOT NULL,
    city     TEXT NOT NULL,
    progress INTEGER NOT NULL DEFAULT 0,
    created  TEXT NOT NULL
);

INSERT INTO clients (name, company, city, progress, created) VALUES
    ('Rebecca Bauch',           'Daugherty-Daniel',                  'South Cory',      79, 'Oct 25, 2020'),
    ('Felicita Yundt',          'Johns-Weissnat',                    'East Ariel',       67, 'Jan 8, 2020'),
    ('Mr. Larry Satterfield V', 'Hyatt Ltd',                         'Windlerburgh',     16, 'Dec 18, 2020'),
    ('Mr. Broderick Kub',       'Kshlerin, Bauch and Ernser',        'New Kirstenport',  71, 'Sep 13, 2020'),
    ('Barry Weber',             'Schulist, Mosciski and Heidenreich','East Violettestad',80, 'Jul 24, 2020');

-- +goose Down
DROP TABLE IF EXISTS clients;
