INSERT INTO `books` (`id`, `name`, `author`, `publication_date`, `library_id`)
VALUES
	('1f27d9b8-f387-11e5-9280-8438355dc0e0','Abstract of North Carolina wills','North Carolina','1910-01-01 10:19:31','cdc100c8-f385-11e5-9280-8438355dc0e0'),
	('37bbc0b6-f387-11e5-9280-8438355dc0e0','Telugu-English dictionary','Percival, P','1862-01-01 10:19:31','cdc100c8-f385-11e5-9280-8438355dc0e0'),
	('57740c24-f386-11e5-9280-8438355dc0e0','All the Rage','Courtney Summers','2015-04-14 12:47:33','a7c1d1d9-f385-11e5-9280-8438355dc0e0'),
	('885c39e0-f387-11e5-9280-8438355dc0e0','Tales of the Greek Heroes','Roger Lancelyn Green','1995-03-01 07:03:57','f023a6c9-f385-11e5-9280-8438355dc0e0'),
	('96fc9801-f386-11e5-9280-8438355dc0e0','Dairy queen','Catherine Gilbert Murdock','2006-05-22 11:23:13','a7c1d1d9-f385-11e5-9280-8438355dc0e0'),
	('a5d62080-f387-11e5-9280-8438355dc0e0','Running Out of Time','Margaret Peterson Haddix','1996-06-01 17:14:01','f023a6c9-f385-11e5-9280-8438355dc0e0'),
	('c9ed2e8e-f387-11e5-9280-8438355dc0e0','Totally Joe','James Howe','2007-01-24 09:56:46','f023a6c9-f385-11e5-9280-8438355dc0e0'),
	('cc4f6dbd-f386-11e5-9280-8438355dc0e0','Kissing in America','Margo Rabb','2015-05-26 16:13:08','a7c1d1d9-f385-11e5-9280-8438355dc0e0');

INSERT INTO `libraries` (`id`, `name`, `address`, `phone`)
VALUES
	('a7c1d1d9-f385-11e5-9280-8438355dc0e0','New York Public Library','5th Ave at 42nd St, New York, NY 10018, United-States','+1 917-275-6975'),
	('cdc100c8-f385-11e5-9280-8438355dc0e0','Library of Congress','101 Independence Ave SE, Washington, DC 20540, United-States','+1 202-707-5000'),
	('f023a6c9-f385-11e5-9280-8438355dc0e0','Tenley-Friendship Neighborhood Library','4450 Wisconsin Ave NW, Washington, DC 20016, United-States','+1 202-727-1488');