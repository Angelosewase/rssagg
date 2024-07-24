-- +gooseUp 
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) NOT NULL  DEFAULT '' ;

-- +gooseDown
ALTER TABLE users DROP COLUMN api_key ;
